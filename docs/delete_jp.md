# Delete
`mgorm.Delete`はDELETE句を呼び出します．

引数にはデータベースのコネクション(`mgorm.Conn`)を指定します．

#### 例
```go
err := mgorm.Delete(db).From("employees").Exec()
// DELETE FROM employees;
```


# Methods
`mgorm.Delete`に使用できるメソッドは以下です．

- [From](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md#from)
- [Where](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md#where)
- [And](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md#and)
- [Or](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md#or)

これらのメソッドは以下のEBNFに従って実行することができます．

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

mgorm.Delete
    .From
    [.Where [{.And} | {.Or}]]
    .Exec
```

例えば以下の実装はコンパイルエラーを吐き出します．

```go
// NG
err := mgorm.Delete(db).
    Where("emp_no = ?", 1001).
    From("employees").Exec()

// NG
err := mgorm.Delete(db).
    From("employees").
    Where("emp_no = ?", 1001).
    Where("emp_no = ?", 1002).Exec()
```


## From
`From`はFROM句を呼び出します．

引数には複数テーブルを指定できます．
また，引数のテーブル名にエイリアスを含めることができます．

#### 例
```go
err := mgorm.Delete(db).From("employees").Exec()
// DELETE FROM employees;

err := mgorm.Delete(db).From("employees", "dept_emp").Exec()
// DELETE FROM employees, dept_emp;

err := mgorm.Delete(db).From("employees AS e").Exec()
// DELETE FROM employees AS e;
```


## Where
`Where`はWHERE句を呼び出します．

第1引数に条件式，第2引数に値を指定できます．
この際，条件式における`?`に値が代入されます．

#### 例
```go
err := mgorm.Delete(db).From("employees").
    Where("emp_no = 1001").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001;

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001;
```


## And
`And`はAND句を呼び出します．
このとき実行されるSQLは，条件式が`()`で括られた形となります．

第1引数に条件式，第2引数に値を指定できます．
この際，条件式における`?`に値が代入されます．

`And`は複数回呼び出すことができます．

#### 例
```go
err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("first_name = ?", "Taro").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (first_name = 'Taro');

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("first_name = ? OR last_name = ?", "Taro", "Sato").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (first_name = 'Taro' OR last_name = 'Sato');

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("first_name = ?", "Taro").
    And("last_name = ?", "Sato").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (first_name = 'Taro');
//      AND (last_name = 'Sato');
```


## Or
`Or`はOR句を呼び出します．
このとき実行されるSQLは，条件式が`()`で括られた形となります．

第1引数に条件式，第2引数に値を指定できます．
この際，条件式における`?`に値が代入されます．

`Or`は複数回呼び出すことができます．

#### 例
```go
err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ? AND first_name = ?", 1002, "Taro").Exec()
// DELETE FROM employees
//  WHERE emp_no = 1001
//  OR (emp_no = 1002 AND first_name = 'Taro');

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).
    Or("emp_no = ?", 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002)
//      OR (emp_no = 1003);
```
