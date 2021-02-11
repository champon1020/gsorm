package errors

import (
	"fmt"
)

// Error is custom error structure.
type Error struct {
	Msg  string
	Code errorCode
	Line int
}

func (e Error) Error() string {
	if e.Line != 0 {
		return fmt.Sprintf("[%s] Line=%d | %v", e.Code.String(), e.Line, e.Msg)
	}
	return fmt.Sprintf("[%s] %v", e.Code.String(), e.Msg)
}

// Is validates whether err is same error value or not.
func (e *Error) Is(_err error) bool {
	err, ok := _err.(*Error)
	if !ok {
		return false
	}
	if e.Msg != err.Msg || e.Code != err.Code || e.Line != err.Line {
		return false
	}
	return true
}

// New creates error object.
func New(msg string, code errorCode) error {
	return &Error{Msg: msg}
}

// NewWith creates error object with options.
func NewWith(msg string, code errorCode, line int) error {
	return &Error{Msg: msg, Code: code, Line: line}
}
