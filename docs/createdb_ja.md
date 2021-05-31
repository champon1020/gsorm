# CreateDB
`CreateDB`はCREATE DATABASE文を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateDB.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#CreateDB)

#### 例
```go
err := gsorm.CreateDB(db, "employees").Migrate()
// CREATE DATABASE employees;
```
