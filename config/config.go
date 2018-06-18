package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/krzyszko/loaddriver/ess"
)

type Config struct {
	Components []ess.ComponentConfig `json:"components"`
}

func GetConfig(fileName string) (*Config, error) {
	config := &Config{}
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
