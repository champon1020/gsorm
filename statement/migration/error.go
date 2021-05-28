package migration

import "github.com/morikuni/failure"

const (
	errInvalidValue  failure.StringCode = "InvalidValue"
	errInvalidClause failure.StringCode = "InvalidClause"
	errInvalidSyntax failure.StringCode = "InvalidSyntax"
	errFailedParse   failure.StringCode = "FailedParse"
)
