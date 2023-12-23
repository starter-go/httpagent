package com

import "github.com/starter-go/httpagent"

// DefaultClient ...
type DefaultClient struct {

	//starter:component
	_as func(httpagent.Client) //starter:as("#")

	Factory httpagent.Clients //starter:inject("#")

	inner httpagent.Client
}

func (inst *DefaultClient) _impl() httpagent.Client { return inst }

// Execute  ...
func (inst *DefaultClient) Execute(req *httpagent.Request) (*httpagent.Response, error) {
	client := inst.inner
	if client == nil {
		client = inst.Factory.GetClient()
		inst.inner = client
	}
	return client.Execute(req)
}
