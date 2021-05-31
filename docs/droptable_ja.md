# DropTable
`gsorm.DropTable`はDROP TABLE文を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#DropTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#DropTable)

#### 例
```go
err := gsorm.DropTable(db, "employees").Migrate()
// DROP TABLE employees;
```
