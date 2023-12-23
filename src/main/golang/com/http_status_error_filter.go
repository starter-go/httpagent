package com

import (
	"fmt"

	"github.com/starter-go/httpagent"
)

// HTTPStatusErrorFilter ...
// 这个过滤器检查 HTTP 响应状态码，如果不成功就返回错误
type HTTPStatusErrorFilter struct {

	//starter:component
	_as func(httpagent.FilterRegistry) //starter:as(".")

}

func (inst *HTTPStatusErrorFilter) _impl() (httpagent.FilterRegistry, httpagent.Filter) {
	return inst, inst
}

// FilterRegistrations ...
func (inst *HTTPStatusErrorFilter) FilterRegistrations() []*httpagent.FilterRegistration {
	r1 := &httpagent.FilterRegistration{
		Name:    "httpStatusErrorFilter",
		Order:   12,
		Enabled: true,
		Filter:  inst,
	}
	return []*httpagent.FilterRegistration{r1}
}

// Handle ...
func (inst *HTTPStatusErrorFilter) Handle(c *httpagent.Context, chain httpagent.FilterChain) error {
	err := chain.Handle(c)
	if err != nil {
		return err
	}
	resp := c.Response
	if resp == nil {
		return fmt.Errorf("no Response")
	}
	return inst.check(resp)
}

func (inst *HTTPStatusErrorFilter) check(resp *httpagent.Response) error {
	code := resp.Status
	if 200 <= code && code < 300 {
		return nil
	}
	return fmt.Errorf("HTTP %s", resp.Message)
}
