# Insert
`mgorm.Insert`はINSERT句を呼び出します．

引数にはデータベースのコネクション(`mgorm.Conn`)，テーブル名，カラム名を指定します．

カラム名は複数指定することができます．
カラム名は空でも問題ありません．

#### 例
```go
err := mgorm.Insert(db, "employees").
    Values(1001, "1996-03-09", "Taro", "Sato", "M", "2020-04-01").Exec()
// INSERT INTO employees
//      VALUES (1001, '1996-03-09', 'Taro', 'Sato', 'M', '2020-04-01');

err := mgorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1001, "Taro").Exec()
// INSERT INTO employees (emp_no, first_name)
//      VALUES (1001, 'Taro');
```


# Methods
`mgorm.Insert`に使用できるメソッドは以下です．

- [Values](https://github.com/champon1020/mgorm/tree/main/docs/insert_ja.md#values)
- [Select](https://github.com/champon1020/mgorm/tree/main/docs/insert_ja.md#select)
- [Model](https://github.com/champon1020/mgorm/tree/main/docs/insert_ja.md#model)

これらのメソッドは以下のEBNFに従って実行することができます．

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

mgorm.Insert
    (.Values {.Values}) | .Select | .Model
    .Exec
```

例えば以下の実装はコンパイルエラーを吐き出します．

```go
// NG
err := mgorm.Insert(db).Exec()

// NG
err := mgorm.Insert(db).
    Model(model1).
    Model(model2).Exec()
```


## Values
`Values`はVALUES句を呼び出します．

引数には複数値を指定します．

`Values`は複数回呼び出すことができます．

#### 例
```go
err := mgorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1001, "Taro").Exec()
// INSERT INTO employees (emp_no, first_name)
//      VALUES (1001, 'Taro');

err := mgorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1001, "Taro").
    Values(1002, "Jiro").Exec()
// INSERT INTO employees (emp_no, first_name)
//      VALUES (1001, 'Taro'), (1002, 'Jiro');
```


## Select
`Select`はINSERT INTO ... SELECT構文を呼び出します．

引数には`mgorm.SelectStmt`を指定します．

#### 例
```go
err := mgorm.Insert(db, "dept_manager").
    Select(mgorm.Select(nil).From("dept_emp")).Exec()
// INSERT INTO dept_manager
//      SELECT * FROM dept_emp;
```


## Model
`Model`は構造体をマッピングします．

引数には構造体のポインタ，マップのポインタ，構造体のスライスのポインタ，マップのスライスのポインタ，事前定義型のスライスのポインタのいずれかを指定します．

構造体もしくは構造体のスライスをマッピングする際，対象のカラム名はフィールド名もしくはフィールドタグから推定されます．

Modelについての詳細は[Model](https://github.com/champon1020/mgorm/blob/main/docs/model_ja.md)に記載されています．

#### 例
```go
type Employee struct {
    ID        int    `mgorm:"emp_no"`
    FirstName string
}

employees := []Employee{{ID: 1001, FirstName: "Taro"}, {ID: 1002, FirstName: "Jiro"}}

err := mgorm.Insert(db, "employees", "emp_no", "first_name").
    Model(&employees).Exec()
// INSERT INTO employees (emp_no, first_name)
//  VALUES (1001, 'Taro'), (1002, 'Jiro');
```
