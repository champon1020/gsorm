# AlterTable
`gsorm.AlterTable`はALTER TABLE句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#AlterTable)

#### 例
```go
err := gsorm.AlterTable(db, "employees").
    AddColumn("nickname", "VARCHAR(64)").NotNull().Migrate()
// ALTER TABLE employees
//      ADD COLUMN nickname VARCHAR(64) NOT NULL;

err := gsorm.AlterTable(db, "employees").
    AddCons("UC_nickname").Primary("nickname").Migrate()
// ALTER TABLE employees
//      ADD CONSTRAINT UC_nickname UNIQUE (nickname);
```


## Methods
`gsorm.AlterTable`に使用できるメソッドは以下です．

- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw_ja.md#rawclause)
- [Rename](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#rename)
- [AddColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#addcolumn)
    - [NotNull](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#addcolumn.notnull)
    - [Default](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#addcolumn.default)
- [DropColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#dropcolumn)
- [RenameColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#renamecolumn)
- [AddCons](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#addcons)
    - [Unique](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#addcons.unique)
    - [Primary](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#addcons.primary)
    - [Foreign](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#addcons.foreign)
        - [Ref](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#addcons.foreign.ref)

これらのメソッドは以下のEBNFに従って実行することができます．
例外として，`RawClause`は任意で呼び出すことができます．

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

gsorm.AlterTable
    (
        .Rename
        | (.AddColumn {.NotNull} {.Default})
        | .DropColumn
        | .RenameColumn
        | (.AddCons (.Unique | .Primary | .Foreign .Ref))
    )
    .Migrate
```

例えば，以下の実装はコンパイルエラーが出力されます．

```go
// NG
err := gsorm.AlterTable(db, "employees").
    Rename("people").
    DropColumn("id").Migrate()

// NG
err := gsorm.AlterTable(db, "employees").
    AddCons("UC_id").Migrate()
```


## Rename
`Rename`はRENAME TO句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Rename)

#### 例
```go
err := gsorm.AlterTable(db, "employees").
    Rename("people").Migrate()
// ALTER TABLE employees
//      RENAME TO people;
```


## AddColumn
`AddColumn`はADD COLUMN句を呼び出します．

`AddColumn`に続けて`NotNull`，`Default`を呼び出すことができます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.AddColumn)

#### 例
```go
err := gsorm.AlterTable(db, "employees").
    AddColumn("nickname", "VARCHAR(64)").Migrate()
// ALTER TABLE employees
//      ADD COLUMN nickname VARCHAR(64);
```


### AddColumn.NotNull
`NotNull`はNOT NULL句を呼び出します．

`NotNull`は`AddColumn`に続けて呼び出すことができます．

また，`NotNull`と`Default`は併用できます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.NotNull)

#### 例
```go
err := gsorm.AlterTable(db, "employees").
    AddColumn("nickname", "VARCHAR(64)").
    NotNull().Migrate()
// ALTER TABLE employees
//      ADD COLUMN nickanme VARCHAR(64) NOT NULL;

err := gsorm.AlterTable(db, "employees").
    AddColumn("nickname", "VARCHAR(64)").
    NotNull().
    Default("none").Migrate()
// ALTER TABLE employees
//      ADD COLUMN nickanme VARCHAR(64) NOT NULL DEFAULT 'none';
```

### AddColumn.Default
`Default`はDEFAULT句を呼び出します．

`Default`は`AddColumn`に続けて呼び出すことができます．

また，`Default`と`NotNull`は併用できます．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Default)

#### 例
```go
err := gsorm.AlterTable(db, "employees").
    AddColumn("nickname", "VARCHAR(64)").
    Default("none").Migrate()
// ALTER TABLE employees
//      ADD COLUMN nickanme VARCHAR(64) DEFAULT 'none';

err := gsorm.AlterTable(db, "employees").
    AddColumn("nickname", "VARCHAR(64)").
    Default("none").
    NotNull().Migrate()
// ALTER TABLE employees
//      ADD COLUMN nickanme VARCHAR(64) DEFAULT 'none'  NOT NULL;
```


## DropColumn
`DropColumn`はDROP COLUMN句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.DropColumn)

#### 例
```go
err := gsorm.AlterTable(db, "employees").
    DropColumn("nickname").Migrate()
// ALTER TABLE employees
//      DROP COLUMN nickname;
```


## RenameColumn
`RenameColumn`はRENAME COLUMN句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.RenameColumn)

#### 例
```go
err := gsorm.AlterTable(db, "employees").
    RenameColumn("emp_no", "id").Migrate()
// ALTER TABLE employees
//      RENAME COLUMN emp_no TO id;
```


## AddCons
`AddCons`はADD CONSTRAINT句を呼び出します．

`AddCons`に続けて`Unique`，`Primary`，`Foreign`のいずれかを呼び出す必要があります．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.AddCons)


### AddCons.Unique
`Unique`はUNIQUE句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Unique)

#### 例
```go
err := gsorm.AlterTable(db, "employees").
    AddCons("UC_nickname").Unique("nickname").Migrate()
// ALTER TABLE employees
//      ADD CONSTRAINT UC_nickname UNIQUE (nickname);

err := gsorm.AlterTable(db, "employees").
    AddCons("UC_nickname").Unique("nickname", "first_name").Migrate()
// ALTER TABLE employees
//      ADD CONSTRAINT UC_nickname UNIQUE (nickname, first_name);
```


### AddCons.Primary
`Primary`はPRIMARY KEY句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Primary)

#### 例
```go
err := gsorm.AlterTable(db, "employees").
    AddCons("PK_emp_no").Primary("emp_no").Migrate()
// ALTER TABLE employees
//      ADD CONSTRAINT PK_emp_no PRIMARY KEY (emp_no);

err := gsorm.AlterTable(db, "employees").
    AddCons("PK_emp_no").Primary("emp_no", "first_name").Migrate()
// ALTER TABLE employees
//      ADD CONSTRAINT PK_emp_no PRIMARY KEY (emp_no, first_name);
```

### AddCons.Foreign
`Foreign`はFOREIGN KEY句を呼び出します．

`Foreign`に続けて`Ref`を呼び出す必要があります．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Foreign)


### AddCons.Foreign.Ref
`Ref`はREFERENCES句を呼び出します．

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Ref)

#### 例
```go
err := gsorm.AlterTable(db, "dept_emp").
    AddCons("FK_emp_no").Foreign("emp_no").Ref("employees", "emp_no").Migrate()
// ALTER TABLE dept_emp
//      ADD CONSTRAINT FK_emp_no FOREIGN KEY (emp_no) REFERENCES employees (emp_no);

err := gsorm.AlterTable(db, "dept_emp").
    AddCons("FK_emp_no").Foreign("emp_no", "from_date").Ref("employees", "emp_no", "hire_date").Migrate()
// ALTER TABLE dept_emp
//      ADD CONSTRAINT FK_emp_no FOREIGN KEY (emp_no, from_date) REFERENCES employees (emp_no, hire_date);
```
