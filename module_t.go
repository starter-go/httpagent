package httpagent

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName    = "github.com/starter-go/httpagent"
	theModuleVersion = "v0.0.4"
	theModuleRev     = 4
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewMainModule ...
func NewMainModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRev)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	return mb
}

// NewTestModule ...
func NewTestModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRev)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}
