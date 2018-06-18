package processors

import (
	"io"

	"github.com/krzyszko/loaddriver/ess"
)

type Processor interface {
	io.Writer
	io.Reader
	ess.Component
}
