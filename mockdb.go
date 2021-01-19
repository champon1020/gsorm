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
	expected [][]*opArgs
	actual   [][]*opArgs
}

// Query is the function for implementing DB.
func (m *MockDB) query(query string, args ...interface{}) (sqlRows, error) { return nil, nil }

// Exec is the function for implementing DB.
func (m *MockDB) exec(query string, args ...interface{}) (sql.Result, error) { return nil, nil }

func (m *MockDB) addExecuted(called []*opArgs) {
	m.actual = append(m.actual, called)
}

// AddExpected adds expected function calls.
func (m *MockDB) AddExpected(stmt *Stmt) {
	m.expected = append(m.expected, stmt.called)
}

// Result returns the difference between expected and actual queries that is executed.
func (m *MockDB) Result() error {
	return nil
}
