package com

import (
	"github.com/starter-go/httpagent"
	"github.com/starter-go/vlog"
)

// LogFilter ...
type LogFilter struct {

	//starter:component
	_as func(httpagent.FilterRegistry) //starter:as(".")

}

func (inst *LogFilter) _impl() (httpagent.FilterRegistry, httpagent.Filter) {
	return inst, inst
}

// FilterRegistrations ...
func (inst *LogFilter) FilterRegistrations() []*httpagent.FilterRegistration {
	r1 := &httpagent.FilterRegistration{
		Name:    "logFilter",
		Order:   11,
		Enabled: true,
		Filter:  inst,
	}
	return []*httpagent.FilterRegistration{r1}
}

// Handle ...
func (inst *LogFilter) Handle(c *httpagent.Context, chain httpagent.FilterChain) error {

	inst.logRequest(c.Request)
	defer func() {
		inst.logResponse(c.Response)
	}()

	return chain.Handle(c)
}

func (inst *LogFilter) logRequest(req *httpagent.Request) {
	if req == nil {
		return
	}
	method := req.Method
	url := req.URL
	vlog.Info("HTTP %s %s", method, url)
}

func (inst *LogFilter) logResponse(resp *httpagent.Response) {
	if resp == nil {
		return
	}
	msg := resp.Message
	vlog.Info("    ............................................. %s", msg)
}
