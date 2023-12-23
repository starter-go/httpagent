package main4httpagent

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p19f7fd0ff4_com_DefaultClient{})
    inst.register(&p19f7fd0ff4_com_DefaultClientFactory{})
    inst.register(&p19f7fd0ff4_com_HTTPStatusErrorFilter{})
    inst.register(&p19f7fd0ff4_com_LogFilter{})


    return nil
}
