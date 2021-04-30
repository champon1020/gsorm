package database

import (
	"github.com/champon1020/mgorm/domain"
	"github.com/google/go-cmp/cmp"
	"github.com/morikuni/failure"
)

// compareStmts compares two statements. If their are different, returns error.
func compareStmts(expected domain.Stmt, actual domain.Stmt) error {
	expectedCalled := expected.Called()
	actualCalled := actual.Called()
	if len(expectedCalled) != len(actualCalled) {
		err := failure.New(errInvalidMockExpectation,
			failure.Context{"expected": expected.FuncString(), "actual": actual.FuncString()},
			failure.Message("invalid mock expectation"))
		return err
	}
	for i, e := range expectedCalled {
		if diff := cmp.Diff(actualCalled[i], e); diff != "" {
			err := failure.New(errInvalidMockExpectation,
				failure.Context{"expected": expected.FuncString(), "actual": actual.FuncString()},
				failure.Message("invalid mock expectation"))
			return err
		}
	}
	return nil
}

// expectation can be implemented by expected operation.
type expectation interface {
	String() string
}

// ExpectedQuery is expectation of executing query.
type ExpectedQuery struct {
	stmt       domain.Stmt
	willReturn interface{}
}

func (e *ExpectedQuery) String() string {
	return e.stmt.FuncString()
}

// expectedBegin is expectation of beginning transaction.
type expectedBegin struct{}

func (e *expectedBegin) String() string {
	return "mgorm.MockDB.Begin"
}

// expectedCommit is expectation of transaction commit.
type expectedCommit struct{}

func (e *expectedCommit) String() string {
	return "mgorm.MockTx.Commit"
}

// expectedRollback is expectation of transaction rollback.
type expectedRollback struct{}

func (e *expectedRollback) String() string {
	return "mgorm.MockTx.Rollback"
}
