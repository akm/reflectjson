package reflectjson

import (
	"os"
	"reflect"
	"regexp"
)

func Process(objectMap map[string][]interface{}, ptn *regexp.Regexp) {
	Output(os.Stdout, SeriazlizableWithCategories(objectMap, func(t reflect.Type) bool {
		return ptn.MatchString(t.PkgPath())
	}))
}
