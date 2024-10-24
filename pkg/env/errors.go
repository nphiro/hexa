package env

import "errors"

var ErrGetCurrWorkingDir = errors.New("cannot get current working directory")

var ErrReadFileEnv = errors.New("cannot read .env file")

var ErrParseConfig = errors.New("cannot parse env config")

var ErrParseEnv = errors.New("cannot parse env value to object")
