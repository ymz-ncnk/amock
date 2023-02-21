package parser

import (
	"reflect"
	"testing"

	testdata_amockgen "github.com/ymz-ncnk/amock/testdata/amockgen"
)

func TestParser(t *testing.T) {
	want := testdata_amockgen.MxTypeDesc
	iDesc, err := Parse(reflect.TypeOf((*testdata_amockgen.Mx)(nil)).Elem())
	if err != nil {
		t.Error(err)
	}
	iDesc.Name = iDesc.Name + "Mock"
	if !reflect.DeepEqual(iDesc, want) {
		t.Errorf("unexpected iDesc, want '%v', actual '%v'", want, iDesc)
	}
}
