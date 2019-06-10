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
	max := flag.Int("max-distributions", 10, "The maximum number of revisions (distributions) to keep for a repo.")

	flag.Parse()

	p, err := publish.NewPublisher(*pub, *dsn)

	if err != nil {
		log.Fatal(err)
	}

	opts, err := publisher.NewDefaultPruneOptions()

	if err != nil {
		log.Fatal(err)
	}

	opts.MaxDistributions = *max

	for _, repo_name := range flag.Args() {

		r, err := repo.NewDataRepoFromString(repo_name)

		if err != nil {
			log.Fatal(err)
		}

		p.Prune(r, opts)
	}
}
