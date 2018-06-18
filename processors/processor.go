package processors

import (
	"io"

	"github.com/krzyszko/loaddriver/imp"
)

type Processor interface {
	io.Writer
	io.Reader
	imp.Component
}
