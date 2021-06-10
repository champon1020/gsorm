# Delete
`gsorm.Delete`はDELETE句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Delete.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Delete)

#### 例
```go
err := gsorm.Delete(db).From("employees").Exec()
// DELETE FROM employees;
```


## Methods
- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw_ja.md#rawclause)
- [From](https://github.com/champon1020/gsorm/tree/main/docs/delete_ja.md#from)
- [Where](https://github.com/champon1020/gsorm/tree/main/docs/delete_ja.md#where)
- [And](https://github.com/champon1020/gsorm/tree/main/docs/delete_ja.md#and)
- [Or](https://github.com/champon1020/gsorm/tree/main/docs/delete_ja.md#or)

これらのメソッドは以下のEBNFに従って実行することができます．
但し，例外として`RawClause`は任意で呼び出すことができます．

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

gsorm.Delete
    .From
    [.Where [{.And} | {.Or}]]
    .Exec
```

例えば以下の実装はコンパイルエラーを吐き出します．

```go
// NG
err := gsorm.Delete(db).
    Where("emp_no = ?", 1001).
    From("employees").Exec()

// NG
err := gsorm.Delete(db).
    From("employees").
    Where("emp_no = ?", 1001).
    Where("emp_no = ?", 1002).Exec()
```


## From
`From`はFROM句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Delete.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#DeleteStmt.From)

#### 例
```go
err := gsorm.Delete(db).From("employees").Exec()
// DELETE FROM employees;

err := gsorm.Delete(db).From("employees", "dept_emp").Exec()
// DELETE FROM employees, dept_emp;

err := gsorm.Delete(db).From("employees AS e").Exec()
// DELETE FROM employees AS e;
```


## Where
`Where`はWHERE句を呼び出します．

クエリが実行されるとき，条件式における`?`に値が代入されます．

代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値がスライスもしくは配列の場合，その要素が展開されます．
- 値が`gsorm.Stmt`型の場合，`gsorm.Stmt`は展開されます．
- 以上の条件に該当しない値はそのまま展開される．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Delete.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#DeleteStmt.Where)

#### 例
```go
err := gsorm.Delete(db).From("employees").
    Where("emp_no = 1001").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001;

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001;

err := gsorm.Delete(db).From("employees").
    Where("first_name = ?", "Taro").Exec()
// DELETE FROM employees
//      WHERE first_name = 'Taro';

err := gsorm.Delete(db).From("employees").
    Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// DELETE FROM employees
//      WHERE birth_date = '2006-01-02 00:00:00';

err := gsorm.Delete(db).From("employees").
    Where("first_name LIKE ?", "%Taro").Exec()
// DELETE FROM employees
//      WHERE first_name LIKE '%Taro';

err := gsorm.Delete(db).From("employees").
    Where("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no BETWEEN 1001 AND 1003;

err := gsorm.Delete(db).From("employees").
    Where("emp_no IN (?)", []int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no IN (1001, 1002);

err := gsorm.Delete(db).From("employees").
    Where("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no IN (1001, 1002);

err := gsorm.Delete(db).From("employees").
    Where("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// DELETE FROM employees
//      WHERE emp_no IN (SELECT emp_no FROM dept_manager);
```


## And
`And`はAND句を呼び出します．

このときAND句は条件式は`()`で括られます．

クエリが実行されるとき，条件式における`?`に値が代入されます．

代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値がスライスもしくは配列の場合，その要素が展開されます．
- 値が`gsorm.Stmt`型の場合，`gsorm.Stmt`は展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`And`は複数回呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Delete.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#DeleteStmt.And)

#### 例
```go
err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no = 1002").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no = ?", 1002).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("first_name = ? OR first_name = ?", "Taro", "Jiro").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (first_name = 'Taro' OR first_name = 'Jiro');

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no = ?", 1002).
    And("emp_no = ?", 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);
//      AND (emp_no = 1003);

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (birth_date = '2006-01-02 00:00:00');

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("first_name LIKE ?", "%Taro").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (first_name LIKE '%Taro');

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no BETWEEN 1001 AND 1003);

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", []int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no IN (1001, 1002));

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no IN (1001, 1002));

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no IN (SELECT emp_no FROM dept_manager));
```


## Or
`Or`はOR句を呼び出します．

このときOR句は条件式は`()`で括られます．

クエリが実行されるとき，条件式における`?`に値が代入されます．

代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値がスライスもしくは配列の場合，その要素が展開されます．
- 値が`gsorm.Stmt`型の場合，`gsorm.Stmt`は展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`Or`は複数回呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Delete.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#DeleteStmt.Or)

#### 例
```go
err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = 1002").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).Exec()
// DELETE FROM employees
//  WHERE emp_no = 1001
//  OR (emp_no = 1002);

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ? AND first_name = ?", 1002, "Taro").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002 AND first_name = 'Taro');

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).
    Or("emp_no = ?", 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002)
//      OR (emp_no = 1003);

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (birth_date = '2006-01-02 00:00:00');

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("first_name LIKE ?", "%Taro").Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (first_name LIKE '%Taro');

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no BETWEEN 1001 AND 1003);

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", []int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no IN (1001, 1002));

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no IN (1001, 1002));

err := gsorm.Delete(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// DELETE FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no IN (SELECT emp_no FROM dept_manager));
```
