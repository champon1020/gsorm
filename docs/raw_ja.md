# Raw
`RawClause`，`RawStmt`は指定された文字列による句や文を呼び出します．


## RawClause
`RawClause`はSQLの汎用性を高めるために作られました．

クエリが実行されるとき，条件式における`?`に値が代入されます．

代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値がスライスもしくは配列の場合，その要素が展開されます．
- 値が`gsorm.Stmt`型の場合，`gsorm.Stmt`は展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`RawClause`は全てのStmt構造体においてサポートされており，任意のタイミングで呼び出すことができます．

しかし，`InsertStmt.Select`, `InsertStmt.Model`, `UpdateStmt.Model`, `CreateTable.Model`などのメソッドと併用することはできません．

#### 例
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
`RawStmt`は文字列で指定したSQLを呼び出します．

クエリが実行されるとき，条件式における`?`に値が代入されます．

代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値がスライスもしくは配列の場合，その要素が展開されます．
- 値が`gsorm.Stmt`型の場合，`gsorm.Stmt`は展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`RawStmt`は`Query`，`Exec`，`Migrate`の全てをサポートしています．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#RawStmt.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#RawStmt)

#### 例
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
