package reflectjson

import (
	// "fmt"
	"image/gif"
	"net/http"

	"testing"
)

func TestSerializeableWithCategories(t *testing.T) {
	categorizedTypes := CategorizedStructs(map[string][]interface{}{
		"http": []interface{}{
			(*http.Request)(nil),
		},
		"image": []interface{}{
			(*gif.GIF)(nil),
		},
	})

	compareStrings := func(name string, expecteds []string) {
		structs := categorizedTypes[name]

		actuals := []string{}
		for _, t := range structs {
			actuals = append(actuals, t.PkgPath+"."+t.Name)
		}

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

	{
		compareStrings("http", []string{
			"crypto/tls.ConnectionState",
			"crypto/x509.Certificate",
			"crypto/x509/pkix.AttributeTypeAndValue",
			"crypto/x509/pkix.Extension",
			"crypto/x509/pkix.Name",
			"math/big.Int",
			"mime/multipart.FileHeader",
			"mime/multipart.Form",
			"net/http.Request",
			"net/http.Response",
			"net/url.URL",
			"net/url.Userinfo",
			"time.Location",
			"time.Time",
			"time.zone",
			"time.zoneTrans",
		})
	}

	{
		compareStrings("image", []string{
			"image.Config",
			"image.Paletted",
			"image.Point",
			"image.Rectangle",
			"image/gif.GIF",
		})
	}

}
