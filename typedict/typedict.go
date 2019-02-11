package typedict

import (
	// "fmt"
	"reflect"
)

type TypeDict map[string]reflect.Type

func NewFromTypes(types []reflect.Type) TypeDict {
	m := TypeDict{}
	for _, t := range types {
		m[KeyOf(t)] = t
	}

	return m
}

func New(objects []interface{}) TypeDict {
	types := []reflect.Type{}
	for _, obj := range objects {
		types = append(types, reflect.TypeOf(obj))
	}
	return NewFromTypes(types)
}

func (m TypeDict) Dig() TypeDict {
	for _, t := range m {
		m.DigType(t)
	}
	return m
}

func (m TypeDict) Types(filters ...func(reflect.Type) bool) []reflect.Type {
	r := []reflect.Type{}
	for _, t := range m {
		if Filters(filters).Match(t) {
			r = append(r, t)
		}
	}
	return r
}

func (m TypeDict) Include(t reflect.Type) bool {
	_, ok := m[KeyOf(t)]
	return ok
}

func (m TypeDict) DigType(t reflect.Type) {
	switch t.Kind() {
	case reflect.Struct:
		m.DigStruct(t)
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		m.DigType(t.Elem())
	}
}

func (m TypeDict) DigStruct(t reflect.Type) {
	numField := t.NumField()
	for i := 0; i < numField; i++ {
		f := t.Field(i)
		ft := f.Type
		if ft.PkgPath() != "" {
			key := KeyOf(ft)
			_, ok := m[key]
			if !ok {
				m[key] = ft
				switch ft.Kind() {
				case reflect.Struct:
					m.DigType(ft)
				}
			}
		}
	}
}
