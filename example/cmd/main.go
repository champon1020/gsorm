package main

import (
	"database/sql"
	"fmt"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/example"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_db, err := sql.Open("mysql", "root:toor@tcp(127.0.0.1:33306)/employees")
	if err != nil {
		fmt.Println("Failed to open database connection")
		return
	}
	db := mgorm.New(_db)

	example.Sample1(db)
	example.Sample2(db)
	example.Sample3(db)
	example.Sample4(db)
}

/*
func sampleGorm() {
	db, err := gorm.Open("mysql", "root:toor@tcp(127.0.0.1:33306)/employees")
	if err != nil {
		fmt.Println("Failed to open database connection")
		return
	}

	type Employee struct {
		EmpNo     int    `gorm:"emp_no"`
		BirthDate string `gorm:"birth_date"`
		FirstName string `gorm:"first_name"`
		LastName  string `gorm:"last_name"`
		Gender    string `gorm:"gender"`
		HireDate  string `gorm:"hire_date"`
	}

	emp := new([]Employee)
	start := time.Now()
	db.Find(emp)
	end := time.Now()
	fmt.Printf("[GORM] Len: %d, Index[0]: %v\n", len(*emp), (*emp)[0])
	fmt.Printf("%f [sec]\n", (end.Sub(start)).Seconds())
}
*/
