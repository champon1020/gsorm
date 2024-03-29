# Function Query
These functions can be used to execute function query easily.

- [Count](https://github.com/champon1020/gsorm/tree/main/docs/function.md#count)
- [Sum](https://github.com/champon1020/gsorm/tree/main/docs/function.md#sum)
- [Avg](https://github.com/champon1020/gsorm/tree/main/docs/function.md#avg)
- [Max](https://github.com/champon1020/gsorm/tree/main/docs/function.md#max)
- [Min](https://github.com/champon1020/gsorm/tree/main/docs/function.md#min)

In fact, `gsorm.Select` is called internally.


## Count
`gsorm.Count` calls SELECT COUNT(...) statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Count.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Count)

#### Example
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
`gsorm.Sum` calls SELECT SUM(...) statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Sum.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Sum)

#### Example
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
`gsorm.Avg` calls SELECT AVG(...) statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Avg.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Avg)

#### Example
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
`gsorm.Max` calls SELECT MAX(...) statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Max.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Max)

#### Example
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
`gsorm.Min` calls SELECT MIN(...) statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#Min.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#Min)

#### Example
```go
err := gsorm.Min(db, "salary").From("salaries").Query(&model)
// SELECT MIN(salary) FROM salaries;

// You can use gsorm.Select like this.
err := gsorm.Select(db, "MIN(salary)").From("salaries").Query(&model)
// SELECT MIN(salary) FROM salaries;

err := gsorm.Min(db, "salary", "CASE WHEN from_date < '1990-01-01' THEN salary END").From("salaries").Query(&model)
// SELECT MIN(salary), MIN(CASE WHEN from_date < '1990-01-01' THEN salary END) FROM salaries;
```
