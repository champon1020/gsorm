package mgorm

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
)

const (
	opCompareTo = "mgorm.MockDB.compareTo"
	opComplete  = "mgorm.MockDB.Complete"
)

// MockDB is the mock databse object.
// This structure stores what query will be executed and what value will be returned.
type MockDB struct {
	// Store expected statements.
	expected []*Stmt

	// Store data which is to be returned with query execution.
	willReturn map[int]interface{}

	// How many times query was executed.
	execCnt int
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
// This is dummy function for implementing mgorm.sqlDB.
func (m *MockDB) Ping() error {
	return nil
}

// Expect adds expected statement to be executed.
func (m *MockDB) Expect(stmt *Stmt) *MockDB {
	m.expected = append(m.expected, stmt)
	return m
}

// Return adds value which is to be returned with query.
func (m *MockDB) Return(v interface{}) {
	if len(m.expected) == 0 {
		return
	}
	idx := len(m.expected) - 1
	m.willReturn[idx] = v
}

// compareTo compares executed statement with expected statement.
// This function is called when some query was executed.
func (m *MockDB) compareTo(stmt *Stmt) (interface{}, error) {
	if len(m.expected) <= m.execCnt {
		err := fmt.Errorf("%s was executed, but not expected", stmt.funcString())
		return nil, internal.NewError(opCompareTo, internal.KindRuntime, err)
	}

	expStmt := m.expected[m.execCnt]
	if len(stmt.called) != len(expStmt.called) {
		err := fmt.Errorf("%s was executed, but %s is expected", stmt.funcString(), expStmt.funcString())
		return nil, internal.NewError(opCompareTo, internal.KindRuntime, err)
	}

	for i, e := range expStmt.called {
		if diff := cmp.Diff(stmt.called[i], e); diff != "" {
			err := fmt.Errorf(
				"%s was executed, but %s is expected", stmt.funcString(), expStmt.funcString())
			return nil, internal.NewError(opCompareTo, internal.KindRuntime, err)
		}
	}

	defer func() { m.execCnt++ }()

	if v, ok := m.willReturn[m.execCnt]; ok {
		return v, nil
	}
	return nil, nil
}

// Complete checks whether all of expected statements was executed or not.
func (m *MockDB) Complete() error {
	if len(m.expected) != 0 {
		err := fmt.Errorf("no query has left, but %s is expected", m.expected[0].funcString())
		return internal.NewError(opComplete, internal.KindRuntime, err)
	}
	return nil
}
