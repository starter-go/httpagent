package com

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/starter-go/httpagent"
)

// CompleteURLFilter ...
type CompleteURLFilter struct {

	//starter:component
	_as func(httpagent.FilterRegistry) //starter:as(".")

	BaseURL string //starter:inject("${httpagent.filters.complete-url-filter.base-url}")

}

func (inst *CompleteURLFilter) _impl() (httpagent.FilterRegistry, httpagent.Filter) {
	return inst, inst
}

// FilterRegistrations ...
func (inst *CompleteURLFilter) FilterRegistrations() []*httpagent.FilterRegistration {
	r1 := &httpagent.FilterRegistration{
		Name:    "CompleteURLFilter",
		Order:   1,
		Enabled: true,
		Filter:  inst,
	}
	return []*httpagent.FilterRegistration{r1}
}

// Handle ...
func (inst *CompleteURLFilter) Handle(c *httpagent.Context, chain httpagent.FilterChain) error {

	url := c.Request.URL

	if strings.HasPrefix(url, "http:") {
		// NOP
	} else if strings.HasPrefix(url, "https:") {
		// NOP
	} else if strings.HasPrefix(url, "/") {
		c.Request.URL = inst.href(url)
	} else if strings.HasPrefix(url, "./") {
		c.Request.URL = inst.href(url)
	} else {
		return fmt.Errorf("bad HTTP URL [%s]", url)
	}

	return chain.Handle(c)
}

func (inst *CompleteURLFilter) href(u string) string {
	base, err := url.Parse(inst.BaseURL)
	if err != nil {
		return u
	}
	ref, err := url.Parse(u)
	if err != nil {
		return u
	}
	url := base.ResolveReference(ref)
	return url.String()
}
