package loader

import (
	"encoding/json"
	"fmt"
	"plugin"
	"reflect"
	"runtime"

	"github.com/apex/log"
	"github.com/krzyszko/loaddriver/config"
	"github.com/krzyszko/loaddriver/ess"
)

type LoaderError struct {
	error string
}

func (l LoaderError) Error() string { return l.error }

func ComponentsFromConfiguration(c *config.Config) ([]ess.Component, error) {
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
	creator, err := p.Lookup("Handler")
	if err != nil {
		return nil, err
	}
	fcreate := creator.(func() ess.Component)
	cpnt := fcreate()
	err = json.Unmarshal(c.Params, cpnt)
	if err != nil {
		return nil, err
	}
	cnfComponents := &struct {
		Components []ess.ComponentConfig
	}{}
	err = json.Unmarshal(c.Params, cnfComponents)
	if err != nil {
		return nil, err
	}
	cpnts := reflect.Indirect(reflect.ValueOf(cpnt)).FieldByName("Cpnts")
	if len(cnfComponents.Components) > 0 {
		log.Debug("Evaluating components")
		if cpnts.IsValid() && cpnts.CanSet() {
			components, err := loadComponents(cnfComponents.Components)
			if err != nil {
				return nil, err
			}
			cpnts.Set(reflect.ValueOf(components))
		} else {
			_, fn, line, _ := runtime.Caller(0)
			log.Errorf("Components configured but not expected by %s:%d", fn, line)
		}
	}
	component, ok := cpnt.(ess.Component)
	if !ok {
		_, fn, line, _ := runtime.Caller(0)
		return nil, &LoaderError{fmt.Sprintf("%s:%d: cpnt doesn't implement interface", fn, line)}
	}
	return component, nil
}
