package main

import (
	"os"

	"github.com/starter-go/httpagent/modules/httpagent"
	"github.com/starter-go/units"
)

func main() {

	a := os.Args
	m := httpagent.ModuleForTest()
	c := new(units.Context)

	c.UsePanic = true
	c.Arguments = a
	c.Module = m

	units.Run(c)
}
