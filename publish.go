package publish

import (
	"github.com/whosonfirst/go-whosonfirst-repo"
	"io"
)

type Publisher interface {
	Publish(io.ReadCloser, string) error
	Prune(repo.Repo) error
}
