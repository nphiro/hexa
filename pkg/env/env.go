package env

import (
	"reflect"
)

func Write[T any](obj T, opts *WriteOptions) (err error) {
	if opts == nil {
		opts = NewWriteOptions()
	}
	loadEnvFile(opts.EnvFile)
	if err := writeStruct(reflect.ValueOf(obj).Elem()); err != nil {
		return err
	}
	return nil
}
