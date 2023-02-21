package amock

import (
	"errors"
	"io"
	"math/big"
	"reflect"
	"testing"

	"github.com/ymz-ncnk/amock/core"
	"github.com/ymz-ncnk/amock/parser"
	testdata_amockgen "github.com/ymz-ncnk/amock/testdata/amockgen"
	"github.com/ymz-ncnk/amock/testdata/mock"
	"github.com/ymz-ncnk/amockgen"
)

func TestAmock(t *testing.T) {

	prepareForGenerate := func(wantConf Conf, t *testing.T) (
		aMockGen mock.AMockGen, persistor mock.Persistor, wantErr error) {
		var (
			wantData  = []byte("package " + wantConf.Package + "\n")
			wantIDesc = func() amockgen.MockImplDesc {
				d := testdata_amockgen.ReaderTypeDesc
				d.Package = wantConf.Package
				d.Name = wantConf.Name
				return d
			}()
		)
		wantErr = errors.New("fail")
		aMockGen = mock.NewAMockGen()
		aMockGen.RegisterGenerate(
			func(iDesc amockgen.MockImplDesc) ([]byte, error) {
				if !reflect.DeepEqual(iDesc, wantIDesc) {
					t.Errorf("unexpected iDesc, want '%v' catual '%v'", iDesc, wantIDesc)
				}
				return wantData, nil
			},
		)
		persistor = mock.NewPersistor()
		persistor.RegisterPersist(func(name string, data []byte,
			path string) error {
			if !reflect.DeepEqual(data, wantData) {
				t.Errorf("unexpected data, want '%v' catual '%v'", data, wantData)
			}
			if path != wantConf.Path {
				t.Errorf("unexpected path, want '%v' catual '%v'", path, wantConf.Path)
			}
			return wantErr
		})
		return
	}

	t.Run("Generate", func(t *testing.T) {
		wantConf := Conf{
			Path:    "testdata/mock",
			Package: "mock",
			Name:    "Reader",
		}
		aMockGen, persistor, wantErr := prepareForGenerate(wantConf, t)
		aMock := NewWith(aMockGen, persistor)
		err := aMock.Generate(reflect.TypeOf((*io.Reader)(nil)).Elem())
		if err != wantErr {
			t.Errorf("unexpected err, want '%v' catual '%v'", err, wantErr)
		}
		result := CheckCalls([]*core.Mock{aMockGen.Mock, persistor.Mock})
		if len(result) > 0 {
			t.Error(result)
		}
	})

	t.Run("GenerateAs", func(t *testing.T) {
		wantConf := Conf{
			Path:    "want/pkg",
			Package: "pkg",
			Name:    "ReaderMock",
		}
		aMockGen, persistor, wantErr := prepareForGenerate(wantConf, t)
		aMock := NewWith(aMockGen, persistor)
		err := aMock.GenerateAs(reflect.TypeOf((*io.Reader)(nil)).Elem(),
			wantConf)
		if err != wantErr {
			t.Errorf("unexpected err, want '%v' catual '%v'", err, wantErr)
		}
		result := CheckCalls([]*core.Mock{aMockGen.Mock, persistor.Mock})
		if len(result) > 0 {
			t.Error(result)
		}
	})

	t.Run("Generate for struct", func(t *testing.T) {
		aMock, err := New()
		if err != nil {
			t.Fatal(err)
		}
		err = aMock.Generate(reflect.TypeOf((*big.Int)(nil)).Elem())
		if err != parser.ErrNotInterface {
			t.Fatal(err)
		}
	})

}
