package ess

import (
	"encoding/json"
)

type Component interface {
	Run() error
	Init(map[string]interface{}) error
}

type ComponentConfig struct {
	Plugin string          `json:"plugin"`
	Params json.RawMessage `json:"params"`
}
