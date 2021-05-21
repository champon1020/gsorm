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

func BenchmarkSelectOne_Struct_standard(b *testing.B) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	rows, err := db.Query("SELECT * FROM employees LIMIT 1")
	if err != nil || rows == nil {
		b.Fatal(err)
	}
	defer rows.Close()

	var e Employee
	for rows.Next() {
		if err := rows.Scan(&e.EmpNo,
			&e.BirthDate,
			&e.FirstName,
			&e.LastName,
			&e.Gender,
			&e.HireDate); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSelectOne_Struct_mgorm(b *testing.B) {
	db, err := mgorm.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	var e Employee
	err = mgorm.Select(db).From("employees").Limit(1).Query(&e)
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkSelectOne_Struct_gorm(b *testing.B) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var e Employee
	err = db.First(&e).Error
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkSelectOne_Struct_sqlx(b *testing.B) {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	var e Employee
	err = db.Get(&e, "SELECT * FROM employees LIMIT 1")
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkSelectOne_Struct_gorp(b *testing.B) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	b.ResetTimer()

	var e Employee
	_, err = dbmap.Select(&e, "SELECT * FROM employees LIMIT 1")
	if err != nil {
		b.Fatal(err)
	}
}
