# CreateDB
`CreateDB`を使用したとき，`Migrate`を呼び出すことでデータベースを作成することができます．

#### 例
```go
err := mgorm.CreateDB(db, "employees").Migrate()
// CREATE DATABASE employees;
```
