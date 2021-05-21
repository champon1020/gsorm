# DropDB
`gsorm.DropDB`はDROP DATABASE句を呼び出します．

引数にはデータベースのコネクション(`gsorm.Conn`)，データベース名を指定します．

#### 例
```go
err := gsorm.DropDB(db, "employees").Migrate()
// DROP DATABASE employees;
```
