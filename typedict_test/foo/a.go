package foo

import (
	"github.com/akm/typedict/typedict_test/bar"
	"github.com/akm/typedict/typedict_test/baz"
)

type A struct {
	Bar []*bar.B
	Baz baz.D
}
