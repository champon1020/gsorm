package main

import (
	"database/sql"
	"fmt"
	"time"

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

	sampleList := []func(*mgorm.DB, *[]example.Employee) (string, error){
		example.SelectSample1,
		example.SelectSample2,
		example.SelectSample3,
		example.SelectSample4,
		example.SelectSample5,
		example.SelectSample6,
		example.SelectSample7,
		example.SelectSample8,
	}

	for i, f := range sampleList {
		fmt.Printf("-------------------------\n")
		fmt.Printf("*** Query Sample %d ***\n", i+1)
		emp := new([]example.Employee)
		start := time.Now()
		sql, err := f(db, emp)
		end := time.Now()

		fmt.Printf("Query: %s\n", sql)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Time: %f[sec]\n", (end.Sub(start)).Seconds())
		fmt.Printf("Len: %d\n", len(*emp))
		fmt.Printf("First row:\n")
		fmt.Printf("  emp_no: %v\n", (*emp)[0].EmpNo)
		fmt.Printf("  birth_date: %v\n", (*emp)[0].BirthDate)
		fmt.Printf("  first_name: %v\n", (*emp)[0].FirstName)
		fmt.Printf("  last_name: %v\n", (*emp)[0].LastName)
		fmt.Printf("  gender: %v\n", (*emp)[0].Gender)
		fmt.Printf("  hire_date: %v\n", (*emp)[0].HireDate)
	}
}
