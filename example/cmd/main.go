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

	// Select samples.
	example.SelectSample1(db)
	example.SelectSample2(db)
	example.SelectSample3(db)
	example.SelectSample4(db)
	example.SelectSample5(db)
	example.SelectSample6(db)
	example.SelectSample7(db)
	example.SelectSample8(db)

	// Insert samples.
	//example.InsertSample1(db)

	// Update samples.
	//example.UpdateSample1(db)
}
