package mgorm

import (
	"fmt"

	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
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

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
// This is dummy function for implementing mgorm.sqlDB.
func (m *MockDB) Ping() error {
	return nil
}

func (m *MockDB) addExecuted(stmt *Stmt) {
	m.actual = append(m.actual, stmt)
}

// AddExpected adds expected function calls.
func (m *MockDB) AddExpected(stmt *Stmt) {
	m.expected = append(m.expected, stmt)
}

// Result returns the difference between expected and actual queries that is executed.
func (m *MockDB) Result() error {
	i := 0
	for ; i < len(m.actual); i++ {
		if len(m.expected) <= i {
			return fmt.Errorf("%v was executed, but not expected", getFunctionString(m.actual[i]))
		}

		j := 0
		for ; j < len(m.actual[i].called); j++ {
			if len(m.expected[i].called) <= j {
				return fmt.Errorf(
					"%v was executed, but %v is expected",
					getFunctionString(m.actual[i]),
					getFunctionString(m.expected[i]),
				)
			}

			if diff := cmp.Diff(m.actual[i].called[j], m.expected[i].called[j]); diff != "" {
				return fmt.Errorf(
					"%v was executed, but %v is expected",
					getFunctionString(m.actual[i]),
					getFunctionString(m.expected[i]),
				)
			}
		}

		if j < len(m.expected[i].called) {
			return fmt.Errorf(
				"%v was executed, but %v is expected",
				getFunctionString(m.actual[i]),
				getFunctionString(m.expected[i]),
			)
		}
	}

	if i < len(m.expected) {
		return fmt.Errorf("no query was executed, but %v is expected", getFunctionString(m.expected[i]))
	}

	return nil
}

func getFunctionString(stmt *Stmt) string {
	s := stmt.cmd.String()
	for _, e := range stmt.called {
		s += fmt.Sprintf(".%s", e.String())
	}
	return s
}
