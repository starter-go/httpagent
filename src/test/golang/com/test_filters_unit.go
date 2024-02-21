package com

import (
	"net/http"

	"github.com/starter-go/httpagent"
	"github.com/starter-go/units"
)

// TestFilters ...
type TestFilters struct {

	//starter:component
	_as func(units.Units) //starter:as(".")

	Client  httpagent.Client  //starter:inject("#")
	Clients httpagent.Clients //starter:inject("#")

}

func (inst *TestFilters) _impl() units.Units { return inst }

// Units ...
func (inst *TestFilters) Units(list []*units.Registration) []*units.Registration {
	r1 := &units.Registration{
		Name:     "testFilters",
		Enabled:  true,
		Priority: 0,
		Test:     inst.testFilters,
	}
	return []*units.Registration{r1}
}

func (inst *TestFilters) testFilters() error {

	req := new(httpagent.Request)
	req.URL = "https://gitee.com/starter-go"
	req.Method = http.MethodGet

	resp, err := inst.Client.Execute(req)

	if err == nil {
		resp.GetEntity()
	}

	return err
}

////////////////////////////////////////////////////////////////////////////////
