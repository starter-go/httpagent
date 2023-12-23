package golang

import (
	"testing"

	"github.com/starter-go/httpagent/modules/httpagent"
	"github.com/starter-go/units"
)

func TestFilters(t *testing.T) {

	args := []string{
		"--debug.enabled=true",
		"--debug.log-properties=true",
		"--test.units.run-all=1",
		"--test.units.run-list=testFilters",
	}

	m := httpagent.ModuleForTest()
	r := units.NewRunner()
	r.Testing(t)
	r.EnablePanic(true)
	r.Dependencies(m)
	r.Run(args)
}
