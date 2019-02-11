package typedict

import (
	"fmt"
	"reflect"
)

type TypeDict map[string]reflect.Type

func NewFromTypes(types []reflect.Type) TypeDict {
	m := TypeDict{}
	for _, t := range types {
		m.DigType(t)
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
	return m
}

func (m TypeDict) Keys() []string {
	r := []string{}
	for k, _ := range m {
		r = append(r, k)
	}
	return r
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

func (m TypeDict) Structs(filters ...func(reflect.Type) bool) []reflect.Type {
	filters = append(filters, func(t reflect.Type) bool {
		return t.Kind() == reflect.Struct
	})
	return m.Types(filters...)
}

func (m TypeDict) DigType(t reflect.Type) {
	key := KeyOf(t)
	fmt.Printf("DigType %s\n", key)
	_, ok := m[key]
	if ok {
		return
	}
	m[key] = t
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
		// fmt.Printf("DigStruct %s.%s [%v]\n", t.PkgPath()+"."+t.Name(), f.Name, ft.String())
		m.DigType(ft)
	}
}
