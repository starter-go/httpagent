package httpagent

import (
	"github.com/starter-go/application"
	"github.com/starter-go/httpagent"
	"github.com/starter-go/units/modules/units"

	"github.com/starter-go/httpagent/gen/main4httpagent"
	"github.com/starter-go/httpagent/gen/test4httpagent"
)

// Module ...
func Module() application.Module {
	return ModuleForMain()
}

func ModuleForMain() application.Module {
	mb := httpagent.NewMainModule()
	mb.Components(main4httpagent.ExportComponents)
	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := httpagent.NewTestModule()

	mb.Components(test4httpagent.ExportComponents)

	mb.Depend(Module())
	mb.Depend(units.Module())

	return mb.Create()
}
