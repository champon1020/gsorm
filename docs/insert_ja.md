# Insert
`gsorm.Insert`はINSERT句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Insert.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Insert)

#### 例
```go
err := gsorm.Insert(db, "employees").
    Values(1001, "1996-03-09", "Taro", "Sato", "M", "2020-04-01").Exec()
// INSERT INTO employees
//      VALUES (1001, '1996-03-09', 'Taro', 'Sato', 'M', '2020-04-01');

err := gsorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1001, "Taro").Exec()
// INSERT INTO employees (emp_no, first_name)
//      VALUES (1001, 'Taro');
```


## Methods
- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw_ja.md#rawclause)
- [Values](https://github.com/champon1020/gsorm/tree/main/docs/insert_ja.md#values)
- [Select](https://github.com/champon1020/gsorm/tree/main/docs/insert_ja.md#select)
- [Model](https://github.com/champon1020/gsorm/tree/main/docs/insert_ja.md#model)

これらのメソッドは以下のEBNFに従って実行することができます．

但し，例外として`RawClause`は任意で呼び出すことができます．

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

gsorm.Insert
    (.Values {.Values}) | .Select | .Model
    .Exec
```

例えば以下の実装はコンパイルエラーを吐き出します．

```go
// NG
err := gsorm.Insert(db).Exec()

// NG
err := gsorm.Insert(db).
    Model(model1).
    Model(model2).Exec()
```


## Values
`Values`はVALUES句を呼び出します．

`Values`は複数回呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Insert.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#InsertStmt.Values)

#### 例
```go
err := gsorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1001, "Taro").Exec()
// INSERT INTO employees (emp_no, first_name)
//      VALUES (1001, 'Taro');

err := gsorm.Insert(db, "employees", "emp_no", "first_name").
    Values(1001, "Taro").
    Values(1002, "Jiro").Exec()
// INSERT INTO employees (emp_no, first_name)
//      VALUES (1001, 'Taro'), (1002, 'Jiro');
```


## Select
`Select`はINSERT INTO ... SELECT文を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Insert.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#InsertStmt.Select)

#### 例
```go
err := gsorm.Insert(db, "dept_manager").
    Select(gsorm.Select(nil).From("dept_emp")).Exec()
// INSERT INTO dept_manager
//      SELECT * FROM dept_emp;
```


## Model
`Model`は構造体をマッピングします．

Modelについての詳細は[Model](https://github.com/champon1020/gsorm/blob/main/docs/model_ja.md)に記載されています．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Insert.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#InsertStmt.Model)

#### 例
```go
type Employee struct {
    ID        int    `gsorm:"emp_no"`
    FirstName string
}

employees := []Employee{{ID: 1001, FirstName: "Taro"}, {ID: 1002, FirstName: "Jiro"}}

err := gsorm.Insert(db, "employees", "emp_no", "first_name").
    Model(&employees).Exec()
// INSERT INTO employees (emp_no, first_name)
//  VALUES (1001, 'Taro'), (1002, 'Jiro');
```
