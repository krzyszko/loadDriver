package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type httpSampler struct {
	ReqPayload  string          `json:"request"`
	RespPayLoad string          `json:"response"`
	MethodName  string          `json:"method_name"`
	URL         string          `json:"url"`
	Components  json.RawMessage `json:"components"`
}

func (h *httpSampler) Init(registry map[string]interface{}) error {
	var httpMethod func(string) (*http.Response, error)
	switch h.MethodName {
	case "GET":
		httpMethod = http.Get
	default:
		httpMethod = http.Get
	}
	h.method = func() error {
		//req := registry[h.reqPayload]
		resp, err := httpMethod(h.URL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		registry[h.RespPayLoad] = body
		return nil
	}
	return nil
}

func (h *httpSampler) Run() error {
	err := h.method()
	if err != nil {
		return err
	}
	return nil
}

var Plugin httpSampler
