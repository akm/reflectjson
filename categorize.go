package typedict

import (
	"reflect"
)

func CategorizedTypes(objectMap map[string][]interface{}, filters ...func(reflect.Type) bool) map[string][]*DataType {
	res := map[string][]*DataType{}

	for key, objects := range objectMap {
		types := New(objects).Types(filters...)
		dataTypes := SerializableTypes(types)
		res[key] = dataTypes
	}

	return res
}
