package main

/*

this is basically the following written in Go:

aws --profile {PROFILE} s3 ls s3://dist.whosonfirst.org/bundles/ | grep -v json | grep bz2 > index.txt
aws --profile {PROFILE} s3 cp --acl public-read index.txt s3://dist.whosonfirst.org/bundles/
rm index.txt

it is included here so that the functionality exists with the same -publisher-dsn that all the other
tools use and so that there is a pure-go (no python dependencies, etc) tool; also the 'index.txt' file
is a legacy file used by some services that consume who's on first (20180822/thisisaaronland)

*/

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-aws/s3"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

func main() {

	dsn := flag.String("publisher-dsn", "", "A valid DSN string for your distribution publisher.")

	flag.Parse()

	cfg, err := s3.NewS3ConfigFromString(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	conn, err := s3.NewS3Connection(cfg)

	if err != nil {
		log.Fatal(err)
	}

	bundles := make([]string, 0)
	mu := new(sync.RWMutex)

	cb := func(obj *s3.S3Object) error {

		if strings.HasSuffix(obj.Key, ".bz2") {

			mu.Lock()
			defer mu.Unlock()

			bundles = append(bundles, obj.Key)
		}

		return nil
	}

	opts := s3.DefaultS3ListOptions()

	conn.List(cb, opts)
	sort.Strings(bundles)

	str_bundles := strings.Join(bundles, "\n")

	r := strings.NewReader(str_bundles)
	fh := ioutil.NopCloser(r)

	key := "index.txt#ACL=public-read"

	err = conn.Put(key, fh)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
