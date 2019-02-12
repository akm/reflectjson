package typedict

import (
	"reflect"
)

type Filters []func(reflect.Type) bool

func (filters Filters) Match(t reflect.Type) bool {
	for _, f := range filters {
		if !f(t) {
			return false
		}
	}
	return true
}

// See https://golang.org/pkg/reflect/#Kind
var SimpleKinds = []reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	// reflect.Uintptr,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
	// reflect.Array,
	// reflect.Chan,
	// reflect.Func,
	// reflect.Interface,
	// reflect.Map,
	// reflect.Ptr,
	// reflect.Slice,
	// reflect.String,
	// reflect.Struct,
	// reflect.UnsafePointer,
}

func KindFilter(kinds ...reflect.Kind) func(reflect.Type) bool {
	return func(t reflect.Type) bool {
		for _, kind := range kinds {
			if t.Kind() == kind {
				return true
			}
		}
		return false
	}
}
