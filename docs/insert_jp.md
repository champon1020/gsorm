# Insert
`mgorm.Insert`を使用したとき，`Exec`を呼び出すことでテーブルにカラムを挿入することができます．

`mgorm.Insert`の第1引数は`mgorm.Conn`の型，第2引数はテーブル名をstring型として，第3引数以降は複数のカラム名をstring型として受け取ることができます．
`mgorm.Conn`を実装した型としては`*mgorm.DB`，`*mgorm.Tx`，`*mgorm.MockDB`，`*mgorm.MockTx`があります．

詳細は[Transaction]()，[Mock]()に記載されています．

#### 例
```go
// INSERT INTO people (id, name) VALUES (10, 'Taro');
mgorm.Insert(db, "people", "id", "name").Values(10, "Taro").Exec()
```
