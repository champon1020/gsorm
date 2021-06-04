package database

import (
	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/morikuni/failure"
)

// mockTx is mock transaction.
type mockTx struct {
	// Parent mock database.
	db domain.MockDB

	// Expected statements.
	expected []expectation
}

// Ping is dummy function.
func (m *mockTx) Ping() error {
	return nil
}

// Exec is dummy function.
func (m *mockTx) Exec(string, ...interface{}) (domain.Result, error) {
	return nil, nil
}

// Query is dummy function.
func (m *mockTx) Query(string, ...interface{}) (domain.Rows, error) {
	return nil, nil
}

// Commit commits the transaction.
func (m *mockTx) Commit() error {
	expected := m.popExpected()
	if expected == nil {
		return failure.New(errInvalidMockExpectation, failure.Message("gsorm.mockTx.Commit is not expected"))
	}
	if _, ok := expected.(*expectedCommit); !ok {
		return failure.New(errInvalidMockExpectation, failure.Message("gsorm.mockTx.Commit is not expected"))
	}
	return nil
}

// Rollback aborts the transaction.
func (m *mockTx) Rollback() error {
	expected := m.popExpected()
	if expected == nil {
		return failure.New(errInvalidMockExpectation, failure.Message("gsorm.mockTx.Rollback is not expected"))
	}
	if _, ok := expected.(*expectedRollback); !ok {
		return failure.New(errInvalidMockExpectation, failure.Message("gsorm.mockTx.Rollback is not expected"))
	}
	return nil
}

// popExpected pops expected operation.
func (m *mockTx) popExpected() expectation {
	if len(m.expected) == 0 {
		return nil
	}
	op := m.expected[0]
	m.expected = m.expected[1:]
	return op
}

// ExpectCommit appends Commit operation to expected.
func (m *mockTx) ExpectCommit() {
	m.expected = append(m.expected, &expectedCommit{})
}

// ExpectRollback appends Rollback operation to expected.
func (m *mockTx) ExpectRollback() {
	m.expected = append(m.expected, &expectedRollback{})
}

// Expect appends expected statement.
func (m *mockTx) Expect(s domain.Stmt) {
	m.expected = append(m.expected, &expectedQuery{stmt: s})
}

// ExpectWithReturn appends expected statement with value which is to be returned with query.
func (m *mockTx) ExpectWithReturn(s domain.Stmt, v interface{}) {
	m.expected = append(m.expected, &expectedQuery{stmt: s, willReturn: v})
}

// Complete checks whether all of expected statements was executed or not.
func (m *mockTx) Complete() error {
	if len(m.expected) != 0 {
		return failure.New(errInvalidMockExpectation,
			failure.Context{"expected": m.expected[0].String(), "actual": "none"},
			failure.Message("invalid mock expectation"))
	}
	return nil
}

// CompareWith compares expected statement with executed statement.
func (m *mockTx) CompareWith(s domain.Stmt) (interface{}, error) {
	expected := m.popExpected()
	if expected == nil {
		err := failure.New(errInvalidMockExpectation,
			failure.Context{"expected": "none", "actual": s.FuncString()},
			failure.Message("invalid mock expectation"))
		return nil, err
	}
	eq, ok := expected.(*expectedQuery)
	if !ok {
		err := failure.New(errInvalidMockExpectation,
			failure.Context{"expected": expected.String(), "actual": s.FuncString()},
			failure.Message("invalid mock expectation"))
		return nil, err
	}
	if err := eq.stmt.CompareWith(s); err != nil {
		return nil, failure.Translate(err, errInvalidMockExpectation)
	}
	return eq.willReturn, nil
}
