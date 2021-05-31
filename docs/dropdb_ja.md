# DropDB
`gsorm.DropDB`はDROP DATABASE文を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#DropDB.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#DropDB)

#### 例
```go
err := gsorm.DropDB(db, "employees").Migrate()
// DROP DATABASE employees;
```
