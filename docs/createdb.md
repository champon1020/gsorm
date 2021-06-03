# CreateDB
`gsorm.CreateDB` calls CREATE DATABASE statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateDB.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#CreateDB)

#### Example
```go
err := gsorm.CreateDB(db, "employees").Migrate()
// CREATE DATABASE employees;
```
