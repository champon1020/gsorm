# CreateDB
`CreateDB`はCREATE DATABASE句を呼び出します．

引数にはデータベースのコネクション(`mgorm.Conn`)，データベース名を指定します．

#### 例
```go
err := mgorm.CreateDB(db, "employees").Migrate()
// CREATE DATABASE employees;
```
