package mgorm

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

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
func (m *MockDB) query(string, ...interface{}) (sqlRows, error) { return nil, nil }

// Exec is the function for implementing DB.
func (m *MockDB) exec(string, ...interface{}) (sql.Result, error) { return nil, nil }

func (m *MockDB) addExecuted(called []*opArgs) {
	m.actual = append(m.actual, called)
}

// AddExpected adds expected function calls.
func (m *MockDB) AddExpected(stmt *Stmt) {
	m.expected = append(m.expected, stmt.called)
}

// Result returns the difference between expected and actual queries that is executed.
func (m *MockDB) Result() error {
	i := 0
	for ; i < len(m.actual); i++ {
		if len(m.expected) <= i {
			return fmt.Errorf("%v was executed, but not expected", opArgsToQueryString(m.actual[i]))
		}

		j := 0
		for ; j < len(m.actual[i]); j++ {
			if len(m.expected[i]) <= j {
				return fmt.Errorf(
					"%v was executed, but %v is expected",
					opArgsToQueryString(m.actual[i]),
					opArgsToQueryString(m.expected[i]),
				)
			}

			if m.actual[i][j].op != m.expected[i][j].op {
				return fmt.Errorf(
					"%v was executed, but %v is expected",
					opArgsToQueryString(m.actual[i]),
					opArgsToQueryString(m.expected[i]),
				)
			}

			if !reflect.DeepEqual(m.actual[i][j].args, m.expected[i][j].args) {
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

func opArgsToQueryString(opArgs []*opArgs) (s string) {
	for _, oa := range opArgs {
		// Get function name.
		sep := strings.Split(string(oa.op), ".")
		funcName := sep[len(sep)-1]
		if s != "" && funcName != "" {
			s += "."
		}

		// Convert oa.args to string.
		var argsStr string
		for _, arg := range oa.args {
			if argsStr != "" {
				argsStr += ", "
			}
			switch arg := arg.(type) {
			case []string, string:
				argsStr += fmt.Sprintf("%+q", arg)
			default:
				argsStr += fmt.Sprintf("%v", arg)
			}
		}

		s += fmt.Sprintf("%s(%v)", funcName, argsStr)
	}
	return
}
