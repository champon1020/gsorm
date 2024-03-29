# Function Query
COUNTやAVGなどのfunction queryを簡単に実装するために，以下のような関数が設けられています．

- [Count](https://github.com/champon1020/gsorm/tree/main/docs/function_ja.md#count)
- [Sum](https://github.com/champon1020/gsorm/tree/main/docs/function_ja.md#sum)
- [Avg](https://github.com/champon1020/gsorm/tree/main/docs/function_ja.md#avg)
- [Max](https://github.com/champon1020/gsorm/tree/main/docs/function_ja.md#max)
- [Min](https://github.com/champon1020/gsorm/tree/main/docs/function_ja.md#min)

実態は`gsorm.Select`と同様です．


## Count
`gsorm.Count`はSELECT COUNT(...)文を実行します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Count.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Count)

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
`gsorm.Sum`はSELECT SUM(...)文を実行します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Sum.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Sum)

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
`gsorm.Avg`はSELECT AVG(...)文を実行します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Avg.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Avg)

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
`gsorm.Max`はSELECT MAX(...)文を実行します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Max.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Max)

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
`gsorm.Min`はSELECT MIN(...)文を実行します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Min.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Min)

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
