# Connection
`gsorm.Conn` is the interface related to the database connection.

`gsorm.Conn` is embeded into these interfaces.
- [gsorm.DB](https://github.com/champon1020/gsorm/tree/main/docs/connection.md#db)
- [gsorm.Tx](https://github.com/champon1020/gsorm/tree/main/docs/connection.md#tx)
- [gsorm.MockDB](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktx)
- [gsorm.MockTx](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktx)


## DB
`gsorm.DB` is the interface of database connection.

The instance that implements `gsorm.DB` interface can be created by `gsorm.Open`.

This is almost the same as `sql.Open`.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#DB.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#DB)

#### Example
```go
db, err := gsorm.Open("mysql", "root:toor@tcp(localhost:3306)/employees?parseTime=true")
if err != nil {
	log.Fatal(err)
}
```


## Tx
`gsorm.Tx` is the interface of database transaction.

The instance that implements `gsorm.Tx` interface can be created by `gsorm.DB.Begin` method.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Tx.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Tx)

#### Example
```go
tx, err := db.Begin()
if err != nil {
	log.Fatal(err)
}
```
