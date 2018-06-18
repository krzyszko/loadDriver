package samplers

import (
	"io"

	"github.com/krzyszko/loaddriver/ess"
)

type Sampler interface {
	ess.Component
	io.Writer
	io.Reader
}
