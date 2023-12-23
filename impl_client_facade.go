package httpagent

import (
	"fmt"

	"github.com/starter-go/application/attributes"
)

type myClientFacade struct {
	clientContext *ClientContext
}

func (inst *myClientFacade) _impl() Client { return inst }

func (inst *myClientFacade) Execute(req *Request) (*Response, error) {

	if req == nil {
		return nil, fmt.Errorf("no Request")
	}

	c := new(Context)
	c.Attributes = attributes.NewTable(nil)
	c.ClientContext = nil
	c.Request = req

	err := inst.clientContext.Chain.Handle(c)
	if err != nil {
		return nil, err
	}
	if c.Response == nil {
		return nil, fmt.Errorf("no Response")
	}
	return c.Response, nil
}
