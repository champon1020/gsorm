# gsorm
[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm.svg)](https://pkg.go.dev/github.com/champon1020/gsorm)
[![Go Report Card](https://goreportcard.com/badge/github.com/champon1020/gsorm)](https://goreportcard.com/report/github.com/champon1020/gsorm)
[![codecov](https://codecov.io/gh/champon1020/gsorm/branch/main/graph/badge.svg?token=7FCUS2VZMV)](https://codecov.io/gh/champon1020/gsorm)

This is new Simple and SQL-like ORM framework written in Go language.
gsorm lets you implement database operation easily and intuitively.

Major features of gsorm are as follows:

- SQL-like implementation
- Provide the gsorm's own mock
- Mapping into struct, map, variable or their slice with high performance
- Smart database migration using field tags of Go structure

You can see the usage in [Quick Start](https://github.com/champon1020/gsorm#quick-start) or [Documents](https://github.com/champon1020/gsorm/blob/main/docs/README.md)


## Benchmark
I measured the benchmark of the mapping query results with [MySQL Employees Dataset](https://dev.mysql.com/doc/employee/en/) which has about 300000 rows.
Also I compared the average of 10 trials to other ORM libraries.

As a result, gsorm is faster than other libraries when mapping the multi rows.

The result are as follows:

| ORM | (ns/op) |
| ---- | ---- |
| standard | 0.40952 |
| gorm | >= 200 ms/op |
| sqlx | 0.49695 |
| gorp | 0.56168 |
| **gsorm** | **0.38252** |

If you want to run the benchmark on your machine, follow the steps below.

```
git clone git@github.com:champon1020/employees_database.git

cd employees_database

docker-compose up -d

cd /path/to/gsorm/benchmark

go test -bench . -benchmem -count 10
```

Benchmark codes are written under `benchmark` directory.


## Installation
```
go get github.com/champon1020/gsorm
```


## Features
### SQL-Like Implementation
You can implement the database query intuitaively like writing SQL.

```go
err := gsorm.Select(db, "emp_no", "first_name").From("employees").
    Where("emp_no > ?", 1002).
    Query(&model)
// SELECT emp_no, first_name FROM employees WHERE emp_no > 1002;

err := gsorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1001, "Taro").
    Values(1002, "Jiro").
    Exec()
// INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro');
```


### Flexible Implementation
gsorm provides `RawStmt` and `RawClause` methods.

Using them, it is possible to implement the database query more flexible.

```go
err := gsorm.Select(db).
    RawClause("INTO new_employees").
    From("employees").
    Where("emp_no > ?", 1002).
    Query(interface{}{})
// SELECT * INTO new_employees FROM employees WHERE emp_no > 1002;

err := gsorm.AlterTable(db, "employees").
    RawClause("MODIFY COLUMN emp_no INT(24)").NotNull().
    Migrate()
// ALTER TABLE employees MODIFY COLUMN emp_no INT(24) NOT NULL;

err := gsorm.RawStmt(db, "REPLACE INTO employees VALUES (?, ?)", 1003, "Saburo").Exec()
// REPLACE INTO employees VALUES (1003, 'Saburo');
```


### Test with Mock
gsorm provides own Mock structure.

You can implement testing code with mock easily.

```go
type Employee struct {
    EmpNo     int
    Firstname string
}

func TestWithMock(t *testing.T) {
    // Open the connection to mock database.
    mock := gsorm.OpenMock()

    // Expect statements that will be executed.
    mock.Expect(gsorm.Insert(nil, "employees", "emp_no", "first_name").
        Values(1001, "Taro").
        Values(1002, "Jiro"))
    mock.ExpectWithReturn(gsorm.Select(nil, "emp_no", "first_name").From("employees"), []Employee{
        {ID: 1001, FirstName: "Taro"},
        {ID: 1002, FirstName: "Jiro"},
    })

    actualProcess := func(db gsorm.DB) error {
        model := []Employee{}

        if err := gsorm.Insert(nil, "employees", "emp_no", "first_name").
            Values(1001, "Taro").
            Values(1002, "Jiro").Exec(); err != nil {
            return err
        }

        if err := gsorm.Select(nil, "emp_no", "first_name").From("employees").Query(&model); err != nil {
            return err
        }

        return nil
    }

    if err := actualProcess(mock); err != nil {
        t.Error(err)
    }

    // Check if all expected statements was executed.
    if err := mock.Complete(); err != nil {
        t.Error(err)
    }
}
```


### Smart Migration
Implementation of database migration tends to be complicated and high cost.

Using gsorm, it is possible to map the Go structure into SQL.

You can determine the properties of database columns using the field tag of the Go structure.

```go
type Employee struct {
	EmpNo     int       `gsorm:"emp_no,pk=PK_emp_no,notnull=t"`
	Name      string    `gsorm:"first_name,typ=VARCHAR(14),notnull=t"`
	BirthDate time.Time `gsorm:"notnull=t"`
}

err := gsorm.CreateTable(db, "employees").Model(&Employee{}).Migrate()
	// CREATE TABLE employees (
	//      emp_no      INT         NOT NULL,
	//      first_name  VARCHAR(14) NOT NULL,
	//      birth_date  DATETIME    NOT NULL,
	//      CONSTRAINT PK_emp_no PRIMARY KEY (emp_no)
	// );
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
