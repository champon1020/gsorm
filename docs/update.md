# Update
`gsorm.Update` calls UPDATE statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Update.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Update)

#### Example
```go
gsorm.Update(db).Set(10, "employees").
    Set("emp_no", 1001).
    Set("birth_date", "1995-07-07").
    Set("first_name", "Hanako").
    Set("last_name", "Suzuki").
    Set("gender", "W").
    Set("hire_date", time.Date(2019, time.September, 1, 0, 0, 0, 0, time.UTC)).Exec()
// UPDATE employees
//      SET emp_no = 1001,
//          birth_date = '1995-07-07',
//          first_name = 'Hanako',
//          last_name = 'Suzuki',
//          gender = 'W',
//          hire_date = '2019-09-01';
```


## Methods
- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw_ja.md#rawclause)
- [Set](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#set)
- [Where](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#where)
- [And](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#and)
- [Or](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#or)
- [Model](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#model)

These methods can be executed according to the following EBNF.

Exceptionally, `RawClause` can be executed at any time.

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

gsorm.Update(DB, table, columns...)
    (.Set {.Set}) | .Model
    [.Where [{.And} | {.Or}]]
    .Exec
```

For example, these implementations output the compile error.

```go
// NG
err := gsorm.Update(db, "employees", "emp_no", "first_name").Exec()

// NG
err := gsorm.Update(db, "employees", "emp_no", "first_name").
    Set("emp_no", 1001).
    Set("first_name", "Hanako").
    And("emp_no < ? AND first_name = ?", 1000, "Taro")
    Where("emp_no > ?", 1000).Exec()
```


## Set
`Set` calls SET clause.

`Set` can be called multiple times.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Update.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement#UpdateStmt.Set)

#### Example
```go
err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").Exec()
// UPDATE employees
//      SET first_name = 'Hanako';

err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").
    Set("last_name", "Suzuki").Exec()
// UPDATE employees
//      SET first_name = 'Hanako',
//          last_name = 'Suzuki';
```


## Where
`Where` calls WHERE clause.

When the query is executed, the values will be assigned to `?` in the expression.

Assignment rules are as follows:
- If the type of value is `string` or `time.Time`, the value is enclosed in single quotes
- If the value is slice or array, its elements are expanded
- If the type of value is `gsorm.Stmt`, `gsorm.Stmt` is built
- If the above conditions are not met, the value is assigned as is

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Update.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement#UpdateStmt.Where)

#### Example
```go
err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").
    Where("emp_no = 1001").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001;

err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001;

err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").
    Where("first_name = ?", "Taro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE first_name = 'Taro';

err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").
    Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE birth_date = '2006-01-02 00:00:00';

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("first_name LIKE ?", "%Taro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE first_name LIKE '%Taro';

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no BETWEEN 1001 AND 1003;

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no IN (?)", []int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no IN (1001, 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no IN (1001, 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
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

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Update.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement#UpdateStmt.And)

#### Example
```go
err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no = 1002").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no = ?", 1002).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("first_name = ? OR first_name = ?", "Taro", "Jiro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (first_name = 'Taro' OR first_name = 'Jiro');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no = ?", 1002).
    And("emp_no = ?", 1003).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);
//      AND (emp_no = 1003);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (birth_date = '2006-01-02 00:00:00');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("first_name LIKE ?", "%Taro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (first_name LIKE '%Taro');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no BETWEEN 1001 AND 1003);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", []int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no IN (1001, 1002));

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no IN (1001, 1002));

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
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

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Update.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement#UpdateStmt.Or)

#### Example
```go
err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no = 1002").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no = ? AND first_name = ?", 1002, "Taro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no = 1002 OR first_name = 'Taro');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).
    Or("emp_no = ?", 1003).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);
//      OR (emp_no = 1003);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (birth_date = '2006-01-02 00:00:00');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("first_name LIKE ?", "%Taro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (first_name LIKE '%Taro');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no BETWEEN 1001 AND 1003);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", []int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no IN (1001, 1002));

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no IN (1001, 1002));

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no IN (SELECT emp_no FROM dept_manager));
```


## Model
`Model` maps the model into SQL.

Details are given in [Model](https://github.com/champon1020/gsorm/blob/main/docs/model.md).

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Update.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement#UpdateStmt.Model)

#### Example
```go
type Employee struct {
    ID        int       `gsorm:"emp_no"`
    BirthDate time.Time
    FirstName string
    LastName  string
    Gender    string
    HireDate  string
}

emp1 := Employee{ID: 1000, FirstName: "Taro"}

gsorm.Update(db, "employees").
    Model(&emp1, "emp_no", "first_name").Exec()
// UPDATE employees
//  SET emp_no = 1000,
//      first_name = 'Taro';

emp2 = Employee{
    EmpNo: 1000,
    BirthDate: time.Date(1965, time.April, 4, 0, 0, 0, 0, time.UTC),
    FirstName: "Taro",
    LastName: "Sato",
    Gender: "M",
    HireDate: "1988-04-01",
}

gsorm.Update(db, "employees").
    Model(&emp2).Exec()
// UPDATE employees
//  SET emp_no = 1000,
//      birth_date = '1965-04-04 00:00:00'
//      first_name = 'Taro',
//      last_name = 'Sato',
//      gender = 'M',
//      hire_date = '1988-04-01';
```
