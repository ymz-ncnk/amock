package parser

import (
	"reflect"
	"regexp"
	"strconv"

	"github.com/ymz-ncnk/amockgen"
)

// These constants are used to name params and return variables in generated
// methods.
const (
	ParamName     = "p"
	ReturnVarName = "r"
)

// Parse creates amockgen.MockImplDesc from the interface type. If tp is not an
// interface returns ErrNotInterface.
func Parse(tp reflect.Type) (iDesc amockgen.MockImplDesc, err error) {
	if tp.Kind() != reflect.Interface {
		return amockgen.MockImplDesc{}, ErrNotInterface
	}
	iDesc = amockgen.MockImplDesc{
		InterfaceType: tp.String(),
		Name:          tp.Name(),
		Package:       pkg(tp),
		Methods:       []amockgen.MethoDesc{},
	}
	for i := 0; i < tp.NumMethod(); i++ {
		iDesc.Methods = append(iDesc.Methods, parseMethod(tp.Method(i)))
	}
	return
}

func parseMethod(method reflect.Method) (mDesc amockgen.MethoDesc) {
	mDesc = amockgen.MethoDesc{
		Name:       method.Name,
		Params:     []amockgen.VarDesc{},
		ReturnVars: []amockgen.VarDesc{},
	}
	for i := 0; i < method.Type.NumIn(); i++ {
		mDesc.Params = append(mDesc.Params, parseParam(i, method.Type.In(i)))
	}
	for i := 0; i < method.Type.NumOut(); i++ {
		mDesc.ReturnVars = append(mDesc.ReturnVars,
			parseReturnValue(i, method.Type.Out(i)))
	}
	return mDesc
}

func parseParam(index int, tp reflect.Type) (vDesc amockgen.VarDesc) {
	return amockgen.VarDesc{
		Name:      ParamName + strconv.Itoa(index),
		Type:      tp.String(),
		Interface: tp.Kind() == reflect.Interface,
	}
}

func parseReturnValue(index int, tp reflect.Type) (vDesc amockgen.VarDesc) {
	return amockgen.VarDesc{
		Name:      ReturnVarName + strconv.Itoa(index),
		Type:      tp.String(),
		Interface: tp.Kind() == reflect.Interface,
	}
}

func pkg(tp reflect.Type) string {
	re := regexp.MustCompile(`^(.*)\.`)
	match := re.FindStringSubmatch(tp.String())
	if len(match) != 2 {
		return ""
	}
	return match[1]
}
