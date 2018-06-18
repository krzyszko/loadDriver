package samplers

import (
	"io"

	"github.com/krzyszko/loaddriver/imp"
)

type Sampler interface {
	imp.Component
	io.Writer
	io.Reader
}
