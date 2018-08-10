package main

// this is work in progress - you should assume that anything and
// everything might change still (20180728/thisisaaronland)

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/tidwall/pretty"
	"github.com/whosonfirst/go-whosonfirst-dist"
	"github.com/whosonfirst/go-whosonfirst-dist-publish"
	"github.com/whosonfirst/go-whosonfirst-dist-publish/publisher"
	"github.com/whosonfirst/go-whosonfirst-repo"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

// PLEASE RECONCILE ME WITH publishers/s3.go NOT TO MENTION
// go-whosonfirst-repo AND go-whosonfirst-dist

var re_distname *regexp.Regexp
var re_disttype *regexp.Regexp

// PLEASE RECONCILE ME WITH THE CODE IN publisher/s3.go

func init() {

	// whosonfirst-data-venue-us-ca-1533149830.tar.bz2
	// whosonfirst-data-venue-us-ny-latest.db.bz2

	re_distname = regexp.MustCompile(`([a-z\-]+)\-(\d+|latest)\.(.*)$`)

	// this needs to be moved in to go-whosonfirst-dist

	re_disttype = regexp.MustCompile(`x\-urn\:([^\:]+)\:([^\:]+)\:([^\#]+)(?:\#(.*))?`)
}

type PublishOptions struct {
	Workdir   string
	Publisher publisher.Publisher
}

func PublishInventory(inv *dist.Inventory, opts *PublishOptions) error {

	t1 := time.Now()

	defer func() {
		log.Printf("time to publish inventory %v\n", time.Since(t1))
	}()

	wg := new(sync.WaitGroup)

	for _, item := range *inv {

		wg.Add(1)

		go func(item *dist.Item) {

			defer wg.Done()

			err := PublishItem(item, opts)

			if err != nil {
				log.Printf("Failed to publish %s %s\n", item.Name, err)
			}

		}(item)
	}

	wg.Wait()

	return nil
}

func PublishItem(item *dist.Item, opts *PublishOptions) error {

	if !shouldPublish(item, opts) {
		log.Println("DO NOT PUBLISH", item.NameCompressed)
		return nil
	}

	nc := item.NameCompressed

	lu, err := time.Parse(time.RFC3339, item.LastUpdate)

	if err != nil {
		return err
	}

	// PLEASE RECONCILE ALL THIS STUFF WITH THE NAME/TYPE PARSING IN shouldPublish
	// (20180809/thisisaaronland)

	suffix := fmt.Sprintf("-%d.", lu.Unix())

	nc_ts := strings.Replace(nc, "-latest.", suffix, -1)

	// what is NewDistributionTypeFromString(t) ...

	t := item.Type
	t = strings.Replace(t, "x-urn:whosonfirst:", "", -1)

	var prefix string

	// this will all be made less-shit...

	if strings.HasPrefix(t, "csv:meta") {
		prefix = "meta"
	} else if strings.HasPrefix(t, "database:sqlite") {
		prefix = "sqlite"
	} else if strings.HasPrefix(t, "fs:bundle") {
		prefix = "bundles"
	} else {
		return errors.New("Invalid or unsupported prefix")
	}

	source := filepath.Join(opts.Workdir, nc)

	dest_ts := filepath.Join(prefix, nc_ts)
	dest_latest := filepath.Join(prefix, nc)

	err = publishFile(source, dest_ts, opts)

	if err != nil {
		return err
	}

	err = publishFile(source, dest_latest, opts)

	if err != nil {
		return err
	}

	// publish the inventory files

	inv_ts := dest_ts + ".json"
	inv_latest := dest_latest + ".json"

	enc_latest, err := json.Marshal(item)

	if err != nil {
		return err
	}

	// quick! look over there!!
	// make sure the pointers in the inventory file point
	// to the relevant distributions

	// only tweak the compressed name since it will still uncompress
	// with the -latest suffix (20180802/thisisaaronland)

	item.NameCompressed = nc_ts

	enc_ts, err := json.Marshal(item)

	if err != nil {
		return err
	}

	enc_latest = pretty.Pretty(enc_latest)
	enc_ts = pretty.Pretty(enc_ts)

	err = publishBytes(enc_ts, inv_ts, opts)

	if err != nil {
		return err
	}

	err = publishBytes(enc_latest, inv_latest, opts)

	if err != nil {
		return err
	}

	return nil
}

// PLEASE RENAME ME...

func shouldPublish(item *dist.Item, opts *PublishOptions) bool {

	publish := true

	/*
		if item latest: copy in to tmp var
		else: create tmp var and replace "-{TIMESTAMP}" with "-latest" <-- this should never happen since the "-{TIMESTAMP}" version is
		      	     	     	 	 		     	       	   cloned below in PublishItem but you get the idea...
		then:
		fetch tmp-latest
		compare fetch-latest sha256_commpressed with item sha256_compressed
		if same: skip
	*/

	// PLEASE RECONCILE THIS WILL THE CODE TO GENERATE PUBLISH KEYS
	// BELOW... (20180809/thisisaaronland)

	m_name := re_distname.FindAllStringSubmatch(item.NameCompressed, -1)
	m_type := re_disttype.FindAllStringSubmatch(item.Type, -1)

	if len(m_name) == 0 || len(m_type) == 0 {

		log.Printf("Failed to parse %s (%d) or %s (%d) which is weird...\n", item.NameCompressed, len(m_name), item.Type, len(m_type))

	} else {

		tmp_name := fmt.Sprintf("%s-latest.%s.json", m_name[0][1], m_name[0][3])
		tmp_minor := m_type[0][3]

		if tmp_minor == "bundle" {
			tmp_minor = "bundles" // ARGHHHHHHHHHHPPPPTPPFFFFFFHHHHTTTT
		}

		tmp_key := fmt.Sprintf("%s/%s", tmp_minor, tmp_name)

		var tmp_item *dist.Item

		tmp_fh, err := opts.Publisher.Fetch(tmp_key)

		ok := true

		if err != nil {

			if !opts.Publisher.IsNotFound(err) {
				log.Printf("failed to fetch %s for comparing with %s: %s\n", tmp_key, item.NameCompressed, err)
			}

			ok = false
		}

		if ok {

			tmp_body, err := ioutil.ReadAll(tmp_fh)

			if err != nil {
				log.Printf("failed to read %s for comparing with %s: %s\n", tmp_key, item.NameCompressed, err)
			} else {

				err = json.Unmarshal(tmp_body, &tmp_item)

				if err != nil {
					log.Printf("failed to parse %s for comparing with %s: %s\n", tmp_key, item.NameCompressed, err)
					ok = false
				}
			}
		}

		if ok {

			if item.Sha256Compressed == tmp_item.Sha256Compressed {
				publish = false
			}

			if !publish {
				log.Printf("sha256 hashes for remote#%s and local#%s match: do not republish\n", tmp_key, item.NameCompressed)
			}
		}
	}

	return publish
}

func publishFile(source string, dest string, opts *PublishOptions) error {

	fh, err := os.Open(source)

	if err != nil {
		return err
	}

	defer fh.Close()

	return opts.Publisher.Publish(fh, dest)
}

func publishBytes(b []byte, dest string, opts *PublishOptions) error {

	r := bytes.NewReader(b)
	fh := ioutil.NopCloser(r)

	return opts.Publisher.Publish(fh, dest)
}

func main() {

	workdir := flag.String("workdir", "", "Where to read build files from. If empty the code will attempt to use the current working directory.")

	pub := flag.String("publisher", "s3", "Valid publishers are: s3")
	dsn := flag.String("publisher-dsn", "", "A valid DSN string for your distribution publisher.")

	flag.Parse()

	p, err := publish.NewPublisher(*pub, *dsn)

	if err != nil {
		log.Fatal(err)
	}

	opts := &PublishOptions{
		Workdir:   *workdir,
		Publisher: p,
	}

	for _, repo_name := range flag.Args() {

		r, err := repo.NewDataRepoFromString(repo_name)

		if err != nil {
			log.Fatal(err)
		}

		// PLEASE FIX ME... this should be in a library...

		fname := fmt.Sprintf("%s-inventory.json", r.Name())
		path := filepath.Join(*workdir, fname)

		fh, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer fh.Close()

		body, err := ioutil.ReadAll(fh)

		if err != nil {
			log.Fatal(err)
		}

		var inv *dist.Inventory

		err = json.Unmarshal(body, &inv)

		if err != nil {
			log.Fatal(err)
		}

		// ctx, cancel := context.WithCancel(context.Background())
		// defer cancel()

		PublishInventory(inv, opts)
	}
}
