package forge

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	// ErrInvalidValue returned when the value passed to Unmarshal is nil or not a pointer to a struct.
	ErrInvalidValue = errors.New("value must be a non-nil pointer to a struct")

	// ErrUnsupportedFieldType returned when a field with tag "env" is unsupported.
	ErrUnsupportedFieldType = errors.New("field is an unsupported type")

	// ErrUnexportedField returned when a field with tag "env" is not exported.
	ErrUnexportedField = errors.New("field must be exported")
)

func NewEnvironment() Environment {
	environment := Environment{}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		environment[pair[0]] = pair[1]
	}

	return environment
}

type Environment map[string]string

func (environment Environment) ImportEnvFileContents(fileContents string) error {
	lines := strings.Split(fileContents, "\n")
	for _, line := range lines {
		// Skip blank lines
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		key := parts[0]
		value := parts[1]

		// Skip the env variable is already set
		if _, alreadySet := environment[key]; alreadySet {
			continue
		}

		// Set the value
		environment[key] = value
	}

	return nil
}

func (environment Environment) Decode(target interface{}) error {
	rv := reflect.ValueOf(target)

	// Make sure it's not a primitive type or nil before calling Elem()
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return ErrInvalidValue
	}
	rv = rv.Elem()

	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fieldValue := rv.Field(i)
		fieldType := t.Field(i)

		// If the field is a struct, call ourselves to keep scanning deeper
		if fieldValue.Kind() == reflect.Struct {
			valueInterface := fieldValue.Addr().Interface()
			err := environment.Decode(valueInterface)
			if err != nil {
				return err
			}
		}

		// Get the tag value
		tag := fieldType.Tag.Get("env")
		if tag == "" {
			continue
		}

		// Confirm the value can be set
		if !fieldValue.CanSet() {
			return ErrUnexportedField
		}

		// Get the existing value from the environment
		newValue, ok := environment[tag]
		if !ok {
			continue
		}

		// Actually modify the target
		err := reflectSet(fieldType.Type, fieldValue, newValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func reflectSet(fieldType reflect.Type, fieldValue reflect.Value, rawValueFromSource string) error {
	switch fieldType.Kind() {
	case reflect.Ptr:
		ptr := reflect.New(fieldType.Elem())
		err := reflectSet(fieldType.Elem(), ptr.Elem(), rawValueFromSource)
		if err != nil {
			return err
		}
		fieldValue.Set(ptr)
	case reflect.String:
		fieldValue.SetString(rawValueFromSource)
	case reflect.Bool:
		newValue, err := strconv.ParseBool(rawValueFromSource)
		if err != nil {
			return err
		}
		fieldValue.SetBool(newValue)
	case reflect.Int:
		newValue, err := strconv.Atoi(rawValueFromSource)
		if err != nil {
			return err
		}
		fieldValue.SetInt(int64(newValue))
	default:
		return ErrUnsupportedFieldType
	}

	return nil
}
