package amock

import (
	"errors"
	"testing"

	"github.com/ymz-ncnk/amock/core"
	testdata_amockgen "github.com/ymz-ncnk/amock/testdata/amockgen"
)

func TestCheckCalls(t *testing.T) {
	reader := testdata_amockgen.NewReaderMock()
	reader.RegisterRead(func(p []byte) (n int, err error) {
		return 0, nil
	})
	reader.RegisterRead(func(p []byte) (n int, err error) {
		return 1, nil
	})
	reader.Read([]byte{})
	arr := reader.CheckCalls()
	if len(arr) != 1 {
		t.Error("unexpected CheckCalls result")
	}
	err := checkMethodCallsInfo(arr[0], 2, 1)
	if err != nil {
		t.Error(err)
	}
	reader.Read(nil)
	arr = reader.CheckCalls()
	if len(arr) != 0 {
		t.Error("unexpected CheckCalls result")
	}
}

func checkMethodCallsInfo(info core.MethodCallsInfo, expectedCalls,
	actualCalls int) error {
	if info.MockName != "ReaderMock" {
		return errors.New("unexpected MockName")
	}
	if info.MethodName != "Read" {
		return errors.New("unexpected MethodName")
	}
	if info.ExpectedCalls != expectedCalls {
		return errors.New("unexpected ExpectedCalls")
	}
	if info.ActualCalls != actualCalls {
		return errors.New("unexpected ActualCalls")
	}
	return nil
}
