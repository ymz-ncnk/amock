package amock

import (
	"reflect"

	"github.com/ymz-ncnk/amock/parser"
	"github.com/ymz-ncnk/amockgen"
	"github.com/ymz-ncnk/amockgen/text_template"
	persistor_mod "github.com/ymz-ncnk/persistor"
	"golang.org/x/tools/imports"
)

// FilenameExtenstion of the generated files.
const FilenameExtenstion = ".gen.go"

// DefConf is the default configuration for AMock.
var DefConf = Conf{Path: "testdata/mock", Package: "mock"}

// New creates a new AMock.
func New() (aMock AMock, err error) {
	aMockGen, err := text_template.New()
	if err != nil {
		return
	}
	aMock = NewWith(aMockGen, persistor_mod.NewHarDrivePersistor())
	return
}

// NewWith creates a new configurable AMock.
func NewWith(aMockGen amockgen.AMockGen,
	persistor persistor_mod.Persistor) AMock {
	return AMock{
		aMockGen:  aMockGen,
		persistor: persistor,
	}
}

// AMock is a mock implementations generator.
type AMock struct {
	aMockGen  amockgen.AMockGen
	persistor persistor_mod.Persistor
}

// Generate generates mock implementation of the interface. If tp is not
// an interface returns parser.ErrNotInterface.
// Filename and mock implementation type will be equal to tp.Name().
// Uses DefConf.
func (aMock AMock) Generate(tp reflect.Type) error {
	return aMock.GenerateAs(tp, DefConf)
}

// GenerateAs performs like Generate. With help of this method you can configure
// the generation process.
func (aMock AMock) GenerateAs(tp reflect.Type, conf Conf) (err error) {
	iDesc, err := parser.Parse(tp)
	if err != nil {
		return
	}
	if len(conf.Package) > 0 {
		iDesc.Package = conf.Package
	}
	if len(conf.Name) > 0 {
		iDesc.Name = conf.Name
	}
	name := iDesc.Name + FilenameExtenstion
	data, err := aMock.aMockGen.Generate(iDesc)
	if err != nil {
		return
	}
	data, err = imports.Process("", data, nil)
	if err != nil {
		return
	}
	return aMock.persistor.Persist(name, data, conf.Path)
}
