package statement

import "github.com/morikuni/failure"

const (
	errInvalidValue  failure.StringCode = "InvalidValue"
	errInvalidClause failure.StringCode = "InvalidClause"
	errInvalidSyntax failure.StringCode = "InvalidSyntax"
	errInvalidType   failure.StringCode = "InvalidType"
	errFailedParse   failure.StringCode = "FailedParse"
)
