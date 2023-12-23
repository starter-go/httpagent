package com

import "github.com/starter-go/httpagent"

// DefaultClientFactory ...
type DefaultClientFactory struct {

	//starter:component
	_as func(httpagent.Clients) //starter:as("#")

	FiltersRegs []httpagent.FilterRegistry //starter:inject(".")

	client httpagent.Client
}

func (inst *DefaultClientFactory) _impl() httpagent.Clients { return inst }

// GetClient ...
func (inst *DefaultClientFactory) GetClient() httpagent.Client {
	c := inst.client
	if c == nil {
		c = inst.NewClient()
		inst.client = c
	}
	return c
}

// NewClient ...
func (inst *DefaultClientFactory) NewClient() httpagent.Client {
	src := inst.FiltersRegs
	dst := make([]*httpagent.FilterRegistration, 0)
	for _, r1 := range src {
		tmp := r1.FilterRegistrations()
		dst = append(dst, tmp...)
	}
	return httpagent.NewClientWithFilters(dst...)
}
