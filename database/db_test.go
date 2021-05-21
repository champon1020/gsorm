package database_test

import (
	"testing"

	"github.com/champon1020/gsorm/database"
	"github.com/morikuni/failure"
	"gotest.tools/v3/assert"
)

func TestDB_Ping(t *testing.T) {
	// Prepare for test.
	db := new(database.ExportedDB)
	sdb := new(SpyDB)
	db.ExportedSetConn(sdb)

	// Actual process.
	if err := db.Ping(); err != nil {
		t.Errorf("Error was occurred: %v", err)
		return
	}

	// Validate if expected error was occurred.
	assert.Equal(t, true, sdb.calledPing)
}

func TestDB_GetDriver(t *testing.T) {
	expected := database.MysqlDriver

	db := new(database.ExportedDB)
	db.ExportedSetDriver(expected)

	assert.Equal(t, expected, db.GetDriver())
}

func TestDB_Ping_Fail(t *testing.T) {
	expectedError := database.ErrFailedDBConnection

	// Prepare for test.
	db := new(database.ExportedDB)

	// Actual process.
	err := db.Ping()

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestDB_Exec(t *testing.T) {
	// Prepare for test.
	db := new(database.ExportedDB)
	sdb := new(SpyDB)
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
	expectedError := database.ErrFailedDBConnection

	// Prepare for test.
	db := new(database.ExportedDB)

	// Actual process.
	_, err := db.Exec("")
	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestDB_Query(t *testing.T) {
	// Prepare for test.
	db := new(database.ExportedDB)
	sdb := new(SpyDB)
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
	expectedError := database.ErrFailedDBConnection

	// Prepare for test.
	db := new(database.ExportedDB)

	// Actual process.
	_, err := db.Query("")

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestDB_SetConnMaxLifetime(t *testing.T) {
	// Prepare for test.
	db := new(database.ExportedDB)
	sdb := new(SpyDB)
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
	expectedError := database.ErrFailedDBConnection

	// Prepare for test.
	db := new(database.ExportedDB)

	// Actual process.
	err := db.SetConnMaxLifetime(0)

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestDB_SetMaxIdleConns(t *testing.T) {
	// Prepare for test.
	db := new(database.ExportedDB)
	sdb := new(SpyDB)
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
	expectedError := database.ErrFailedDBConnection

	// Prepare for test.
	db := new(database.ExportedDB)

	// Actual process.
	err := db.SetMaxIdleConns(0)

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestDB_SetMaxOpenConns(t *testing.T) {
	// Prepare for test.
	db := new(database.ExportedDB)
	sdb := new(SpyDB)
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
	expectedError := database.ErrFailedDBConnection

	// Prepare for test.
	db := new(database.ExportedDB)

	// Actual process.
	err := db.SetMaxOpenConns(0)

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestDB_Close(t *testing.T) {
	// Prepare for test.
	db := new(database.ExportedDB)
	sdb := new(SpyDB)
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
	expectedError := database.ErrFailedDBConnection

	// Prepare for test.
	db := new(database.ExportedDB)

	// Actual process.
	err := db.Close()

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}

func TestDB_Begin(t *testing.T) {
	// Prepare for test.
	db := new(database.ExportedDB)
	sdb := new(SpyDB)
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
	expectedError := database.ErrFailedDBConnection

	// Prepare for test.
	db := new(database.ExportedDB)

	// Actual process.
	_, err := db.Begin()

	// Validate if expected error was occurred.
	if !failure.Is(err, expectedError) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %+v", expectedError)
		t.Errorf("  Actual:   %+v", err)
	}
}
