package mgorm

import (
	"database/sql"

	"github.com/champon1020/mgorm/internal"
)

// queryArgs store the pair of sql query and arguments.
type opArgs struct {
	op   internal.Op
	args []interface{}
}

// MockDB is the mock databse object that implements DB.
type MockDB struct {
	Expected [][]*opArgs
	Actual   [][]*opArgs
}

// Query is the function for implementing DB.
func (m *MockDB) query(query string, args ...interface{}) (Rows, error) { return nil, nil }

// Exec is the function for implementing DB.
func (m *MockDB) exec(query string, args ...interface{}) (sql.Result, error) { return nil, nil }

func (m *MockDB) addExecuted(called []*opArgs) {
	m.Actual = append(m.Actual, called)
}

// AddExpected adds expected function calls.
func (m *MockDB) AddExpected(stmt *Stmt) {
	m.Expected = append(m.Expected, stmt.called)
}

// Result returns the difference between expected and actual queries that is executed.
func (m *MockDB) Result() error {
	return nil
}
