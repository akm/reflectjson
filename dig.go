package reflectjson

import (
	"reflect"
)

func DigTypes(types []reflect.Type) []reflect.Type {
	m := map[string]reflect.Type{}
	for _, t := range types {
		key := t.PkgPath() + "." + t.Name()
		m[key] = t
	}

	for _, t := range types {
		DigType(m, t)
	}

	r := []reflect.Type{}
	for _, t := range m {
		r = append(r, t)
	}
	return r
}

func DigType(m map[string]reflect.Type, t reflect.Type) {
	switch t.Kind() {
	case reflect.Struct:
		DigStruct(m, t)
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		DigType(m, t.Elem())
	}
}

func DigStruct(m map[string]reflect.Type, t reflect.Type) {
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
					DigType(m, ft)
				}
			}
		}
	}
}
