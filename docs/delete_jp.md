# Delete
`mgorm.Delete`を使用したとき，`Exec`メソッドを呼び出すことでカラムを消去することができます．

#### 例
```go
err := mgorm.Delete(db).From("employees").Exec()
// DELETE FROM employees;
```


# Methods
`mgorm.Delete`で使用できるメソッドを以下に示します．

- [From](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md#from)
- [Where](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md#where)
- [And / Or](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md#and--or)

```
[]: optional, |: or, **: able to used many times

mgorm.Delete(DB)
    .From(table)
    [.Where(expression, values...)]
    [.And(expression, values) | .Or(expression, values)]**
```

上の図において，上に行くほど実行優先度が高いです．
例えば，以下のようなことはできません．

```go
// NG
err := mgorm.Delete(db).
    Where("emp_no = ?", 10000).
    From("employees").Exec()
```

これに反した場合，コンパイルエラーを吐き出します．


## From
`From`は複数のテーブル名を受け取ります．
必要であれば，テーブル名にはエイリアスを含めることができます．

#### 例
```go
err := mgorm.Delete(db).From("employees").Exec()
// DELETE FROM employees;

err := mgorm.Delete(db).From("employees", "dept_emp").Exec()
// DELETE FROM employees, dept_emp;

err := mgorm.Delete(db).From("employees AS e").
    Where("e.emp_no = ?", 10000).Exec()
// DELETE FROM employees AS e
//  WHERE e.emp_no = 10000;
```


## Where
`Where`は引数に条件式を受け取ります．

詳しい使用方法は`mgorm.Select`における[Where]()に記載されています．

#### 例
```go
err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 10000).Exec()
// DELETE FROM employees
//  WHERE emp_no = 10000;
```


## And / Or
`And`，`Or`は引数に条件式を受け取ります．

詳しい使用方法は`mgorm.Select`における[And / Or]()に記載されています．

#### 例
```go
err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 10000).
    And("first_name = ? OR last_name = ?", "Taro", "Sato").Exec()
// DELETE FROM employees
//  WHERE emp_no = 10000
//  AND (first_name = 'Taro' AND last_name = 'Sato');

err := mgorm.Delete(db).From("employees").
    Where("emp_no > ?", 10000).
    And("emp_no < ? AND first_name = ?", 10000, "Taro").Exec()
// DELETE FROM employees
//  WHERE emp_no > 10000
//  OR (emp_no < 10000 AND first_name = 'Taro');
```
