package mgorm

import (
	"database/sql"

	"github.com/champon1020/mgorm/internal"
)

type opArgs struct {
	op   internal.Op
	args []interface{}
}

// MockDB is the mock databse object that implements DB.
type MockDB struct {
	expected []*Stmt
	actual   []*Stmt
}

// Query is the function for implementing DB.
func (m *MockDB) query(string, ...interface{}) (sqlRows, error) { return nil, nil }

// Exec is the function for implementing DB.
func (m *MockDB) exec(string, ...interface{}) (sql.Result, error) { return nil, nil }

func (m *MockDB) addExecuted(stmt *Stmt) {
	m.actual = append(m.actual, stmt)
}

// AddExpected adds expected function calls.
func (m *MockDB) AddExpected(stmt *Stmt) {
	m.expected = append(m.expected, stmt)
}

/*
// Result returns the difference between expected and actual queries that is executed.
func (m *MockDB) Result() error {
	i := 0
	for ; i < len(m.actual); i++ {
		if len(m.expected) <= i {
			return fmt.Errorf("%v was executed, but not expected", opArgsToQueryString(m.actual[i]))
		}

		j := 0
		for ; j < len(m.actual[i].called); j++ {
			if len(m.expected[i].called) <= j {
				return fmt.Errorf(
					"%v was executed, but %v is expected",
					opArgsToQueryString(m.actual[i]),
					opArgsToQueryString(m.expected[i]),
				)
			}

			if m.actual[i][j] != m.expected[i][j] {
				return fmt.Errorf(
					"%v was executed, but %v is expected",
					opArgsToQueryString(m.actual[i]),
					opArgsToQueryString(m.expected[i]),
				)
			}
		}

		if j < len(m.expected[i]) {
			return fmt.Errorf(
				"%v was executed, but %v is expected",
				opArgsToQueryString(m.actual[i]),
				opArgsToQueryString(m.expected[i]),
			)
		}
	}

	if i < len(m.expected) {
		return fmt.Errorf("no query was executed, but %v is expected", opArgsToQueryString(m.expected[i]))
	}

	return nil
}
*/
