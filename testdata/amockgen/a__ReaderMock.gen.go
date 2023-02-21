// Code generated by amockgen. DO NOT EDIT.

package amockgen

import amock_core "github.com/ymz-ncnk/amock/core"

// New creates a new ReaderMock.
func NewReaderMock() ReaderMock {
	return ReaderMock{
		Mock: amock_core.New("ReaderMock"),
	}
}

// ReaderMock is a mock implementation of the io.Reader.
type ReaderMock struct {
	*amock_core.Mock
}

// RegisterRead registers a function as a single Read() method call.
func (mock ReaderMock) RegisterRead(
	fn func(p0 []uint8) (r0 int, r1 error)) ReaderMock {
	mock.Register("Read", fn)
	return mock
}

// RegisterRead registers a function as n Read() method calls.
func (mock ReaderMock) RegisterNRead(n int,
	fn func(p0 []uint8) (r0 int, r1 error)) ReaderMock {
	mock.RegisterN("Read", n, fn)
	return mock
}

// UnregisterRead unregisters Read() method calls.
func (mock ReaderMock) UnregisterRead() ReaderMock {
	mock.Unregister("Read")
	return mock
}

func (mock ReaderMock) Read(p0 []uint8) (r0 int, r1 error) {
	result, err := mock.Call("Read", p0)
	if err != nil {
		panic(err)
	}
	r0 = result[0].(int)
	r1, _ = result[1].(error)
	return
}