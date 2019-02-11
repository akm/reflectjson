package reflectjson

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"
)

func Process(objectMap map[string][]interface{}, ptn *regexp.Regexp) {
	res := map[string][]*DataType{}

	for key, objects := range objectMap {
		types := []reflect.Type{}
		for _, obj := range objects {
			types = append(types, reflect.TypeOf(obj))
		}

		types = NewTypeDict(types).Dig().Types(func(t reflect.Type) bool {
			return ptn.MatchString(t.PkgPath())
		})

		dataTypes := []*DataType{}
		for _, t := range types {
			dt := NewDataType(t)
			dataTypes = append(dataTypes, dt)
		}

		res[key] = dataTypes
	}

	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to json.MarshalIndent because of %v\n", err)
		return
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write because of %v\n", err)
		return
	}
}
