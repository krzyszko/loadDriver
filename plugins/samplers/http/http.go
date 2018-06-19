package main

import (
	"io/ioutil"
	"net/http"

	"github.com/krzyszko/loaddriver/ess"

	"github.com/apex/log"
)

type httpSampler struct {
	ReqPayload  string `json:"request"`
	RespPayLoad string `json:"response"`
	MethodName  string `json:"method_name"`
	URL         string `json:"url"`
	method      func() error
	Cmpts       []ess.Component
}

func (h *httpSampler) Init(registry map[string]interface{}) error {
	for _, cpnt := range h.Cmpts {
		cpnt.Init(registry)
	}
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

		log.Debugf("%s\n", body)
		return nil
	}
	return nil
}

func (h *httpSampler) Run() error {
	for _, cpnt := range h.Cmpts {
		cpnt.Run()
	}
	err := h.method()
	if err != nil {
		return err
	}
	return nil
}

func Handler() ess.Component {
	return new(httpSampler)
}
