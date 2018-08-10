package publisher

import (
	"github.com/whosonfirst/go-whosonfirst-repo"
	"io"
)

type Publisher interface {
	Publish(io.ReadCloser, string) error
	Fetch(string) (io.ReadCloser, error)
	Prune(repo.Repo) error // most likely a string rather than a repo.Repo
	Index(repo.Repo) error // most likely a string rather than a repo.Repo
	IsNotFound(error) bool
}
