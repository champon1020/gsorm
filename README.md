# gsorm
This is new Simple and SQL-like ORM framework written in Go language.
gsorm lets you implement database operation easily and intuitively.

Major features of gsorm are as follows:

- SQL-like implementation
- Provide the gsorm's own mock
- Mapping into struct, map, variable or their slice with high performance

You can see the usage in [Quick Start](https://github.com/champon1020/gsorm#quick-start) or [Documents](https://github.com/champon1020/gsorm/blob/main/docs/README.md)


## Installation
```
go get github.com/champon1020/gsorm
```


## Quick Start
```go
package main

import (
	"log"
	"time"

	"github.com/champon1020/gsorm"
	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	EmpNo     int       `gsorm:"emp_no,pk=PK_emp_no,notnull=t"`
	Name      string    `gsorm:"first_name,typ=VARCHAR(14),notnull=t"`
	BirthDate time.Time `gsorm:"notnull=t"`
}

func main() {
	db, err := gsorm.Open("mysql", "root:toor@tcp(localhost:3306)/employees?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	e := Employee{}

	// CREATE TABLE employees (
	//      emp_no      INT         NOT NULL,
	//      first_name  VARCHAR(14) NOT NULL,
	//      birth_date  DATETIME    NOT NULL,
	//      CONSTRAINT PK_emp_no PRIMARY KEY (emp_no)
	// );
	err = gsorm.CreateTable(db, "employees").Model(&e).Migrate()
	if err != nil {
		log.Fatal(err)
	}

	e = Employee{EmpNo: 1001, Name: "Taro", BirthDate: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)}

	// INSERT INTO employees VALUES (1001, 'Taro', '2006-01-02 15:04:05');
	err = gsorm.Insert(db, "employees").Model(&e).Exec()
	if err != nil {
		log.Fatal(err)
	}

	// INSERT INTO employees
	//      VALUES (1002, 'Jiro', '2007-01-02 15:04:05'), (1003, 'Saburo', '2006-01-02 15:04:05')
	err = gsorm.Insert(db, "employees").
		Values(1002, "Jiro", time.Date(2007, time.January, 2, 15, 4, 5, 0, time.UTC)).
		Values(1003, "Saburo", time.Date(2008, time.January, 2, 15, 4, 5, 0, time.UTC)).
		Exec()
	if err != nil {
		log.Fatal(err)
	}

	es := []Employee{}

	// SELECT first_name FROM employees WHERE id >= 1001;
	err = gsorm.Select(db, "first_name").From("employees").Query(&es)
	if err != nil {
		log.Fatal(err)
	}

	// UPDATE employees SET first_name = 'Kotaro' WHERE emp_no = 1001;
	err = gsorm.Update(db, "employees").Set("first_name", "Kotaro").Where("emp_no = ?", 1001).Exec()
	if err != nil {
		log.Fatal(err)
	}

	// DELETE FROM employees WHERE first_name = 'Kotaro';
	err = gsorm.Delete(db).From("employees").Where("first_name = ?", "Kotaro").Exec()
	if err != nil {
		log.Fatal(err)
	}

	// DROP TABLE employees;
	err = gsorm.DropTable(db, "employees").Migrate()
	if err != nil {
		log.Fatal(err)
	}
}
```
