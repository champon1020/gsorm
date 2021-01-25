# mgorm
This is new ORM framework written in Go language.
`mgorm` lets you implement database operation easily and intuitively.

## Examples
`SELECT`, `INSERT`, `UPDATE` and `DELETE` statements are supported in `mgorm`.

```go
// SELECT * FROM employees;
err := mgorm.Select(db, "*").From("employees").Query(emp)

// INSERT INTO employees (emp_no, birth_date, first_name, last_name, gender, hire_date) 
// VALUES (9999, "1997-01-01", "Taro", "Japan", "M", "2020-01-01");
err := mgorm.Insert(db, "employees", "emp_no", "birth_date", "first_name", "last_name", "gender", "hire_date").
    Values(9999, "1997-01-01", "Taro", "Japan", "M", "2020-01-01").
    Exec()
    
// UPDATE employees
// SET emp_no=10000, birth_date="1997-01-02", first_name="Hanako", last_name="Nihon", gender="F", hire_date="2020-01-02"
// WHERE emp_no = 9999;
err := mgorm.Update(db, "employees", "emp_no", "birth_date", "first_name", "last_name", "gender", "hire_date").
    Set(10000, "1997-01-02", "Hanako", "Nihon", "F", "2020-01-02").
    Where(emp_no = ?, 9999).
    Exec()
    
// DELETE employees FROM employees WHERE emp_no = 10000; 
err := mgorm.Delete(db, "employees").
    Where(emp_no = ?, 10000).
    Exec()
```

## Specification
Framework specification is written in following files.

- [SELECT.md](https://github.com/champon1020/mgorm/tree/main/specification/SELECT.md)
