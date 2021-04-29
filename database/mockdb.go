package database

import (
	"database/sql"
	"time"

	"github.com/champon1020/mgorm/interfaces"
	"github.com/champon1020/mgorm/internal"
	"github.com/morikuni/failure"
)

type MockDB interface {
	Mock
	Ping() error
	SetConnMaxLifetime(n time.Duration) error
	SetMaxIdleConns(n int) error
	SetMaxOpenConns(n int) error
	Close() error
	Begin() (MockTx, error)
	ExpectBegin() MockTx
	Expect(stmt interfaces.Stmt) MockDB
	Return(v interface{})
}

// mockDB is mock databse connection pool.
// This structure stores mainly what query will be executed and what value will be returned.
type mockDB struct {
	// Expected statements.
	expected []expectation

	// Begun transactions.
	tx []MockTx

	// How many times transaction has begun.
	txItr int
}

// NewMockDB creates MockDB instance.
func NewMockDB() MockDB {
	return &mockDB{}
}

// GetDriver returns sql driver.
func (m *mockDB) GetDriver() internal.SQLDriver {
	return 0
}

// Ping is dummy function.
func (m *mockDB) Ping() error {
	return nil
}

// Exec is dummy function.
func (m *mockDB) Exec(string, ...interface{}) (sql.Result, error) {
	return nil, nil
}

// Query is dummy function.
func (m *mockDB) Query(string, ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

// SetConnMaxLifetime is dummy function.
func (m *mockDB) SetConnMaxLifetime(n time.Duration) error {
	return nil
}

// SetMaxIdleConns is dummy function.
func (m *mockDB) SetMaxIdleConns(n int) error {
	return nil
}

// SetMaxOpenConns is dummy function.
func (m *mockDB) SetMaxOpenConns(n int) error {
	return nil
}

// Close is dummy function.
func (m *mockDB) Close() error {
	return nil
}

// Begin starts the mock transaction.
func (m *mockDB) Begin() (MockTx, error) {
	expected := m.popExpected()
	tx := m.nextTx()
	if tx == nil || expected == nil {
		err := failure.New(errInvalidMockExpectation,
			failure.Context{"expected": "none"},
			failure.Message("mgorm.mockDB.Begin is not expected"))
		return nil, err
	}
	_, ok := expected.(*expectedBegin)
	if !ok {
		err := failure.New(errInvalidMockExpectation,
			failure.Context{"expected": expected.String()},
			failure.Message("mgorm.mockDB.Begin is not expected"))
		return nil, err
	}
	return tx, nil
}

// nextTx pops begun transaction.
func (m *mockDB) nextTx() MockTx {
	if len(m.tx) <= m.txItr {
		return nil
	}
	defer m.incrementTx()
	return m.tx[m.txItr]
}

// incrementTx increments txItr.
func (m *mockDB) incrementTx() {
	m.txItr++
}

// ExpectBegin appends operation of beginning transaction to expected.
func (m *mockDB) ExpectBegin() MockTx {
	tx := &mockTx{db: m}
	m.tx = append(m.tx, tx)
	m.expected = append(m.expected, &expectedBegin{})
	return tx
}

// Expect appends expected statement.
func (m *mockDB) Expect(s interfaces.Stmt) MockDB {
	m.expected = append(m.expected, &ExpectedQuery{stmt: s})
	return m
}

// Return appends value which is to be returned with query.
func (m *mockDB) Return(v interface{}) {
	if e, ok := m.expected[len(m.expected)-1].(*ExpectedQuery); ok {
		e.willReturn = v
	}
}

// Complete checks whether all of expected statements was executed or not.
func (m *mockDB) Complete() error {
	if len(m.expected) != 0 {
		return failure.New(errInvalidMockExpectation,
			failure.Context{"expected": m.expected[0].String(), "actual": "none"},
			failure.Message("invalid mock expectation"))
	}
	for _, tx := range m.tx {
		if err := tx.Complete(); err != nil {
			return err
		}
	}
	return nil
}

// CompareWith compares expected statement with executed statement.
func (m *mockDB) CompareWith(s interfaces.Stmt) (interface{}, error) {
	expected := m.popExpected()
	if expected == nil {
		err := failure.New(errInvalidMockExpectation,
			failure.Context{"expected": "none", "actual": s.FuncString()},
			failure.Message("invalid mock expectation"))
		return nil, err
	}
	eq, ok := expected.(*ExpectedQuery)
	if !ok {
		err := failure.New(errInvalidMockExpectation,
			failure.Context{"expected": expected.String(), "actual": s.FuncString()},
			failure.Message("invalid mock expectation"))
		return nil, err
	}
	return eq.willReturn, compareStmts(eq.stmt, s)
}

// popExpected pops expected operation.
func (m *mockDB) popExpected() expectation {
	if len(m.expected) == 0 {
		return nil
	}
	op := m.expected[0]
	m.expected = m.expected[1:]
	return op
}
