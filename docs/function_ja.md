# Function Query
COUNTやAVGなどのfunction queryを簡単に実装するために，以下のような関数が設けられています．

扱い方は`gsorm.Select`と同様です．

- [Count](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md#count)
- [Sum](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md#sum)
- [Avg](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md#avg)
- [Max](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md#max)
- [Min](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md#min)


## Count
`gsorm.Count`を使用することでSELECT COUNT(column)...を実行できます．

`gsorm.Count`には複数のカラム名をstring型で受け取ります．

#### 例
```go
err := gsorm.Count(db, "emp_no").From("employees").Query(&model)
// SELECT COUNT(emp_no) FROM employees;

// You can use gsorm.Select like this.
err := gsorm.Select(db, "COUNT(emp_no)").From("employees").Query(&model)
// SELECT COUNT(emp_no) FROM employees;

err := gsorm.Count(db, "emp_no", "CASE WHEN birth_date < '1960-01-01' THEN 1 END").From("employees").Query(&model)
// SELECT COUNT(emp_no), COUNT(CASE WHEN birth_date < '1960-01-01' THEN 1 END) FROM employees;
```


## Sum
`gsorm.Sum`を使用することでSELECT SUM(column)...を実行できます．

`gsorm.Sum`には複数のカラム名をstring型で受け取ります．

#### 例
```go
err := gsorm.Sum(db, "salary").From("salaries").Query(&model)
// SELECT SUM(salary) FROM salaries;

// You can use gsorm.Select like this.
err := gsorm.Select(db, "SUM(salary)").From("salaries").Query(&model)
// SELECT SUM(salary) FROM salaries;

err := gsorm.Sum(db, "salary", "CASE WHEN from_date < '1990-01-01' THEN salary END").From("salaries").Query(&model)
// SELECT SUM(salary), SUM(CASE WHEN from_date < '1990-01-01' THEN salary END) FROM salaries;
```


## Avg
`gsorm.Avg`を使用することでSELECT AVG(column)...を実行できます．

`gsorm.Avg`には複数のカラム名をstring型で受け取ります．

#### 例
```go
err := gsorm.Avg(db, "salary").From("salaries").Query(&model)
// SELECT AVG(salary) FROM salaries;

// You can use gsorm.Select like this.
err := gsorm.Select(db, "AVG(salary)").From("salaries").Query(&model)
// SELECT AVG(salary) FROM salaries;

err := gsorm.Avg(db, "salary", "CASE WHEN from_date < '1990-01-01' THEN salary END").From("salaries").Query(&model)
// SELECT AVG(salary), AVG(CASE WHEN from_date < '1990-01-01' THEN salary END) FROM salaries;
```


## Max
`gsorm.Max`を使用することでSELECT MAX(column)...を実行できます．

`gsorm.Max`には複数のカラム名をstring型で受け取ります．

#### 例
```go
err := gsorm.Max(db, "salary").From("salaries").Query(&model)
// SELECT MAX(salary) FROM salaries;

// You can use gsorm.Select like this.
err := gsorm.Select(db, "MAX(salary)").From("salaries").Query(&model)
// SELECT MAX(salary) FROM salaries;

err := gsorm.Max(db, "salary", "CASE WHEN from_date < '1990-01-01' THEN salary END").From("salaries").Query(&model)
// SELECT MAX(salary), MAX(CASE WHEN from_date < '1990-01-01' THEN salary END) FROM salaries;
```


## Min
`gsorm.Min`を使用することでSELECT MIN(column)...を実行できます．

`gsorm.Min`には複数のカラム名をstring型で受け取ります．

#### 例
```go
err := gsorm.Min(db, "salary").From("salaries").Query(&model)
// SELECT MIN(salary) FROM salaries;

// You can use gsorm.Select like this.
err := gsorm.Select(db, "MIN(salary)").From("salaries").Query(&model)
// SELECT MIN(salary) FROM salaries;

err := gsorm.Min(db, "salary", "CASE WHEN from_date < '1990-01-01' THEN salary END").From("salaries").Query(&model)
// SELECT MIN(salary), MIN(CASE WHEN from_date < '1990-01-01' THEN salary END) FROM salaries;
```
