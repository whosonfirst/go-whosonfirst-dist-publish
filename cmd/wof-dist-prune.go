package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-dist-publish"
	"github.com/whosonfirst/go-whosonfirst-dist-publish/publisher"
	"github.com/whosonfirst/go-whosonfirst-repo"
	"log"
)

func PruneInventory(p publish.Publisher, r repo.Repo) error {
	return nil
}

func main() {

	pub := flag.String("publisher", "s3", "Valid publishers are: s3")
	dsn := flag.String("publisher-dsn", "", "A valid DSN string for your distribution publisher.")

	flag.Parse()

	var p publish.Publisher

	switch *pub {

	case "s3":

		s3_p, err := publisher.NewS3PublisherFromDSN(*dsn)

		if err != nil {
			log.Fatal(err)
		}

		p = s3_p

	default:
		log.Fatal("Invalid publisher")
	}

	for _, repo_name := range flag.Args() {

		r, err := repo.NewDataRepoFromString(repo_name)

		if err != nil {
			log.Fatal(err)
		}

		p.Prune(r)
	}
}
