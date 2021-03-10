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


## Values
`mgorm.Insert`を用いてカラムを挿入する際，`Values`を用いることで値を挿入することができます．
`Values`は連続で複数回使用することができます．

#### 例
```go
// INSERT INTO people (id, name) VALUES (10, 'Taro');
mgorm.Insert(db, "people", "id", "name").Values(10, "Taro").Exec()

// INSERT INTO people (id, name) VALUES (10, 'Taro'), (20, 'Jiro');
mgorm.Insert(db, "people", "id", "name").Values(10, "Taro").Values(20, "Jiro").Exec()
```


## Select
`Select`を用いることでINSERT INTO ... SELECTという文を実行することができます．
これは`mgorm.Select`とは異なる関数(メソッド)です．

`Select`は引数に`mgorm.SelectStmt`を受け取ります．

#### 例
```go
// INSERT INTO people (id, name) SELECT (id, name) FROM others;
mgorm.Insert(db, "people", "id", "name").Select(mgorm.Select(nil, "id", "name").From("others")).Exec()
```
