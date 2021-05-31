# CreateTable
`gsorm.CreateTable` calls CREATE TABLE statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#CreateTable)

#### Example
```go
err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Column("birth_date", "DATE").NotNull().
    Column("first_name", "VARCHAR(14)").NotNull().
    Column("last_name", "VARCHAR(16)").NotNull().
    Column("gender", "ENUM('M', 'F')").NotNull().
    Column("hire_date", "DATE").NotNull().
    Cons("PK_employees").Primary("emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no      INT             NOT NULL,
//      birth_date  DATE            NOT NULL,
//      first_name  VARCHAR(14)     NOT NULL,
//      last_name   VARCHAR(16)     NOT NULL,
//      gender      ENUM('M', 'F')  NOT NULL,
//      hire_date   DATE            NOT NULL,
//      CONSTRAINT PK_employees PRIMARY KEY (emp_no)
// );
```


## Methods
- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw.md#rawclause)
- [Column](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#column)
  - [NotNull](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#columnnotnull)
  - [Default](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#columndefault)
- [Cons](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#cons)
  - [Unique](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#consunique)
  - [Primary](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#consprimary)
  - [Foreign](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#consforeign)
    - [Ref](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#consforeignref)
- [Model](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#model)

These methods can be executed according to the following EBNF.

Exceptionally, `RawClause` can be executed at any time.

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

Error = gsorm.CreateTable ColumnStmt { ColumnStmt } { ConstraintStmt } Migrate

ColumnStmt = Column [ NotNull ] [ Default ]
ConstraintStmt = Cons ( Unique | Primary | Foreign Ref )
```

For example, these implementations output the compile error.

```go
// NG
err := gsorm.CreateTable(db, "employees").
    Cons("PK_employees").Primary("emp_no").
    Column("emp_no", "INT").NotNull().Migrate()

// NG
err := gsorm.CreateTable(db, "employees").
    NotNull().Column("emp_no", "INT").Migrate()

// NG
err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().Primary("PK_employees").Migrate()

// NG
err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("PK_employees").Primary("emp_no").
    Column("birth_date", "DATE").NotNull().Migrate()
```


## Column
`Column` calls the column definition.

`Column` can be called multiple times.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTable.Column)

#### Example
```go
err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").Migrate()
// CREATE TABLE employees (
//      emp_no INT
// );

err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").
    Column("birth_date", "DATE").Migrate()
// CREATE TABLE employees (
//      emp_no      INT,
//      birth_date  DATE
// );

```


### Column.NotNull
`NotNull` calls NOT NULL clause.

`NotNull` is called after `Column`.

`NotNull` and `Default` can be used together.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTable.NotNull)

#### Example
```go
err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL
// );

err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Column("birth_date", "DATE").Migrate()
// CREATE TABLE employees (
//      emp_no      INT     NOT NULL,
//      birth_date  DATE
// );
```


### Column.Default
`Default` calls DEFAULT clause.

`Default` is called after `Default`

`Default` and `NotNull` can be used together.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTable.Default)

#### Example
```go
err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").Default(1).Migrate()
// CREATE TABLE employees (
//      emp_no INT DEFAULT 1
// );

err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().Default(1).Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL DEFAULT 1
// );

err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").Default(1).NotNull().Migrate()
// CREATE TABLE employees (
//      emp_no INT DEFAULT 1 NOT NULL
// );

err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().Default(1).
    Column("birth_date", "DATE").NotNull().Migrate()
// CREATE TABLE employees (
//      emp_no      INT     NOT NULL DEFAULT 1,
//      birth_date  DATE    NOT NULL
// );
```


## Cons
`Cons` calls CONSTRAINT clause.

`Unique`, `Primary` or `Foreign` must be called after `Cons`.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTable.Cons)

#### Example
```go
err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("UC_emp_no").Unique("emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT UC_emp_no UNIQUE (emp_no)
// );

err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("PK_employees").Primary("emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT PK_employees PRIMARY KEY (emp_no)
// );

err := gsorm.CreateTable(db, "dept_emp").
    Column("emp_no", "INT").NotNull().
    Cons("FK_dept_emp_emp_no").Foreign("emp_no").Ref("employees", "emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT FK_dept_emp_emp_no FOREIGN KEY (emp_no) REFERENCES employees (emp_no)
// );
```


### Cons.Unique
`Unique` calls UNIQUE clause.

`Unique` is called after `Cons`.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTable.Unique)

#### Example
```go
err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("UC_emp_no").Unique("emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT UC_emp_no UNIQUE (emp_no)
// );

err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Column("first_name", "VARCHAR(14)").NotNull().
    Cons("UC_emp_no_first_name").Unique("emp_no", "first_name").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT UC_emp_no_first_name UNIQUE (emp_no, first_name)
// );
```


### Cons.Primary
`Primary` calls PRIMARY KEY clause.

`Primary` is called after `Cons`.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTable.Primary)

#### Example
```go
err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("PK_employees").Primary("emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT PK_employees PRIMARY KEY (emp_no)
// );

err := gsorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Column("first_name", "VARCHAR(14)").NotNull().
    Cons("PK_employees").Primary("emp_no", "first_name").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT PK_employees PRIMARY KEY (emp_no, first_name)
// );
```


### Cons.Foreign
`Foreign` calls FOREIGN KEY clause.

`Foreign` is called after `Cons`.

`Ref` must be called after `Foreign`.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTable.Foreign)


### Cons.Foreign.Ref
`Ref` calls REFERENCES clause.

`Ref` is called after `Foreign`.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTable.Ref)

#### Example
```go
err := gsorm.CreateTable(db, "dept_emp").
    Column("emp_no", "INT").NotNull().
    Cons("FK_dept_emp").Foreign("emp_no").Ref("employees(emp_no)").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no) REFERENCES employees (emp_no)
// );

err := gsorm.CreateTable(db, "dept_emp").
    Column("emp_no", "INT").NotNull().
    Column("first_name", "VARCHAR(14)").NotNull().
    Cons("FK_dept_emp").Foreign("emp_no", "first_name").Ref("employees", "emp_no", "first_name").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no, first_name) REFERENCES employees (emp_no, first_name)
// );
```


## Model
`Model` maps the model into table schema.

Details are given in [Model](https://github.com/champon1020/gsorm/tree/main/docs/model.md).

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTable.Model)

#### Example
```go
type Employee struct {
    ID          int         `gsorm:"emp_no typ=INT notnull=t"`
    BirthDate   time.Time   `gsorm:"typ=DATE notnull=t"`
    FirstName   string      `gsorm:"typ=VARCHAR(16) notnull=t"`
    LastName    string      `gsorm:"typ=VARCHAR(14) notnull=t"`
    Gender      string      `gsorm:"typ=ENUM('M', 'F') notnull=t"`
    HireDate    string      `gsorm:"typ=DATE notnull=t"`
}

err := gsorm.CreateTable(db, "employees").
    Model(&Employee{}).Migrate()
// CREATE TABLE employees (
//      emp_no      INT             NOT NULL,
//      birth_date  DATE            NOT NULL,
//      first_name  VARCHAR(14)     NOT NULL,
//      last_name   VARCHAR(16)     NOT NULL,
//      gender      ENUM('M', 'F')  NOT NULL,
//      hire_date   DATE            NOT NULL,
//      CONSTRAINT PK_employees PRIMARY KEY (emp_no)
// );
```
