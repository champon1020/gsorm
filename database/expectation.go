package database

import "github.com/champon1020/mgorm/interfaces/domain"

// expectation can be implemented by expected operation.
type expectation interface {
	String() string
}

// expectedQuery is expectation of executing query.
type expectedQuery struct {
	stmt       domain.Stmt
	willReturn interface{}
}

func (e *expectedQuery) String() string {
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
