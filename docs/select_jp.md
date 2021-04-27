# Select
`mgorm.Select`はSELECT句を呼び出します．

引数にはデータベースのコネクション(`mgorm.Conn`)，カラム名を指定します．

カラム名は複数指定することができます．カラム名はからでも問題ありません．

#### 例
```go
err := mgorm.Select(db, "emp_no").From("employees").Query(&model)
// SELECT emp_no FROM people;

err := mgorm.Select(db).From("employees").Query(&model)
// SELECT * FROM people;

err := mgorm.Select(db, "emp_no", "first_name").From("employees").Query(&model)
// SELECT emp_no, first_name FROM people;

err := mgorm.Select(db, "emp_no, first_name").From("employees").Query(&model)
// SELECT emp_no, first_name FROM people;

err := mgorm.Select(db, "emp_no, first_name", "last_name").From("employees").Query(&model)
// SELECT emp_no, first_name, last_name FROM people;
```


# Methods
`mgorm.Select`に使用できるメソッドは以下です．

- [From](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#from)
- [Join](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#join)
- [LeftJoin](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#leftjoin)
- [RightJoin](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#rightjoin)
- [Where](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#where)
- [And](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#and)
- [Or](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#or)
- [GroupBy](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#groupby)
- [Having](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#having)
- [Union](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#union)
- [UnionAll](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#unionall)
- [OrderBy](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#orderby)
- [Limit](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#limit)
- [Offset](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#offset)

これらのメソッドは以下のEBNFに従って実行することができます．

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

mgorm.Select
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
err := mgorm.Select(db).
    Where("emp_no = ?", 10000).
    From("employees").Query(&model)

// NG
err := mgorm.Select(db).
    Join("dept_manager AS d").Query(&model)
```


## From
`From`はFROM句を呼び出します．

引数には複数のテーブル名を指定します．
これらのテーブル名にはエイリアスを含めることができます．

#### 例
```go
err := mgorm.Select(db, "emp_no").From("employees").Query(&model)
// SELECT emp_no FROM employees;

err := mgorm.Select(db, "e.emp_no").From("employees AS e").Query(&model)
// SELECT e.emp_no FROM employees AS e;

err := mgorm.Select(db, "e.emp_no").From("employees as e").Query(&model)
// SELECT e.emp_no FROM employees as e;

err := mgorm.Select(db, "emp_no", "dept_no").From("employees", "departments").Query(&model)
// SELECT emp_no, dept_no FROM employees, departments;
```


## Join
`Join`はINNERT JOIN句を呼び出します．

引数にはテーブル名を指定します．
このテーブル名にはエイリアスを含めることができます．

`Join`，`LeftJoin`，`RightJoin`は複数回呼び出すことができます．

#### 例
```go
err := mgorm.Select(db, "e.emp_no", "d.dept_no").
    From("employees AS e").
    Join("dept_manager AS d").
    On("e.emp_no = d.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no FROM employees AS e
//      INNER JOIN dept_manager AS d
//      ON e.emp_no = d.emp_no;

err := mgorm.Select(db, "e.emp_no", "d.dept_no", "s.salary").
    From("employees AS e").
    Join("dept_manager AS d").On("e.emp_no = d.emp_no").
    LeftJoin("salaries AS s").On("e.emp_no = s.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no, s.salary FROM employees AS e
//      INNER JOIN dept_manager AS d ON e.emp_no = d.emp_no;
//      LEFT  JOIN salaries     AS s ON e.emp_no = s.emp_no;
```


## LeftJoin
`LeftJoin`はLEFT JOIN句を呼び出します．

引数にはテーブル名を指定します．
このテーブル名にはエイリアスを含めることができます．

`Join`，`LeftJoin`，`RightJoin`は複数回呼び出すことができます．

#### 例
```go
err := mgorm.Select(db, "e.emp_no", "d.dept_no").
    From("employees AS e").
    LeftJoin("dept_manager AS d").
    On("e.emp_no = d.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no FROM employees AS e
//      LEFT JOIN dept_manager AS d
//      ON e.emp_no = d.emp_no;

err := mgorm.Select(db, "e.emp_no", "d.dept_no", "s.salary").
    From("employees AS e").
    LeftJoin("dept_manager AS d").On("e.emp_no = d.emp_no").
    RightJoin("salaries AS s").On("e.emp_no = s.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no, s.salary FROM employees AS e
//      LEFT  JOIN dept_manager AS d ON e.emp_no = d.emp_no;
//      RIGHT JOIN salaries     AS s ON e.emp_no = s.emp_no;
```


## RightJoin
`RightJoin`はRIGHT JOIN句を呼び出します．

引数にはテーブル名を指定します．
このテーブル名にはエイリアスを含めることができます．

`Join`，`LeftJoin`，`RightJoin`は複数回呼び出すことができます．

#### 例
```go
err := mgorm.Select(db, "e.emp_no", "d.dept_no").
    From("employees AS e").
    RightJoin("dept_manager AS d").
    On("e.emp_no = d.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no FROM employees AS e
//      RIGHT JOIN dept_manager AS d
//      ON e.emp_no = d.emp_no;

err := mgorm.Select(db, "e.emp_no", "d.dept_no", "s.salary").
    From("employees AS e").
    RightJoin("dept_manager AS d").On("e.emp_no = d.emp_no").
    Join("salaries AS s").On("e.emp_no = s.emp_no").Query(&model)
// SELECT e.emp_no, d.dept_no, s.salary FROM employees AS e
//      RIGHT JOIN dept_manager AS d ON e.emp_no = d.emp_no;
//      INNER JOIN salaries     AS s ON e.emp_no = s.emp_no;
```


## Where
`Where`はWHERE句を呼び出します．

第1引数に条件式，第2引数以降に複数値を指定できます．
この際，条件式における`?`に値が代入されます．
また，代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が`()`内に展開されます．
- 値が`*mgorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない値はそのまま展開される．

#### 例
```go
err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001;

err := mgorm.Select(db).From("employees").
    Where("emp_no = ?", 1001).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001;

err := mgorm.Select(db).From("employees").
    Where("first_name = ?", "Taro").Query(&model)
// SELECT * FROM employees
//      WHERE first_name = 'Taro';

err := mgorm.Select(db).From("employees").
    Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Query(&model)
// SELECT * FROM employees
//      WHERE birth_date = '2006-01-02 00:00:00';

err := mgorm.Select(db).From("employees").
    Where("first_name LIKE ?", "%Taro").Query(&model)
// SELECT * FROM employees
//      WHERE first_name LIKE '%Taro';

err := mgorm.Select(db).From("employees").
    Where("birth_date BETWEEN ? AND ?", 1001, 1002).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no BETWEEN 1001 AND 1002;

err := mgorm.Select(db).From("employees").
    Where("emp_no IN (?)", []int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no IN (1001, 1002);

err := mgorm.Select(db).From("employees").
    Where("emp_no IN (?)", [2]int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no IN (1001, 1002);

err := mgorm.Select(db).From("employees").
    Where("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no IN (SELECT emp_no FROM dept_manager);
```


## And
`And`はAND句を呼び出します．
この時実行されるSQLは，条件式が`()`で括られた形となります．

第1引数に条件式，第2引数以降に複数値を指定できます．
この際，条件式における`?`に値が代入されます．
また，代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が`()`内に展開されます．
- 値が`*mgorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない場合，値はそのまま展開される．

`And`は複数回呼び出すことができます．

#### 例
```go
err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    And("emp_no = 1002").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND emp_no = 1002;

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    And("emp_no = ?", 1002).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND emp_no = 1002;

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    And("first_name = ?", "Taro").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND first_name = 'Taro';

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND birth_date = '2006-01-02 00:00:00';

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    And("first_name LIKE ?", "%Taro").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND first_name LIKE '%Taro';

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    And("birth_date BETWEEN ? AND ?", 1001, 1002).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND emp_no BETWEEN 1001 AND 1002;

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    And("emp_no IN (?)", []int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND emp_no IN (1001, 1002);

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    And("emp_no IN (?)", [2]int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND emp_no IN (1001, 1002);

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    And("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      AND emp_no IN (SELECT emp_no FROM dept_manager);
```


## Or
`Or`はOR句を呼び出します．
この時実行されるSQLは，条件式が`()`で括られた形となります．

第1引数に条件式，第2引数以降に複数値を指定できます．
この際，条件式における`?`に値が代入されます．
また，代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が`()`内に展開されます．
- 値が`*mgorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない場合，値はそのまま展開される．

`Or`は複数回呼び出すことができます．

#### 例
```go
err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    Or("emp_no = 1002").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR emp_no = 1002;

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    Or("emp_no = ?", 1002).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR emp_no = 1002;

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    Or("first_name = ?", "Taro").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR first_name = 'Taro';

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR birth_date = '2006-01-02 00:00:00';

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    Or("first_name LIKE ?", "%Taro").Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR first_name LIKE '%Taro';

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    Or("birth_date BETWEEN ? AND ?", 1001, 1002).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR emp_no BETWEEN 1001 AND 1002;

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    Or("emp_no IN (?)", []int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR emp_no IN (1001, 1002);

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    Or("emp_no IN (?)", [2]int{1001, 1002}).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR emp_no IN (1001, 1002);

err := mgorm.Select(db).From("employees").
    Where("emp_no = 1001").
    Or("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).Query(&model)
// SELECT * FROM employees
//      WHERE emp_no = 1001
//      OR emp_no IN (SELECT emp_no FROM dept_manager);
```


## GroupBy
`GroupBy`はGROUP BY句を呼び出します．

引数には複数カラムを指定できます．

#### 例
```go
err := mgorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").Query(&model)
// SELECT emp_no, AVG(salary) FROM salaries
//      GROUP BY emp_no;
```


## Having
`Having`はHAVING句を呼び出します．

第1引数に条件式，第2引数以降に複数値を指定できます．
この際，条件式における`?`に値が代入されます．
また，代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が`()`内に展開されます．
- 値が`*mgorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない場合，値はそのまま展開される．

#### 例
```go
err := mgorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").
    Having("AVG(salary) > 130000").Query(&model)
// SELECT emp_no, AVG(salary) FROM salaries
//      GROUP BY emp_no
//      HAVING AVG(salary) > 130000;

err := mgorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").
    Having("AVG(salary) > ?", 130000).Query(&model)
// SELECT emp_no, AVG(salary) FROM salaries
//      GROUP BY emp_no
//      HAVING AVG(salary) > 130000;
```


## Union
`Union`は引数に`*mgorm.SelectStmt`型を受け取ります．
つまり，`mgorm.Select`による文を受け取ることができます．

`Union`と同様に`UnionAll`も使用することができます．

#### 例
```go
mgorm.Select(db, "emp_no", "dept_no").From("dept_manager").
    Union(mgorm.Select(db, "emp_no", "dept_no").From("dept_emp")).Query(&model)
// SELECT * FROM employees
//  UNION (SELECT * FROM departments);

mgorm.Select(db, "emp_no", "dept_no").From("dept_manager").
    UnionAll(mgorm.Select(db, "emp_no", "dept_no").From("dept_emp")).Query(&model)
// SELECT * FROM employees
//  UNION ALL (SELECT * FROM departments);
```


## OrderBy
`OrderBy`は引数に複数のカラム名をstring型で受け取ります．

必要であれば，カラム名に`DESC`や`ASC`などの順序の方向も含めてください．
`DESC`や`ASC`は小文字でも問題ありません．

#### 例
```go
err := mgorm.Select(db).From("employees").
    OrderBy("birth_date").Query(&model)
// SELECT * FROM employees
//  ORDER BY birth_date;

err := mgorm.Select(db).From("employees").
    OrderBy("hire_date DESC").Query(&model)
// SELECT * FROM employees
//  ORDER BY hire_date DESC;

err := mgorm.Select(db).From("employees").
    OrderBy("birth_date").
    OrderBy("hire_date DESC").Query(&model)
// SELECT id FROM people
//  ORDER BY birth_date
//  ORDER BY hire_date DESC;
```


## Limit
`Limit`は引数にint型を受け取ります．

#### 例
```go
err := mgorm.Select(db).From("employees").
    Limit(10).Query(&model)
// SELECT * FROM employees
//  LIMIT 10;
```


## Offset
`Offset`は引数にint型を受け取ります．
`Offset`は`Limit`の直後のみ呼び出すことができます．

#### 例
```go
err := mgorm.Select(db).From("employees").
    Limit(10).
    Offset(5).Query(&model)
// SELECT * FROM employees
//  LIMIT 10
//  OFFSET 5;
```
