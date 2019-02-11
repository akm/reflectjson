package typedict

import (
	"reflect"
)

func KeyOf(t reflect.Type) string {
	name := t.Name()
	if name == "" {
		return KeyOf(t.Elem()) + "/" + t.Kind().String()
	} else {
		pkgPath := t.PkgPath()
		if pkgPath == "" {
			return name
		} else {
			return pkgPath + "." + name
		}
	}
}
