package syntax

// Error code list.
const (
	ErrInvalid = iota
	ErrInvalidType
	ErrInvalidLen
	ErrQueryFailed
	ErrExecFailed
	ErrScanFailed
	ErrUnknown
)

// Error struct.
type Error struct {
	Code int
	Msg  string
}

func (e Error) Error() string {
	return e.Msg
}

func newError(code int, msg string) error {
	return Error{Code: code, Msg: msg}
}
