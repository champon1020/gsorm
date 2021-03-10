# Update
`mgorm.Update`を使用したとき，`Exec`を呼び出すことでテーブル上のカラムを更新することができます．

`mgorm.Update`の第1引数は`mgorm.Conn`の型，第2引数はテーブル名をstring型として，第3引数以降は複数のカラム名をstring型として受け取ることができます．
カラム名を指定しない場合は，全てのカラムとして適用されます．

`mgorm.Conn`を実装した型としては`*mgorm.DB`，`*mgorm.Tx`，`*mgorm.MockDB`，`*mgorm.MockTx`があります．

詳細は[Transaction]()，[Mock]()に記載されています．

#### 例
```go
// UPDATE people SET id=10, name='Taro';
mgorm.Update(db, "people", "id", "name").Set(10, "Taro").Exec()
```


## Set
`mgorm.Update`を用いてカラムを更新するとき，`Set`を用いることで値を更新することができます．

#### 例
```go
// UPDATE people SET id=10, name='Taro';
mgorm.Update(db, "people", "id", "name").Set(10, "Taro").Exec()
```


## Where
`Where`は引数に条件式を受け取ります．

詳しい使用方法は`mgorm.Select`における[Where]()と同様です．

#### 例
```go
// UPDATE people SET name='Taro' WHERE id = 10;
mgorm.Update(db, "people", "name").Set("Taro").Where("id = ?", 10).Exec()
```
