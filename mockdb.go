package mgorm

import (
	"database/sql"
)

type mockRows struct{}

func (mr *mockRows) Close() error               { return nil }
func (mr *mockRows) Columns() ([]string, error) { return []string{}, nil }
func (mr *mockRows) Next() bool                 { return false }
func (mr *mockRows) Scan(...interface{}) error  { return nil }

// queryArgs store the pair of sql query and arguments.
type queryArgs struct {
	query string
	args  []interface{}
}

// MockDB is the mock databse object that implements DB.
type MockDB struct {
	Expected []*Stmt
	Actual   []queryArgs
}

// addQuery adds the pair of query and arguments to mock structure.
func (m *MockDB) addQuery(query string, args ...interface{}) {
	qa := queryArgs{query: query, args: args}
	m.Actual = append(m.Actual, qa)
}

// Query is the function for implementing DB.
func (m *MockDB) query(query string, args ...interface{}) (Rows, error) {
	m.addQuery(query, args...)
	rows := &mockRows{}
	return rows, nil
}

// Exec is the function for implementing DB.
func (m *MockDB) exec(query string, args ...interface{}) (sql.Result, error) {
	m.addQuery(query, args...)
	return nil, nil
}

// AddExpected adds expected function calls.
func (m *MockDB) AddExpected(stmt *Stmt) {
	m.Expected = append(m.Expected, stmt)
}

// Result returns the difference between expected and actual queries that is executed.
func (m *MockDB) Result() error {
	return nil
}
