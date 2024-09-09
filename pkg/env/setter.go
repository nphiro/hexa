package env

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

func writeStruct(v reflect.Value) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ErrWriteStruct
		}
	}()

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.Struct {
			writeStruct(field)
		} else {
			writeStructField(field, fieldType)
		}
	}
	return nil
}

func writeStructField(field reflect.Value, fieldType reflect.StructField) {
	tag := fieldType.Tag.Get("env")
	if tag == "" {
		return
	}
	envValue := getEnvValue(tag, fieldType.Tag.Get("envDefault"))
	if envValue == "" {
		return
	}

	writeValue(field, envValue)
}

func getEnvValue(key, defaultValue string) string {
	envValue := os.Getenv(key)
	if envValue == "" {
		envValue = defaultValue
	}
	return strings.TrimSpace(envValue)
}

func writeValue(field reflect.Value, envValue string) error {
	switch field.Kind() {
	case reflect.String:
		return setString(field, envValue)
	case reflect.Bool:
		return setBool(field, envValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setInt(field, envValue)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUint(field, envValue)
	case reflect.Float32, reflect.Float64:
		return setFloat(field, envValue)
	case reflect.Slice:
		return setSlice(field, envValue)
	case reflect.Map:
		return setMap(field, envValue)
	}
	return nil
}

func setString(field reflect.Value, envValue string) error {
	field.SetString(envValue)
	return nil
}

func setBool(field reflect.Value, envValue string) error {
	boolValue, err := strconv.ParseBool(envValue)
	if err != nil {
		return err
	}
	field.SetBool(boolValue)
	return nil
}

func setInt(field reflect.Value, envValue string) error {
	intValue, err := strconv.ParseInt(envValue, 10, 64)
	if err != nil {
		return err
	}
	field.SetInt(intValue)
	return nil
}

func setUint(field reflect.Value, envValue string) error {
	uintValue, err := strconv.ParseUint(envValue, 10, 64)
	if err != nil {
		return err
	}
	field.SetUint(uintValue)
	return nil
}

func setFloat(field reflect.Value, envValue string) error {
	floatValue, err := strconv.ParseFloat(envValue, 64)
	if err != nil {
		return err
	}
	field.SetFloat(floatValue)
	return nil
}

func setSlice(field reflect.Value, envValue string) error {
	values := strings.Split(envValue, ",")
	slice := reflect.MakeSlice(field.Type(), len(values), len(values))
	for i, v := range values {
		if err := writeValue(slice.Index(i), v); err != nil {
			return err
		}
	}
	field.Set(slice)
	return nil
}

func setMap(field reflect.Value, envValue string) error {
	values := strings.Split(envValue, ",")
	m := reflect.MakeMap(field.Type())
	for _, v := range values {
		kv := strings.Split(v, ":")
		if len(kv) != 2 {
			continue
		}
		key := reflect.New(field.Type().Key()).Elem()
		if err := writeValue(key, kv[0]); err != nil {
			return err
		}
		value := reflect.New(field.Type().Elem()).Elem()
		if err := writeValue(value, kv[1]); err != nil {
			return err
		}
		m.SetMapIndex(key, value)
	}
	field.Set(m)
	return nil
}
