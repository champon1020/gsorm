# CreateIndex
`mgorm.CreateIndex`はCREATE INDEX句を呼び出します．

引数にはデータベースのコネクション(`mgorm.Conn`)，インデックス名を指定します．

#### 例
```go
err := mgorm.CreateIndex(db, "IDX_emp_no").
    On("employees", "emp_no").Migrate()
// CREATE INDEX IDX_emp
//  ON employees (emp_no);

err := mgorm.CreateIndex(db, "IDX_emp_no").
    On("employees", "emp_no", "first_name").Migrate()
// CREATE INDEX IDX_emp
//  ON employees (emp_no, first_name);
```
