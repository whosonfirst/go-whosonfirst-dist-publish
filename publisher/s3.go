package publisher

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-aws/s3"
	"github.com/whosonfirst/go-whosonfirst-dist-publish"
	"github.com/whosonfirst/go-whosonfirst-repo"
	"io"
	"log"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"
)

var re_distname *regexp.Regexp

func init() {

	// whosonfirst-data-venue-us-ca-1533149830.tar.bz2
	// whosonfirst-data-venue-us-ny-latest.db.bz2
	// see the we're excluding -latest ?

	re_distname = regexp.MustCompile(`([a-z\-]+)\-(\d+)\..*$`)
}

type S3Publisher struct {
	publish.Publisher
	conn *s3.S3Connection
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

func (p *S3Publisher) Prune(r repo.Repo) error {

	max_pubs := 1 // sudo make me a config option somewhere...

	// See this? It's ugly as all get out

	grouped := make(map[string]map[string][]*s3.S3Object)

	// Basically what we're after is something like this:
	//
	// + whosonfirst-data-venue-us-ca
	//   + 123455
	//    - ...csv.bz2
	//    - ...db.bz2

	cb := func(obj *s3.S3Object) error {

		fname := filepath.Base(obj.Key)

		m := re_distname.FindAllStringSubmatch(fname, -1)

		if len(m) == 0 {
			return nil
		}

		group := m[0][1]

		// FIX ME TO ACCOUNT FOR REPO-{PLACETYPE}...

		if group != r.Name() {
			return nil
		}

		str_ts := m[0][2]

		by_ts, ok := grouped[group]

		if !ok {
			by_ts = make(map[string][]*s3.S3Object, 0)
		}

		by_ts[str_ts] = append(by_ts[str_ts], obj)

		grouped[group] = by_ts

		log.Println(obj.Key, group, str_ts)
		return nil
	}

	err := p.conn.List(cb)

	if err != nil {
		return err
	}

	for repo_name, details := range grouped {

		pubdates := make([]int, 0)

		for str_ts, _ := range details {

			ts, err := strconv.Atoi(str_ts)

			if err != nil {
				return err
			}

			pubdates = append(pubdates, ts)
		}

		log.Println(repo_name, pubdates)

		count := len(pubdates)

		if count <= max_pubs {
			continue
		}

		sort.Sort(sort.Reverse(sort.IntSlice(pubdates)))

		for i := max_pubs; i < count; i++ {

			ts := pubdates[i]
			str_ts := strconv.Itoa(ts)

			for _, obj := range details[str_ts] {
				log.Println("PRUNE", repo_name, obj.Key)
			}
		}

	}

	return nil
}
