package internal

import (
	"errors"
	"fmt"
)

// Kind is kind of error.
type Kind int

// Kinds of error.
const (
	KindBasic Kind = iota
	KindType
	KindMemory
	KindDatabase
	KindUnknown
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
	return e.Err.Error()
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
		diff += fmt.Sprintf("\nKind:\n  got : %d\n  want: %d", actual.Kind, expected.Kind)
	}
	if errors.Is(actual.Err, expected.Err) {
		diff += fmt.Sprintf("\nErr:\n  got : %s\n  want: %s", actual.Err, expected.Err)
	}
	return diff
}
