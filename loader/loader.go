package loader

import (
	"encoding/json"
	"plugin"

	"github.com/krzyszko/loaddriver/config"
	"github.com/krzyszko/loaddriver/ess"
)

type LoaderError struct {
	error string
}

func (l LoaderError) Error() string { return l.error }

func LoadFromConfiguration(c *config.Config) ([]ess.Component, error) {
	components, err := loadComponents(c.Components)
	if err != nil {
		return nil, err
	}
	return components, nil
}

func loadComponents(c []ess.ComponentConfig) ([]ess.Component, error) {
	var components []ess.Component
	for _, componentConfig := range c {
		component, err := loadComponent(componentConfig)
		if err != nil {
			return nil, err
		}
		components = append(components, component)
	}
	return components, nil
}

func loadComponent(c ess.ComponentConfig) (ess.Component, error) {
	p, err := plugin.Open(c.Plugin)
	if err != nil {
		return nil, err
	}
	cpnt, err := p.Lookup("Plugin")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(c.Params, cpnt)
	if err != nil {
		return nil, err
	}
	//to do evaluate components

	component, ok := cpnt.(ess.Component)
	if !ok {
		return nil, &LoaderError{"Unable to cast to ess.Component"}
	}
	return component, nil
}
