package mgorm_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
	"github.com/stretchr/testify/assert"
)

type SpyDB struct {
	sql.DB

	calledPing               bool
	calledExec               bool
	calledQuery              bool
	calledSetConnMaxLifetime bool
	calledSetMaxIdleConns    bool
	calledSetMaxOpenConns    bool
	calledClose              bool
	calledBegin              bool
}

func (d *SpyDB) Ping() error {
	d.calledPing = true
	return nil
}

func (d *SpyDB) Exec(string, ...interface{}) (sql.Result, error) {
	d.calledExec = true
	return nil, nil
}

func (d *SpyDB) Query(string, ...interface{}) (*sql.Rows, error) {
	d.calledQuery = true
	return nil, nil
}

func (d *SpyDB) SetConnMaxLifetime(time.Duration) {
	d.calledSetConnMaxLifetime = true
}

func (d *SpyDB) SetMaxIdleConns(int) {
	d.calledSetMaxIdleConns = true
}

func (d *SpyDB) SetMaxOpenConns(int) {
	d.calledSetMaxOpenConns = true
}

func (d *SpyDB) Close() error {
	d.calledClose = true
	return nil
}

func (d *SpyDB) Begin() (*sql.Tx, error) {
	d.calledBegin = true
	return nil, nil
}

func TestDB_Ping(t *testing.T) {
	// Prepare for test.
	db := new(mgorm.DB)
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

func TestDB_Ping_Fail(t *testing.T) {
	expectedErr := errors.New("DB conn is nil", errors.InvalidValueError).(*errors.Error)

	// Prepare for test.
	db := new(mgorm.DB)

	// Actual process.
	err := db.Ping()
	if err == nil {
		t.Errorf("Error was not occurred")
		return
	}

	// Validate if expected error was occurred.
	actualErr, ok := err.(*errors.Error)
	if !ok {
		t.Errorf("Error type is invalid")
		return
	}
	if !actualErr.Is(expectedErr) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
		t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
	}
}

func TestDB_Exec(t *testing.T) {
	// Prepare for test.
	db := new(mgorm.DB)
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
	expectedErr := errors.New("DB conn is nil", errors.InvalidValueError).(*errors.Error)

	// Prepare for test.
	db := new(mgorm.DB)

	// Actual process.
	_, err := db.Exec("")
	if err == nil {
		t.Errorf("Error was not occurred")
		return
	}

	// Validate if expected error was occurred.
	actualErr, ok := err.(*errors.Error)
	if !ok {
		t.Errorf("Error type is invalid")
		return
	}
	if !actualErr.Is(expectedErr) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
		t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
	}
}

func TestDB_Query(t *testing.T) {
	// Prepare for test.
	db := new(mgorm.DB)
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
	expectedErr := errors.New("DB conn is nil", errors.InvalidValueError).(*errors.Error)

	// Prepare for test.
	db := new(mgorm.DB)

	// Actual process.
	_, err := db.Query("")
	if err == nil {
		t.Errorf("Error was not occurred")
		return
	}

	// Validate if expected error was occurred.
	actualErr, ok := err.(*errors.Error)
	if !ok {
		t.Errorf("Error type is invalid")
		return
	}
	if !actualErr.Is(expectedErr) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
		t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
	}
}

func TestDB_SetConnMaxLifetime(t *testing.T) {
	// Prepare for test.
	db := new(mgorm.DB)
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
	expectedErr := errors.New("DB conn is nil", errors.InvalidValueError).(*errors.Error)

	// Prepare for test.
	db := new(mgorm.DB)

	// Actual process.
	err := db.SetConnMaxLifetime(0)
	if err == nil {
		t.Errorf("Error was not occurred")
		return
	}

	// Validate if expected error was occurred.
	actualErr, ok := err.(*errors.Error)
	if !ok {
		t.Errorf("Error type is invalid")
		return
	}
	if !actualErr.Is(expectedErr) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
		t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
	}
}

func TestDB_SetMaxIdleConns(t *testing.T) {
	// Prepare for test.
	db := new(mgorm.DB)
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
	expectedErr := errors.New("DB conn is nil", errors.InvalidValueError).(*errors.Error)

	// Prepare for test.
	db := new(mgorm.DB)

	// Actual process.
	err := db.SetMaxIdleConns(0)
	if err == nil {
		t.Errorf("Error was not occurred")
		return
	}

	// Validate if expected error was occurred.
	actualErr, ok := err.(*errors.Error)
	if !ok {
		t.Errorf("Error type is invalid")
		return
	}
	if !actualErr.Is(expectedErr) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
		t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
	}
}

func TestDB_SetMaxOpenConns(t *testing.T) {
	// Prepare for test.
	db := new(mgorm.DB)
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
	expectedErr := errors.New("DB conn is nil", errors.InvalidValueError).(*errors.Error)

	// Prepare for test.
	db := new(mgorm.DB)

	// Actual process.
	err := db.SetMaxOpenConns(0)
	if err == nil {
		t.Errorf("Error was not occurred")
		return
	}

	// Validate if expected error was occurred.
	actualErr, ok := err.(*errors.Error)
	if !ok {
		t.Errorf("Error type is invalid")
		return
	}
	if !actualErr.Is(expectedErr) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
		t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
	}
}

func TestDB_Close(t *testing.T) {
	// Prepare for test.
	db := new(mgorm.DB)
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
	expectedErr := errors.New("DB conn is nil", errors.InvalidValueError).(*errors.Error)

	// Prepare for test.
	db := new(mgorm.DB)

	// Actual process.
	err := db.Close()
	if err == nil {
		t.Errorf("Error was not occurred")
		return
	}

	// Validate if expected error was occurred.
	actualErr, ok := err.(*errors.Error)
	if !ok {
		t.Errorf("Error type is invalid")
		return
	}
	if !actualErr.Is(expectedErr) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
		t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
	}
}

func TestDB_Begin(t *testing.T) {
	// Prepare for test.
	db := new(mgorm.DB)
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
	expectedErr := errors.New("DB conn is nil", errors.InvalidValueError).(*errors.Error)

	// Prepare for test.
	db := new(mgorm.DB)

	// Actual process.
	_, err := db.Begin()
	if err == nil {
		t.Errorf("Error was not occurred")
		return
	}

	// Validate if expected error was occurred.
	actualErr, ok := err.(*errors.Error)
	if !ok {
		t.Errorf("Error type is invalid")
		return
	}
	if !actualErr.Is(expectedErr) {
		t.Errorf("Different error was occurred")
		t.Errorf("  Expected: %s, Code: %d", expectedErr.Error(), expectedErr.Code)
		t.Errorf("  Actual:   %s, Code: %d", actualErr.Error(), actualErr.Code)
	}
}
