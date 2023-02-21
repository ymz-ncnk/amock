package core

import (
	"fmt"
	"reflect"
	"sync"
)

// MockName is a type for a mock name.
type MockName string

// Func should have a function value. Used to represent one method call.
type Func interface{}

// New creates new Mock.
func New(name MockName) *Mock {
	return &Mock{name: name}
}

// Mock helps you to mock interfaces.
type Mock struct {
	name MockName
	m    sync.Map
}

// Register registers a method. A function is registered as one method call.
// You could chain Register calls:
// mock.Register("Handle", ...).Register("Handle", ...)
func (mock *Mock) Register(name MethodName, fn Func) *Mock {
	if !isFunc(fn) {
		panic(ErrNotFunction)
	}
	method, _ := mock.m.LoadOrStore(name, NewMethod())
	method.(*Method).AddMethodCall(fn)
	return mock
}

// RegisterN registers a method. A function is registered as several method
// calls.
func (mock *Mock) RegisterN(name MethodName, n int, fn Func) *Mock {
	for i := 0; i < n; i++ {
		mock.Register(name, fn)
	}
	return mock
}

// Unregister unregisters a method.
func (mock *Mock) Unregister(name MethodName) *Mock {
	mock.m.Delete(name)
	return mock
}

// Call calls a method with specified parameters. Uses reflection to execute
// functions registered as method calls. Note that the reflect.Value parameters
// are passed to these functions as is.
// If no method was registered, UnknownMethodCallError is returned. If all
// registered method calls have already been made, UnexpectedMethodCallError is
// returned.
func (mock *Mock) Call(name MethodName, params ...interface{}) (
	[]interface{}, error) {
	method, pst := mock.m.Load(name)
	if !pst {
		return nil, NewUnknownMethodCallError(mock.name, name)
	}
	vals, err := method.(*Method).Call(params)
	if err != nil {
		if err == ErrUnexpectedCall {
			return nil, NewUnexpectedMethodCallError(mock.name, name)
		} else {
			panic(fmt.Sprintf("unepxected '%v' err", err))
		}
	}
	return vals, nil
}

// CheckCalls checks method calls. If all registered methods were called the
// estimated number of times, an empty array is returned.
func (mock *Mock) CheckCalls() []MethodCallsInfo {
	arr := []MethodCallsInfo{}
	mock.m.Range(func(key, value interface{}) bool {
		methodName := key.(MethodName)
		method := value.(*Method)
		info, ok := method.CheckCalls(mock.name, methodName)
		if !ok {
			arr = append(arr, info)
		}
		return true
	})
	return arr
}

func isFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}
