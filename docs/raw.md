# Raw
`Rawclause` or `RawStmt` calls clause or statement with raw string.


## RawClause
`RawClause` makes the implementation using gsorm generic.

When the query is executed, the values will be assigned to `?` in the expression.

Assignment rules are as follows:
- If the type of value is `string` or `time.Time`, the value is enclosed in single quotes
- If the value is slice or array, its elements are expanded
- If the type of value is `gsorm.Stmt`, `gsorm.Stmt` is built
- If the above conditions are not met, the value is assigned as is

`RawsClause` is supported on all Stmt structs and can be called at any time.

However, `RawClause` cannot be used with `InsertStmt.Select`, `InsertStmt.Model`, `UpdateStmt.Model` or `CreateTable.Model` methods.

#### Example
```go
err := gsorm.CreateTable(db, "dept_emp").
    Column("dept_no", "INT").NotNull().RawClause("AUTO_INCREMENT").
    Column("emp_no", "INT").NotNull().
    Cons("FK_dept_emp").Foreign("emp_no").Ref("employees", "emp_no").RawClause("ON UPDATE CASCADE").
    Migrate()
// CREATE TABLE employees(
//      dept_no INT NOT NULL AUTO_INCREMENT,
//      emp_no  INT NOT NULL,
//      CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no) REFERENCES employees (emp_no) ON UPDATE CASCADE
// );
```


## RawStmt
`RawStmt` calls the SQL with raw string.

When the query is executed, the values will be assigned to `?` in the expression.

Assignment rules are as follows:
- If the type of value is `string` or `time.Time`, the value is enclosed in single quotes
- If the value is slice or array, its elements are expanded
- If the type of value is `gsorm.Stmt`, `gsorm.Stmt` is built
- If the above conditions are not met, the value is assigned as is

`RawStmt` has `Query`, `Exec` and `Migrate` methods.

#### Example
```go
err := gsorm.RawStmt("SELECT * FROM employees").Query(&model)
// SELECT * FROM employees;

err := gsorm.RawStmt("SELECT * FROM employees WHERE emp_no = ?", 1001).Query(&model)
// SELECT * FROM employees WHERE emp_no = 1001;

err := gsorm.RawStmt("SELECT * FROM employees WHERE first_name = ?", "Taro").Query(&model)
// SELECT * FROM employees WHERE first_name = 'Taro';

err := gsorm.RawStmt("SELECT * FROM employees WHERE birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Query(&model)
// SELECT * FROM employees WHERE birth_date = '2006-01-02 00:00:00';

err := gsorm.RawStmt("SELECT * FROM employees WHERE emp_no IN (?)", []int{1001, 1002}).Query(&model)
// SELECT * FROM employees WHERE emp_no IN (1001, 1002);

err := gsorm.RawStmt("SELECT * FROM employees WHERE emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Query(&model)
// SELECT * FROM employees WHERE emp_no IN (SELECT emp_no FROM dept_manager);

err := gsorm.RawStmt("DELETE FROM employees").Exec()
// DELETE FROM employees;

err := gsorm.RawStmt("ALTER TABLE employees DROP PRIMARY KEY PK_emp_no").Migrate()
// ALTER TABLE employees DROP PRIMARY KEY PK_emp_no;
```
