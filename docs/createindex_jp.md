# CreateIndex
`mgorm.CreateIndex`を使用したとき，`Migrate`を呼び出すことでインデックスを作成することができます．

#### 例
```go
err := mgorm.CreateIndex(db, "IDX_emp_no").
    On("employees", "emp_no").Migrate()
// CREATE INDEX IDX_emp_no
//  ON employees (emp_no);

err := mgorm.CreateIndex(db, "IDX_emp_no").
    On("employees", "emp_no", "first_name").Migrate()
// CREATE INDEX IDX_emp_no
//  ON employees (emp_no, first_name);
```
