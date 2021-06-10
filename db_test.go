package gsorm_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/stretchr/testify/assert"
)

func TestDB_Ping(t *testing.T) {
	// Prepare for test.
	db := gsorm.ExportedDB{}
	sdb := &SpyDB{}
	db.ExportedSetConn(sdb)

	// Actual process.
	if err := db.Ping(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, sdb.calledPing)
}

func TestDB_Ping_Fail(t *testing.T) {
	expectedErr := "gsorm.db.conn is nil"

	// Prepare for test.
	db := gsorm.ExportedDB{}

	// Actual process.
	err := db.Ping()

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestDB_Exec(t *testing.T) {
	// Prepare for test.
	db := gsorm.ExportedDB{}
	sdb := &SpyDB{}
	db.ExportedSetConn(sdb)

	// Actual process.
	if _, err := db.Exec(""); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, sdb.calledExec)
}

func TestDB_Exec_Fail(t *testing.T) {
	expectedErr := "gsorm.db.conn is nil"

	// Prepare for test.
	db := gsorm.ExportedDB{}

	// Actual process.
	_, err := db.Exec("")

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestDB_Query(t *testing.T) {
	// Prepare for test.
	db := gsorm.ExportedDB{}
	sdb := &SpyDB{}
	db.ExportedSetConn(sdb)

	// Actual process.
	if _, err := db.Query(""); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, sdb.calledQuery)
}

func TestDB_Query_Fail(t *testing.T) {
	expectedErr := "gsorm.db.conn is nil"

	// Prepare for test.
	db := gsorm.ExportedDB{}

	// Actual process.
	_, err := db.Query("")

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestDB_SetConnMaxLifetime(t *testing.T) {
	// Prepare for test.
	db := gsorm.ExportedDB{}
	sdb := &SpyDB{}
	db.ExportedSetConn(sdb)

	// Actual process.
	if err := db.SetConnMaxLifetime(0); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, sdb.calledSetConnMaxLifetime)
}

func TestDB_SetConnMaxLifetime_Fail(t *testing.T) {
	expectedErr := "gsorm.db.conn is nil"

	// Prepare for test.
	db := gsorm.ExportedDB{}

	// Actual process.
	err := db.SetConnMaxLifetime(0)

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestDB_SetMaxIdleConns(t *testing.T) {
	// Prepare for test.
	db := gsorm.ExportedDB{}
	sdb := &SpyDB{}
	db.ExportedSetConn(sdb)

	// Actual process.
	if err := db.SetMaxIdleConns(0); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, sdb.calledSetMaxIdleConns)
}

func TestDB_SetMaxIdleConns_Fail(t *testing.T) {
	expectedErr := "gsorm.db.conn is nil"

	// Prepare for test.
	db := gsorm.ExportedDB{}

	// Actual process.
	err := db.SetMaxIdleConns(0)

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestDB_SetMaxOpenConns(t *testing.T) {
	// Prepare for test.
	db := gsorm.ExportedDB{}
	sdb := &SpyDB{}
	db.ExportedSetConn(sdb)

	// Actual process.
	if err := db.SetMaxOpenConns(0); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, sdb.calledSetMaxOpenConns)
}

func TestDB_SetMaxOpenConns_Fail(t *testing.T) {
	expectedErr := "gsorm.db.conn is nil"

	// Prepare for test.
	db := gsorm.ExportedDB{}

	// Actual process.
	err := db.SetMaxOpenConns(0)

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestDB_Close(t *testing.T) {
	// Prepare for test.
	db := gsorm.ExportedDB{}
	sdb := &SpyDB{}
	db.ExportedSetConn(sdb)

	// Actual process.
	if err := db.Close(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, sdb.calledClose)
}

func TestDB_Close_Fail(t *testing.T) {
	expectedErr := "gsorm.db.conn is nil"

	// Prepare for test.
	db := gsorm.ExportedDB{}

	// Actual process.
	err := db.Close()

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestDB_Begin(t *testing.T) {
	// Prepare for test.
	db := gsorm.ExportedDB{}
	sdb := &SpyDB{}
	db.ExportedSetConn(sdb)

	// Actual process.
	if _, err := db.Begin(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, sdb.calledBegin)
}

func TestDB_Begin_Fail(t *testing.T) {
	expectedErr := "gsorm.db.conn is nil"

	// Prepare for test.
	db := gsorm.ExportedDB{}

	// Actual process.
	_, err := db.Begin()

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestTx_Ping(t *testing.T) {
	db := &gsorm.ExportedDB{}
	sdb := &SpyDB{}
	db.ExportedSetConn(sdb)

	tx := &gsorm.ExportedTx{}
	tx.ExportedSetDB(db)

	err := tx.Ping()
	if err != nil {
		t.Errorf("error was occured: %v", err)
	}

	assert.Equal(t, true, sdb.calledPing)
}

func TestTx_Ping_Fail(t *testing.T) {
	expectedErr := "gsorm.tx.conn is nil"

	// Prepare for test.
	tx := &gsorm.ExportedTx{}

	// Actual process.
	err := tx.Ping()

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestTx_Exec(t *testing.T) {
	// Prepare for test.
	tx := &gsorm.ExportedTx{}
	stx := &SpyTx{}
	tx.ExportedSetConn(stx)

	// Actual process.
	if _, err := tx.Exec(""); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, stx.calledExec)
}

func TestTx_Exec_Fail(t *testing.T) {
	expectedErr := "gsorm.tx.conn is nil"

	// Prepare for test.
	tx := &gsorm.ExportedTx{}

	// Actual process.
	_, err := tx.Exec("")

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestTx_Query(t *testing.T) {
	// Preparep for test.
	tx := &gsorm.ExportedTx{}
	stx := &SpyTx{}
	tx.ExportedSetConn(stx)

	// Actual process.
	if _, err := tx.Query(""); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, stx.calledQuery)
}

func TestTx_Query_Fail(t *testing.T) {
	expectedErr := "gsorm.tx.conn is nil"

	// Prepare for test.
	tx := &gsorm.ExportedTx{}

	// Actual process.
	_, err := tx.Query("")

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestTx_Commit(t *testing.T) {
	// Prepare for test.
	tx := &gsorm.ExportedTx{}
	stx := &SpyTx{}
	tx.ExportedSetConn(stx)

	// Actual process.
	if err := tx.Commit(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, stx.calledCommit)
}

func TestTx_Commit_Fail(t *testing.T) {
	expectedErr := "gsorm.tx.conn is nil"

	// Prepare for test.
	tx := &gsorm.ExportedTx{}

	// Actual process.
	err := tx.Commit()

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}

func TestTx_Rollback(t *testing.T) {
	// Prepare for test.
	tx := &gsorm.ExportedTx{}
	stx := &SpyTx{}
	tx.ExportedSetConn(stx)

	// Actual process.
	if err := tx.Rollback(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, stx.calledRollback)
}

func TestTx_Rollback_Fail(t *testing.T) {
	expectedErr := "gsorm.tx.conn is nil"

	// Prepare for test.
	tx := &gsorm.ExportedTx{}

	// Actual process.
	err := tx.Rollback()

	// Validate if expected error was occurred.
	assert.EqualError(t, err, expectedErr)
}
