# Connection
`gsorm.Conn`はデータベースコネクションに関するインタフェースです．

`gsorm.Conn`は以下のインタフェースに埋め込まれています．
- [gsorm.DB](https://github.com/champon1020/gsorm/tree/main/docs/connection_ja.md#db)
- [gsorm.Tx](https://github.com/champon1020/gsorm/tree/main/docs/connection_ja.md#tx)
- [gsorm.MockDB](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mocktx)
- [gsorm.MockTx](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mocktx)


# DB
`gsorm.DB`を実装する構造体は`gsorm.Open`関数で生成されます．

第一引数にはデータベースドライバー名，第二引数にはDSNをを指定します．

使用方法は`sql.Open`とほとんど変わりません．

#### 例
```go
db, err := gsorm.Open("mysql", "root:toor@tcp(localhost:3306)/employees?parseTime=true")
if err != nil {
	log.Fatal(err)
}
```


## Tx
`gsorm.Tx`を実装する構造体は`gsorm.DB`の`Begin`メソッドによって生成されます．

#### 例
```go
tx, err := db.Begin()
if err != nil {
	log.Fatal(err)
}
```
