package database

// SQLDriver is SQL driver.
type SQLDriver int

// SQL driver list.
const (
	PsqlDriver SQLDriver = iota
	MysqlDriver
)
