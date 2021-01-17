package integration

import (
	"database/sql"
)

type mockRows struct{}

func (mr *mockRows) Close() error               { return nil }
func (mr *mockRows) Columns() ([]string, error) { return []string{}, nil }
func (mr *mockRows) Next() bool                 { return false }
func (mr *mockRows) Scan(...interface{}) error  { return nil }

type QueryArgs struct {
	query string
	args  []interface{}
}

// MockDB is the mock databse object that implements DbIface.
type MockDB struct {
	Expected []QueryArgs
	Actual   []QueryArgs
}

func (m *MockDB) addQuery(query string, args ...interface{}) {
	qa := QueryArgs{query: query, args: args}
	m.Actual = append(m.Actual, qa)
}

// Query is the function for implementing DbIface.
func (m *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	m.addQuery(query, args...)
	return nil, nil
}

// Exec is the function for implementing DbIface.
func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.addQuery(query, args...)
	return nil, nil
}
