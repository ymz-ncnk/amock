package amock

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestAmockIntegration(t *testing.T) {
	dname, err := os.MkdirTemp("", "amock")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dname)
	// dname = "testdata"

	aMock, err := New()
	if err != nil {
		t.Fatal(err)
	}
	conf := Conf{Package: "testdata", Name: "ReaderMock", Path: dname}
	err = aMock.GenerateAs(reflect.TypeOf((*io.Reader)(nil)).Elem(),
		conf)
	if err != nil {
		t.Fatal(err)
	}
	wantPath := conf.Path + string(os.PathSeparator) + conf.Name +
		FilenameExtenstion
	if _, err := os.Stat(wantPath); err != nil {
		t.Error("file was not generated")
	}
}
