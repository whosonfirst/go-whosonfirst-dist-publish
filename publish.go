package publish

import (
	"errors"
	"github.com/whosonfirst/go-whosonfirst-dist-publish/publisher"
	"strings"
)

func NewPublisher(name string, dsn string) (publisher.Publisher, error) {

	var p publisher.Publisher
	var err error

	switch strings.ToUpper(name) {

	case "S3":

		p, err = publisher.NewS3PublisherFromDSN(dsn)

	default:
		err = errors.New("Unknown or invalid publisher name")
	}

	return p, err
}
