package internal

import (
	"fmt"
)

// Kind is kind of error.
type Kind string

// Kinds of error.
const (
	KindBasic    Kind = "BasicError"
	KindType     Kind = "TypeError"
	KindMemory   Kind = "MemoryError"
	KindDatabase Kind = "DatabaseError"
	KindUnknown  Kind = "UnknownError"
)

// Op stores location information where error is occurred.
type Op string

// Error is custom error structure.
type Error struct {
	Op   Op
	Kind Kind
	Err  error
}

func (e Error) Error() string {
	msg := fmt.Sprintf("[%s] %s: %v", e.Kind, e.Op, e.Err)
	return msg
}

// NewError generates error object.
func NewError(op Op, kind Kind, err error) error {
	return &Error{Op: op, Kind: kind, Err: err}
}

// CmpError compare the twe errors.
func CmpError(actual Error, expected Error) string {
	var diff string
	if actual.Op != expected.Op {
		diff += fmt.Sprintf("\nOp:\n  got : %s\n  want: %s", actual.Op, expected.Op)
	}
	if actual.Kind != expected.Kind {
		diff += fmt.Sprintf("\nKind:\n  got : %s\n  want: %s", actual.Kind, expected.Kind)
	}
	if actual.Err.Error() != expected.Err.Error() {
		diff += fmt.Sprintf("\nErr:\n  got : %s\n  want: %s", actual.Err.Error(), expected.Err.Error())
	}
	return diff
}
