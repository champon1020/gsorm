package internal

// SQLDriver is SQL driver.
type SQLDriver int

// SQL driver list.
const (
	PSQL SQLDriver = iota
	MySQL
)
