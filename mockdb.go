package mgorm

import (
	"fmt"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/google/go-cmp/cmp"
)

const (
	opCompareTo      = "mgorm.MockDB.compareTo"
	opComplete       = "mgorm.MockDB.Complete"
	opMockDBBegin    = "mgorm.MockDB.Begin"
	opMockTxCommit   = "mgorm.MockTx.Commit"
	opMockTxRollback = "mgorm.MockTx.Rollback"
)

// MockDB is the mock databse object.
// This structure stores mainly what query will be executed and what value will be returned.
type MockDB struct {
	// Expected statements.
	expected []expectation

	// Begun transactions.
	tx []*MockTx

	// How many times transaction has begun.
	txItr int
}

// Ping is dummy function.
func (m *MockDB) Ping() error {
	return nil
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
	tx := m.nextTx()
	if tx == nil {
		err := fmt.Errorf("%s was executed but not expected", opMockDBBegin)
		return nil, internal.NewError(opMockDBBegin, internal.KindRuntime, err)
	}
	return tx, nil
}

// Expect appends expected statement.
func (m *MockDB) Expect(s *Stmt) *MockDB {
	m.expected = append(m.expected, &expectedQuery{stmt: s})
	return m
}

// ExpectBegin appends operation of beginning transaction to expected.
func (m *MockDB) ExpectBegin() *MockTx {
	tx := &MockTx{db: m}
	m.tx = append(m.tx, tx)
	return tx
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
		err := fmt.Errorf("No query was executed, but %s is expected", m.expected[0].String())
		return internal.NewError(opComplete, internal.KindRuntime, err)
	}
	for _, tx := range m.tx {
		if err := tx.Complete(); err != nil {
			return err
		}
	}
	return nil
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

// nextTx pops begun transaction.
func (m *MockDB) nextTx() *MockTx {
	if len(m.tx) <= m.txItr {
		return nil
	}
	defer func() { m.txItr++ }()
	return m.tx[m.txItr]
}

// MockTx is mock transaction.
type MockTx struct {
	// Parent mock database.
	db *MockDB

	// Expected statements.
	expected []expectation
}

// Ping is dummy function.
func (m *MockTx) Ping() error {
	return nil
}

// Commit commits the transaction.
func (m *MockTx) Commit() error {
	expected := m.popExpected()
	if expected == nil {
		err := fmt.Errorf("%s was executed but not expected", opMockTxCommit)
		return internal.NewError(opMockTxCommit, internal.KindRuntime, err)
	}
	if _, ok := expected.(*expectedCommit); !ok {
		err := fmt.Errorf("%s was executed but not expected", opMockTxCommit)
		return internal.NewError(opMockTxCommit, internal.KindRuntime, err)
	}
	return nil
}

// Rollback aborts the transaction.
func (m *MockTx) Rollback() error {
	expected := m.popExpected()
	if expected == nil {
		err := fmt.Errorf("%s was executed but not expected", opMockTxRollback)
		return internal.NewError(opMockTxRollback, internal.KindRuntime, err)
	}
	if _, ok := expected.(*expectedRollback); !ok {
		err := fmt.Errorf("%s was executed but not expected", opMockTxRollback)
		return internal.NewError(opMockTxCommit, internal.KindRuntime, err)
	}
	return nil
}

// Expect appends expected statement.
func (m *MockTx) Expect(s *Stmt) *MockTx {
	m.expected = append(m.expected, &expectedQuery{stmt: s})
	return m
}

// ExpectCommit appends Commit operation to expected.
func (m *MockTx) ExpectCommit() {
	m.expected = append(m.expected, &expectedCommit{})
}

// ExpectRollback appends Rollback operation to expected.
func (m *MockTx) ExpectRollback() {
	m.expected = append(m.expected, &expectedRollback{})
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
		err := fmt.Errorf("No query was executed, but %s is expected", m.expected[0].String())
		return internal.NewError(opComplete, internal.KindRuntime, err)
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

// compareTo compares executed statement with expected statement.
// This function is called when some query was executed.
func compareTo(mock Mock, stmt *Stmt) (interface{}, error) {
	expected := mock.popExpected()
	if expected == nil {
		err := fmt.Errorf("%s was executed but not expected", stmt.funcString())
		return nil, internal.NewError(opCompareTo, internal.KindRuntime, err)
	}

	e, ok := expected.(*expectedQuery)
	if !ok {
		err := fmt.Errorf("%s was executed but not expected", expected.String())
		return nil, internal.NewError(opCompareTo, internal.KindRuntime, err)
	}

	if len(stmt.called) != len(e.stmt.called) {
		err := fmt.Errorf("%s was executed, but %s is expected", stmt.funcString(), e.stmt.funcString())
		return nil, internal.NewError(opCompareTo, internal.KindRuntime, err)
	}

	for i, expr := range e.stmt.called {
		if diff := cmp.Diff(stmt.called[i], expr); diff != "" {
			err := fmt.Errorf("%s was executed, but %s is expected", stmt.funcString(), e.stmt.funcString())
			return nil, internal.NewError(opCompareTo, internal.KindRuntime, err)
		}
	}

	if e.willReturn != nil {
		return e.willReturn, nil
	}

	return nil, nil
}

// expectation can be implemented by expected operation.
type expectation interface {
	String() string
	completed() bool
}

// expectedOp is common expectation.
type expectedOp struct {
	done bool
}

func (e *expectedOp) completed() bool {
	return e.done
}

// expectedQuery is expectation of executing query.
type expectedQuery struct {
	stmt       *Stmt
	willReturn interface{}
	expectedOp
}

func (e *expectedQuery) String() string {
	return e.stmt.funcString()
}

// expectedCommit is expectation of transaction commit.
type expectedCommit struct {
	expectedOp
}

func (e *expectedCommit) String() string {
	return "mgorm.(*MockTx).Commit()"
}

// expectedRollback is expectation of transaction rollback.
type expectedRollback struct {
	expectedOp
}

func (e *expectedRollback) String() string {
	return "mgorm.(*MockTx).Rollback()"
}
