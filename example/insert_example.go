package example

import (
	"fmt"
	"time"

	"github.com/champon1020/mgorm"
)

// InsertSample1 is.
func InsertSample1(db *mgorm.DB) {
	start := time.Now()

	// INSERT INTO employees ("emp_no", "birth_date", "first_name", "last_name", "gender", "hire_date")
	// VALUES (10000, "1997-04-30", "Taro", "Yokohama", "M", "2020-01-02");
	err := mgorm.Insert(db, "employees", "emp_no", "birth_date", "first_name", "last_name", "gender", "hire_date").
		Values(10000, "1997-04-30", "Taro", "Yokohama", "M", "2020-01-02").
		Exec()

	end := time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Time: %f[sec]\n", (end.Sub(start)).Seconds())
}
