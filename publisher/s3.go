package publisher

import (
	"errors"
	"fmt"
)


func NewS3PublisherFromDSN(dsn string) (Publisher, error) {

	return nil, errors.New("Please parse DSN with strings...")
}

func NewS3Publisher(cfg *s3.S3Config) (Publisher, error) {

	uri := fmt.Sprintf("s3://%s?region=%s&prefix=%s&credentials=%s", cfg.Bucket, cfg.Region, cfg.Prefix, cfg.Credentials)
	return NewBlobPublisher(uri)
}
