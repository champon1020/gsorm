package benchmark

import "time"

type Employee struct {
	EmpNo     int       `db:"emp_no"`
	BirthDate time.Time `db:"birth_date"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Gender    string    `db:"gender"`
	HireDate  time.Time `db:"hire_date"`
}

var dsn = "root:toor@tcp(localhost:33306)/employees?parseTime=true"
