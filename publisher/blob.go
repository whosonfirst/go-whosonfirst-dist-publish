package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-dist"
	"github.com/whosonfirst/go-whosonfirst-repo"
	"github.com/aaronland/gocloud-blob-bucket"
	"gocloud.dev/blob"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var re_distname *regexp.Regexp
var re_disttype *regexp.Regexp

func init() {

	// whosonfirst-data-venue-us-ca-1533149830.tar.bz2
	// whosonfirst-data-venue-us-ny-latest.db.bz2

	re_distname = regexp.MustCompile(`([a-z\-0-9]+)\-(\d+|latest)\..*$`)

	// this needs to be moved in to go-whosonfirst-dist

	re_disttype = regexp.MustCompile(`x\-urn\:([^\:]+)\:([^\:]+)\:([^\#]+)(?:\#(.*))?`)
}

type BlobPublisher struct {
	Publisher
	bucket *blob.Bucket
}

func NewBlobPublisherFromDSN(dsn string) (Publisher, error) {

	ctx := context.Background()
	b, err := bucket.OpenBucket(ctx, dsn)
	
	if err != nil {
		return nil, err
	}

	p := BlobPublisher{
		bucket: b,
	}

	return &p, nil
}

func (p *BlobPublisher) IsNotFound(err error) bool {
	return true	 // PLEASE FIX ME
}

func (p *BlobPublisher) Fetch(key string) (io.ReadCloser, error) {

	ctx := context.Background()
	return p.bucket.NewReader(ctx, key, nil)
}

func (p *BlobPublisher) Publish(fh io.ReadCloser, dest string) error {

	// ARGH... making S3 things public with Go Cloud... SAD FACE
	
	key := fmt.Sprintf("%s#ACL=public-read", dest)

	t1 := time.Now()

	defer func() {
		log.Printf("time to publish %s: %v\n", key, time.Since(t1))
	}()

	ctx := context.Background()
	
	wr, err := p.bucket.NewWriter(ctx, key, nil)

	if err != nil {
		return err
	}

	_, err = io.Copy(wr, fh)

	if err != nil {
		return err
	}

	return wr.Close()
}

// THIS NEEDS OPTIONS AND FLAGS
// IT IS NOT CLEAR THIS NEEDS OR SHOULD BE A repo.Repo THINGY

func (p *BlobPublisher) Prune(r repo.Repo, opts *PruneOptions) error {

	// grouped is a make(map[string]map[string][]*s3.BlobObject)
	// which is not ideal but it's also too soon to optimize...

	grouped, err := p.group(r)

	if err != nil {
		return err
	}

	to_prune := make([]*s3.BlobObject, 0)

	for _, details := range grouped {

		pubdates := make([]int, 0)

		for str_ts, _ := range details {

			if str_ts == "latest" {
				continue
			}

			ts, err := strconv.Atoi(str_ts)

			if err != nil {
				return err
			}

			pubdates = append(pubdates, ts)
		}

		count := len(pubdates)

		if count <= opts.MaxDistributions {
			continue
		}

		sort.Sort(sort.Reverse(sort.IntSlice(pubdates)))
		// log.Println(repo_name, pubdates)

		for i := opts.MaxDistributions; i < count; i++ {

			ts := pubdates[i]
			str_ts := strconv.Itoa(ts)

			for _, obj := range details[str_ts] {
				to_prune = append(to_prune, obj)
			}
		}

	}

	// we are using a waitgroup rather than channels so if there's a
	// problem then it will only be logged and not stop the execution
	// of other deletions - obviously the code will need to be changed
	// if that's a problem some day... (20180804/thisisaaronland)

	wg := new(sync.WaitGroup)

	for _, obj := range to_prune {

		wg.Add(1)

		go func(obj *s3.BlobObject) {

			defer wg.Done()

			key := obj.Key // remember this is *s3.BlobObject Key and _not_ KeyRaw (because of p.conn.prefix)

			ctx := ctx.Background()
			
			err := p.bucket.Delete(ctx, key)

			if err != nil {
				log.Printf("Failed to delete %s because %s", key, err)
			}

		}(obj)
	}

	wg.Wait()

	return nil
}

func (p *BlobPublisher) BuildIndex(r repo.Repo) (map[string][]*dist.Item, error) {

	items := make(map[string][]*dist.Item)

	grouped, err := p.group(r)

	if err != nil {
		return nil, err
	}

	repos := make([]string, 0)

	for repo_name, _ := range grouped {
		repos = append(repos, repo_name)
	}

	sort.Strings(repos)

	for _, repo_name := range repos {

		details, ok := grouped[repo_name]

		if !ok {
			continue // how would this even be possible... ?
		}

		var latest []*s3.BlobObject

		pubdates := make([]int, 0)

		for str_ts, _ := range details {

			if str_ts == "latest" {
				latest = details[str_ts]
				continue
			}

			ts, err := strconv.Atoi(str_ts)

			if err != nil {
				return nil, err
			}

			pubdates = append(pubdates, ts)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(pubdates)))

		objects := make([][]*s3.BlobObject, 0)

		objects = append(objects, latest)

		for _, ts := range pubdates {

			str_ts := strconv.Itoa(ts)
			objects = append(objects, details[str_ts])
		}

		for _, o := range objects {

			o_items, err := p.appendObjectsToItems(o)

			if err != nil {
				return nil, err
			}

			for _, i := range o_items {

				// again, this is all code that should be in go-whosonfirst-dist proper...

				m := re_disttype.FindAllStringSubmatch(i.Type, -1)

				if len(m) == 0 {
					log.Println("Unable to parse distribution type", i.Type)
					continue
				}

				// [[x-urn:whosonfirst:database:sqlite#common whosonfirst database sqlite common]]

				// major := m[0][2]
				minor := m[0][3]

				t := minor // should this be major/minor leaving the details to some other method?

				t_items, ok := items[t]

				if !ok {
					t_items = make([]*dist.Item, 0)
				}

				t_items = append(t_items, i)
				items[t] = t_items
			}
		}
	}

	return items, nil
}

func (p *BlobPublisher) appendObjectsToItems(objects []*s3.BlobObject) ([]*dist.Item, error) {

	items := make([]*dist.Item, 0)

	for _, o := range objects {

		ext := filepath.Ext(o.Key)

		if ext != ".json" {
			continue
		}

		r, err := p.conn.Get(o.Key)

		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(r)

		if err != nil {
			return nil, err
		}

		var item dist.Item

		err = json.Unmarshal(body, &item)

		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	return items, nil
}

// this is its own method because we'll probably need and want it for generating
// index pages

// mmmmmaybe? pre-sort everything before we return things?

func (p *BlobPublisher) group(r repo.Repo) (map[string]map[string][]*s3.BlobObject, error) {

	grouped := make(map[string]map[string][]*s3.BlobObject)

	// Basically what we're after is something like this:
	//
	// + whosonfirst-data-venue-us-ca
	//   + 123455
	//    - ...csv.bz2
	//    - ...db.bz2

	mu := new(sync.RWMutex)

	cb := func(obj *s3.BlobObject) error {

		fname := filepath.Base(obj.Key)

		m := re_distname.FindAllStringSubmatch(fname, -1)

		if len(m) == 0 {
			return nil
		}

		group := m[0][1]

		if !strings.HasPrefix(group, r.Name()) {
			return nil
		}

		str_ts := m[0][2]

		mu.Lock()
		defer mu.Unlock()

		by_ts, ok := grouped[group]

		if !ok {
			by_ts = make(map[string][]*s3.BlobObject, 0)
		}

		by_ts[str_ts] = append(by_ts[str_ts], obj)

		grouped[group] = by_ts
		return nil
	}

	opts := s3.DefaultBlobListOptions()
	err := p.conn.List(cb, opts)

	if err != nil {
		return nil, err
	}

	return grouped, nil
}
