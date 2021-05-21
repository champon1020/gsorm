# DropTable
`gsorm.DropTable`はDROP TABLE句を呼び出します．

引数にはデータベースのコネクション(`gsorm.Conn`)，データベース名を指定します．

#### 例
```go
err := gsorm.DropTable(db, "employees").Migrate()
// DROP TABLE employees;
```
