package controlers

import (
	"github.com/krzyszko/loaddriver/plan"
)

type Shaper interface {
	SetDuration(int64) error
	plan.Component
}
