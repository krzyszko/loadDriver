package samplers

import (
	"io"

	"github.com/krzyszko/loaddriver/plan"
)

type Sampler interface {
	plan.Component
	io.Writer
	io.Reader
}
