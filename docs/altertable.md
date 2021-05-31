# AlterTable
`gsorm.AlterTable` calls ALTER TALBE statement.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm#AlterTable)

#### Example
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
- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw.md#rawclause)
- [Rename](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#rename)
- [AddColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#addcolumn)
  - [NotNull](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#addcolumnnotnull)
  - [Default](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#addcolumndefault)
- [DropColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#dropcolumn)
- [RenameColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#renamecolumn)
- [AddCons](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#addcons)
  - [Unique](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#addconsunique)
  - [Primary](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#addconsprimary)
  - [Foreign](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#addconsforeign)
    - [Ref](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#addconsforeignref)

These methods can be executed according to the following EBNF.

Exceptionally, `RawClause` can be executed at any time.

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

For example, these implementations output the compile error.

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
`Rename` calls RENAME TO caluse.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Rename)

#### Example
```go
err := gsorm.AlterTable(db, "employees").
    Rename("people").Migrate()
// ALTER TABLE employees
//      RENAME TO people;
```


## AddColumn
`AddColumn` calls ADD COLUMN clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.AddColumn)

#### 例
```go
err := gsorm.AlterTable(db, "employees").
    AddColumn("nickname", "VARCHAR(64)").Migrate()
// ALTER TABLE employees
//      ADD COLUMN nickname VARCHAR(64);
```


### AddColumn.NotNull
`NotNull` calls NOT NULL clause.

`NotNull` can be called after `AddColumn`.

`NotNull` and `Default` can be used together.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.NotNull)

#### Example
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
`Default` calls DEFAULT clause.

`Default` can be called after `AddColumn`.

`Default` and `NotNull` can be used together.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Default)

#### Example
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
`DropColumn` calls DROP COLUMN clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.DropColumn)

#### Example
```go
err := gsorm.AlterTable(db, "employees").
    DropColumn("nickname").Migrate()
// ALTER TABLE employees
//      DROP COLUMN nickname;
```


## RenameColumn
`RenameColumn` calls RENAME COLUMN clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.RenameColumn)

#### Example
```go
err := gsorm.AlterTable(db, "employees").
    RenameColumn("emp_no", "id").Migrate()
// ALTER TABLE employees
//      RENAME COLUMN emp_no TO id;
```


## AddCons
`AddCons` calls ADD CONSTRAINT clause.

`Unique`, `Primary` or `Foreign` is called after `AddCons`.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.AddCons)

#### Example
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


### AddCons.Unique
`Unique` calls UNIQUE clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Unique)

#### Example
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
`Primary` calls PRIMARY KEY clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Primary)

#### Example
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


### AddCons.Foregin
`Foreign` calls FOREIGN KEY clause.

`Ref` is called after `Foreign`.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Foreign)


### AddCons.Foreing.Ref
`Ref` calls REFERENCES clause.

[![Go Reference](https://pkg.go.dev/badge/github.com/champon1020/gsorm#AlterTable.svg)](https://pkg.go.dev/github.com/champon1020/gsorm/statement/migration#AlterTableStmt.Ref)

#### Example
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
