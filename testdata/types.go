package testdata

import (
	"io"

	"github.com/ymz-ncnk/amockgen"
)

type IntAlias int

type Mx interface {
	M1(p0 int) (r0 float32)
	M2(p0 *[3]string, p1 []bool) (r0 []*uint, r1 [10]IntAlias)
	M3(chan error)

	M4(p0 io.Reader)
	M5(p0 io.Reader, p1 io.Writer) (r0 interface{}, r1 io.ReadCloser)
	M6(p interface{})

	M7(p0 chan int, p1 io.Writer) (r0 map[int]IntAlias, r1 error)
	M8(p0 map[IntAlias]int, p1 *io.Reader, p2 interface{}) (
		r0 *io.WriteCloser, r1 error, r2 error)
	M9(p0 *chan int, p1 io.Reader)

	M10()
}

var MxTypeDesc = amockgen.MockImplDesc{
	Package: "mock",
	Name:    "MxMock",
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
				{Name: "r1", Type: "[10]testdata.IntAlias"},
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
				{Name: "r0", Type: "map[int]testdata.IntAlias"},
				{Name: "r1", Type: "error", Interface: true},
			},
		},
		{
			Name: "M8",
			Params: []amockgen.VarDesc{
				{Name: "p0", Type: "map[testdata.IntAlias]int"},
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

// type Doer interface {
// 	Do(n io.Reader, p IntAlias, m io.Writer)
// 	DoNothing()
// }

// // Doer
// var DoerTypeDesc = amockgen.MockImplDesc{
// 	Package: "amockgen",
// 	Name:    "DoerMock",
// 	Methods: []amockgen.MethoDesc{
// 		{
// 			Name: "Do",
// 			Params: []amockgen.VarDesc{
// 				{
// 					Name:      "n",
// 					Type:      "interface{}",
// 					Interface: true,
// 				},
// 				{
// 					Name: "p",
// 					Type: "IntAlias",
// 				},
// 				{
// 					Name:      "m",
// 					Type:      "io.Writer",
// 					Interface: true,
// 				},
// 			},
// 		},
// 		{
// 			Name:         "DoNothing",
// 			Params:       []amockgen.VarDesc{},
// 			ReturnVars: []amockgen.VarDesc{},
// 		},
// 	},
// }

// io.Reader
var ReaderTypeDesc = amockgen.MockImplDesc{
	Package: "amockgen",
	Name:    "ReaderMock",
	Methods: []amockgen.MethoDesc{
		{
			Name: "Read",
			Params: []amockgen.VarDesc{
				{
					Name: "p",
					Type: "[]byte",
				},
			},
			ReturnVars: []amockgen.VarDesc{
				{
					Name: "n",
					Type: "int",
				},
				{
					Name: "err",
					Type: "error",
					// Error:     true,
					Interface: true,
				},
			},
		},
	},
}

// // io.WriterTo
// var WriterToTypeDesc = amockgen.MockImplDesc{
// 	Package: "amockgen",
// 	Name:    "WriterToMock",
// 	Methods: []amockgen.MethoDesc{
// 		{
// 			Name: "WriteTo",
// 			Params: []amockgen.VarDesc{
// 				{
// 					Name:      "w",
// 					Type:      "io.Writer",
// 					Interface: true,
// 				},
// 			},
// 			ReturnVars: []amockgen.VarDesc{
// 				{
// 					Name: "n",
// 					Type: "int64",
// 				},
// 				{
// 					Name: "err",
// 					Type: "error",
// 					// Error:     true,
// 					Interface: true,
// 				},
// 			},
// 		},
// 	},
// }

// // func some() {
// // 	io.ReadWriteCloser
// // }

// // io.ReadWriteCloser
// var ReadWriteCloserTypeDesc = amockgen.MockImplDesc{
// 	Package: "amockgen",
// 	Name:    "ReadWriteCloserMock",
// 	Methods: []amockgen.MethoDesc{
// 		{
// 			Name: "Read",
// 			Params: []amockgen.VarDesc{
// 				{
// 					Name: "p",
// 					Type: "[]byte",
// 				},
// 			},
// 			ReturnVars: []amockgen.VarDesc{
// 				{
// 					Name: "n",
// 					Type: "int",
// 				},
// 				{
// 					Name: "err",
// 					Type: "error",
// 					// Error:     true,
// 					Interface: true,
// 				},
// 			},
// 		},
// 		{
// 			Name: "Write",
// 			Params: []amockgen.VarDesc{
// 				{
// 					Name: "p",
// 					Type: "[]byte",
// 				},
// 			},
// 			ReturnVars: []amockgen.VarDesc{
// 				{
// 					Name: "n",
// 					Type: "int64",
// 				},
// 				{
// 					Name: "err",
// 					Type: "error",
// 					// Error:     true,
// 					Interface: true,
// 				},
// 			},
// 		},
// 		{
// 			Name: "Close",
// 			ReturnVars: []amockgen.VarDesc{
// 				{
// 					Name: "err",
// 					Type: "error",
// 					// Error:     true,
// 					Interface: true,
// 				},
// 			},
// 		},
// 	},
// }
