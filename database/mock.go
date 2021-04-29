package database

import (
	"github.com/champon1020/mgorm/interfaces"
	"github.com/google/go-cmp/cmp"
	"github.com/morikuni/failure"
)

// Mock is mock database conneciton pool.
type Mock interface {
	Conn
	Complete() error
	CompareWith(interfaces.Stmt) (interface{}, error)
}

// compareStmts compares two statements. If their are different, returns error.
func compareStmts(expected interfaces.Stmt, actual interfaces.Stmt) error {
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
	stmt       interfaces.Stmt
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
