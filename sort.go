package typedict

import (
	"reflect"
)

func DataTypeSorter(dataTypes []*DataType) func(int, int) bool {
	return func(i, j int) bool {
		return (dataTypes[i].PkgPath + "." + dataTypes[i].Name) < (dataTypes[j].PkgPath + "." + dataTypes[j].Name)
	}
}

func ReflectTypeSorter(types []reflect.Type) func(int, int) bool {
	return func(i, j int) bool {
		return (types[i].PkgPath() + "." + types[i].Name()) < (types[j].PkgPath() + "." + types[j].Name())
	}
}
