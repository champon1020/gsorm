package gsorm

import (
	"time"

	"github.com/champon1020/gsorm/interfaces"
	"golang.org/x/xerrors"
)

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

// Ping is dummy function.
func (m *mockDB) Ping() error {
	return nil
}

// Exec is dummy function.
func (m *mockDB) Exec(string, ...interface{}) (iresult, error) {
	return nil, nil
}

// Query is dummy function.
func (m *mockDB) Query(string, ...interface{}) (irows, error) {
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
func (m *mockDB) Begin() (Tx, error) {
	expected := m.popExpected()
	tx := m.nextTx()
	if tx == nil || expected == nil {
		return nil, xerrors.New("gsorm.mockDB.Begin is not expected")
	}
	_, ok := expected.(*expectedBegin)
	if !ok {
		return nil, xerrors.New("gsorm.mockDB.Begin is not expected")
	}
	return tx, nil
}

// nextTx pops begun transaction.
func (m *mockDB) nextTx() Tx {
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
func (m *mockDB) Expect(s interfaces.Stmt) {
	m.expected = append(m.expected, &expectedQuery{stmt: s})
}

// ExpectWithReturn appends expected statement with value which is to be returned with query.
func (m *mockDB) ExpectWithReturn(s interfaces.Stmt, v interface{}) {
	m.expected = append(m.expected, &expectedQuery{stmt: s, willReturn: v})
}

// Complete checks whether all of expected statements was executed or not.
func (m *mockDB) Complete() error {
	if len(m.expected) != 0 {
		return xerrors.Errorf("%s is expected but not executed", m.expected[0].String())
	}
	for _, tx := range m.tx {
		if err := tx.Complete(); err != nil {
			return err
		}
	}
	return nil
}

// compareWith compares expected statement with executed statement.
func (m *mockDB) compareWith(s interfaces.Stmt) (interface{}, error) {
	expected := m.popExpected()
	if expected == nil {
		return nil, xerrors.Errorf("%s is not expected but executed", s.String())
	}
	eq, ok := expected.(*expectedQuery)
	if !ok {
		return nil, xerrors.Errorf("statements comparison was failed:\nexpected: %s\nactual:   %s\n",
			expected.String(), s.String())
	}
	if err := eq.stmt.CompareWith(s); err != nil {
		return nil, err
	}
	return eq.willReturn, nil
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

// mockTx is mock transaction.
type mockTx struct {
	// Parent mock database.
	db MockDB

	// Expected statements.
	expected []expectation
}

// Ping is dummy function.
func (m *mockTx) Ping() error {
	return nil
}

// Exec is dummy function.
func (m *mockTx) Exec(string, ...interface{}) (iresult, error) {
	return nil, nil
}

// Query is dummy function.
func (m *mockTx) Query(string, ...interface{}) (irows, error) {
	return nil, nil
}

// Commit commits the transaction.
func (m *mockTx) Commit() error {
	expected := m.popExpected()
	if expected == nil {
		return xerrors.New("gsorm.mockTx.Commit is not expected")
	}
	if _, ok := expected.(*expectedCommit); !ok {
		return xerrors.New("gsorm.mockTx.Commit is not expected")
	}
	return nil
}

// Rollback aborts the transaction.
func (m *mockTx) Rollback() error {
	expected := m.popExpected()
	if expected == nil {
		return xerrors.New("gsorm.mockTx.Rollback is not expected")
	}
	if _, ok := expected.(*expectedRollback); !ok {
		return xerrors.New("gsorm.mockTx.Rollback is not expected")
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
func (m *mockTx) Expect(s interfaces.Stmt) {
	m.expected = append(m.expected, &expectedQuery{stmt: s})
}

// ExpectWithReturn appends expected statement with value which is to be returned with query.
func (m *mockTx) ExpectWithReturn(s interfaces.Stmt, v interface{}) {
	m.expected = append(m.expected, &expectedQuery{stmt: s, willReturn: v})
}

// Complete checks whether all of expected statements was executed or not.
func (m *mockTx) Complete() error {
	if len(m.expected) != 0 {
		return xerrors.Errorf("%s is expected but not executed", m.expected[0].String())
	}
	return nil
}

// compareWith compares expected statement with executed statement.
func (m *mockTx) compareWith(s interfaces.Stmt) (interface{}, error) {
	expected := m.popExpected()
	if expected == nil {
		return nil, xerrors.Errorf("%s is not expected but executed", s.String())
	}
	eq, ok := expected.(*expectedQuery)
	if !ok {
		return nil, xerrors.Errorf("statements comparison was failed:\nexpected: %s\nactual:   %s\n",
			expected.String(), s.String())
	}
	if err := eq.stmt.CompareWith(s); err != nil {
		return nil, err
	}
	return eq.willReturn, nil
}
