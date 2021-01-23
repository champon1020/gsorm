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

// SelectSample1 is.
func SelectSample1(db *mgorm.DB) {
	fmt.Println("Query Sample 1")
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

// SelectSample2 is.
func SelectSample2(db *mgorm.DB) {
	fmt.Println("Query Sample 2")
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

// SelectSample3 is.
func SelectSample3(db *mgorm.DB) {
	fmt.Println("Query Sample 3")
	emp := new([]Employee)
	start := time.Now()

	// SELECT * FROM employees WHERE emp_no > 15000;
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

// SelectSample4 is.
func SelectSample4(db *mgorm.DB) {
	fmt.Println("Query Sample 4")
	emp := new([]Employee)
	start := time.Now()

	// SELECT * FROM employees WHERE emp_no > 15000 AND birth_date < "1960-10-20";
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

// SelectSample5 is.
func SelectSample5(db *mgorm.DB) {
	fmt.Println("Query Sample 5")
	emp := new([]Employee)
	start := time.Now()

	// SELECT * FROM employees WHERE emp_no < 15000 OR 20000 < emp_no;
	err := mgorm.Select(db, "*").
		From("employees").
		Where("emp_no < ?", 15000).
		Or("emp_no > ?", 20000).
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

// SelectSample6 is.
func SelectSample6(db *mgorm.DB) {
	fmt.Println("Query Sample 6")
	emp := new([]Employee)
	start := time.Now()

	// SELECT * FROM employees LIMIT 5;
	err := mgorm.Select(db, "*").
		From("employees").
		Limit(5).
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

// SelectSample7 is.
func SelectSample7(db *mgorm.DB) {
	fmt.Println("Query Sample 7")
	emp := new([]Employee)
	start := time.Now()

	// SELECT * FROM employees OFFSET 6;
	err := mgorm.Select(db, "*").
		From("employees").
		Offset(6).
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
