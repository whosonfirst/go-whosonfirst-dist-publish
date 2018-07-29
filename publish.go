package publish

import (
	"io"
)

type Publisher interface {
	Publish(io.ReadCloser, string) error
}
