package httpserver

import "errors"

var ErrDecodeKey = ErrDecodeKeyArgument{}

var ErrInvalidCertKeyPair = errors.New("invalid cert")

type ErrDecodeKeyArgument struct {
	Type string
}

func (e ErrDecodeKeyArgument) Error() string {
	return "failed to decode " + e.Type + " pem from base64"
}

func (e ErrDecodeKeyArgument) Is(target error) bool {
	_, ok := target.(ErrDecodeKeyArgument)
	return ok
}
