# Raw
`RawClause`，`RawStmt`は指定された文字列による句や文を呼び出します．


## RawClause
`RawClause`はSQLの汎用性を高めるために作られました．

第1引数に文字列，第2引数以降に複数値を指定できます．
この際，文字列における`?`に値が代入されいます．
また，代入規則は以下に従います．
- 値が`string`もしくは`time.Time`の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が展開されます．
- 値が`*mgorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない値はそのまま展開される．

使用者は好きなタイミングで`RawClause`を呼び出すことができます．
ただし，`InsertStmt`，`UpdateStmt`，`CreateTable`のメソッドリストに含まれる`Model`と併用することはできません．

`RawClause`は全てのStmt構造体においてサポートされています．

#### 例
```go
err := mgorm.CreateTable(db, "dept_emp").
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

引数には文字列を指定します．

`RawStmt`は`Query`，`Exec`，`Migrate`の全てをサポートしています．

#### 例
```go
err := mgorm.RawStmt("SELECT * FROM employees").Query(&model)
// SELECT * FROM employees;

err := mgorm.RawStmt("DELETE FROM employees").Exec()
// DELETE FROM employees;

err := mgorm.RawStmt("ALTER TABLE employees DROP PRIMARY KEY PK_emp_no").Migrate()
// ALTER TABLE employees DROP PRIMARY KEY PK_emp_no;
```
