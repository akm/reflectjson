package typedict

import (
	"reflect"
)

func CategorizedStructs(objectMap map[string][]interface{}, filters ...func(reflect.Type) bool) map[string][]*DataType {
	res := map[string][]*DataType{}

	for key, objects := range objectMap {
		structs := New(objects).Structs(filters...)
		dataTypes := SerializableTypes(structs)
		res[key] = dataTypes
	}

	return res
}
