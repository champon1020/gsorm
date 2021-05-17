# DropTable
`mgorm.DropTable`はDROP TABLE句を呼び出します．

引数にはデータベースのコネクション(`mgorm.Conn`)，データベース名を指定します．

#### 例
```go
err := mgorm.DropTable(db, "employees").Migrate()
// DROP TABLE employees;
```
