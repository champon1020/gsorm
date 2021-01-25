# SELECT statement.
`mgorm.Select` creates instance of `*mgorm.Stmt` and assigns `SELECT` command to it.
First argument of `mgorm.Select` is instance of `*mgorm.DB`.
Second and later arguments are column names that you want to find.

Finally, you must call `mgorm.Stmt.Query` to execute SQL query.
Then, the argument is pointer to array or slice of struct which you want to store.
`mgorm.Stmt.Query` returns runtime error.
If no error was occurred, it would return `nil`.

In following examples, we use [genschsa/mysql-employees](https://hub.docker.com/r/genschsa/mysql-employees) as mock database.
And `db` is instance of `*mgorm.DB` and `emp` is pointer to slice of `Employee`.
Struct of `Employee` is declared as bellow.

```go
type Employee struct {
	EmpNo       int       `mgorm:"emp_no"`
	BirthDate   time.Time `mgorm:"birth_date" layout:"2006-01-02"`
	FirstName   string    `mgorm:"first_name"`
	LastName    string    `mgorm:"last_name"`
	Gender      string    `mgorm:"gender"`
	HireDate    time.Time `mgorm:"hire_date" layout:"2006-01-02"`
}

emp := new([]Employee)
```

About `Employee`, it is possible to select column name with struct tag `mgorm`.
If `mgorm` tag is not declared, the snake case of field name is used as column name.

When field type is `time.Time`, you can define the layout with `layout` tag.
Of course you can set variables like `time.ANSIC` or `time.RFC3339` to the tag.

## Simple Usage
You can set multi columns to `mgorm.Select`.

```go
// SELECT * FROM employees;
err := mgorm.Select(db, "*").From("employees").Query(emp)

// SELECT emp_no FROM employees;
err := mgorm.Select(db, "emp_no").From("employees").Query(emp)

// SELECT emp_no, first_name FROM employees;
err := mgorm.Select(db, "emp_no", "first_name").From("employees").Query(emp)
```

## WHERE statement.
When you set variable to `mgorm.Stmt.Where`, you must write `?` to first argument where you want to assign.

```go
// SELECT * FROM employees WHERE emp_no = 10001;
err := mgorm.Select(db, "*").
    From("employees").
    Where("emp = 10001").
    Query(emp)
    
// SELECT * FROM employees WHERE emp_no = 10001;
err := mgorm.Select(db, "*").
    From("employees").
    Where("emp = ?", 10001).
    Query(emp)
    
// SELECT * FROM employees WHERE emp_no > 10001;
err := mgorm.Select(db, "*").
    From("employees").
    Where("emp > ?", 10001).
    Query(emp)
```

## AND, OR, NOT statement.
It is possible to add `AND`, `OR` and `NOT` to first argument of `mgorm.Stmt.Where` by yourself.
When you want to enclose the statement with parentheses like `... OR (emp_no = 10001 AND first_name = "Georgi")`, you can use `mgorm.Stmt.Or` or `mgorm.Stmt.And`.

```go
// SELECT * FROM employees WHERE emp_no = 10001 AND first_name = "Georgi";
err := mgorm.Select(db, "*").
    From("employees").
    Where("emp_no = ? AND first_name = ?", 10001, "Georgi").
    Query(emp)
    
// SELECT * FROM employees WHERE emp_no = 10001 OR emp_no = 15000;
err := mgorm.Select(db, "*").
    From("employees").
    Where("emp_no = ? OR first_name = ?", 10001, 15000).
    Query(emp)
    
// SELECT * FROM employees WHERE NOT emp_no = 10001;
err := mgorm.Select(db, "*").
    From("employees").
    Where("NOT emp_no = ?", 10001).
    Query(emp)
    
// SELECT * FROM employees WHERE NOT emp_no = 15000 OR (emp_no = 10001 AND first_name = "Georgi");
err := mgorm.Select(db, "*").
    From("employees").
    Where("NOT emp_no = ?", 15000).
    Or("emp_no = ? AND first_name = ?", 10001, "Georgi").
    Query(emp)
```

## ORDER BY statement.
You can choose `DESC` by assining `true` to second argument of `mgorm.Stmt.OrderBy`.
When you call `mgorm.Stmt.OrderBy` multi time, the order of calling reflects to actual SQL.

```go
// SELECT * FROM employees ORDER BY birth_date.
err := mgorm.Select(db, "*").
    From("employees").
    OrderBy("birth_date", false).
    Query(emp)

// SELECT * FROM employees ORDER BY birth_date DESC.
err := mgorm.Select(db, "*").
    From("employees").
    OrderBy("birth_date", true).
    Query(emp)
    
// SELECT * FROM employees ORDER BY hire_date, birth_date DESC.
err := mgorm.Select(db, "*").
    From("employees").
    OrderBy("hire_date").
    OrderBy("birth_date", true).
    Query(emp)
```

## LIMIT and OFFSET statement.
You cannot use `mgorm.Stmt.Offset` without `mgorm.Stmt.Limit` because the syntax is not allowed in SQL.

```go
// SELECT * FROM employees LIMIT 5;
err := mgorm.Select(db, "*").
    From("employees").
    Limit(5).
    Query(emp)
    
// SELECT * FROM employees LIMIT 5 OFFSET 6;
err := mgorm.Select(db, "*").
    From("employees").
    Limit(5).
    Offset(6).
    Query(emp)
```
