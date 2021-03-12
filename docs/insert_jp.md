# Insert
`mgorm.Insert`を使用したとき，`Exec`を呼び出すことでテーブルにカラムを挿入することができます．

#### 例
```go
mgorm.Insert(db, "employees").
    Values(1000, "1996-03-09", "Taro", "Sato", "M", "2020-04-01").Exec()
// INSERT INTO employees
//  VALUES (1000, '1996-03-09', 'Taro', 'Sato', 'M', '2020-04-01');

mgorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1000, "Taro").Exec()
// INSERT INTO employees (emp_no, first_name)
//  VALUES (1000, 'Taro');
```


# Methods
`mgorm.Insert`で使用できるメソッドを以下に示します．

- [Values](https://github.com/champon1020/mgorm/tree/main/docs/insert_jp.md#values)
- [Select](https://github.com/champon1020/mgorm/tree/main/docs/insert_jp.md#select)
- [Model](https://github.com/champon1020/mgorm/tree/main/docs/insert_jp.md#model)

```
[]: optional, |: or, {}: block, **: able to use many times

mgorm.Insert(DB, table, columns...)
    {.Values(values...)** | .Select(*mgorm.SelectStmt) | .Model(model)}
    .Exec()
```


## Values
`mgorm.Insert`を用いてカラムを挿入するとき，`Values`を用いることで値を挿入することができます．
`Values`は連続で複数回使用することができます．

#### 例
```go
mgorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1000, "Taro").Exec()
// INSERT INTO employees (emp_no, first_name)
//  VALUES (1000, 'Taro');

mgorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1000, "Taro").Values(2000, "Jiro").Exec()
// INSERT INTO employees (emp_no, first_name)
//  VALUES (1000, 'Taro'), (2000, 'Jiro');
```


## Select
`Select`を用いることでINSERT INTO ... SELECTという文を実行することができます．
これは`mgorm.Select`とは異なる関数(メソッド)です．

`Select`は引数に`mgorm.SelectStmt`を受け取ります．

#### 例
```go
mgorm.Insert(db, "dept_manager").
    Select(mgorm.Select(nil).From("dept_emp")).Exec()
// INSERT INTO dept_manager
//  SELECT * FROM dept_emp;
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

mgorm.Insert(db, "employees", "emp_no", "first_name").
    Model(&employees).Exec()
// INSERT INTO employees (emp_no, first_name)
//  VALUES (1000, 'Taro'), (2000, 'Jiro');
```
