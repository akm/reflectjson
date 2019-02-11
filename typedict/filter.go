package typedict

import (
	"reflect"
)

type Filters []func(reflect.Type) bool

func (filters Filters) Match(t reflect.Type) bool {
	for _, f := range filters {
		if !f(t) {
			return false
		}
	}
	return true
}
