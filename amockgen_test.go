package amock

import (
	"errors"
	"io"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/ymz-ncnk/amock/core"
	testdata_amockgen "github.com/ymz-ncnk/amock/testdata/amockgen"
)

func TestMxMock(t *testing.T) {

	t.Run("M1 method", func(t *testing.T) {
		var (
			wantP0         = 4
			wantR  float32 = 10.4
		)
		mx := testdata_amockgen.NewMxMock()
		mx.RegisterM1(func(p0 int) (r float32) {
			if p0 != wantP0 {
				t.Errorf("unexpected p, want '%v', actual '%v'", wantP0, p0)
			}
			return wantR
		})
		r := mx.M1(wantP0)
		if r != wantR {
			t.Errorf("unexpected p, want '%v', actual '%v'", wantR, r)
		}
	})

	t.Run("M2 method", func(t *testing.T) {
		var (
			arr    = [3]string{"a", "b", "c"}
			wantP0 = &arr
			wantP1 = []bool{true, false, true}
			wantR0 = []*uint{}
			wantR1 = [10]big.Int{}
		)

		mx := testdata_amockgen.NewMxMock()
		mx.RegisterM2(func(p0 *[3]string, p1 []bool) (r0 []*uint,
			r1 [10]big.Int) {
			if p0 != wantP0 {
				t.Errorf("unexpected p0, want '%v', actual '%v'", wantP0, p0)
			}
			if !reflect.DeepEqual(p1, wantP1) {
				t.Errorf("unexpected p1, want '%v', actual '%v'", wantP1, p1)
			}
			return wantR0, wantR1
		}).RegisterM2(func(p0 *[3]string, p1 []bool) (r0 []*uint, r1 [10]big.Int) {
			return
		})

		r0, r1 := mx.M2(wantP0, wantP1)
		if !reflect.DeepEqual(r0, wantR0) {
			t.Errorf("unexpected r0, want '%v', actual '%v'", wantR0, r0)
		}
		if !reflect.DeepEqual(r1, wantR1) {
			t.Errorf("unexpected r1, want '%v', actual '%v'", wantR1, r1)
		}

		mx.M2(nil, nil)
	})

	t.Run("M3 method", func(t *testing.T) {
		var (
			wantP0 = make(chan error)
		)
		mx := testdata_amockgen.NewMxMock()
		mx.RegisterM3(func(p0 chan error) {
			if p0 != wantP0 {
				t.Errorf("unexpected p, want '%v', actual '%v'", wantP0, p0)
			}
		})
		mx.RegisterM3(func(p0 chan error) {})
		mx.M3(wantP0)
		mx.M3(nil)
	})

	t.Run("M4 method", func(t *testing.T) {
		var (
			wantP0 io.Reader = &os.File{}
		)
		mx := testdata_amockgen.NewMxMock()
		mx.RegisterM4(func(p0 io.Reader) {
			if p0 != wantP0 {
				t.Errorf("unexpected p, want '%v', actual '%v'", wantP0, p0)
			}
		}).RegisterM4(func(p0 io.Reader) {})

		mx.M4(wantP0)
		mx.M4(nil)
	})

	t.Run("M5 method", func(t *testing.T) {
		var (
			wantP0             = &os.File{}
			wantP1             = &os.File{}
			wantR0 interface{} = "string"
			wantR1             = &os.File{}
		)
		mx := testdata_amockgen.NewMxMock()
		mx.RegisterM5(func(p0 io.Reader, p1 io.Writer) (r0 interface{},
			r1 io.ReadCloser) {
			if p0 != wantP0 {
				t.Errorf("unexpected p0, want '%v', actual '%v'", wantP0, p0)
			}
			if p1 != wantP1 {
				t.Errorf("unexpected p1, want '%v', actual '%v'", wantP1, p1)
			}
			return wantR0, wantR1
		}).RegisterM5(func(p0 io.Reader, p1 io.Writer) (r0 interface{},
			r1 io.ReadCloser) {
			return
		})

		r0, r1 := mx.M5(wantP0, wantP1)
		if r0 != wantR0 {
			t.Errorf("unexpected r0, want '%v', actual '%v'", wantR0, r0)
		}
		if r1 != wantR1 {
			t.Errorf("unexpected r1, want '%v', actual '%v'", wantR1, r1)
		}
		mx.M5(nil, nil)
	})

	t.Run("M6 method", func(t *testing.T) {
		var (
			wantP0 interface{} = 5
		)
		mx := testdata_amockgen.NewMxMock()
		mx.RegisterM6(func(p0 interface{}) {
			if p0 != wantP0 {
				t.Errorf("unexpected p, want '%v', actual '%v'", wantP0, p0)
			}
		}).RegisterM6(func(p0 interface{}) {})
		mx.M6(wantP0)
		mx.M6(nil)
	})

	t.Run("M7 method", func(t *testing.T) {
		var (
			wantP0 = make(chan int)
			wantP1 = &os.File{}
			wantR0 = map[int]big.Int{10: {}}
			wantR1 = errors.New("fail")
		)
		mx := testdata_amockgen.NewMxMock()
		mx.RegisterM7(func(p0 chan int, p1 io.Writer) (
			r0 map[int]big.Int, r1 error) {
			if p0 != wantP0 {
				t.Errorf("unexpected p0, want '%v', actual '%v'", wantP0, p0)
			}
			if p1 != wantP1 {
				t.Errorf("unexpected p1, want '%v', actual '%v'", wantP1, p1)
			}
			return wantR0, wantR1
		}).RegisterM7(func(p0 chan int, p1 io.Writer) (
			r0 map[int]big.Int, r1 error) {
			return
		})

		r0, r1 := mx.M7(wantP0, wantP1)
		if !reflect.DeepEqual(r0, wantR0) {
			t.Errorf("unexpected r0, want '%v', actual '%v'", wantR0, r0)
		}
		if r1 != wantR1 {
			t.Errorf("unexpected r1, want '%v', actual '%v'", wantR1, r1)
		}
		mx.M7(nil, nil)
	})

	t.Run("M8 method", func(t *testing.T) {
		var (
			wantP0                = map[string]int{"str": 5}
			wantP1 io.Reader      = &os.File{}
			wantP3 interface{}    = 88
			wantR0 io.WriteCloser = &os.File{}
			wantR1                = errors.New("fail1")
			wantR3                = errors.New("fail2")
		)
		mx := testdata_amockgen.NewMxMock()
		mx.RegisterM8(func(p0 map[string]int,
			p1 *io.Reader, p3 interface{}) (r0 *io.WriteCloser, r1 error, r3 error) {
			if !reflect.DeepEqual(p0, wantP0) {
				t.Errorf("unexpected p0, want '%v', actual '%v'", wantP0, p0)
			}
			if p1 != &wantP1 {
				t.Errorf("unexpected p1, want '%v', actual '%v'", wantP1, p1)
			}
			if p3 != wantP3 {
				t.Errorf("unexpected p3, want '%v', actual '%v'", wantP3, p3)
			}
			return &wantR0, wantR1, wantR3
		}).RegisterM8(func(p0 map[string]int,
			p1 *io.Reader, p3 interface{}) (r0 *io.WriteCloser, r1 error, r3 error) {
			return
		})

		r0, r1, r3 := mx.M8(wantP0, &wantP1, wantP3)
		if !reflect.DeepEqual(r0, &wantR0) {
			t.Errorf("unexpected r0, want '%v', actual '%v'", wantR0, r0)
		}
		if r1 != wantR1 {
			t.Errorf("unexpected r1, want '%v', actual '%v'", wantR1, r1)
		}
		if r3 != wantR3 {
			t.Errorf("unexpected r3, want '%v', actual '%v'", wantR3, r3)
		}
		mx.M8(nil, nil, nil)
	})

	t.Run("M9 method", func(t *testing.T) {
		var (
			wantP0 = make(chan int)
			wantP1 = &os.File{}
		)
		mx := testdata_amockgen.NewMxMock()
		mx.RegisterM9(func(p0 *chan int, p1 io.Reader) {
			if p0 != &wantP0 {
				t.Errorf("unexpected p0, want '%v', actual '%v'", wantP0, p0)
			}
			if p1 != wantP1 {
				t.Errorf("unexpected p1, want '%v', actual '%v'", wantP1, p1)
			}
		}).RegisterM9(func(p0 *chan int, p1 io.Reader) {})

		mx.M9(&wantP0, wantP1)
		mx.M9(nil, nil)
	})

	t.Run("M10 method", func(t *testing.T) {
		mx := testdata_amockgen.NewMxMock()
		mx.RegisterM10(func() {})
		mx.M10()
	})

}

func TestUnknownMethodCall(t *testing.T) {
	want := core.NewUnknownMethodCallError("ReaderMock", "Read")
	reader := testdata_amockgen.NewReaderMock()
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				if err.Error() != want.Error() {
					t.Errorf("unexpected error, want '%v', actual '%v'", want, err)
				}
			}
		}
	}()
	reader.Read(nil)
}

func TestUnexpectedMethodCall(t *testing.T) {
	want := core.NewUnexpectedMethodCallError("ReaderMock", "Read")
	reader := testdata_amockgen.NewReaderMock()
	reader.RegisterRead(func(p0 []byte) (n int, err error) {
		return 0, nil
	})
	reader.Read(nil)
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				if err.Error() != want.Error() {
					t.Errorf("unexpected error, want '%v', actual '%v'", want, err)
				}
			}
		}
	}()
	reader.Read(nil)
}
