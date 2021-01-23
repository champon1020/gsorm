package example

import (
	"fmt"
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

// Sample1 is.
func Sample1(db *mgorm.DB) {
	emp := new([]Employee)
	start := time.Now()

	// SELECT * FROM employees;
	err := mgorm.Select(db, "*").
		From("employees").
		Query(emp)

	end := time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Time: %f[sec], Len: %d\n", (end.Sub(start)).Seconds(), len(*emp))
	if len(*emp) > 0 {
		fmt.Printf("Index[0]: %v\n", (*emp)[0])
	}
}

// Sample2 is.
func Sample2(db *mgorm.DB) {
	emp := new([]Employee)
	start := time.Now()

	// SELECT emp_no, first_name FROM employees;
	err := mgorm.Select(db, "emp_no", "first_name").
		From("employees").
		Query(emp)

	end := time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Time: %f[sec], Len: %d\n", (end.Sub(start)).Seconds(), len(*emp))
	if len(*emp) > 0 {
		fmt.Printf("Index[0]: %v\n", (*emp)[0])
	}
}

// Sample3 is.
func Sample3(db *mgorm.DB) {
	emp := new([]Employee)
	start := time.Now()

	// SELECT * FROM employees WHERE id > 15000;
	err := mgorm.Select(db, "*").
		From("employees").
		Where("emp_no > ?", 15000).
		Query(emp)

	end := time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Time: %f[sec], Len: %d\n", (end.Sub(start)).Seconds(), len(*emp))
	if len(*emp) > 0 {
		fmt.Printf("Index[0]: %v\n", (*emp)[0])
	}
}

// Sample4 is.
func Sample4(db *mgorm.DB) {
	emp := new([]Employee)
	start := time.Now()

	err := mgorm.Select(db, "*").
		From("employees").
		Where("emp_no > ?", 15000).
		And("birth_date < ?", "1960-10-20").
		Query(emp)

	end := time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Time: %f[sec], Len: %d\n", (end.Sub(start)).Seconds(), len(*emp))
	if len(*emp) > 0 {
		fmt.Printf("Index[0]: %v\n", (*emp)[0])
	}
}
