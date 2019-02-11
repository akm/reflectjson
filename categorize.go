package reflectjson

import (
	"reflect"

	"github.com/akm/reflectjson/typedict"
)

func SeriazlizableWithCategories(objectMap map[string][]interface{}, filters ...func(reflect.Type) bool) map[string][]*DataType {
	res := map[string][]*DataType{}

	for key, objects := range objectMap {
		types := typedict.New(objects).Dig().Types(filters...)

		dataTypes := []*DataType{}
		for _, t := range types {
			dt := NewDataType(t)
			dataTypes = append(dataTypes, dt)
		}

		res[key] = dataTypes
	}

	return res
}
