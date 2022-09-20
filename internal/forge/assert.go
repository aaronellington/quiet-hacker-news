package forge

import (
	"reflect"
	"testing"
)

// Assert is foobar
func Assert(t *testing.T, expected interface{}, actual interface{}) bool {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Got %v, Expected: %v", actual, expected)
		return false
	}

	return true
}
