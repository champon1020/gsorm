package mgorm

import (
	"database/sql"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
	"github.com/morikuni/failure"
)

// Mock is mock database conneciton pool.
type Mock interface {
	Conn
	Complete() error
	CompareWith(Stmt) (interface{}, error)
}

// MockDB is mock databse connection pool.
// This structure stores mainly what query will be executed and what value will be returned.
type MockDB struct {
	// Expected statements.
	expected []expectation

	// Begun transactions.
	tx []*MockTx

	// How many times transaction has begun.
	txItr int
}

func (m *MockDB) getDriver() internal.SQLDriver {
	return 0
}

// Ping is dummy function.
func (m *MockDB) Ping() error {
	return nil
}

// Exec is dummy function.
func (m *MockDB) Exec(string, ...interface{}) (sql.Result, error) {
	return nil, nil
}

// Query is dummy function.
func (m *MockDB) Query(string, ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

// SetConnMaxLifetime is dummy function.
func (m *MockDB) SetConnMaxLifetime(n time.Duration) error {
	return nil
}

// SetMaxIdleConns is dummy function.
func (m *MockDB) SetMaxIdleConns(n int) error {
	return nil
}

// SetMaxOpenConns is dummy function.
func (m *MockDB) SetMaxOpenConns(n int) error {
	return nil
}

// Close is dummy function.
func (m *MockDB) Close() error {
	return nil
}

// Begin starts the mock transaction.
func (m *MockDB) Begin() (*MockTx, error) {
	expected := m.popExpected()
	tx := m.nextTx()
	if tx == nil || expected == nil {
		err := failure.New(errInvalidMockExpectation,
			failure.Context{"expected": "none"},
			failure.Message("mgorm.MockDB.Begin is not expected"))
		return nil, err
	}
	_, ok := expected.(*expectedBegin)
	if !ok {
		err := failure.New(errInvalidMockExpectation,
			failure.Context{"expected": expected.String()},
			failure.Message("mgorm.MockDB.Begin is not expected"))
		return nil, err
	}
	return tx, nil
}

// nextTx pops begun transaction.
func (m *MockDB) nextTx() *MockTx {
	if len(m.tx) <= m.txItr {
		return nil
	}
	defer m.incrementTx()
	return m.tx[m.txItr]
}

// incrementTx increments txItr.
func (m *MockDB) incrementTx() {
	m.txItr++
}

// ExpectBegin appends operation of beginning transaction to expected.
func (m *MockDB) ExpectBegin() *MockTx {
	tx := &MockTx{db: m}
	m.tx = append(m.tx, tx)
	m.expected = append(m.expected, &expectedBegin{})
	return tx
}

// Expect appends expected statement.
func (m *MockDB) Expect(s Stmt) *MockDB {
	m.expected = append(m.expected, &expectedQuery{stmt: s})
	return m
}

// Return appends value which is to be returned with query.
func (m *MockDB) Return(v interface{}) {
	if e, ok := m.expected[len(m.expected)-1].(*expectedQuery); ok {
		e.willReturn = v
	}
}

// Complete checks whether all of expected statements was executed or not.
func (m *MockDB) Complete() error {
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
func (m *MockDB) CompareWith(s Stmt) (interface{}, error) {
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
	return eq.willReturn, compareStmts(eq.stmt, s)
}

// popExpected pops expected operation.
func (m *MockDB) popExpected() expectation {
	if len(m.expected) == 0 {
		return nil
	}
	op := m.expected[0]
	m.expected = m.expected[1:]
	return op
}

// MockTx is mock transaction.
type MockTx struct {
	// Parent mock database.
	db *MockDB

	// Expected statements.
	expected []expectation
}

func (m *MockTx) getDriver() internal.SQLDriver {
	return 0
}

// Ping is dummy function.
func (m *MockTx) Ping() error {
	return nil
}

// Exec is dummy function.
func (m *MockTx) Exec(string, ...interface{}) (sql.Result, error) {
	return nil, nil
}

// Query is dummy function.
func (m *MockTx) Query(string, ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

// Commit commits the transaction.
func (m *MockTx) Commit() error {
	expected := m.popExpected()
	if expected == nil {
		return failure.New(errInvalidMockExpectation, failure.Message("mgorm.MockTx.Commit is not expected"))
	}
	if _, ok := expected.(*expectedCommit); !ok {
		return failure.New(errInvalidMockExpectation, failure.Message("mgorm.MockTx.Commit is not expected"))
	}
	return nil
}

// Rollback aborts the transaction.
func (m *MockTx) Rollback() error {
	expected := m.popExpected()
	if expected == nil {
		return failure.New(errInvalidMockExpectation, failure.Message("mgorm.MockTx.Rollback is not expected"))
	}
	if _, ok := expected.(*expectedRollback); !ok {
		return failure.New(errInvalidMockExpectation, failure.Message("mgorm.MockTx.Rollback is not expected"))
	}
	return nil
}

// popExpected pops expected operation.
func (m *MockTx) popExpected() expectation {
	if len(m.expected) == 0 {
		return nil
	}
	op := m.expected[0]
	m.expected = m.expected[1:]
	return op
}

// ExpectCommit appends Commit operation to expected.
func (m *MockTx) ExpectCommit() {
	m.expected = append(m.expected, &expectedCommit{})
}

// ExpectRollback appends Rollback operation to expected.
func (m *MockTx) ExpectRollback() {
	m.expected = append(m.expected, &expectedRollback{})
}

// Expect appends expected statement.
func (m *MockTx) Expect(s Stmt) *MockTx {
	m.expected = append(m.expected, &expectedQuery{stmt: s})
	return m
}

// Return appends value which is to be returned with query.
func (m *MockTx) Return(v interface{}) {
	if e, ok := m.expected[len(m.expected)-1].(*expectedQuery); ok {
		e.willReturn = v
	}
}

// Complete checks whether all of expected statements was executed or not.
func (m *MockTx) Complete() error {
	if len(m.expected) != 0 {
		return failure.New(errInvalidMockExpectation,
			failure.Context{"expected": m.expected[0].String(), "actual": "none"},
			failure.Message("invalid mock expectation"))
	}
	return nil
}

// CompareWith compares expected statement with executed statement.
func (m *MockTx) CompareWith(s Stmt) (interface{}, error) {
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
	return eq.willReturn, compareStmts(eq.stmt, s)
}

// compareStmts compares two statements. If their are different, returns error.
func compareStmts(expected Stmt, actual Stmt) error {
	expectedCalled := expected.Called()
	actualCalled := actual.Called()
	if len(expectedCalled) != len(actualCalled) {
		err := failure.New(errInvalidMockExpectation,
			failure.Context{"expected": expected.FuncString(), "actual": actual.FuncString()},
			failure.Message("invalid mock expectation"))
		return err
	}
	for i, e := range expectedCalled {
		if diff := cmp.Diff(actualCalled[i], e); diff != "" {
			err := failure.New(errInvalidMockExpectation,
				failure.Context{"expected": expected.FuncString(), "actual": actual.FuncString()},
				failure.Message("invalid mock expectation"))
			return err
		}
	}
	return nil
}

// expectation can be implemented by expected operation.
type expectation interface {
	String() string
}

// expectedQuery is expectation of executing query.
type expectedQuery struct {
	stmt       Stmt
	willReturn interface{}
}

func (e *expectedQuery) String() string {
	return e.stmt.FuncString()
}

// expectedBegin is expectation of beginning transaction.
type expectedBegin struct{}

func (e *expectedBegin) String() string {
	return "mgorm.MockDB.Begin"
}

// expectedCommit is expectation of transaction commit.
type expectedCommit struct{}

func (e *expectedCommit) String() string {
	return "mgorm.MockTx.Commit"
}

// expectedRollback is expectation of transaction rollback.
type expectedRollback struct{}

func (e *expectedRollback) String() string {
	return "mgorm.MockTx.Rollback"
}
