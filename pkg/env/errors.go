package env

import "errors"

var ErrGetCurrWorkingDir = errors.New("cannot get current working directory")

var ErrReadFileEnv = errors.New("cannot read .env file")

var ErrWriteStruct = errors.New("cannot write env value to struct")
