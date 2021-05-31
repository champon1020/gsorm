# CreateIndex
`gsorm.CreateIndex`はCREATE INDEX文を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateIndex.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#CreateIndex)

#### 例
```go
err := gsorm.CreateIndex(db, "IDX_emp_no").
    On("employees", "emp_no").Migrate()
// CREATE INDEX IDX_emp
//  ON employees (emp_no);

err := gsorm.CreateIndex(db, "IDX_emp_no").
    On("employees", "emp_no", "first_name").Migrate()
// CREATE INDEX IDX_emp
//  ON employees (emp_no, first_name);
```
