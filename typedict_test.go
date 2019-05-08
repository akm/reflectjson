package typedict

import (
	// "fmt"
	"net/http"
	"reflect"
	"sort"

	"testing"

	"github.com/akm/typedict/typedict_test/foo"
)

func TestTypeDict(t *testing.T) {
	dict := New([]interface{}{
		http.Request{},
	})

	{
		actualKeys := dict.Keys()
		sort.Strings(actualKeys)

		expecteds := []string{
			"net/url.URL",
			// "net/url.URL:ptr",
			// "net/url.Values",
		}

		for _, expected := range expecteds {
			if _, ok := dict[expected]; !ok {
				t.Errorf("Expects %s but wasn't included\n", expected)
			}
		}
	}

	{
		structs := dict.Types(KindFilter(reflect.Struct))
		structNames := []string{}
		for _, t := range structs {
			structNames = append(structNames, KeyOf(t))
		}

		expecteds := []string{
			"net/http.Request",
			"net/http.Response",
			"net/url.URL",
			"mime/multipart.Form",
		}

		for _, expected := range expecteds {
			found := false
			for _, structName := range structNames {
				if expected == structName {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("%s not found as a struct", expected)
			}
		}
	}

}

func compareStrings(t *testing.T, name string, actuals, expecteds []string) {
	if len(expecteds) != len(actuals) {
		t.Errorf("%s's length expects %d but was %d\nexpected: %v\nactual: %v\n", name, len(expecteds), len(actuals), expecteds, actuals)
		return
	}
	for i, expected := range expecteds {
		if expected != actuals[i] {
			t.Errorf("%s [%d] expects %s but was %s\nexpected: %v\nactual: %v\n", name, i, expected, actuals[i], expecteds, actuals)
		}
	}
}


type TestEnumB int

const (
	TestEnumB1 TestEnumB = 1
	TestEnumB2 TestEnumB = 2
)

type TestStruct1 struct {
	Name   string
	Status TestEnumB
}

func TestTypeDictWithCustomStruct(t *testing.T) {
	dict := New([]interface{}{
		(*TestStruct1)(nil),
	})

	{
		actualKeys := dict.Keys()
		sort.Strings(actualKeys)

		compareStrings(t, "keys", actualKeys, []string{
			"github.com/akm/typedict.TestEnumB",
			"github.com/akm/typedict.TestStruct1",
			"string",
		})

		types := dict.Types(KindFilter(append([]reflect.Kind{reflect.Struct}, SimpleKinds...)...))
		typeNames := []string{}
		for _, t := range types {
			typeNames = append(typeNames, t.PkgPath()+"."+t.Name())
		}
		compareStrings(t, "keys", typeNames, []string{
			"github.com/akm/typedict.TestEnumB",
			"github.com/akm/typedict.TestStruct1",
		})
	}
}


func TestTypeDictDIgType(t *testing.T) {
	dict := New([]interface{}{
		(*foo.A)(nil),
	})

	actualKeys := dict.Keys()
	sort.Strings(actualKeys)

	compareStrings(t, "keys", actualKeys, []string{
		"github.com/akm/typedict/typedict_test/bar.B",
		"github.com/akm/typedict/typedict_test/baz.C",
		"github.com/akm/typedict/typedict_test/baz.D",
		"github.com/akm/typedict/typedict_test/foo.A",
		"string",
	})
}
