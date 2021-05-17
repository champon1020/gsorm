# DropDB
`mgorm.DropDB`はDROP DATABASE句を呼び出します．

引数にはデータベースのコネクション(`mgorm.Conn`)，データベース名を指定します．

#### 例
```go
err := mgorm.DropDB(db, "employees").Migrate()
// DROP DATABASE employees;
```
