# Function Query
COUNTやAVGなどのfunction queryを簡単に実装するために，以下のような関数が設けられています．

扱い方は`mgorm.Select`と同様です．

- [Count](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md#count)
- [Sum](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md#sum)
- [Avg](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md#avg)
- [Max](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md#max)
- [Min](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md#min)


## Count
`mgorm.Count`を使用することでSELECT COUNT(column)...を実行できます．

`mgorm.Count`には複数のカラム名をstring型で受け取ります．

#### 例
```go
err := mgorm.Count(db, "emp_no").From("employees").Query(&model)
// SELECT COUNT(emp_no) FROM employees;

// You can use mgorm.Select like this.
err := mgorm.Select(db, "COUNT(emp_no)").From("employees").Query(&model)
// SELECT COUNT(emp_no) FROM employees;

err := mgorm.Count(db, "emp_no", "CASE WHEN birth_date < '1960-01-01' THEN 1 END").From("employees").Query(&model)
// SELECT COUNT(emp_no), COUNT(CASE WHEN birth_date < '1960-01-01' THEN 1 END) FROM employees;
```


## Sum
`mgorm.Sum`を使用することでSELECT SUM(column)...を実行できます．

`mgorm.Sum`には複数のカラム名をstring型で受け取ります．

#### 例
```go
err := mgorm.Sum(db, "salary").From("salaries").Query(&model)
// SELECT SUM(salary) FROM salaries;

// You can use mgorm.Select like this.
err := mgorm.Select(db, "SUM(salary)").From("salaries").Query(&model)
// SELECT SUM(salary) FROM salaries;

err := mgorm.Sum(db, "salary", "CASE WHEN from_date < '1990-01-01' THEN salary END").From("salaries").Query(&model)
// SELECT SUM(salary), SUM(CASE WHEN from_date < '1990-01-01' THEN salary END) FROM salaries;
```


## Avg
`mgorm.Avg`を使用することでSELECT AVG(column)...を実行できます．

`mgorm.Avg`には複数のカラム名をstring型で受け取ります．

#### 例
```go
err := mgorm.Avg(db, "salary").From("salaries").Query(&model)
// SELECT AVG(salary) FROM salaries;

// You can use mgorm.Select like this.
err := mgorm.Select(db, "AVG(salary)").From("salaries").Query(&model)
// SELECT AVG(salary) FROM salaries;

err := mgorm.Avg(db, "salary", "CASE WHEN from_date < '1990-01-01' THEN salary END").From("salaries").Query(&model)
// SELECT AVG(salary), AVG(CASE WHEN from_date < '1990-01-01' THEN salary END) FROM salaries;
```


## Max
`mgorm.Max`を使用することでSELECT MAX(column)...を実行できます．

`mgorm.Max`には複数のカラム名をstring型で受け取ります．

#### 例
```go
err := mgorm.Max(db, "salary").From("salaries").Query(&model)
// SELECT MAX(salary) FROM salaries;

// You can use mgorm.Select like this.
err := mgorm.Select(db, "MAX(salary)").From("salaries").Query(&model)
// SELECT MAX(salary) FROM salaries;

err := mgorm.Max(db, "salary", "CASE WHEN from_date < '1990-01-01' THEN salary END").From("salaries").Query(&model)
// SELECT MAX(salary), MAX(CASE WHEN from_date < '1990-01-01' THEN salary END) FROM salaries;
```


## Min
`mgorm.Min`を使用することでSELECT MIN(column)...を実行できます．

`mgorm.Min`には複数のカラム名をstring型で受け取ります．

#### 例
```go
err := mgorm.Min(db, "salary").From("salaries").Query(&model)
// SELECT MIN(salary) FROM salaries;

// You can use mgorm.Select like this.
err := mgorm.Select(db, "MIN(salary)").From("salaries").Query(&model)
// SELECT MIN(salary) FROM salaries;

err := mgorm.Min(db, "salary", "CASE WHEN from_date < '1990-01-01' THEN salary END").From("salaries").Query(&model)
// SELECT MIN(salary), MIN(CASE WHEN from_date < '1990-01-01' THEN salary END) FROM salaries;
```
