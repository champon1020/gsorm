# Select
`mgorm.Select`を使用したとき，`Query`を呼び出すことでデータベースからの検索結果をマッピングすることができます．

`mgorm.Select`の第1引数は`mgorm.Conn`の型，第2引数以降は複数のカラム名をstring型として受け取ることができます．
`mgorm.Conn`を実装した型としては`*mgorm.DB`，`*mgorm.Tx`，`*mgorm.MockDB`，`*mgorm.MockTx`があります．

詳細は[Transaction]()，[Mock]()に記載されています．

#### 例
```go
// SELECT emp_no FROM people;
err := mgorm.Select(db, "emp_no").From("employees").Query(&model)

// SELECT * FROM people;
err := mgorm.Select(db).From("employees").Query(&model)

// SELECT emp_no, first_name FROM people;
err := mgorm.Select(db, "emp_no", "first_name").From("employees").Query(&model)

// SELECT emp_no, first_name FROM people;
err := mgorm.Select(db, "emp_no, first_name").From("employees").Query(&model)

// SELECT emp_no, first_name, last_name FROM people;
err := mgorm.Select(db, "emp_no, first_name", "last_name").From("employees").Query(&model)
```


## メソッド
`mgorm.Select`を使用する際の文法を以下に示します．
各メソッドは上に行くほど呼び出しの優先度が高いです．

各メソッドについては順に説明いたします．

```
[]: optional,  |: or,  {}: one of them

mgorm.Select(DB, columns...).From(tables...)
    [{ .Join(tables) | .LeftJoin(table) | .RightJoin(table) }.On(expression)]
    [.Where(expression, values...)]
    [.Or(expression, values...) | .And(expression, values...)]
    [.GroupBy(columns...)]
    [.Having(expression, values...)]
    [.Union(*mgorm.SelectStmt) | .UnionAll(*mgorm.SelectStmt)]
    [.OrderBy(columns...)]
    [.Limit(number)]
    [.Offset(number)]
    .Query(*model)
```


## From
`From`は複数のテーブル名をstring方で受け取ります．
必要であれば，テーブル名にエイリアスを含めることができます．

#### 例
```go
// SELECT emp_no FROM employees;
err := mgorm.Select(db, "emp_no").From("employees").Query(&model)

// SELECT e.emp_no FROM employees AS e;
err := mgorm.Select(db, "e.emp_no").From("employees AS e").Query(&model)

// SELECT e.emp_no FROM employees as e;
err := mgorm.Select(db, "e.emp_no").From("employees as e").Query(&model)

// SELECT emp_no, dept_no FROM employees, departments;
err := mgorm.Select(db, "emp_no", "dept_no").From("employees", "departments").Query(&model)
```


## Join
JOIN句で使用できるのは`Join`，`LeftJoin`，`RightJoin`の3種類です．
ここで，`Join`はINNER JOIN句としてSQLを実行します．

`Join`は結合したいテーブル名を引数としてstring型で受け取ります．
このとき受け取ることができるテーブル名は1つのみです．
複数テーブルを結合したい場合は，連続して`Join`メソッドを呼び出してください．

また，これらJOIN句に関するメソッドの後には`On`を呼び出す必要があります．
`On`には結合条件となる式を引数としてstring型で受け取ります．

これらの使用方法は，`LeftJoin`や`RightJoin`を使用する場合も同様です．

#### 例
```go
// SELECT e.emp_no, d.dept_no FROM employees AS e
//  INNER JOIN dept_manager AS d ON e.emp_no = d.emp_no;
err := mgorm.Select(db, "e.emp_no", "d.dept_no").From("employees AS e").
    Join("dept_manager AS d").On("e.emp_no = d.emp_no").Query(&model)

// SELECT e.emp_no, d.dept_no FROM employees AS e
//  LEFT JOIN dept_manager AS d ON e.emp_no = d.emp_no;
err := mgorm.Select(db, "e.emp_no", "d.dept_no").From("employees AS e").
    LeftJoin("dept_manager AS d").On("e.emp_no = d.emp_no").Query(&model)

// SELECT e.emp_no, d.dept_no FROM employees AS e
//  RIGHT JOIN dept_manager AS d ON e.emp_no = d.emp_no;
err := mgorm.Select(db, "e.emp_no", "d.dept_no").From("employees AS e").
    RightJoin("dept_manager AS d").On("e.emp_no = d.emp_no").Query(&model)

// SELECT e.emp_no, d.dept_no, s.salary FROM employees AS e
//  INNER JOIN dept_manager AS d ON e.emp_no = d.emp_no;
//  LEFT  JOIN salaries     AS s ON e.emp_no = s.emp_no;
err := mgorm.Select(db, "e.emp_no", "d.dept_no", "s.salary").From("employees AS e").
    Join("dept_manager AS d").On("e.emp_no = d.emp_no").
    LeftJoin("salaries AS s").On("e.emp_no = s.emp_no").Query(&model)
```


## Where
`Where`は引数に条件式を受け取ります．

条件式自体はstring型で受け取りますが，式の中に`?`を書くことで値を置き換えることができます．
このとき，値は複数置き換えることができます．

#### 例
```go
// SELECT * FROM employees
//  WHERE emp_no = 20000;
err := mgorm.Select(db).From("employees").
    Where("emp_no = ?", 20000).Query(&model)

// SELECT * FROM employees
//  WHERE first_name LIKE '%Taro';
err := mgorm.Select(db).From("employees").
    Where("first_name LIKE ?", "%Taro").Query(&model)

// SELECT * FROM employees
//  WHERE emp_no IN (10000, 20000);
err := mgorm.Select(db).From("employees").
    Where("emp_no IN (?)", []int{10000, 20000}).Query(&model)

// SELECT * FROM employees
//  WHERE emp_no BETWEEN 10000 AND 20000;
err := mgorm.Select(db).From("employees").
    Where("birth_date BETWEEN ?", [2]int{10000, 20000}).Query(&model)
```


また，`Where`を用いることで副問合せを実行することもできます．
これは，値として`mgorm.Select`による文を渡すことで実現できます．

#### 例
```go
// SELECT * FROM employees
//  WHERE emp_no IN (SELECT emp_no FROM dept_manager);
err := mgorm.Select(db).From("employees").
    Where("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).Query(&model)
```


## Or / And
`Or`，`And`は`Where`とほぼ同様の使用方法になります．
ただし，これらを使用した場合の条件式は`(`と`)`で囲まれたものになります．

#### 例
```go
// SELECT * FROM employees
//  WHERE emp_no > 20000
//  AND (first_name = 'Taro' OR first_name = 'Jiro');
err := mgorm.Select(db).From("employees").
    Where("emp_no > ?" 20000).
    And("first_name = ? OR first_name = ?", "Taro", "Jiro").Query(&model)

// SELECT * FROM employees
//  WHERE emp_no > 20000
//  OR (emp_no <= 20000 AND first_name = 'Saburo');
err := mgorm.Select(db).From("employees").
    Where("emp_no > ?", 20000).
    Or("emp_no <= ? AND first_name = ?", 20000, "Saburo").Query(&model)
```

また，`Where`と同様に副問合せも用いることができます．

#### 例
```go
// SELECT * FROM employees
//  WHERE emp_no > 20000
//  OR (first_name = 'Saburo' AND emp_no IN (SELECT emp_no FROM dept_manager));
err := mgorm.Select(db).From("employees").
    Where("emp_no > ?", 20000).
    Or("first_name = ? AND emp_no IN (?)", "Saburo", mgorm.Select(nil, "emp_no").From("dept_manager")).Query(&model)
```


## GroupBy
`GroupBy`は引数に複数のカラム名をstring型で受け取ります．


#### 例
```go
// SELECT emp_no, AVG(salary) FROM salaries
//  GROUP BY emp_no;
err := mgorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").Query(&model)
```


## Having
`Having`は引数に条件式を受け取ります．
条件式の受け取り方は`Where`や`And`，`Or`と同様です．

基本的には`GroupBy`と共に使用しますが，`Having`のみで使用することも可能です．

#### 例
```go
// SELECT emp_no, AVG(salary) FROM salaries
//  GROUP BY emp_no
//  HAVING AVG(salary) > 130000;
err := mgorm.Select(db, "emp_no", "AVG(salary)").From("salaries").
    GroupBy("emp_no").
    Having("AVG(salary) > ?", 130000).Query(&model)

// SELECT SUM(salary) FROM salaries
//  HAVING SUM(salary) > 1000000;
err := mgorm.Sum(db, "salary").From("salaries").
    Having("SUM(salary) > ?", 1000000).Query(&model);
```

`mgorm.Sum`については[Function]()に記載されています．


## Union
`Union`は引数に`*mgorm.SelectStmt`型を受け取ります．
つまり，`mgorm.Select`による文を受け取ることができます．

`Union`と同様に`UnionAll`も使用することができます．

#### 例
```go
// SELECT * FROM employees
//  UNION (SELECT * FROM departments);
mgorm.Select(db, "emp_no", "dept_no").From("dept_manager").
    Union(mgorm.Select(db, "emp_no", "dept_no").From("dept_emp")).Query(&model)

// SELECT * FROM employees
//  UNION ALL (SELECT * FROM departments);
mgorm.Select(db, "emp_no", "dept_no").From("dept_manager").
    UnionAll(mgorm.Select(db, "emp_no", "dept_no").From("dept_emp")).Query(&model)
```


## OrderBy
`OrderBy`は引数に複数のカラム名をstring型で受け取ります．

必要であれば，カラム名に`DESC`や`ASC`などの順序の方向も含めてください．
`DESC`や`ASC`は小文字でも問題ありません．

#### 例
```go
// SELECT * FROM employees
//  ORDER BY birth_date;
err := mgorm.Select(db).From("employees").
    OrderBy("birth_date").Query(&model)

// SELECT * FROM employees
//  ORDER BY hire_date DESC;
err := mgorm.Select(db).From("employees").
    OrderBy("hire_date DESC").Query(&model)

// SELECT id FROM people
//  ORDER BY birth_date
//  ORDER BY hire_date DESC;
err := mgorm.Select(db).From("employees").
    OrderBy("birth_date").
    OrderBy("hire_date DESC").Query(&model)
```


## Limit
`Limit`は引数にint型を受け取ります．

#### 例
```go
// SELECT * FROM employees
//  LIMIT 10;
err := mgorm.Select(db).From("employees").
    Limit(10).Query(&model)
```


## Offset
`Offset`は引数にint型を受け取ります．
`Offset`は`Limit`の直後のみ呼び出すことができます．

#### 例
```go
// SELECT * FROM employees
//  LIMIT 10
//  OFFSET 5;
err := mgorm.Select(db).From("employees").
    Limit(10).
    Offset(5).Query(&model)
```
