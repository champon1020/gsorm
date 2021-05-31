# Delete
`gsorm.Delete` calls DELETE statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Delete.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Delete)

#### Example
```go
err := gsorm.Delete(db).From("employees").Exec()
// DELETE FROM employees;
```


## Methods
- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw.md#rawclause)
- [From](https://github.com/champon1020/gsorm/tree/main/docs/delete.md#from)
- [Where](https://github.com/champon1020/gsorm/tree/main/docs/delete.md#where)
- [And](https://github.com/champon1020/gsorm/tree/main/docs/delete.md#and)
- [Or](https://github.com/champon1020/gsorm/tree/main/docs/delete.md#or)

These methods can be executed according to the following EBNF.

Exceptionally, `RawClanuse` can be executed at any time.

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

For example, these implementations output the compile error.

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
`From` calls FROM clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Delete.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement#DeleteStmt.From)

#### Example
```go
err := gsorm.Delete(db).From("employees").Exec()
// DELETE FROM employees;

err := gsorm.Delete(db).From("employees", "dept_emp").Exec()
// DELETE FROM employees, dept_emp;

err := gsorm.Delete(db).From("employees AS e").Exec()
// DELETE FROM employees AS e;
```


## Where
`Where` calls WHERE clause.

When the query is executed, the values will be assigned to `?` in the expression.

Assignment rules are as follows:
- If the type of value is `string` or `time.Time`, the value is enclosed in single quotes
- If the value is slice or array, its elements are expanded
- If the type of value is `gsorm.Stmt`, `gsorm.Stmt` is built
- If the above conditions are not met, the value is assigned as is

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Delete.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement#DeleteStmt.Where)

#### Example
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
`And` calls AND clause.

The expression of AND clause is enclused in parenthesises.

When the query is executed, the values will be assigned to `?` in the expression.

Assignment rules are as follows:
- If the type of value is `string` or `time.Time`, the value is enclosed in single quotes
- If the value is slice or array, its elements are expanded
- If the type of value is `gsorm.Stmt`, `gsorm.Stmt` is built
- If the above conditions are not met, the value is assigned as is

`And` can be called mutliple times.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Delete.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement#DeleteStmt.And)

#### Example
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
`Or` calls OR clause.

When the query is executed, the values will be assigned to `?` in the expression.

Assignment rules are as follows:
- If the type of value is `string` or `time.Time`, the value is enclosed in single quotes
- If the value is slice or array, its elements are expanded
- If the type of value is `gsorm.Stmt`, `gsorm.Stmt` is built
- If the above conditions are not met, the value is assigned as is

`Or` can be called mutliple times.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Delete.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement#DeleteStmt.Or)

#### Example
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
