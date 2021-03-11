# Insert
`mgorm.Insert`を使用したとき，`Exec`を呼び出すことでテーブルにカラムを挿入することができます．

`mgorm.Insert`の第1引数は`mgorm.Conn`の型，第2引数はテーブル名をstring型として，第3引数以降は複数のカラム名をstring型として受け取ることができます．
カラム名を指定しない場合は，全てのカラムとして適用されます．

`mgorm.Conn`を実装した型としては`*mgorm.DB`，`*mgorm.Tx`，`*mgorm.MockDB`，`*mgorm.MockTx`があります．

詳細は[Transaction]()，[Mock]()に記載されています．

#### 例
```go
// INSERT INTO employees
//  VALUES (1000, '1996-03-09', 'Taro', 'Sato', 'M', '2020-04-01');
mgorm.Insert(db, "employees").
    Values(1000, "1996-03-09", "Taro", "Sato", "M", "2020-04-01").Exec()

// INSERT INTO employees (emp_no, first_name)
//  VALUES (1000, 'Taro');
mgorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1000, "Taro").Exec()
```


## Values
`mgorm.Insert`を用いてカラムを挿入するとき，`Values`を用いることで値を挿入することができます．
`Values`は連続で複数回使用することができます．

#### 例
```go
// INSERT INTO employees (emp_no, first_name)
//  VALUES (1000, 'Taro');
mgorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1000, "Taro").Exec()

// INSERT INTO employees (emp_no, first_name)
//  VALUES (1000, 'Taro'), (2000, 'Jiro');
mgorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1000, "Taro").Values(2000, "Jiro").Exec()
```


## Select
`Select`を用いることでINSERT INTO ... SELECTという文を実行することができます．
これは`mgorm.Select`とは異なる関数(メソッド)です．

`Select`は引数に`mgorm.SelectStmt`を受け取ります．

#### 例
```go
// INSERT INTO dept_manager
//  SELECT * FROM dept_emp;
mgorm.Insert(db, "dept_manager").
    Select(mgorm.Select(nil).From("dept_emp")).Exec()
```


## Model
`mgorm.Insert`を使用すとき，`Model`を使用することで構造体をマッピングしてカラムを挿入することができます．

`Model`は引数として構造体のポインタ，構造体スライスのポインタ，マップ型のポインタなどを受け取ることができます．

また，フィールドタグを変更することで対応するカラム名を変更することができます．
指定しない場合は，フィールド名のスネークケースとなります．

Modelの型やタグについての詳細は[Model]()に記載されています．

#### 例
```go
type Employee struct {
    ID        int    `mgorm:"emp_no"`
    FirstName string
}

employees := []Employee{{ID: 1000, FirstName: "Taro"}, {ID: 2000, FirstName: "Jiro"}}

// INSERT INTO employees (emp_no, first_name) VALUES (1000, 'Taro'), (2000, 'Jiro');
mgorm.Insert(db, "employees", "emp_no", "first_name").Model(&employees).Exec()
```
