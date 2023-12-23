package test4httpagent
import (
    pdea5a0f47 "github.com/starter-go/httpagent"
    p0fa9b76ea "github.com/starter-go/httpagent/src/test/golang/com"
     "github.com/starter-go/application"
)

// type p0fa9b76ea.TestFilters in package:github.com/starter-go/httpagent/src/test/golang/com
//
// id:com-0fa9b76eacbd9771-com-TestFilters
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p0fa9b76eac_com_TestFilters struct {
}

func (inst* p0fa9b76eac_com_TestFilters) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0fa9b76eacbd9771-com-TestFilters"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0fa9b76eac_com_TestFilters) new() any {
    return &p0fa9b76ea.TestFilters{}
}

func (inst* p0fa9b76eac_com_TestFilters) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0fa9b76ea.TestFilters)
	nop(ie, com)

	
    com.Client = inst.getClient(ie)


    return nil
}


func (inst*p0fa9b76eac_com_TestFilters) getClient(ie application.InjectionExt)pdea5a0f47.Client{
    return ie.GetComponent("#alias-dea5a0f47697e78c03558bf00bc7ff9c-Client").(pdea5a0f47.Client)
}


