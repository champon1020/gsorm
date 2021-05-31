# CreateTable
`gsorm.CreateTable`はCREATE TABLE文を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#CreateTableStmt)

#### 例
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
// Same as the previous example.
```


## Methods
- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw_ja.md#rawclause)
- [Column](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#column)
  - [NotNull](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#columnnotnull)
  - [Default](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#columndefault)
- [Cons](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#cons)
  - [Unique](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#consunique)
  - [Primary](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#consprimary)
  - [Foreign](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#consforeign)
    - [Ref](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#consforeignref)
- [Model](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#model)

これらのメソッドは以下のEBNFに従って実行することができます．

但し，例外として`RawClause`は任意で呼び出すことができます．

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

Error = gsorm.CreateTable ColumnStmt { ColumnStmt } { ConstraintStmt } Migrate

ColumnStmt = Column [ NotNull ] [ Default ]
ConstraintStmt = Cons ( Unique | Primary | Foreign Ref )
```

例えば以下の実装はコンパイルエラーを吐き出します．

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
`Column`はカラムの定義を呼び出します．

`Column`は複数回び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTableStmt.Column)

#### 例
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
`NotNull`はNOTNULL句を呼び出します．

`NotNull`は`Column`に続けて呼び出すことができます．

`NotNull`と`Default`は併用できます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTableStmt.NotNull)

#### 例
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
`Default`はDEFAULT句を呼び出します．

`Default`は`Column`に続けて呼び出すことができます．

`Default`と`NotNull`は併用できます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTableStmt.Default)

#### 例
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
`Cons`はCONSTRAINT句を呼び出します．

`Cons`に続けて`Unique`，`Primary`，`Foreign`のいずれかを呼び出す必要があります．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTableStmt.Cons)

#### 例
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
`Unique`はUNIQUE句を呼び出します．

`Unique`は`Cons`に続けて呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTableStmt.Unique)

#### 例
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
`Primary`はPRIMARY KEY句を呼び出します．

`Primary`は`Cons`に続けて呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTableStmt.Primary)

#### 例
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
`Foreign`はFOREIGN KEY句を呼び出します．

`Foreign`は`Cons`に続けて呼び出すことができます．

`Foreign`に続けて`Ref`を呼び出す必要があります．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTableStmt.Foreign)


### Cons.Foreign.Ref
`Ref`はREFERENCES句を呼び出します．

第1引数に参照テーブル名，第2引数以降に複数の参照カラム名を指定します．

`Ref`は`Foreign`に続けて呼び出すことができます．

#### 例
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
`Model`は構造体をマッピングします．

Modelについての詳細は[Model](https://github.com/champon1020/gsorm/tree/main/docs/model_ja.md)に記載されています．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#CreateTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#CreateTableStmt.Model)

#### 例
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
