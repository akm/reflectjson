package reflectjson

import (
	"reflect"
)

type TypeDict map[string]reflect.Type

func NewTypeDict(types []reflect.Type) TypeDict {
	m := TypeDict{}
	for _, t := range types {
		key := t.PkgPath() + "." + t.Name()
		m[key] = t
	}

	return m.Dig()
}

func (m TypeDict) Dig() TypeDict {
	for _, t := range m {
		m.DigType(t)
	}
	return m
}

func DigTypes(types []reflect.Type) []reflect.Type {
	m := NewTypeDict(types)

	r := []reflect.Type{}
	for _, t := range m {
		r = append(r, t)
	}
	return r
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
			key := ft.PkgPath() + "." + ft.Name()
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
