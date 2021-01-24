package example

import (
	"time"

	"github.com/champon1020/mgorm"
)

// Employee structure.
type Employee struct {
	EmpNo     int       `mgorm:"emp_no"`
	BirthDate time.Time `mgorm:"birth_date" layout:"2006-01-02"`
	FirstName string    `mgorm:"first_name"`
	LastName  string    `mgorm:"last_name"`
	Gender    string    `mgorm:"gender"`
	HireDate  time.Time `mgorm:"hire_date" layout:"2006-01-02"`
}

// SelectSample1 is.
func SelectSample1(db *mgorm.DB, emp *[]Employee) (string, error) {
	// SELECT * FROM employees;
	s := mgorm.Select(db, "*").
		From("employees")

	return s.String(), s.Query(emp)
}

// SelectSample2 is.
func SelectSample2(db *mgorm.DB, emp *[]Employee) (string, error) {
	// SELECT emp_no, first_name FROM employees;
	s := mgorm.Select(db, "emp_no", "first_name").
		From("employees")

	return s.String(), s.Query(emp)
}

// SelectSample3 is.
func SelectSample3(db *mgorm.DB, emp *[]Employee) (string, error) {
	// SELECT * FROM employees WHERE emp_no > 15000;
	s := mgorm.Select(db, "*").
		From("employees").
		Where("emp_no > ?", 15000)

	return s.String(), s.Query(emp)
}

// SelectSample4 is.
func SelectSample4(db *mgorm.DB, emp *[]Employee) (string, error) {
	// SELECT * FROM employees WHERE emp_no > 15000 AND birth_date < "1960-10-20";
	s := mgorm.Select(db, "*").
		From("employees").
		Where("emp_no > ?", 15000).
		And("birth_date < ?", "1960-10-20")

	return s.String(), s.Query(emp)
}

// SelectSample5 is.
func SelectSample5(db *mgorm.DB, emp *[]Employee) (string, error) {
	// SELECT * FROM employees WHERE emp_no < 15000 OR 20000 < emp_no;
	s := mgorm.Select(db, "*").
		From("employees").
		Where("emp_no < ?", 15000).
		Or("emp_no > ?", 20000)

	return s.String(), s.Query(emp)
}

// SelectSample6 is.
func SelectSample6(db *mgorm.DB, emp *[]Employee) (string, error) {
	// SELECT * FROM employees LIMIT 5;
	s := mgorm.Select(db, "*").
		From("employees").
		Limit(5)

	return s.String(), s.Query(emp)
}

// SelectSample7 is.
func SelectSample7(db *mgorm.DB, emp *[]Employee) (string, error) {
	// SELECT * FROM employees LIMIT 5 OFFSET 6;
	s := mgorm.Select(db, "*").
		From("employees").
		Limit(5).
		Offset(6)

	return s.String(), s.Query(emp)
}

// SelectSample8 is.
func SelectSample8(db *mgorm.DB, emp *[]Employee) (string, error) {
	// SELECT * FROM employees ORDER BY emp_no DESC;
	s := mgorm.Select(db, "*").
		From("employees").
		OrderBy("emp_no", true)

	return s.String(), s.Query(emp)
}

// SelectSample9 is.
func SelectSample9(db *mgorm.DB, emp *[]Employee) (string, error) {
	s := mgorm.Select(db, "*").
		From("employees").
		Not("emp_no = ?", 10001)

	return s.String(), s.Query(emp)
}
