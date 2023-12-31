package httpagent

import (
	"fmt"
	"testing"
)

func TestFilterChain(t *testing.T) {

	b := new(FilterChainBuilder)

	b.AddRegistration(&FilterRegistration{
		Order:   8,
		Enabled: true,
		Name:    "f8",
		Filter:  &testingFilter{order: 8},
	})

	b.AddRegistration(&FilterRegistration{
		Order:   6,
		Enabled: true,
		Name:    "f6",
		Filter:  &testingFilter{order: 6},
	})

	b.AddRegistration(&FilterRegistration{
		Order:   5,
		Enabled: true,
		Name:    "f5",
		Filter:  &testingFilter{order: 5},
	})

	b.AddRegistration(&FilterRegistration{
		Order:   7,
		Enabled: true,
		Name:    "f7",
		Filter:  &testingFilter{order: 7},
	})

	req := new(Context)
	chain := b.Create()
	chain.Handle(req)
}

////////////////////////////////////////////////////////////////////////////////

type testingFilter struct {
	order int
}

func (inst *testingFilter) Handle(c *Context, chain FilterChain) error {

	fmt.Println("", inst.order)

	return chain.Handle(c)
}
