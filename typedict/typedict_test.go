package typedict

import (
	// "fmt"
	"net/http"
	"sort"

	"testing"
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
			"net/url.URL:ptr",
			"net/url.Values",
		}

		for _, expected := range expecteds {
			if _, ok := dict[expected]; !ok {
				t.Errorf("Expects %s but wasn't included\n", expected)
			}
		}
	}

	{
		structs := dict.Structs()
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
