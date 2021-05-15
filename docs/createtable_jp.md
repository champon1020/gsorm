# CreateTable
`mgorm.CreateTable`はCREATE TABLE句を呼び出します．

引数にはデータベースのコネクション(`mgorm.Conn`)，テーブル名を指定します．

#### 例
```go
err := mgorm.CreateTable(db, "employees").
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
    ID          int         `mgorm:"emp_no typ=INT notnull=t"`
    BirthDate   time.Time   `mgorm:"typ=DATE notnull=t"`
    FirstName   string      `mgorm:"typ=VARCHAR(16) notnull=t"`
    LastName    string      `mgorm:"typ=VARCHAR(14) notnull=t"`
    Gender      string      `mgorm:"typ=ENUM('M', 'F') notnull=t"`
    HireDate    string      `mgorm:"typ=DATE notnull=t"`
}

err := mgorm.CreateTable(db, "employees").
    Model(&Employee{}).Migrate()
// Same as the previous example.
```


# Methods
`mgorm.CreateTable`で使用できるメソッドを以下に示します．
- [Column](https://github.com/champon1020/mgorm/tree/main/docs/createtable_jp.md#column)
  - [NotNull](https://github.com/champon1020/mgorm/tree/main/docs/createtable_jp.md#notnull)
  - [Default](https://github.com/champon1020/mgorm/tree/main/docs/createtable_jp.md#default)
- [Cons](https://github.com/champon1020/mgorm/tree/main/docs/createtable_jp.md#cons)
  - [Unique](https://github.com/champon1020/mgorm/tree/main/docs/createtable_jp.md#unique)
  - [Primary](https://github.com/champon1020/mgorm/tree/main/docs/createtable_jp.md#primary)
  - [Foreign](https://github.com/champon1020/mgorm/tree/main/docs/createtable_jp.md#foreign)
    - [Ref](https://github.com/champon1020/mgorm/tree/main/docs/createtable_jp.md#ref)

これらのメソッドは以下のEBNFに従って実行することができます．

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

Error = mgorm.CreateTable ColumnStmt { ColumnStmt } { ConstraintStmt } Migrate

ColumnStmt = Column [ NotNull ] [ Default ]
ConstraintStmt = Cons ( Unique | Primary | Foreign Ref )
```

例えば以下の実装はコンパイルエラーを吐き出します．

```go
// NG
err := mgorm.CreateTable(db, "employees").
    Cons("PK_employees").Primary("emp_no").
    Column("emp_no", "INT").NotNull().Migrate()

// NG
err := mgorm.CreateTable(db, "employees").
    NotNull().Column("emp_no", "INT").Migrate()

// NG
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().Primary("PK_employees").Migrate()

// NG
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("PK_employees").Primary("emp_no").
    Column("birth_date", "DATE").NotNull().Migrate()
```


## Column
`Column`はカラムの定義を呼び出します．

引数にはカラム名と型名を指定します．

`Column`は複数回び出すことができます．

#### 例
```go
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").Migrate()
// CREATE TABLE employees (
//      emp_no INT
// );

err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").
    Column("birth_date", "DATE").Migrate()
// CREATE TABLE employees (
//      emp_no      INT,
//      birth_date  DATE
// );
```

## NotNull
`NotNull`はNOTNULL句を呼び出します．

`NotNull`は`Column`に続けて呼び出すことができます．

#### 例
```go
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL
// );

err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Column("birth_date", "DATE").Migrate()
// CREATE TABLE employees (
//      emp_no      INT     NOT NULL,
//      birth_date  DATE
// );
```


## Default
`Default`はDEFAULT句を呼び出します．

引数には値を指定します．
このとき，値は以下のルールに従ってビルドされます．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれる．
- 以上の条件に該当しない場合，値はそのまま展開される．

`Default`は`Column`に続けて呼び出すことができます．

#### 例
```go
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").Default(1).Migrate()
// CREATE TABLE employees (
//      emp_no INT DEFAULT 1
// );

err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().Default(1).Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL DEFAULT 1
// );

err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().Default(1).
    Column("birth_date", "DATE").NotNull().Migrate()
// CREATE TABLE employees (
//      emp_no      INT     NOT NULL DEFAULT 1,
//      birth_date  DATE    NOT NULL
// );
```


## Cons
`Cons`はCONSTRAINT句を呼び出します．

引数には制約名を指定します．

`Cons`は`Column`，`NotNull`，`Default`のいずれかに続けて呼び出すことができます．

`Cons`に続けて`Unique`，`Primary`，`Foreign`のいずれかを呼び出す必要があります．

#### 例
```go
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("UC_emp_no").Unique("emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT UC_emp_no UNIQUE (emp_no)
// );

err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("PK_employees").Primary("emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT PK_employees PRIMARY KEY (emp_no)
// );

err := mgorm.CreateTable(db, "dept_emp").
    Column("emp_no", "INT").NotNull().
    Cons("FK_dept_emp_emp_no").Foreign("emp_no").Ref("employees", "emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT FK_dept_emp_emp_no FOREIGN KEY (emp_no) REFERENCES employees(emp_no)
// );
```


## Unique
`Unique`はUNIQUE句を呼び出します．

引数には複数カラム名を指定します．

`Unique`は`Cons`に続けて呼び出すことができます．

#### 例
```go
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("UC_emp_no").Unique("emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT UC_emp_no UNIQUE (emp_no)
// );

err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Column("first_name", "VARCHAR(14)").NotNull().
    Cons("UC_emp_no_first_name").Unique("emp_no", "first_name").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT UC_emp_no_first_name UNIQUE (emp_no, first_name)
// );
```


## Primary
`Primary`はPRIMARY KEY句を呼び出します．

引数には複数カラム名を指定します．

`Primary`は`Cons`に続けて呼び出すことができます．

#### 例
```go
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("PK_employees").Primary("emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT PK_employees PRIMARY KEY (emp_no)
// );

err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Column("first_name", "VARCHAR(14)").NotNull().
    Cons("PK_employees").Primary("emp_no", "first_name").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT PK_employees PRIMARY KEY (emp_no, first_name)
// );
```


## Foreign
`Foreign`はFOREIGN KEY句を呼び出します．

引数には複数カラム名を指定します．

`Foreign`は`Cons`に続けて呼び出すことができます．

`Foreign`に続けて`Ref`を呼び出す必要があります．

#### 例
```go
err := mgorm.CreateTable(db, "dept_emp").
    Column("emp_no", "INT").NotNull().
    Cons("FK_dept_emp").Foreign("emp_no").Ref("employees(emp_no)").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no) REFERENCES employees(emp_no)
// );

err := mgorm.CreateTable(db, "dept_emp").
    Column("emp_no", "INT").NotNull().
    Column("first_name", "VARCHAR(14)").NotNull().
    Cons("FK_dept_emp").Foreign("emp_no", "first_name").Ref("employees", "emp_no", "first_name").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no, first_name) REFERENCES employees(emp_no, first_name)
// );
```


## Ref
`Ref`はREFERENCES句を呼び出します．

第1引数に参照テーブル名，第2引数以降に複数の参照カラム名を指定します．

`Ref`は`Foreign`に続けて呼び出すことができます．

#### 例
```go
err := mgorm.CreateTable(db, "dept_emp").
    Column("emp_no", "INT").NotNull().
    Cons("FK_dept_emp").Foreign("emp_no").Ref("employees(emp_no)").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no) REFERENCES employees(emp_no)
// );

err := mgorm.CreateTable(db, "dept_emp").
    Column("emp_no", "INT").NotNull().
    Column("first_name", "VARCHAR(14)").NotNull().
    Cons("FK_dept_emp").Foreign("emp_no", "first_name").Ref("employees", "emp_no", "first_name").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no, first_name) REFERENCES employees(emp_no, first_name)
// );
```
