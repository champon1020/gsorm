# CreateIndex
`gsorm.CreateIndex` calls CREATE INDEX statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateIndex.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#CreateIndex)

#### Example
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
