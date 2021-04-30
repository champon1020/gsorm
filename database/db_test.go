package database_test

import (
	"database/sql"
	"time"
)

type SpyDB struct {
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
