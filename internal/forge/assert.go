package forge

import (
	"fmt"
	"reflect"
)

// Assert is foobar
func Assert(expected interface{}, actual interface{}) error {
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("not equal... Got: %v, Expected: %v", actual, expected)
	}

	return nil
}
