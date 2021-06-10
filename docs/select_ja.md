# Select
`gsorm.Select`はSELECT文を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Select)

#### 例
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
- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw_ja.md#rawclause)
- [From](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#from)
- [Join](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#join)
- [LeftJoin](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#leftjoin)
- [RightJoin](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#rightjoin)
- [Where](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#where)
- [And](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#and)
- [Or](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#or)
- [GroupBy](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#groupby)
- [Having](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#having)
- [Union](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#union)
- [UnionAll](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#unionall)
- [OrderBy](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#orderby)
- [Limit](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#limit)
- [Offset](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#offset)
- [Query](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#query)

これらのメソッドは以下のEBNFに従って実行することができます．

但し，例外として`RawClause`は任意で呼び出すことができます．

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

例えば以下の実装はコンパイルエラーを吐き出します．

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
`From`はFROM句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.From)

#### 例
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
`Join`はINNERT JOIN句を呼び出します．

`Join`，`LeftJoin`，`RightJoin`は複数回呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Join)

#### 例
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
`LeftJoin`はLEFT JOIN句を呼び出します．

`Join`，`LeftJoin`，`RightJoin`は複数回呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.LeftJoin)

#### 例
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
`RightJoin`はRIGHT JOIN句を呼び出します．

`Join`，`LeftJoin`，`RightJoin`は複数回呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.RightJoin)

#### 例
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
`Where`はWHERE句を呼び出します．

クエリが実行されるとき，条件式における`?`に値が代入されます．

代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値がスライスもしくは配列の場合，その要素が展開されます．
- 値が`gsorm.Stmt`型の場合，`gsorm.Stmt`は展開されます．
- 以上の条件に該当しない値はそのまま展開される．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Where)

#### 例
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
`And`はAND句を呼び出します．

このときAND句は条件式は`()`で括られます．

クエリが実行されるとき，条件式における`?`に値が代入されます．

代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値がスライスもしくは配列の場合，その要素が展開されます．
- 値が`gsorm.Stmt`型の場合，`gsorm.Stmt`は展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`And`は複数回呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.And)

#### 例
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
`Or`はOR句を呼び出します．

このときOR句は条件式は`()`で括られます．

クエリが実行されるとき，条件式における`?`に値が代入されます．

代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値がスライスもしくは配列の場合，その要素が展開されます．
- 値が`gsorm.Stmt`型の場合，`gsorm.Stmt`は展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`Or`は複数回呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Or)

#### 例
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
`GroupBy`はGROUP BY句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.GroupBy)

#### 例
```go
err := gsorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").Query(&model)
// SELECT emp_no, AVG(salary) FROM salaries
//      GROUP BY emp_no;
```


## Having
`Having`はHAVING句を呼び出します．

クエリが実行されるとき，条件式における`?`に値が代入されます．

代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値がスライスもしくは配列の場合，その要素が展開されます．
- 値が`gsorm.Stmt`型の場合，`gsorm.Stmt`は展開されます．
- 以上の条件に該当しない値はそのまま展開される．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Having)

#### 例
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
`Union`はUNION句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Union)

#### 例
```go
gsorm.Select(db, "emp_no", "dept_no").From("dept_manager").
    Union(gsorm.Select(db, "emp_no", "dept_no").From("dept_emp")).Query(&model)
// SELECT emp_no, dept_no FROM dept_manager
//      UNION (SELECT emp_no, dept_no FROM dept_emp);
```


## UnionAll
`UnionAll`はUNION ALL句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.UnionAll)

#### 例
```go
gsorm.Select(db, "emp_no", "dept_no").From("dept_manager").
    UnionAll(gsorm.Select(db, "emp_no", "dept_no").From("dept_emp")).Query(&model)
// SELECT emp_no, dept_no FROM dept_manager
//  UNION ALL (SELECT emp_no, dept_no FROM dept_emp);
```


## OrderBy
`OrderBy`はORDER BY句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.OrderBy)

#### 例
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
`Limit`はLIMIT句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Limit)

#### 例
```go
err := gsorm.Select(db).From("employees").
    Limit(10).Query(&model)
// SELECT * FROM employees
//      LIMIT 10;
```


## Offset
`Offset`はOFFSET句を呼び出します．

`Offset`は`Limit`に続けて呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Offset)

#### 例
```go
err := gsorm.Select(db).From("employees").
    Limit(10).
    Offset(5).Query(&model)
// SELECT * FROM employees
//      LIMIT 10
//      OFFSET 5;
```


## Query
`Query`はSQLを実行して，結果をmodelにマッピングします．

modelには[sql.Rows.Scan](https://golang.org/pkg/database/sql/#Rows.Scan)に使用できる型を利用できます．

また，`struct{}`，`[]struct{}`，`map[string]interface{}`，`[]map[string]interface{}`を使用することもできます．
このとき，マップの要素，構造体のフィールドの型は`sql.Rows.Scan`に使用できる型である必要があります．

構造体を使用する場合は，フィールドがエクスポートされている必要があります．
構造体フィールド名とデータベースカラムの対応は以下の規則で決定されます．

- フィールドの`gsorm`タグによってカラム名を指定しているとき，そのカラム名が使用されます
- フィールドに`json`が付いているとき，その名前が使用されます
- `gsorm`と`json`の両方が付いているとき，`gsorm`の規則が優先されます
- `gsorm`と`json`の両方とも付いていないとき，フィールド名のスネークケースが使用されます

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Select.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#SelectStmt.Query)

#### 例
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
