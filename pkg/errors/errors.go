package errors

import (
	"errors"
	"fmt"
)

type myError struct {
	kind      Kind
	message   string
	prevError error
}

type Kind string

const (
	BadRequest           Kind = "builtin.BadRequest"
	Internal             Kind = "builtin.Internal"
	ResourceNotFound     Kind = "builtin.ResourceNotFound"
	ResourceAlreadyExist Kind = "builtin.ResourceAlreadyExist"
)

func (k Kind) String() string {
	return string(k)
}

func New(kind Kind, message string) error {
	return &myError{kind: kind, message: message}
}

func Wrap(kind Kind, message string, err error) error {
	return &myError{kind: kind, message: message, prevError: err}
}

func (e *myError) Error() string {
	if e.prevError != nil {
		return fmt.Sprintf("%s: %s", e.message, e.prevError.Error())
	}
	return e.message
}

func (e *myError) Kind() Kind {
	return e.kind
}

func KindIs(err error, kind Kind) bool {
	e, ok := err.(*myError)
	if !ok {
		return false
	}
	return e.Kind() == kind
}

// Is is a wrapper of errors.Is.
func Is(err error, target error) bool {
	return errors.Is(err, target)
}
