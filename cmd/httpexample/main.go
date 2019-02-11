package main

import (
	"net/http"
	"os"
	"reflect"
	"regexp"

	"github.com/akm/typedict"
)

// $ go run cmd/httpexample/main.go
func main() {
	dict := typedict.New([]interface{}{
		(*http.Request)(nil),
	})

	ptn := regexp.MustCompile(`\Anet/`)

	// TypeDict.Structs returns a silce of reflect.Type of struct
	structs := typedict.SerializableTypes(dict.Structs(func(t reflect.Type) bool {
		return ptn.MatchString(t.PkgPath())
	}))
	typedict.WriteJson(os.Stdout, structs)
}
