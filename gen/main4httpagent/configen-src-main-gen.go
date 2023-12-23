package main4httpagent
import (
    pdea5a0f47 "github.com/starter-go/httpagent"
    p19f7fd0ff "github.com/starter-go/httpagent/src/main/golang/com"
     "github.com/starter-go/application"
)

// type p19f7fd0ff.DefaultClient in package:github.com/starter-go/httpagent/src/main/golang/com
//
// id:com-19f7fd0ff45d1301-com-DefaultClient
// class:
// alias:alias-dea5a0f47697e78c03558bf00bc7ff9c-Client
// scope:singleton
//
type p19f7fd0ff4_com_DefaultClient struct {
}

func (inst* p19f7fd0ff4_com_DefaultClient) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-19f7fd0ff45d1301-com-DefaultClient"
	r.Classes = ""
	r.Aliases = "alias-dea5a0f47697e78c03558bf00bc7ff9c-Client"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p19f7fd0ff4_com_DefaultClient) new() any {
    return &p19f7fd0ff.DefaultClient{}
}

func (inst* p19f7fd0ff4_com_DefaultClient) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p19f7fd0ff.DefaultClient)
	nop(ie, com)

	
    com.Factory = inst.getFactory(ie)


    return nil
}


func (inst*p19f7fd0ff4_com_DefaultClient) getFactory(ie application.InjectionExt)pdea5a0f47.Clients{
    return ie.GetComponent("#alias-dea5a0f47697e78c03558bf00bc7ff9c-Clients").(pdea5a0f47.Clients)
}



// type p19f7fd0ff.DefaultClientFactory in package:github.com/starter-go/httpagent/src/main/golang/com
//
// id:com-19f7fd0ff45d1301-com-DefaultClientFactory
// class:
// alias:alias-dea5a0f47697e78c03558bf00bc7ff9c-Clients
// scope:singleton
//
type p19f7fd0ff4_com_DefaultClientFactory struct {
}

func (inst* p19f7fd0ff4_com_DefaultClientFactory) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-19f7fd0ff45d1301-com-DefaultClientFactory"
	r.Classes = ""
	r.Aliases = "alias-dea5a0f47697e78c03558bf00bc7ff9c-Clients"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p19f7fd0ff4_com_DefaultClientFactory) new() any {
    return &p19f7fd0ff.DefaultClientFactory{}
}

func (inst* p19f7fd0ff4_com_DefaultClientFactory) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p19f7fd0ff.DefaultClientFactory)
	nop(ie, com)

	
    com.FiltersRegs = inst.getFiltersRegs(ie)


    return nil
}


func (inst*p19f7fd0ff4_com_DefaultClientFactory) getFiltersRegs(ie application.InjectionExt)[]pdea5a0f47.FilterRegistry{
    dst := make([]pdea5a0f47.FilterRegistry, 0)
    src := ie.ListComponents(".class-dea5a0f47697e78c03558bf00bc7ff9c-FilterRegistry")
    for _, item1 := range src {
        item2 := item1.(pdea5a0f47.FilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p19f7fd0ff.HTTPStatusErrorFilter in package:github.com/starter-go/httpagent/src/main/golang/com
//
// id:com-19f7fd0ff45d1301-com-HTTPStatusErrorFilter
// class:class-dea5a0f47697e78c03558bf00bc7ff9c-FilterRegistry
// alias:
// scope:singleton
//
type p19f7fd0ff4_com_HTTPStatusErrorFilter struct {
}

func (inst* p19f7fd0ff4_com_HTTPStatusErrorFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-19f7fd0ff45d1301-com-HTTPStatusErrorFilter"
	r.Classes = "class-dea5a0f47697e78c03558bf00bc7ff9c-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p19f7fd0ff4_com_HTTPStatusErrorFilter) new() any {
    return &p19f7fd0ff.HTTPStatusErrorFilter{}
}

func (inst* p19f7fd0ff4_com_HTTPStatusErrorFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p19f7fd0ff.HTTPStatusErrorFilter)
	nop(ie, com)

	


    return nil
}



// type p19f7fd0ff.LogFilter in package:github.com/starter-go/httpagent/src/main/golang/com
//
// id:com-19f7fd0ff45d1301-com-LogFilter
// class:class-dea5a0f47697e78c03558bf00bc7ff9c-FilterRegistry
// alias:
// scope:singleton
//
type p19f7fd0ff4_com_LogFilter struct {
}

func (inst* p19f7fd0ff4_com_LogFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-19f7fd0ff45d1301-com-LogFilter"
	r.Classes = "class-dea5a0f47697e78c03558bf00bc7ff9c-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p19f7fd0ff4_com_LogFilter) new() any {
    return &p19f7fd0ff.LogFilter{}
}

func (inst* p19f7fd0ff4_com_LogFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p19f7fd0ff.LogFilter)
	nop(ie, com)

	


    return nil
}


