package benchmark

import (
	"database/sql"
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/go-gorp/gorp"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func BenchmarkSelectAll_Struct_standard(b *testing.B) {
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

func BenchmarkSelectAll_Struct_mgorm(b *testing.B) {
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

func BenchmarkSelectAll_Struct_gorm(b *testing.B) {
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

func BenchmarkSelectAll_Struct_sqlx(b *testing.B) {
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

func BenchmarkSelectAll_Struct_gorp(b *testing.B) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	b.ResetTimer()

	var emp []Employee
	_, err = dbmap.Select(&emp, "SELECT * FROM employees")
	if err != nil {
		b.Fatal(err)
	}
}
