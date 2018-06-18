package ess

import (
	"encoding/json"
)

type Component interface {
	Run() error
	Init(map[string][]byte) error
}

type ComponentConfig struct {
	Plugin string          `json:"plugin"`
	Params json.RawMessage `json:"params"`
}
