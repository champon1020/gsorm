package example

import (
	"time"

	"github.com/champon1020/mgorm"
)

// Employee structure.
type Employee struct {
	EmpNo       int       `mgorm:"emp_no"`
	BirthDate   time.Time `mgorm:"birth_date" layout:"2006-01-02"`
	FirstName   string    `mgorm:"first_name"`
	LastName    string    `mgorm:"last_name"`
	Gender      string    `mgorm:"gender"`
	HireDate    time.Time `mgorm:"hire_date" layout:"2006-01-02"`
	ResultInt   int       `mgorm:"res_int"`
	ResultFloat float32   `mgorm:"res_float"`
}

// QuerySamples returns a sample by index.
func QuerySamples(db *mgorm.DB, model interface{}, i int) (string, bool, error) {
	querySamples := []mgorm.ExecutableStmt{
		// SELECT * FROM employees;
		mgorm.Select(db, "*").
			From("employees"),

		// SELECT emp_no, first_name FROM employees;
		mgorm.Select(db, "emp_no", "first_name").
			From("employees"),

		// SELECT * FROM employees WHERE emp_no > 15000;
		mgorm.Select(db, "*").
			From("employees").
			Where("emp_no > ?", 15000),

		// SELECT * FROM employees WHERE emp_no > 15000 AND birth_date < "1960-10-20";
		mgorm.Select(db, "*").
			From("employees").
			Where("emp_no > ?", 15000).
			And("birth_date < ?", "1960-10-20"),

		// SELECT * FROM employees WHERE emp_no < 15000 OR 20000 < emp_no;
		mgorm.Select(db, "*").
			From("employees").
			Where("emp_no < ?", 15000).
			Or("emp_no > ?", 20000),

		// SELECT * FROM employees WHERE NOT (emp_no = 10001);
		mgorm.Select(db, "*").
			From("employees").
			Where("NOT emp_no = ?", 10001),

		// SELECT * FROM employees WHERE emp_no IN (SELECT * FROM dept_manager WHERE dept_no = "d001");
		mgorm.Select(db, "*").
			From("employees").
			Where("emp_no IN (SELECT emp_no FROM dept_manager WHERE dept_no = ?)", "d001"),

		// SELECT * FROM employees WHERE emp_no IN (SELECT * FROM dept_manager WHERE dept_no = "d001");
		mgorm.Select(db, "*").
			From("employees").
			Where("emp_no IN ?", mgorm.Select(db, "emp_no").
				From("dept_manager").
				Where("dept_no = ?", "d001").
				Var(),
			),

		// SELECT * FROM employees LIMIT 5;
		mgorm.Select(db, "*").
			From("employees").
			Limit(5),

		// SELECT * FROM employees LIMIT 5 OFFSET 6;
		mgorm.Select(db, "*").
			From("employees").
			Limit(5).
			Offset(6),

		// SELECT * FROM employees ORDER BY emp_no DESC;
		mgorm.Select(db, "*").
			From("employees").
			OrderBy("emp_no", true),

		// SELECT COUNT(birth_date) AS res_int FROM employees;
		mgorm.Count(db, "birth_date", "res_int").
			From("employees"),

		// SELECT AVG(emp_no) AS res_float FROM employees;
		mgorm.Avg(db, "emp_no", "res_float").
			From("employees"),

		// SELECT SUM(birth_date) AS res_int FROM employees;
		mgorm.Sum(db, "emp_no", "res_int").
			From("employees"),

		// SELECT MIN(birth_date) AS res_int FROM employees;
		mgorm.Min(db, "emp_no", "res_int").
			From("employees"),

		// SELECT MAX(birth_date) AS res_int FROM employees;
		mgorm.Max(db, "emp_no", "res_int").
			From("employees"),

		// SELECT first_name FROM employees AS e INNER JOIN dept_manager AS d ON e.emp_no = d.emp_no;
		mgorm.Select(db, "first_name").
			From("employees AS e").
			Join("dept_manager AS d").
			On("e.emp_no = d.emp_no"),

		// SELECT emp_no, first_name FROM employees UNION SELECT emp_no, first_name FROM v_full_employees;
		mgorm.Select(db, "emp_no", "first_name").
			From("employees").
			Union(mgorm.Select(nil, "emp_no", "first_name").
				From("v_full_employees").
				Var(),
			),

		// SELECT COUNT(first_name) AS res_int, last_name FROM employees GROUP BY last_name;
		mgorm.Select(db, "COUNT(first_name) AS res_int", "last_name").
			From("employees").
			GroupBy("last_name"),

		// SELECT COUNT(first_name) AS res_int, last_name FROM employees GROUP BY last_name HAVING res_int < 200;
		mgorm.Select(db, "COUNT(first_name) AS res_int", "last_name").
			From("employees").
			GroupBy("last_name").
			Having("res_int < ?", 200),
	}

	next := true
	if len(querySamples) <= i+1 {
		next = false
	}

	return querySamples[i].String(), next, querySamples[i].Query(model)
}

// ExecSamples returns a sample by index.
func ExecSamples(db *mgorm.DB, i int) (string, bool, error) {
	execSamples := []mgorm.ExecutableStmt{
		// INSERT INTO employees ("emp_no", "birth_date", "first_name", "last_name", "gender", "hire_date")
		// VALUES (10000, "1997-04-30", "Taro", "Yokohama", "M", "2020-01-02");
		mgorm.Insert(db, "employees", "emp_no", "birth_date", "first_name", "last_name", "gender", "hire_date").
			Values(10000, "1997-04-30", "Taro", "Yokohama", "M", "2020-01-02"),

		// UPDATE employees
		// SET emp_no=9999, birth_date="1997-03-30", first_name="Hanako", last_name="Tokyo", gender="F",
		//     hire_date="2020-02-02"
		// WHERE emp_no = 10000;
		mgorm.Update(db, "employees", "emp_no", "birth_date", "first_name", "last_name", "gender", "hire_date").
			Set(9999, "1997-03-30", "Hanako", "Tokyo", "F", "2020-02-02").
			Where("emp_no = ?", 10000),
	}
	next := true
	if len(execSamples) <= i+1 {
		next = false
	}

	return execSamples[i].String(), next, execSamples[i].Exec()
}
