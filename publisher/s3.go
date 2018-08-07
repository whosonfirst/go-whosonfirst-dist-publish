package publisher

import (
	"encoding/json"
	"fmt"
	"github.com/whosonfirst/go-bindata-html-template"
	"github.com/whosonfirst/go-whosonfirst-aws/s3"
	"github.com/whosonfirst/go-whosonfirst-dist"
	"github.com/whosonfirst/go-whosonfirst-dist-publish"
	"github.com/whosonfirst/go-whosonfirst-dist-publish/assets/html"
	"github.com/whosonfirst/go-whosonfirst-repo"
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

func init() {

	// whosonfirst-data-venue-us-ca-1533149830.tar.bz2
	// whosonfirst-data-venue-us-ny-latest.db.bz2
	// re_distname_dated = regexp.MustCompile(`([a-z\-]+)\-(\d+)\..*$`)

	re_distname = regexp.MustCompile(`([a-z\-]+)\-(\d+|latest)\..*$`)
}

type HTMLVars struct {
	Date  string
	Items []*dist.Item
}

type S3Publisher struct {
	publish.Publisher
	conn *s3.S3Connection
	cfg  *s3.S3Config
}

func NewS3PublisherFromDSN(dsn string) (publish.Publisher, error) {

	cfg, err := s3.NewS3ConfigFromString(dsn)

	if err != nil {
		return nil, err
	}

	return NewS3Publisher(cfg)
}

func NewS3Publisher(cfg *s3.S3Config) (publish.Publisher, error) {

	conn, err := s3.NewS3Connection(cfg)

	if err != nil {
		return nil, err
	}

	p := S3Publisher{
		conn: conn,
		cfg:  cfg,
	}

	return &p, nil
}

func (p *S3Publisher) Publish(fh io.ReadCloser, dest string) error {

	key := fmt.Sprintf("%s#ACL=public-read", dest)

	t1 := time.Now()

	defer func() {
		log.Printf("time to publish %s: %v\n", key, time.Since(t1))
	}()

	return p.conn.Put(key, fh)
}

// THIS NEEDS OPTIONS AND FLAGS
// IT IS NOT CLEAR THIS NEEDS OR SHOULD BE A repo.Repo THINGY

func (p *S3Publisher) Prune(r repo.Repo) error {

	max_pubs := 1 // sudo make me a config option somewhere...

	// grouped is a make(map[string]map[string][]*s3.S3Object)
	// which is not ideal but it's also too soon to optimize...

	grouped, err := p.group(r)

	if err != nil {
		return err
	}

	to_prune := make([]*s3.S3Object, 0)

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

		if count <= max_pubs {
			continue
		}

		sort.Sort(sort.Reverse(sort.IntSlice(pubdates)))
		// log.Println(repo_name, pubdates)

		for i := max_pubs; i < count; i++ {

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

		go func(obj *s3.S3Object) {

			defer wg.Done()

			key := obj.Key // remember this is *s3.S3Object Key and _not_ KeyRaw (because of p.conn.prefix)

			err := p.conn.Delete(key)

			if err != nil {
				log.Printf("Failed to delete %s because %s", key, err)
			}

		}(obj)
	}

	wg.Wait()

	return nil
}

// this interface will likely change to be ([]*dist.Item, error)
// but not today... (20180807/thisisaaronland)

func (p *S3Publisher) Index(r repo.Repo, wr io.Writer) error {

	items := make([]*dist.Item, 0)

	grouped, err := p.group(r)

	if err != nil {
		return err
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

		var latest []*s3.S3Object

		pubdates := make([]int, 0)

		for str_ts, _ := range details {

			if str_ts == "latest" {
				latest = details[str_ts]
				continue
			}

			ts, err := strconv.Atoi(str_ts)

			if err != nil {
				return err
			}

			pubdates = append(pubdates, ts)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(pubdates)))

		// here is the part where we need to start generating some data
		// but the question is: how to share all this code between html
		// representations and syndication feeds and and and - meanwhile
		// as this is written we are not creating per-type indices (meta,
		// sqlite, bundles...)

		// out := fmt.Sprintf("%s latest %v\n", repo_name, latest)
		// wr.Write([]byte(out))

		objects := make([][]*s3.S3Object, 0)

		objects = append(objects, latest)

		for _, ts := range pubdates {

			str_ts := strconv.Itoa(ts)
			objects = append(objects, details[str_ts])
		}

		for _, o := range objects {

			ts_items, err := p.appendObjectsToItems(o)

			if err != nil {
				return err
			}

			for _, i := range ts_items {
				items = append(items, i)
			}
		}
	}

	log.Println("ITEMS", len(items))

	// It seems likely that all of this template code will get
	// moved in to some more generic package since there's nothing
	// S3 specific about it and could be shared across "publishers"
	// ... but not today (20180807/thisisaaronland)

	// although it is true that all this template stuff could
	// be method-chained I find that it doesn't take long for
	// method-chaining to become inpenetrable gibberish so why
	// start now (20180807/thisisaaronland)

	t := template.New("inventory", html.Asset)

	funcs := template.FuncMap{}
	t = t.Funcs(funcs)

	t, err = t.ParseFiles("html/inventory.html")

	if err != nil {
		return err
	}

	now := time.Now()

	vars := HTMLVars{
		Date:  now.Format(time.RFC3339),
		Items: items,
	}

	log.Println("GO", now)

	err = t.Execute(wr, vars)

	if err != nil {
		return err
	}

	return nil
}

func (p *S3Publisher) appendObjectsToItems(objects []*s3.S3Object) ([]*dist.Item, error) {

	items := make([]*dist.Item, 0)

	for _, o := range objects {

		ext := filepath.Ext(o.Key)

		if ext != ".json" {
			continue
		}

		// filter by "type": "x-urn:whosonfirst:fs:{TYPE} here...

		// out := fmt.Sprintf("%s %s %v\n", repo_name, str_ts, o)
		// wr.Write([]byte(out))

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

func (p *S3Publisher) group(r repo.Repo) (map[string]map[string][]*s3.S3Object, error) {

	grouped := make(map[string]map[string][]*s3.S3Object)

	// Basically what we're after is something like this:
	//
	// + whosonfirst-data-venue-us-ca
	//   + 123455
	//    - ...csv.bz2
	//    - ...db.bz2

	mu := new(sync.RWMutex)

	cb := func(obj *s3.S3Object) error {

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
			by_ts = make(map[string][]*s3.S3Object, 0)
		}

		by_ts[str_ts] = append(by_ts[str_ts], obj)

		grouped[group] = by_ts
		return nil
	}

	opts := s3.DefaultS3ListOptions()
	err := p.conn.List(cb, opts)

	if err != nil {
		return nil, err
	}

	return grouped, nil
}
