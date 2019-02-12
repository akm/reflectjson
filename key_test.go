package typedict

import (
	"net/http"
	"net/url"
	"reflect"

	"testing"
)

type EnumTypeA int

const (
	EnumTypeA1 EnumTypeA = 1
	EnumTypeA2 EnumTypeA = 2
)

func TestKeyOf(t *testing.T) {
	patterns := map[string]interface{}{
		"string":                "",
		"string:ptr":            (*string)(nil),
		"string:ptr:slice":      []*string{},
		"string:slice":          []string{},
		"string:slice:slice":    [][]string{},
		"map[string]string":     map[string]string{},
		"map[string]string:ptr": map[string]*string{},
		"map[string:ptr]string": map[*string]string{},
		"int":                                   0,
		"int:ptr":                               (*int)(nil),
		"int:ptr:slice":                         []*int{},
		"int:slice":                             []int{},
		"int:slice:slice":                       [][]int{},
		"net/http.Request":                      http.Request{},
		"net/http.Request:ptr":                  (*http.Request)(nil),
		"net/url.Values":                        (url.Values)(nil),
		"map[string]string:slice":               (map[string][]string)(nil),
		"github.com/akm/typedict.EnumTypeA":     EnumTypeA1,
		"github.com/akm/typedict.EnumTypeA:ptr": (*EnumTypeA)(nil),
	}

	for expected, obj := range patterns {
		actual := KeyOf(reflect.TypeOf(obj))
		if expected != actual {
			t.Errorf("Expects %v but was %v\n", expected, actual)
		}
	}

}
