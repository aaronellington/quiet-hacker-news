package forge

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// AssertFailure is foobar
type AssertFailure struct {
	Expected interface{}
	Actual   interface{}
}

// Assert is foobar
func Assert(expected interface{}, actual interface{}) error {
	if !reflect.DeepEqual(expected, actual) {
		actualBytes, _ := json.Marshal(AssertFailure{
			Expected: expected,
			Actual:   actual,
		})

		return fmt.Errorf("not equal... %s", string(actualBytes))
	}

	return nil
}
