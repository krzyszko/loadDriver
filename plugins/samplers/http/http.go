package main

import (
	"io/ioutil"
	"net/http"
)

type HTTP struct {
	reqPayload  string `json:request`
	respPayLoad string `json:response`
	methodName  string `json:method_name`
	URL         string `json:url`
	method      func() error
}

func (h *HTTP) Init(registry map[string]interface{}) error {
	var httpMethod func(string) (*http.Response, error)
	switch h.methodName {
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

		registry[h.respPayLoad] = body
		return nil
	}
	return nil
}

func (h *HTTP) Run() error {
	err := h.method()
	if err != nil {
		return err
	}
	return nil
}
