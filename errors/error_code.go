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

func (e errorCode) String() string {
	switch e {
	case UnknownError:
		return "Unknown"
	case InvalidTypeError:
		return "Invalid Type"
	case InvalidFormatError:
		return "Invalid Format"
	case InvalidValueError:
		return "Invalid Value"
	case UnchangeableError:
		return "Unchangeable"
	case DBColumnError, DBQueryError, DBScanError:
		return "Database Error"
	case MockError:
		return "Mock Error"
	}
	return ""
}
