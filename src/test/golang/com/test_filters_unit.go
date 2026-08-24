package com

import (
	"context"
	"net/http"

	"github.com/starter-go/httpagent"
	"github.com/starter-go/units"
)

// TestFilters ...
type TestFilters struct {

	//starter:component
	_as func(units.Unit) //starter:as(".")

	Client  httpagent.Client  //starter:inject("#")
	Clients httpagent.Clients //starter:inject("#")

}

// ListRegistrations implements [units.Unit].
func (inst *TestFilters) ListRegistrations(list []*units.Registration) []*units.Registration {

	r1 := &units.Registration{
		Name:     "testFilters",
		Enabled:  true,
		Priority: 0,
		Do:       inst.testFilters,
	}

	return append(list, r1)
}

func (inst *TestFilters) testFilters(c context.Context) error {

	req := new(httpagent.Request)
	req.URL = "https://gitee.com/starter-go"
	req.Method = http.MethodGet

	resp, err := inst.Client.Execute(req)

	if err == nil {
		resp.GetEntity()
	}

	return err
}

func (inst *TestFilters) _impl() units.Unit {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
