# Raw
`RawClause`，`RawStmt`は指定された文字列による句や文を呼び出します．


## RawClause
`RawClause`はSQLの汎用性を高めるために作られました．

第1引数に文字列，第2引数以降に複数値を指定できます．
この際，文字列における`?`に値が代入されいます．
また，代入規則は以下に従います．
- 値が`string`もしくは`time.Time`の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が展開されます．
- 値が`*gsorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない値はそのまま展開される．

使用者は好きなタイミングで`RawClause`を呼び出すことができます．
ただし，`InsertStmt`，`UpdateStmt`，`CreateTable`のメソッドリストに含まれる`Model`や`InsertStmt`のメソッドリストに含まれる`Select`と併用することはできません．

`RawClause`は全てのStmt構造体においてサポートされています．

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

第1引数に文字列，第2引数以降に複数値を指定できます．
この際，文字列における`?`に値が代入されいます．
また，代入規則は以下に従います．
- 値が`string`もしくは`time.Time`の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が展開されます．
- 値が`*gsorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`RawStmt`は`Query`，`Exec`，`Migrate`の全てをサポートしています．

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
