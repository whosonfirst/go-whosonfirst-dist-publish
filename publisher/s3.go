package publisher

import (
	"github.com/whosonfirst/go-whosonfirst-aws/s3"
	"github.com/whosonfirst/go-whosonfirst-dist-publish"
	"io"
)

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

	return p.conn.Put(dest, fh)
}
