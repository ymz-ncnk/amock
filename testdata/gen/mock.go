//go:build ignore

package main

import (
	"github.com/ymz-ncnk/amock"
	testdata_amockgen "github.com/ymz-ncnk/amock/testdata/amockgen"
	"github.com/ymz-ncnk/amockgen"
	"github.com/ymz-ncnk/amockgen/text_template"
	persistor_mod "github.com/ymz-ncnk/persistor"
	"golang.org/x/tools/imports"
)

func main() {
	aMockGen, err := text_template.New()
	if err != nil {
		panic(err)
	}
	descs := []amockgen.MockImplDesc{
		testdata_amockgen.MxTypeDesc,
		testdata_amockgen.ReaderTypeDesc,
	}

	for i := 0; i < len(descs); i++ {
		err := generate(aMockGen, descs[i])
		if err != nil {
			panic(err)
		}
	}
}

func generate(aMockGen text_template.AMockGen, iDesc amockgen.MockImplDesc) (
	err error) {
	name := "a__" + iDesc.Name + amock.FilenameExtenstion
	data, err := aMockGen.Generate(iDesc)
	if err != nil {
		return
	}
	data, err = imports.Process("", data, nil)
	if err != nil {
		return
	}
	path := "testdata/amockgen"
	persistor := persistor_mod.NewHarDrivePersistor()
	return persistor.Persist(name, data, path)
}
