package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-dist-publish"
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

	r, err := repo.NewDataRepoFromString("whosonfirst-data")

	if err != nil {
		log.Fatal(err)
	}

	// THE NEW NEW... BUT NOT YET
	// err = publisher.Index(p, r)
	
	err = p.Index(r)

	if err != nil {
		log.Fatal(err)
	}

}
