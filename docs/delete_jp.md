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

第1引数に条件式，第2引数以降に複数値を指定できます． この際，条件式における`?`に値が代入されます． また，代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が展開されます．
- 値が`*mgorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない値はそのまま展開される．

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

err := mgorm.Delete(db).From("employees").
    Where("first_name = ?", "Taro").Exec()
// DELETE FROM employees
//      WHERE first_name = 'Taro';

err := mgorm.Delete(db).From("employees").
    Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// DELETE FROM employees
//      WHERE birth_date = '2006-01-02 00:00:00';

err := mgorm.Delete(db).From("employees").
    Where("first_name LIKE ?", "%Taro").Exec()
// DELETE FROM employees
//      WHERE first_name LIKE '%Taro';

err := mgorm.Delete(db).From("employees").
    Where("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no BETWEEN 1001 AND 1003;

err := mgorm.Delete(db).From("employees").
    Where("emp_no IN (?)", []int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no IN (1001, 1002);

err := mgorm.Delete(db).From("employees").
    Where("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no IN (1001, 1002);

err := mgorm.Delete(db).From("employees").
    Where("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// DELETE FROM employees
//      WHERE emp_no IN (SELECT emp_no FROM dept_manager);
```


## And
`And`はAND句を呼び出します．
このとき実行されるSQLは，条件式が`()`で括られた形となります．

第1引数に条件式，第2引数以降に複数値を指定できます． この際，条件式における`?`に値が代入されます． また，代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が展開されます．
- 値が`*mgorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`And`は複数回呼び出すことができます．

#### 例
```go
err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no = 1002").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no = ?", 1002).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("first_name = ? OR first_name = ?", "Taro", "Jiro").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (first_name = 'Taro' OR first_name = 'Jiro');

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no = ?", 1002).
    And("emp_no = ?", 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);
//      AND (emp_no = 1003);

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (birth_date = '2006-01-02 00:00:00');

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("first_name LIKE ?", "%Taro").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (first_name LIKE '%Taro');

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no BETWEEN 1001 AND 1003);

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", []int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no IN (1001, 1002));

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no IN (1001, 1002));

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no IN (SELECT emp_no FROM dept_manager));
```


## Or
`Or`はOR句を呼び出します．
このとき実行されるSQLは，条件式が`()`で括られた形となります．

第1引数に条件式，第2引数以降に複数値を指定できます． この際，条件式における`?`に値が代入されます． また，代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が展開されます．
- 値が`*mgorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`Or`は複数回呼び出すことができます．

#### 例
```go
err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = 1002").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).Exec()
// DELETE FROM employees
//  WHERE emp_no = 1001
//  OR (emp_no = 1002);

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ? AND first_name = ?", 1002, "Taro").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002 AND first_name = 'Taro');

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).
    Or("emp_no = ?", 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002)
//      OR (emp_no = 1003);

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (birth_date = '2006-01-02 00:00:00');

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("first_name LIKE ?", "%Taro").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (first_name LIKE '%Taro');

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no BETWEEN 1001 AND 1003);

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", []int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no IN (1001, 1002));

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no IN (1001, 1002));

err := mgorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no IN (SELECT emp_no FROM dept_manager));
```
