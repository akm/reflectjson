package typedict

import (
	"net/http"
	"reflect"

	"testing"
)

func TestKeyOf(t *testing.T) {
	patterns := map[string]interface{}{
		"string":               "",
		"string/ptr":           (*string)(nil),
		"string/ptr/slice":     []*string{},
		"string/slice":         []string{},
		"string/slice/slice":   [][]string{},
		"int":                  0,
		"int/ptr":              (*int)(nil),
		"int/ptr/slice":        []*int{},
		"int/slice":            []int{},
		"int/slice/slice":      [][]int{},
		"net/http.Request":     http.Request{},
		"net/http.Request/ptr": (*http.Request)(nil),
	}

	for expected, obj := range patterns {
		actual := KeyOf(reflect.TypeOf(obj))
		if expected != actual {
			t.Errorf("Expects %v but was %v\n", expected, actual)
		}
	}

}
