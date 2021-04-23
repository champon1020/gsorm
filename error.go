package mgorm

import "github.com/morikuni/failure"

const (
	errInvalidValue           failure.StringCode = "InvalidValue"
	errInvalidClause          failure.StringCode = "InvalidClause"
	errInvalidSyntax          failure.StringCode = "InvalidSyntax"
	errInvalidType            failure.StringCode = "InvalidType"
	errFailedParse            failure.StringCode = "FailedParse"
	errInvalidMockExpectation failure.StringCode = "InvalidMockExpectation"
	errFailedDBConnection     failure.StringCode = "FailedDBConnection"
	errFailedTxConnection     failure.StringCode = "FailedTxConnection"
)
