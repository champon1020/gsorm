# CreateTable
`mgorm.CreateTable`を使用したとき，`Migrate`を呼び出すことでテーブルを作成することができる．

#### 例
```go
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Column("birth_date", "DATE").NotNull().
    Column("first_name", "VARCHAR(14)").NotNull().
    Column("last_name", "VARCHAR(16)").NotNull().
    Column("gender", "ENUM('M', 'F')").NotNull().
    Column("hire_date", "DATE").NotNull().
    Cons("PK_employees").PrimaryKey("emp_no").Migrate()
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
// Equal to previous example.
```


# Methods
`mgorm.CreateTable`で使用できるメソッドを以下に示します．
- [Column]()
- [NotNull]()
- [Default]()
- [Cons]()
- [Unique]()
- [PrimaryKey]()
- [ForeignKey]()

```
[]: optional, |: or, {}: block, **: able to call many times

mgorm.CreateTable(DB, table)
    {.Column()
        [.NotNull()]
        [.Default(value)]}**
    [.Cons(key)
        {.Unique(columns...) | .PrimaryKey(columns...) | ForeignKey(columns...).Ref(ref)}]**
```

上図において，上のメソッドほど実行優先度が高いです．例えば以下はコンパイルエラーとなります．
```go
// NG
err := mgorm.CreateTable(db, "employees").
    Cons("PK_employees").PrimaryKey("emp_no").
    Column("emp_no", "INT").NotNull().Migrate()

// NG
err := mgorm.CreateTable(db, "employees").
    NotNull().Column("emp_no", "INT").Migrate()

// NG
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().PrimaryKey("PK_employees").Migrate()

// NG
err := mgorm.CreateTable(db, "employees").
    Column("emp_no", "INT").NotNull().
    Cons("PK_employees").PrimaryKey("emp_no").
    Column("birth_date", "DATE").NotNull().Migrate()
```


## Column
`Column`はカラム名とデータベース型をstring型で受け取ります．

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
`Default`は`Column`もしくは`NotNull`に続けて呼び出すことができます．

`Default`は引数に値を受け取ります．
この値の型は，`mgorm`で許されている型のみ受け取ることができます．
型についての詳細は[Type]()に記載されています．

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
`Cons`は引数に鍵名をstring型で受け取り，CONSTRAINT句を呼び出します．

`Cons`のみでは文を完結することができないため，[Unique]()，[PrimaryKey]()，[ForeignKey]()のいずれかを続けて呼び出す必要があります．

また，`Cons`を呼び出すには1回以上`Column`が呼び出されている必要があります．

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
    Cons("PK_employees").PrimaryKey("emp_no").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT PK_employees PRIMARY KEY (emp_no)
// );

err := mgorm.CreateTable(db, "dept_emp").
    Column("emp_no", "INT").NotNull().
    Cons("FK_dept_emp_emp_no").ForeignKey("emp_no").Ref("employees(emp_no)").Migrate()
// CREATE TABLE employees (
//      emp_no INT NOT NULL,
//      CONSTRAINT FK_dept_emp_emp_no FOREIGN KEY (emp_no) REFERENCES employees(emp_no)
// );
```


## Unique


## PrimaryKey


## ForeignKey
