package errors

// errorCode is code of errors.
type errorCode int

// Error codes list.
const (
	UnknownError errorCode = iota
	InvalidTypeError
	InvalidFormatError
	InvalidValueError
	UnchangeableError
	DBColumnError
	DBQueryError
	DBScanError
	MockError
)
