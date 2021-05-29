# DropDB
`gsorm.DropDB` calls DROP DATABASE statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#DropDB.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#DropDB)

#### Example
```go
err := gsorm.DropDB(db, "employees").Migrate()
// DROP DATABASE employees;
```
