# DropTable
`gsorm.DropTable` calls DROP TABLE statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#DropTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#DropTable)

#### Example
```go
err := gsorm.DropTable(db, "employees").Migrate()
// DROP TABLE employees;
```
