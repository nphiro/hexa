package crypter

import "errors"

var ErrDecodeKey = ErrDecodeKeyArgument{}

var ErrParseKey = ErrParseKeyArgument{}

var ErrInvalidToken = errors.New("invalid token")

type ErrDecodeKeyArgument struct {
	Type string
}

func (e ErrDecodeKeyArgument) Error() string {
	return "failed to decode " + e.Type + " key from base64"
}

func (e ErrDecodeKeyArgument) Is(target error) bool {
	_, ok := target.(ErrDecodeKeyArgument)
	return ok
}

type ErrParseKeyArgument struct {
	Type string
}

func (e ErrParseKeyArgument) Error() string {
	return "failed to parse " + e.Type + " key"
}

func (e ErrParseKeyArgument) Is(target error) bool {
	_, ok := target.(ErrParseKeyArgument)
	return ok
}
