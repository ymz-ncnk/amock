# AMock
With help of AMock you can mock any interface you want.

# Tests
Test coverage is about 90%.

# How to use
First, you should download and install Go, version 1.19 or later.

Create in your home directory a `foo` folder with the following structure:
```
foo/
 |‒‒‒reader_mock.go
 |‒‒‒reader_mock_test.go
```

Create a mock implementation of the `io.Reader` interface.

__reader_mock.go__
```go
package foo

import "github.com/ymz-ncnk/amock"

func NewReaderMock() ReaderMock {
  return ReaderMock{amock.New("Reader")}
}

// ReaderMock is a mock implementation of the io.Reader. It simply uses
// amock.Mock as delegate.
type ReaderMock struct {
  *amock.Mock
}

// RegisterRead registers a function as a one Read() method call.
func (reader ReaderMock) RegisterRead(
  fn func(p []byte) (n int, err error)) ReaderMock {
  reader.Register("Read", fn)
  return reader
}

// RegisterReadN registers a function as n Read() method calls.
func (reader ReaderMock) RegisterReadN(n int,
  fn func(p []byte) (n int, err error)) ReaderMock {
  reader.RegisterN("Read", n, fn)
  return reader
}

// UnregisterRead unregisters Read() method calls.
func (reader ReaderMock) UnregisterRead() ReaderMock {
  reader.Unregister("Read")
  return reader
}

func (reader ReaderMock) Read(p []byte) (n int, err error) {
  result, err := reader.Call("Read", p)
  // err here could be one of the amock.UnexpectedMethodCallError or
  // amock.UnknownMethodCallError
  if err != nil {
    return
  }
  n = result[0].(int)
  err, _ = result[1].(error) 
  return
}
```

Run from the command line:
```bash
$ cd ~/foo
$ go mod init foo
$ go get github.com/ymz-ncnk/amock
```

Now to see how the mock implementation works, let's test it.

__reader_mock_test.go__
```go
package foo

import (
  "io"
  "testing"

  "github.com/ymz-ncnk/amock"
)

func TestSeveralCalls(t *testing.T) {
  // Here we register several calls to Read() method, and then call it several
  // times as well.
  var (
    reader = func() ReaderMock {
      reader := NewReaderMock()
      // We should register all expected method calls. Every method call is
      // simply a functions.
      reader.RegisterRead(func(p []byte) (n int, err error) {
        p[0] = 1
        return 1, nil
      }).RegisterRead(func(p []byte) (n int, err error) {
        p[0] = 2
        p[1] = 2
        return 2, nil
      })
      // If we want to register one function for several calls, we can use
      // RegisterN() method. This is especially useful for concurrent method
      // calls.
      return reader.RegisterReadN(2, func(p []byte) (n int, err error) {
        return 0, io.EOF
      })
    }()
    b = make([]byte, 2)
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

  // If we call the Read() method again we will receive
  // amock.UnexpectedMethodCallError.
  _, err = reader.Read(b)
  if err == nil {
    t.Error("unexpected nil error")
  } else {
    want := amock.NewUnexpectedMethodCallError("Reader", "Read")
    if err.Error() != want.Error() {
      t.Errorf("unexpected error, want '%v', actual '%v'", want, err)
    }
  }
}

func TestUnregisteredCall(t *testing.T) {
  // If no calls have been registered for the method and we call it, we will
  // receive amock.UnknownMethodCallError.
  var (
    reader = NewReaderMock()
    b      []byte
  )
  _, err := reader.Read(b)
  if err == nil {
    t.Error("unexpected nil error")
  } else {
    want := amock.NewUnknownMethodCallError("Reader", "Read")
    if err.Error() != want.Error() {
      t.Errorf("unexpected error, want '%v', actual '%v'", want, err)
    }
  }
}

func TestCheckCallsFunction(t *testing.T) {
  // With help of amock.CheckCalls() we could check if all registered method
  // calls have been called.
  var (
    reader = func() ReaderMock {
      return NewReaderMock().RegisterRead(
        func(p []byte) (n int, err error) {
          p[0] = 1
          return 1, nil
        })
    }()
  )
  m := amock.CheckCalls([]*amock.Mock{reader.Mock})
  if len(m) != 1 {
    t.Fatal("unexpected CheckCalls result")
  }
  arr, pst := m[0]
  if !pst {
    t.Fatal("no 0 key in CheckCalls result")
  }
  if len(arr) != 1 {
    t.Fatal("number of the MethodCallsInfo not equal to 1")
  }
  info := arr[0]
  // test info...
  _ = info
}
```

# Thread safety
Mock implementation is fully threadsafe. You can register, unregister, call
methods and check calls number concurrently.

# Mock implementation caveats
Calling `amock.Call()` method with a `nil` parameter can cause a panic like:
```
  panic: reflect: Call using zero Value argument
  ...
```
To avoid this, we can pass `reflect.Value` to the `amock.Call()` function 
instead of `nil`. For example:
```go
// Mock implementation of the io.WriterTo.
type WriterToMock struct {
  *amock.Mock
}

func (writer WriterToMock) RegisterWriteTo(fn func(w io.Writer) (n int64,
  err error)) WriterToMock {
  writer.Register("WriteTo", fn)
  return writer
}

func (writer WriterToMock) WriteTo(w io.Writer) (n int64, err error) {
  // w param here may be nil, so we have to use wVal
  var wVal reflect.Value
  if w == nil {
    wVal = reflect.Zero(reflect.TypeOf((*io.Writer)(nil)).Elem())
  } else {
    wVal = reflect.ValueOf(w)
  }
  // Call() method can accept reflect.Value-s too.
  vals, err := writer.Call("WriteTo", wVal)
    if err != nil {
    return
  }
  n = vals[0].(int64)
  err, _ = vals[1].(error)
  return
}
```