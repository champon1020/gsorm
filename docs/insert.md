# Insert
`gsorm.Insert` calls INSERT statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Insert.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Insert)

#### Example
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

These methods can be executed according to the following EBNF.

Exceptionally, `RawClause` can be executed at any time.

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

gsorm.Insert
    (.Values {.Values}) | .Select | .Model
    .Exec
```

For example, these implementations output the compile error.

```go
// NG
err := gsorm.Insert(db).Exec()

// NG
err := gsorm.Insert(db).
    Model(model1).
    Model(model2).Exec()
```


## Values
`Values` calls VALUES clause.

`Values` can be called mutiple times.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Insert.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#InsertStmt.Values)

#### Example
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
`Select` calls INSERT INTO ... SELECT statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Insert.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#InsertStmt.Select)

#### Example
```go
err := gsorm.Insert(db, "dept_manager").
    Select(gsorm.Select(nil).From("dept_emp")).Exec()
// INSERT INTO dept_manager
//      SELECT * FROM dept_emp;
```


## Model
`Model` maps the model into SQL.

Details are given in[Model](https://github.com/champon1020/gsorm/blob/main/docs/model.md).

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Insert.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#InsertStmt.Model)

#### Example
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
