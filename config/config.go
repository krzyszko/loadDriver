package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Component struct {
	Plugin string                 `json:"plugin"`
	Params map[string]interface{} `json:"params"`
}

type Config struct {
	Components []*Component `json:"components"`
}

func GetComponents(fileName string) (*Config, error) {
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

func GetComponentParams(params map[string]interface{}) []byte {
	content := ""
	for k, v := range params {
		content += fmt.Sprintf("\"%s\":\"%s\",", k, v)
	}
	return []byte(fmt.Sprintf("{%s}", content))
}
