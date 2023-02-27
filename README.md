# AMock
AMock is a simple thread-safe mocking library for Golang. It helps you generate 
mock implementations of interfaces.

# Tests
Test coverage is about 85%.

# How to use
First, you should download and install Go, version 1.14 or later.

Create in your home directory a `foo` folder with the following structure:
```
foo/
 |‒‒‒gen/
 |    |‒‒‒mock.go
 |‒‒‒testdata/
 |      |‒‒‒mock/
 |‒‒‒foo.go
```

__foo.go__
```go
//go:generate go run gen/mock.go
package foo
```

__gen/mock.go__
```go
//go:build ignore

package main

import (
  "io"
  "reflect"

  "github.com/ymz-ncnk/amock"
)

func main() {
  aMock, err := amock.New()
  if err != nil {
    panic(err)
  }
  tp := reflect.TypeOf((*io.Reader)(nil)).Elem()
  // Generated filename and mock implementation type will be equal to tp.Name().
  // Also generated file will be placed into the "testdata/mock" folder.
  // If you want to change these defaults use AMock.GenerateAs() method.
  err = aMock.Generate(tp)
  if err != nil {
    panic(err)
  }
}
```

Run from the command line:
```bash
$ cd ~/foo
$ go mod init foo
$ go get github.com/ymz-ncnk/amock
$ go generate
```

Now you can see `Reader.gen.go` file in the `testdata/mock` folder, which is 
simply uses `*amock_core.Mock` as a delegate. To see how this mock 
implementation works, let's test it. Create a `reader_mock_test.go` file:
```
foo/
 |‒‒‒...
 |‒‒‒reader_mock_test.go
```

__reader_mock_test.go__
```go
package foo

import (
  "io"
  "testing"

  "foo/testdata/mock"

  "github.com/ymz-ncnk/amock"
  amock_core "github.com/ymz-ncnk/amock/core"
)

func TestSeveralCalls(t *testing.T) {
  // Here we register several calls to the Read() method, and then call it 
  // several times as well.	
  var (
    b      = make([]byte, 2)
    reader = func() mock.Reader {
      reader := mock.NewReader()
      // We must register all expected method calls. Each method call is just a 
      // function.			
      reader.RegisterRead(func(p []byte) (n int, err error) {
        p[0] = 1
        return 1, nil
      }).RegisterRead(func(p []byte) (n int, err error) {
        p[0] = 2
        p[1] = 2
        return 2, nil
      })
      // If we want to register one function for multiple calls, we can use the 
      // RegisterN() method. This is especially useful for concurrent method 
      // calls.
      return reader.RegisterNRead(2, func(p []byte) (n int, err error) {
        return 0, io.EOF
      })
    }()
  )

  // In total, we registered 4 calls for the Read() method.

  // First call.
  n, err := reader.Read(b)
  if err != nil {
    panic(err)
  }
  // We are expecting to read 1 byte.
  if n != 1 {
    t.Errorf("unexpected n, want '%v', actual '%v'", 1, n)
  }
  // Here we could test err and b values ...

  // Second call.
  n, err = reader.Read(b)
  if err != nil {
    panic(err)
  }
  // We are expecting to read 2 bytes.
  if n != 2 {
    t.Errorf("unexpected n, want '%v', actual '%v'", 2, n)
  }
  // ...

  // Third call.
  _, err = reader.Read(b)
  // We are expecting to receive io.EOF error
  if err != io.EOF {
    t.Errorf("unexpected err, want '%v', actual '%v'", io.EOF, err)
  }
  // ...

  // Forth call.
  _, err = reader.Read(b)
  // We are expecting to receive io.EOF error
  if err != io.EOF {
    t.Errorf("unexpected err, want '%v', actual '%v'", io.EOF, err)
  }
  // ...

  // If we call the Read() method again we will receive a panic with
  // amock.UnexpectedMethodCallError.
  defer func() {
    want := amock_core.NewUnexpectedMethodCallError("Reader", "Read")
    if r := recover(); r != nil {
      if err, ok := r.(error); ok {
        if err.Error() != want.Error() {
          t.Errorf("unexpected error, want '%v', actual '%v'", want, err)
        }
      }
    }
  }()
  reader.Read(b)
  t.Error("no panic")
}

func TestUnknownMethodCall(t *testing.T) {
  // If no calls have been registered for the method and we call it, we will
  // receive a panic with amock.UnknownMethodCallError.
  reader := mock.NewReader()

  // Handle panic.
  defer func() {
    want := amock_core.NewUnknownMethodCallError("Reader", "Read")
    if r := recover(); r != nil {
      if err, ok := r.(error); ok {
        if err.Error() != want.Error() {
          t.Errorf("unexpected error, want '%v', actual '%v'", want, err)
        }
      }
    }
  }()
  reader.Read([]byte{})
  t.Error("no panic")
}

func TestUnregister(t *testing.T) {
  reader := mock.NewReader()

  // Register two Read() method calls.
  reader.RegisterNRead(2,
    func(p0 []uint8) (r0 int, r1 error) {
      return
    },
  )
  // Unregister all Read() method calls.
  reader.Unregister("Read")

  // Handle panic.
  defer func() {
    want := amock_core.NewUnknownMethodCallError("Reader", "Read")
    if r := recover(); r != nil {
      if err, ok := r.(error); ok {
        if err.Error() != want.Error() {
          t.Errorf("unexpected error, want '%v', actual '%v'", want, err)
        }
      }
    }
  }()

  reader.Read([]byte{})
  t.Error("no panic")
}

func TestCheckCalls(t *testing.T) {
  // With the CheckCalls() method we can check if all the registered method 
  // calls have been executed.
  var (
    want = []amock_core.MethodCallsInfo{
      {
        MockName:      "Reader",
        MethodName:    "Read",
        ExpectedCalls: 1,
        ActualCalls:   0,
      },
    }
    reader = func() mock.Reader {
      return mock.NewReader().RegisterRead(
        func(p []byte) (n int, err error) {
          p[0] = 1
          return 1, nil
        })
    }()
  )
  if callsInfo := reader.CheckCalls(); !reflect.DeepEqual(callsInfo, want) {
    t.Errorf("unexpected callsInfo, want '%v' actual '%v'", want, callsInfo)
  }
  // There is also exists amock.CheckCalls([]*core.Mock) function, which you can
  // use for many mocks.
}
```

## In concurrent test
Let's see how we can use the `Reader` mock in concurrent test. Create a 
`concurrent_test.go` file:
```
foo/
 |‒‒‒...
 |‒‒‒concurrent_test.go
```

__concurrent_test.go__
```go
package foo

import (
  "bytes"
  "foo/testdata/mock"
  "sync"
  "testing"
)

func TestConcurrent(t *testing.T) {
  // Here we want to read twice, concurrently.
  var (
    want = map[string]struct {
      inB    []byte // Input param for Reader.Read().
      outErr error  // Output param.
      expB   []byte // Expected b value.
    }{
      string([]byte{1}):    {[]byte{1}, nil, []byte{10}},
      string([]byte{2, 2}): {[]byte{2, 2}, nil, []byte{20, 20}},
    }

    // Register two calls to the Read() method for the Reader mock.
    // We cannot predict the order of these calls, so we use the RegisterN()
    // method instead of the two Register() methods.
    reader = mock.NewReader().RegisterNRead(2,
      func(b []byte) (n int, err error) {
        if v, pst := want[string(b)]; pst {
          copy(b, v.expB)
          return len(v.expB), v.outErr
        }
        t.Errorf("unexpected input, b = '%v'", b)
        return
      },
    )
    // Create a WaitGroup to know when all the calls are done.
    wg = func() *sync.WaitGroup {
      wg := sync.WaitGroup{}
      wg.Add(len(want))
      return &wg
    }()
  )

  // Each call to the Read() method we make in a separate gorountine.
  for _, v := range want {
    go func(b []byte, wantErr error, wantB []byte) {
      n, err := reader.Read(b)
      // Let's test the results.
      if n != len(wantB) || err != wantErr || !bytes.Equal(b[:n], wantB) {
        t.Errorf("unexpected result n = '%v' err = '%v' b = '%v'", n, err, b)
      }
      wg.Done()
    }(v.inB, v.outErr, v.expB)
  }

  wg.Wait()

  // And finally check the calls that have been made.
  if callsInfo := reader.CheckCalls(); len(callsInfo) > 0 {
    t.Error(callsInfo)
  }
}
```

# Thread safety
Mock implementation is fully threadsafe. You can register, unregister, call
methods and check calls number concurrently.