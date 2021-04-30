package database

import "github.com/morikuni/failure"

const (
	errFailedDBConnection     failure.StringCode = "FailedDBConnection"
	errFailedTxConnection     failure.StringCode = "FailedTxConnection"
	errInvalidMockExpectation failure.StringCode = "InvalidMockExpectation"
)
