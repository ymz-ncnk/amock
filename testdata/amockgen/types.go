package amockgen

import (
	"io"
	"math/big"

	"github.com/ymz-ncnk/amockgen"
)

type Mx interface {
	M1(p0 int) (r0 float32)
	M2(p0 *[3]string, p1 []bool) (r0 []*uint, r1 [10]big.Int)
	M3(chan error)

	M4(p0 io.Reader)
	M5(p0 io.Reader, p1 io.Writer) (r0 interface{}, r1 io.ReadCloser)
	M6(p interface{})

	M7(p0 chan int, p1 io.Writer) (r0 map[int]big.Int, r1 error)
	M8(p0 map[string]int, p1 *io.Reader, p2 interface{}) (
		r0 *io.WriteCloser, r1 error, r2 error)
	M9(p0 *chan int, p1 io.Reader)

	M10()
}

var MxTypeDesc = amockgen.MockImplDesc{
	InterfaceType: "amockgen.Mx",
	Package:       "amockgen",
	Name:          "MxMock",
	Methods: []amockgen.MethoDesc{
		{
			Name: "M1",
			Params: []amockgen.VarDesc{
				{Name: "p0", Type: "int"},
			},
			ReturnVars: []amockgen.VarDesc{
				{Name: "r0", Type: "float32"},
			},
		},
		{
			Name:       "M10",
			Params:     []amockgen.VarDesc{},
			ReturnVars: []amockgen.VarDesc{},
		},
		{
			Name: "M2",
			Params: []amockgen.VarDesc{
				{Name: "p0", Type: "*[3]string"},
				{Name: "p1", Type: "[]bool"},
			},
			ReturnVars: []amockgen.VarDesc{
				{Name: "r0", Type: "[]*uint"},
				{Name: "r1", Type: "[10]big.Int"},
			},
		},
		{
			Name: "M3",
			Params: []amockgen.VarDesc{
				{Name: "p0", Type: "chan error"},
			},
			ReturnVars: []amockgen.VarDesc{},
		},
		{
			Name: "M4",
			Params: []amockgen.VarDesc{
				{Name: "p0", Type: "io.Reader", Interface: true},
			},
			ReturnVars: []amockgen.VarDesc{},
		},
		{
			Name: "M5",
			Params: []amockgen.VarDesc{
				{Name: "p0", Type: "io.Reader", Interface: true},
				{Name: "p1", Type: "io.Writer", Interface: true},
			},
			ReturnVars: []amockgen.VarDesc{
				{Name: "r0", Type: "interface {}", Interface: true},
				{Name: "r1", Type: "io.ReadCloser", Interface: true},
			},
		},
		{
			Name: "M6",
			Params: []amockgen.VarDesc{
				{Name: "p0", Type: "interface {}", Interface: true},
			},
			ReturnVars: []amockgen.VarDesc{},
		},
		{
			Name: "M7",
			Params: []amockgen.VarDesc{
				{Name: "p0", Type: "chan int"},
				{Name: "p1", Type: "io.Writer", Interface: true},
			},
			ReturnVars: []amockgen.VarDesc{
				{Name: "r0", Type: "map[int]big.Int"},
				{Name: "r1", Type: "error", Interface: true},
			},
		},
		{
			Name: "M8",
			Params: []amockgen.VarDesc{
				{Name: "p0", Type: "map[string]int"},
				{Name: "p1", Type: "*io.Reader"},
				{Name: "p2", Type: "interface {}", Interface: true},
			},
			ReturnVars: []amockgen.VarDesc{
				{Name: "r0", Type: "*io.WriteCloser"},
				{Name: "r1", Type: "error", Interface: true},
				{Name: "r2", Type: "error", Interface: true},
			},
		},
		{
			Name: "M9",
			Params: []amockgen.VarDesc{
				{Name: "p0", Type: "*chan int"},
				{Name: "p1", Type: "io.Reader", Interface: true},
			},
			ReturnVars: []amockgen.VarDesc{},
		},
	},
}

// io.Reader
var ReaderTypeDesc = amockgen.MockImplDesc{
	InterfaceType: "io.Reader",
	Package:       "amockgen",
	Name:          "ReaderMock",
	Methods: []amockgen.MethoDesc{
		{
			Name: "Read",
			Params: []amockgen.VarDesc{
				{
					Name: "p0",
					Type: "[]uint8",
				},
			},
			ReturnVars: []amockgen.VarDesc{
				{
					Name: "r0",
					Type: "int",
				},
				{
					Name:      "r1",
					Type:      "error",
					Interface: true,
				},
			},
		},
	},
}
