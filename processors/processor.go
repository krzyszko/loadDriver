package processors

import (
	"io"

	"github.com/krzyszko/loaddriver/plan"
)

type Processor interface {
	io.Writer
	io.Reader
	plan.Component
}
