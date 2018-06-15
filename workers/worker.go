package workers

import (
	"github.com/krzyszko/loaddriver/configs"
	"github.com/krzyszko/loaddriver/shapers"
	"github.com/krzyszko/loaddriver/transports"
)

type Worker interface {
	SetShaper(shapers.Shaper)
	IsMaster() bool
	Start() error
	Stop() error
	SetTransport(transports.Transport)
	SetMaster(Worker) error
	SetConfig(configs.Config)
}
