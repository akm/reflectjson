package typedict

import (
	"os"
	"reflect"
	"regexp"
)

func Process(objectMap map[string][]interface{}, ptn *regexp.Regexp) {
	WriteJson(os.Stdout, CategorizedTypes(objectMap, KindFilter(reflect.Struct), func(t reflect.Type) bool {
		return ptn.MatchString(t.PkgPath())
	}))
}
