package plan

import (
	"github.com/apex/log"
	"github.com/krzyszko/loaddriver/ess"
)

type Plan struct {
	components []ess.Component
	registry   map[string]interface{}
}

func (p *Plan) AddComponent(c ess.Component) {
	p.components = append(p.components, c)
}

func (p *Plan) Run() error {
	log.Info("Running plan...")
	p.registry = make(map[string]interface{})
	defer func() { p.registry = nil }()
	for _, c := range p.components {
		err := c.Init(p.registry)
		if err != nil {
			return err
		}
	}
	for _, c := range p.components {
		err := c.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
