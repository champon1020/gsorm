package example

import (
	"fmt"
	"time"

	"github.com/champon1020/mgorm"
)

// UpdateSample1 is.
func UpdateSample1(db *mgorm.DB) {
	start := time.Now()

	// UPDATE employees
	// SET emp_no=9999, birth_date="1997-03-30", first_name="Hanako", last_name="Tokyo", gender="F",
	//     hire_date="2020-02-02"
	// WHERE emp_no = 10000;
	err := mgorm.Update(db, "employees", "emp_no", "birth_date", "first_name", "last_name", "gender", "hire_date").
		Set(9999, "1997-03-30", "Hanako", "Tokyo", "F", "2020-02-02").
		Where("emp_no = ?", 10000).
		Exec()

	end := time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Time: %f[sec]\n", (end.Sub(start)).Seconds())
}
