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

	flag.Parse()

	p, err := publish.NewPublisher(*pub, *dsn)

	if err != nil {
		log.Fatal(err)
	}

	opts, err := publisher.DefaultIndexerOptions()

	if err != nil {
		log.Fatal(err)
	}

	for _, repo_name := range flag.Args() {

		r, err := repo.NewDataRepoFromString(repo_name)

		if err != nil {
			log.Fatal(err)
		}

		err = publisher.Index(p, r, opts)

		if err != nil {
			log.Fatal(err)
		}
	}
}
