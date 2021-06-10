# Select
`gsorm.Select` calls SELECT statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Select)

#### Example
```go
err := gsorm.Select(db, "emp_no").From("employees").Query(&model)
// SELECT emp_no FROM people;

err := gsorm.Select(db).From("employees").Query(&model)
// SELECT * FROM people;

err := gsorm.Select(db, "emp_no", "first_name").From("employees").Query(&model)
// SELECT emp_no, first_name FROM people;

err := gsorm.Select(db, "emp_no, first_name").From("employees").Query(&model)
// SELECT emp_no, first_name FROM people;

err := gsorm.Select(db, "emp_no, first_name", "last_name").From("employees").Query(&model)
// SELECT emp_no, first_name, last_name FROM people;
```


## Methods
- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw.md#rawclause)
- [From](https://github.com/champon1020/gsorm/tree/main/docs/select.md#from)
- [Join](https://github.com/champon1020/gsorm/tree/main/docs/select.md#join)
- [LeftJoin](https://github.com/champon1020/gsorm/tree/main/docs/select.md#leftjoin)
- [RightJoin](https://github.com/champon1020/gsorm/tree/main/docs/select.md#rightjoin)
- [Where](https://github.com/champon1020/gsorm/tree/main/docs/select.md#where)
- [And](https://github.com/champon1020/gsorm/tree/main/docs/select.md#and)
- [Or](https://github.com/champon1020/gsorm/tree/main/docs/select.md#or)
- [GroupBy](https://github.com/champon1020/gsorm/tree/main/docs/select.md#groupby)
- [Having](https://github.com/champon1020/gsorm/tree/main/docs/select.md#having)
- [Union](https://github.com/champon1020/gsorm/tree/main/docs/select.md#union)
- [UnionAll](https://github.com/champon1020/gsorm/tree/main/docs/select.md#unionall)
- [OrderBy](https://github.com/champon1020/gsorm/tree/main/docs/select.md#orderby)
- [Limit](https://github.com/champon1020/gsorm/tree/main/docs/select.md#limit)
- [Offset](https://github.com/champon1020/gsorm/tree/main/docs/select.md#offset)
- [Query](https://github.com/champon1020/gsorm/tree/main/docs/select.md#query)

These methods is executed according to the following EBNF.

Exceptionally, `RawClause` can be executed at any time.

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

gsorm.Select
    .From
    [(.Join | .LeftJoin | .RightJoin) .On {(.Join | .LeftJoin | .RightJoin) .On}]
    [.Where [{.And} | {.Or}]]
    [.GroupBy]
    [.Having]
    [.Union | .UnionAll]
    [.OrderBy]
    [.Limit [.Offset]]
    .Query
```

For example, these implementations will output the compile error.

```go
// NG
err := gsorm.Select(db).
    Where("emp_no = ?", 10000).
    From("employees").Query(&model)

// NG
err := gsorm.Select(db).
    Join("dept_manager AS d").Query(&model)
```


## From
`From` calls FROM clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.From)

#### Example
```go
err := gsorm.Select(db, "emp_no").From("employees").Query(&model)
// SELECT emp_no FROM employees;

err := gsorm.Select(db, "e.emp_no").From("employees AS e").Query(&model)
// SELECT e.emp_no FROM employees AS e;

err := gsorm.Select(db, "e.emp_no").From("employees as e").Query(&model)
// SELECT e.emp_no FROM employees AS e;

err := gsorm.Select(db, "emp_no", "dept_no").From("employees", "departments").Query(&model)
// SELECT emp_no, dept_no FROM employees, departments;
```


## Join
`Join` calls INNER JOIN clause.

`Join`, `LeftJoin` and `RightJoin` can be called multiple times.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Join)

#### Example
```go
err := gsorm.Select(db, "e.emp_no", "d.dept_no").
    From("employees AS e").
    Join("dept_manager AS d").
    On("e.emp_no = d.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no FROM employees AS e
//      INNER JOIN dept_manager AS d
//      ON e.emp_no = d.emp_no;

err := gsorm.Select(db, "e.emp_no", "d.dept_no", "s.salary").
    From("employees AS e").
    Join("dept_manager AS d").On("e.emp_no = d.emp_no").
    LeftJoin("salaries AS s").On("e.emp_no = s.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no, s.salary FROM employees AS e
//      INNER JOIN dept_manager AS d ON e.emp_no = d.emp_no;
//      LEFT  JOIN salaries     AS s ON e.emp_no = s.emp_no;
```


## LeftJoin
`LeftJoin` calls LEFT JOIN clause.

`Join`, `LeftJoin` and `RightJoin` can be called multiple times.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.LeftJoin)

#### Example
```go
err := gsorm.Select(db, "e.emp_no", "d.dept_no").
    From("employees AS e").
    LeftJoin("dept_manager AS d").
    On("e.emp_no = d.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no FROM employees AS e
//      LEFT JOIN dept_manager AS d
//      ON e.emp_no = d.emp_no;

err := gsorm.Select(db, "e.emp_no", "d.dept_no", "s.salary").
    From("employees AS e").
    LeftJoin("dept_manager AS d").On("e.emp_no = d.emp_no").
    RightJoin("salaries AS s").On("e.emp_no = s.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no, s.salary FROM employees AS e
//      LEFT  JOIN dept_manager AS d ON e.emp_no = d.emp_no;
//      RIGHT JOIN salaries     AS s ON e.emp_no = s.emp_no;
```


## RightJoin
`RightJoin` calls RIGHT JOIN clause.

`Join`, `LeftJoin` and `RightJoin` can be called multiple times.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.RightJoin)

#### Example
```go
err := gsorm.Select(db, "e.emp_no", "d.dept_no").
    From("employees AS e").
    RightJoin("dept_manager AS d").
    On("e.emp_no = d.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no FROM employees AS e
//      RIGHT JOIN dept_manager AS d
//      ON e.emp_no = d.emp_no;

err := gsorm.Select(db, "e.emp_no", "d.dept_no", "s.salary").
    From("employees AS e").
    RightJoin("dept_manager AS d").On("e.emp_no = d.emp_no").
    Join("salaries AS s").On("e.emp_no = s.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no, s.salary FROM employees AS e
//      RIGHT JOIN dept_manager AS d ON e.emp_no = d.emp_no;
//      INNER JOIN salaries     AS s ON e.emp_no = s.emp_no;
```


## Where
`Where` calls WHERE clause.

When the query is executed, the values will be assigned to `?` in the expression.

Assignment rules are as follows:
- If the type of value is `string` or `time.Time`, the value is enclosed in single quotes
- If the value is slice or array, its elements are expanded
- If the type of value is `gsorm.Stmt`, `gsorm.Stmt` is built
- If the above conditions are not met, the value is assigned as is

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Where)

#### Example
```go
err := gsorm.Select(db).From("employees").
    Where("emp_no = 1001").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001;

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001;

err := gsorm.Select(db).From("employees").
    Where("first_name = ?", "Taro").Query(&model)
// SELECT * FROM employees
//      WHERE first_name = 'Taro';

err := gsorm.Select(db).From("employees").
    Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Query(&model)
// SELECT * FROM employees
//      WHERE birth_date = '2006-01-02 00:00:00';

err := gsorm.Select(db).From("employees").
    Where("first_name LIKE ?", "%Taro").Query(&model)
// SELECT * FROM employees
//      WHERE first_name LIKE '%Taro';

err := gsorm.Select(db).From("employees").
    Where("emp_no BETWEEN ? AND ?", 1001, 1003).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no BETWEEN 1001 AND 1003;

err := gsorm.Select(db).From("employees").
    Where("emp_no IN (?)", []int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no IN (1001, 1002);

err := gsorm.Select(db).From("employees").
    Where("emp_no IN (?)", [2]int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no IN (1001, 1002);

err := gsorm.Select(db).From("employees").
    Where("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Query(&model)
// SELECT * FROM employees
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

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.And)

#### Example
```go
err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no = 1002").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no = ?", 1002).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    And("first_name = ? OR first_name = ?", "Taro", "Jiro").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND (first_name = 'Taro' OR first_name = 'Jiro');

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no = ?", 1002).
    And("emp_no = ?", 1003).Exec()
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);
//      AND (emp_no = 1003);

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND (birth_date = '2006-01-02 00:00:00');

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    And("first_name LIKE ?", "%Taro").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND (first_name LIKE '%Taro');

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no BETWEEN ? AND ?", 1001, 1003).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no BETWEEN 1001 AND 1003);

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", []int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no IN (1001, 1002));

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", [2]int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND (emp_no IN (1001, 1002));

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Query(&model)
// SELECT * FROM employees
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

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Or)

#### Example
```go
err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = 1002").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ? AND first_name = ?", 1002, "Taro").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002 AND first_name = 'Taro');

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).
    Or("emp_no = ?", 1003).Exec()
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no = 1002)
//      OR (emp_no = 1003);

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR (birth_date = '2006-01-02 00:00:00');

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("first_name LIKE ?", "%Taro").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR (first_name LIKE '%Taro');

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no BETWEEN ? AND ?", 1001, 1003).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no BETWEEN 1001 AND 1003);

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", []int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no IN (1001, 1002));

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", [2]int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no IN (1001, 1002));

err := gsorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR (emp_no IN (SELECT emp_no FROM dept_manager));
```


## GroupBy
`GroupBy` calls GROUP BY clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.GroupBy)

#### Example
```go
err := gsorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").Query(&model)
// SELECT emp_no, AVG(salary) FROM salaries
//      GROUP BY emp_no;
```


## Having
`Having` calls HAVING clause.

When the query is executed, the values will be assigned to `?` in the expression.

Assignment rules are as follows:
- If the type of value is `string` or `time.Time`, the value is enclosed in single quotes
- If the value is slice or array, its elements are expanded
- If the type of value is `gsorm.Stmt`, `gsorm.Stmt` is built
- If the above conditions are not met, the value is assigned as is

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Having)

#### Example
```go
err := gsorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").
    Having("AVG(salary) > 130000").Query(&model)
// SELECT emp_no, AVG(salary) FROM salaries
//      GROUP BY emp_no
//      HAVING AVG(salary) > 130000;

err := gsorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").
    Having("AVG(salary) > ?", 130000).Query(&model)
// SELECT emp_no, AVG(salary) FROM salaries
//      GROUP BY emp_no
//      HAVING AVG(salary) > 130000;

err := gsorm.Select(db).From("employees").
    Having("first_name = ?", "Taro").Query(&model)
// SELECT * FROM employees
//      HAVING first_name = 'Taro';

err := gsorm.Select(db).From("employees").
    Having("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Query(&model)
// SELECT * FROM employees
//      HAVING birth_date = '2006-01-02 00:00:00';

err := gsorm.Select(db).From("employees").
    Having("first_name LIKE ?", "%Taro").Query(&model)
// SELECT * FROM employees
//      HAVING first_name LIKE '%Taro';

err := gsorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").
    Having("AVG(salary) BETWEEN ? AND ?", 100000, 130000).Query(&model)
// SELECT emp_no, AVG(salary) FROM salaries
//      GROUP BY emp_no
//      HAVING AVG(salary) BETWEEN 100000 AND 130000;

err := gsorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").
    Having("AVG(salary) IN (?)", []int{100000, 130000}).Query(&model)
// SELECT emp_no, AVG(salary) FROM salaries
//      GROUP BY emp_no
//      HAVING AVG(salary) IN (100000, 130000);

err := gsorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").
    Having("AVG(salary) IN (?)", [2]int{100000, 130000}).Query(&model)
// SELECT emp_no, AVG(salary) FROM salaries
//      GROUP BY emp_no
//      HAVING AVG(salary) IN (100000, 130000);

err := gsorm.Select(db).From("employees").
    Having("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Query(&model)
// SELECT * FROM employees
//      HAVING emp_no IN (SELECT emp_no FROM dept_manager);
```


## Union
`Union` calls UNION clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Union)

#### Example
```go
gsorm.Select(db, "emp_no", "dept_no").From("dept_manager").
    Union(gsorm.Select(db, "emp_no", "dept_no").From("dept_emp")).Query(&model)
// SELECT emp_no, dept_no FROM dept_manager
//      UNION (SELECT emp_no, dept_no FROM dept_emp);
```


## UnionAll
`UnionAll` calls UNION ALL clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.UnionAll)

#### Example
```go
gsorm.Select(db, "emp_no", "dept_no").From("dept_manager").
    UnionAll(gsorm.Select(db, "emp_no", "dept_no").From("dept_emp")).Query(&model)
// SELECT emp_no, dept_no FROM dept_manager
//  UNION ALL (SELECT emp_no, dept_no FROM dept_emp);
```


## OrderBy
`OrderBy` calls ORDER BY clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.OrderBy)

#### Example
```go
err := gsorm.Select(db).From("employees").
    OrderBy("birth_date").Query(&model)
// SELECT * FROM employees
//      ORDER BY birth_date;

err := gsorm.Select(db).From("employees").
    OrderBy("birth_date DESC").Query(&model)
// SELECT * FROM employees
//      ORDER BY birth_date DESC;

err := gsorm.Select(db).From("employees").
    OrderBy("birth_date desc").Query(&model)
// SELECT * FROM employees
//      ORDER BY birth_date desc;

err := gsorm.Select(db).From("employees").
    OrderBy("birth_date").
    OrderBy("hire_date DESC").Query(&model)
// SELECT id FROM people
//      ORDER BY birth_date
//      ORDER BY hire_date DESC;
```


## Limit
`Limit` calls LIMIT clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Limit)

#### Example
```go
err := gsorm.Select(db).From("employees").
    Limit(10).Query(&model)
// SELECT * FROM employees
//      LIMIT 10;
```


## Offset
`Offset` calls OFFSET clause.

`Offset` is called after `Limit`.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Offset)

#### Example
```go
err := gsorm.Select(db).From("employees").
    Limit(10).
    Offset(5).Query(&model)
// SELECT * FROM employees
//      LIMIT 10
//      OFFSET 5;
```


## Query
`Query` executes the SQL and maps the results into the model.

The types available for [sql.Rows.Scan](https://golang.org/pkg/database/sql/#Rows.Scan) can be used as the model.

Also `struct`, `[]struct`, `map[string]interface{}`, and `[]map[string]interface{}` can be used as the model.
In this case, the type of slice element or array element must be the types available for `sql.Rows.Scan`.

Using struct as the model, the fileds must be exported.
The correspondance of the field names and the database column names are determined by the following rules.

- If the field is tagged `gsorm` and column name is specified, the name is used
- If the field is tagged `json`, the name is used
- If the field is tagged both `gsorm` and `json`, `gsorm` rule is applied
- If the field isn't tagged `gsorm` nor `json`, the snake case of the field name is used

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Query)

#### Example
```go
type Employee struct {
	EmpNo     int       `gsorm:"id"`
	FirstName string
	BirthDate time.Time
}

model := &[]Employee{}

err := gsorm.Select(db, "emp_no AS id", "first_name", "birth_date").From("employees").Query(&model)
// SELECT emp_no AS id, first_name, birth_date FROM employees;
```
