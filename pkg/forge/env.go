package forge

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	// ErrInvalidValue returned when the value passed to Unmarshal is nil or not a
	// pointer to a struct.
	ErrInvalidValue = errors.New("value must be a non-nil pointer to a struct")

	// ErrUnsupportedType returned when a field with tag "env" is unsupported.
	ErrUnsupportedType = errors.New("field is an unsupported type")

	// ErrUnexportedField returned when a field with tag "env" is not exported.
	ErrUnexportedField = errors.New("field must be exported")
)

// Env sources .env files and parses variables into a existing struct
func Env(target interface{}, dotEnvFilePaths ...string) error {
	EnvReadFiles(dotEnvFilePaths...)

	if err := EnvUnmarshal(target); err != nil {
		return err
	}

	return nil
}

// EnvReadFiles locates and parses .env files
func EnvReadFiles(dotEnvFilePaths ...string) {
	// Default dotEnvFilePaths
	if len(dotEnvFilePaths) < 1 {
		dotEnvFilePaths = []string{
			".env",
			".env.local",
		}
	}

	scratchPad := map[string]string{}

	for _, dotEnvFilePath := range dotEnvFilePaths {
		dotEnvVars := readDotEnvFile(dotEnvFilePath)
		for key, value := range dotEnvVars {
			scratchPad[key] = value
		}
	}

	for key, value := range scratchPad {
		os.Setenv(key, value)
	}
}

func readDotEnvFile(filePath string) map[string]string {
	results := map[string]string{}

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return results
	}

	fileString := string(fileBytes)

	lines := strings.Split(fileString, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		value := parts[1]
		if _, alreadyExists := os.LookupEnv(key); !alreadyExists {
			results[key] = value
		}
	}

	return results
}

// EnvUnmarshal into a target struct
func EnvUnmarshal(target interface{}) error {
	rv := reflect.ValueOf(target)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return ErrInvalidValue
	}

	rv = rv.Elem()

	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		valueField := rv.Field(i)
		switch valueField.Kind() {
		case reflect.Struct:
			valueInterface := valueField.Addr().Interface()
			err := Env(valueInterface)
			if err != nil {
				return err
			}
		}

		typeField := t.Field(i)
		tag := typeField.Tag.Get("env")
		if tag == "" {
			continue
		}

		if !valueField.CanSet() {
			return ErrUnexportedField
		}

		envVar, ok := os.LookupEnv(tag)
		if !ok {
			continue
		}

		err := reflectSet(typeField.Type, valueField, envVar)
		if err != nil {
			return err
		}
	}

	return nil
}

func reflectSet(t reflect.Type, f reflect.Value, value string) error {
	switch t.Kind() {
	case reflect.Ptr:
		ptr := reflect.New(t.Elem())
		err := reflectSet(t.Elem(), ptr.Elem(), value)
		if err != nil {
			return err
		}
		f.Set(ptr)
	case reflect.String:
		f.SetString(value)
	case reflect.Bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		f.SetBool(v)
	case reflect.Int:
		v, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		f.SetInt(int64(v))
	default:
		return ErrUnsupportedType
	}

	return nil
}
