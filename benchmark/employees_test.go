package benchmark

import (
	"database/sql"
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	EmpNo     int    `db:"emp_no"`
	BirthDate string `db:"birth_date"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Gender    string `db:"gender"`
	HireDate  string `db:"hire_date"`
}

var dsn = "root:toor@tcp(localhost:33306)/employees"

func BenchmarkSelect_standard(b *testing.B) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	rows, err := db.Query("SELECT * FROM employees")
	if err != nil || rows == nil {
		b.Fatal(err)
	}
	defer rows.Close()

	var emp []Employee
	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.EmpNo,
			&e.BirthDate,
			&e.FirstName,
			&e.LastName,
			&e.Gender,
			&e.HireDate); err != nil {
			b.Fatal(err)
		}
		emp = append(emp, e)
	}
}

func BenchmarkSelect_mgorm(b *testing.B) {
	db, err := mgorm.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var emp []Employee
	err = mgorm.Select(db).From("employees").Query(&emp)
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkSelect_gorm(b *testing.B) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var emp []Employee
	err = db.Find(&emp).Error
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkSelect_sqlx(b *testing.B) {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var emp []Employee
	err = db.Select(&emp, "SELECT * FROM employees")
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkSelect_gorp(b *testing.B) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	b.ResetTimer()

	var emp []Employee
	_, err = dbmap.Select(&emp, "SELECT * FROM employees")
	if err != nil {
		b.Fatal(err)
	}
}
