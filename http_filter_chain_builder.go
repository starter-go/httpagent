package httpagent

import (
	"fmt"
	"sort"
)

////////////////////////////////////////////////////////////////////////////////

// FilterChainBuilder 用于创建 FilterChain
type FilterChainBuilder struct {
	items []*FilterRegistration
}

// AddRegistration ...
func (inst *FilterChainBuilder) AddRegistration(list ...*FilterRegistration) *FilterChainBuilder {
	inst.items = append(inst.items, list...)
	return inst
}

// AddRegistry ...
func (inst *FilterChainBuilder) AddRegistry(list ...FilterRegistry) *FilterChainBuilder {
	for _, r2 := range list {
		list2 := r2.FilterRegistrations()
		inst.items = append(inst.items, list2...)
	}
	return inst
}

// Create ...
func (inst *FilterChainBuilder) Create() FilterChain {

	all := inst.items
	use := make([]*FilterRegistration, 0)

	for _, r := range all {
		if r == nil {
			continue
		}
		if r.Enabled && r.Filter != nil {
			use = append(use, r)
		}
	}

	sorter := new(filterRegistrationSorter)
	sorter.sort(use)
	var chain FilterChain
	chain = new(filterChainEnding)

	for _, r := range use {
		node := &filterChainNode{
			next:   chain,
			filter: r.Filter,
		}
		chain = node
	}

	return chain
}

////////////////////////////////////////////////////////////////////////////////

type filterChainNode struct {
	next   FilterChain
	filter Filter
}

func (inst *filterChainNode) _impl() FilterChain { return inst }

func (inst *filterChainNode) Execute(req *Request) (*Response, error) {
	return inst.filter.Handle(req, inst.next)
}

////////////////////////////////////////////////////////////////////////////////

type filterChainEnding struct{}

func (inst *filterChainEnding) _impl() FilterChain { return inst }

func (inst *filterChainEnding) Execute(req *Request) (*Response, error) {
	return nil, fmt.Errorf("httpagent.FilterChain: 没有能处理该请求的过滤器")
}

////////////////////////////////////////////////////////////////////////////////

type filterRegistrationSorter struct {
	items []*FilterRegistration
}

func (inst *filterRegistrationSorter) sort(list []*FilterRegistration) {
	inst.items = list
	sort.Sort(inst)
}

func (inst *filterRegistrationSorter) Len() int {
	return len(inst.items)
}

func (inst *filterRegistrationSorter) Less(i1, i2 int) bool {
	o1 := inst.items[i1].Order
	o2 := inst.items[i2].Order
	return o1 > o2
}

func (inst *filterRegistrationSorter) Swap(i1, i2 int) {
	list := inst.items
	list[i1], list[i2] = list[i2], list[i1]
}

////////////////////////////////////////////////////////////////////////////////
