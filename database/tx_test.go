package database_test

import (
	"testing"

	"github.com/champon1020/mgorm/database"
	"github.com/morikuni/failure"
	"gotest.tools/v3/assert"
)

func TestTx_Ping(t *testing.T) {
	db := new(database.ExportedDB)
	sdb := new(SpyDB)
	db.ExportedSetConn(sdb)

	tx := new(database.ExportedTx)
	tx.ExportedSetDB(db)

	err := tx.Ping()
	if err != nil {
		t.Errorf("error was occured: %v", err)
	}

	assert.Equal(t, true, sdb.calledPing)
}

func TestTx_Ping_Fail(t *testing.T) {
	expectedError := database.ErrFailedTxConnection

	// Prepare for test.
	tx := new(database.ExportedTx)

	// Actual process.
	err := tx.Ping()

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestTx_Exec(t *testing.T) {
	// Prepare for test.
	tx := new(database.ExportedTx)
	stx := new(SpyTx)
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
	expectedError := database.ErrFailedTxConnection

	// Prepare for test.
	tx := new(database.ExportedTx)

	// Actual process.
	_, err := tx.Exec("")

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestTx_Query(t *testing.T) {
	// Prepare for test.
	tx := new(database.ExportedTx)
	stx := new(SpyTx)
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
	expectedError := database.ErrFailedTxConnection

	// Prepare for test.
	tx := new(database.ExportedTx)

	// Actual process.
	_, err := tx.Query("")

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestTx_Commit(t *testing.T) {
	// Prepare for test.
	tx := new(database.ExportedTx)
	stx := new(SpyTx)
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
	expectedError := database.ErrFailedTxConnection

	// Prepare for test.
	tx := new(database.ExportedTx)

	// Actual process.
	err := tx.Commit()

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestTx_Rollback(t *testing.T) {
	// Prepare for test.
	tx := new(database.ExportedTx)
	stx := new(SpyTx)
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
	expectedError := database.ErrFailedTxConnection

	// Prepare for test.
	tx := new(database.ExportedTx)

	// Actual process.
	err := tx.Rollback()

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}
