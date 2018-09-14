package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-dist-publish"
	"github.com/whosonfirst/go-whosonfirst-dist-publish/publisher"
	"github.com/whosonfirst/go-whosonfirst-repo"
	"log"
)

func main() {

	pub := flag.String("publisher", "s3", "Valid publishers are: s3")
	dsn := flag.String("publisher-dsn", "", "A valid DSN string for your distribution publisher.")

	dist_name := flag.String("distribution-name", "Who's On First", "...")
	dist_root := flag.String("distribution-root-url", "https://dist.whosonfirst.org/", "...")
	dist_blurb := flag.String("distribution-blurb", `Who's On First is a gazetter of all the places. Note: As of this writing "alt" (or "alternative") files are not included in any of the distributions. If you need that data you will need to clone it directly from the https://github.com/whosonfirst-data GitHub organization.`, "...")

	flag.Parse()

	p, err := publish.NewPublisher(*pub, *dsn)

	if err != nil {
		log.Fatal(err)
	}

	r, err := repo.NewDataRepoFromString("whosonfirst-data")

	if err != nil {
		log.Fatal(err)
	}

	opts, err := publisher.NewDefaultIndexOptions()

	if err != nil {
		log.Fatal(err)
	}

	opts.DistributionName = *dist_name
	opts.DistributionRootURL = *dist_root
	opts.DistributionBlurb = *dist_blurb

	err = publisher.Index(p, r, opts)

	if err != nil {
		log.Fatal(err)
	}

}
