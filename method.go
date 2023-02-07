package amock

import (
	"fmt"
	"reflect"
	"sync"
)

// MethodName is a type for a method name.
type MethodName string

// -----------------------------------------------------------------------------
// MethodCallsInfo holds an information about method calls.
type MethodCallsInfo struct {
	MockName      MockName
	MethodName    MethodName
	ExpectedCalls int
	ActualCalls   int
}

func (info MethodCallsInfo) String() string {
	return fmt.Sprintf("%v.%v() calls count: want %v, actual %v", info.MockName,
		info.MethodName,
		info.ExpectedCalls,
		info.ActualCalls)
}

// -----------------------------------------------------------------------------
// NewMethod creates new Method.
func NewMethod() *Method {
	return &Method{fns: []reflect.Value{}, mu: sync.Mutex{}}
}

// Method represents a struct method.
type Method struct {
	callsCount int
	fns        []reflect.Value
	mu         sync.Mutex
}

// AddMethodCall to the method. Each method call should be a function.
func (method *Method) AddMethodCall(fn Func) {
	method.mu.Lock()
	defer method.mu.Unlock()
	method.fns = append(method.fns, reflect.ValueOf(fn))
}

// Call calls a method once. With help of reflection calls a function,
// registered as a method call, with the given params.
// reflect.Value param is passed to the corresponding function as is.
// If all registered method calls have already been executed, returns
// ErrUnexpectedCall error.
// Threadsafe.
func (method *Method) Call(params []interface{}) ([]interface{}, error) {
	var fn reflect.Value
	method.mu.Lock()
	if len(method.fns) < method.callsCount+1 {
		method.mu.Unlock()
		return nil, ErrUnexpectedCall
	}
	fn = method.fns[method.callsCount]
	method.increaseCallsCount()
	method.mu.Unlock()

	result := fn.Call(toReflectValues(params))
	return fromReflectValues(result), nil
}

// CheckCalls checks method calls. If number of the added method calls is not
// equal to the number of calls, returns ok == false.
func (method *Method) CheckCalls(mockName MockName, methodName MethodName) (
	info MethodCallsInfo, ok bool) {
	method.mu.Lock()
	defer method.mu.Unlock()
	if len(method.fns) != method.callsCount {
		return MethodCallsInfo{mockName, methodName, len(method.fns),
			method.callsCount}, false
	}
	return MethodCallsInfo{}, true
}

func (method *Method) increaseCallsCount() {
	method.callsCount++
}

func toReflectValues(vals []interface{}) []reflect.Value {
	rvals := make([]reflect.Value, len(vals))
	for i := 0; i < len(vals); i++ {
		if rval, ok := vals[i].(reflect.Value); ok {
			rvals[i] = rval
		} else {
			rvals[i] = reflect.ValueOf(vals[i])
		}
	}
	return rvals
}

func fromReflectValues(rvals []reflect.Value) []interface{} {
	vals := make([]interface{}, len(rvals))
	for i := 0; i < len(vals); i++ {
		vals[i] = rvals[i].Interface()
	}
	return vals
}
