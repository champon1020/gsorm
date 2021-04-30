# Alter Table
`mgorm.AlterTable`はALTER TABLE句を呼び出します．

引数にはデータベースのコネクション(`mgorm.Conn`)，テーブル名を指定します．

#### 例
```go
err := mgorm.AlterTable(db, "employees").
    AddColumn("nickname", "VARCHAR(64)").NotNull().Migrate()
// ALTER TABLE clients
//      ADD COLUMN nickname VARCHAR(64) NOT NULL;

err := mgorm.AlterTable(db, "employees").
    AddCons("UC_nickname").Primary("nickname").Migrate()
// ALTER TABLE clients
//      ADD CONSTRAINT UC_nickname UNIQUE (nickname);
```


# Methods
`mgorm.AlterTable`に使用できるメソッドは以下です．

- [Rename](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#rename)
- [AddColumn](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#addcolumn)
    - [NotNull](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#addcolumn.notnull)
    - [Default](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#addcolumn.default)
- [DropColumn](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#dropcolumn)
- [RenameColumn](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#renamecolumn)
- [AddCons](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#addcons)
    - [Unique](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#addcons.unique)
    - [Primary](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#addcons.primary)
    - [Foreign](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#addcons.foreign)
        - [Ref](https://github.com/champon1020/mgorm/tree/main/docs/altertable_jp.md#addcons.foreign.ref)

これらのメソッドは以下のEBNFに従って実行することができます．

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

mgorm.AlterTable
    (
        .Rename
        | (.AddColumn {.NotNull} {.Default})
        | .DropColumn
        | .RenameColumn
        | (.AddCons (.Unique | .Primary | .Foreign .Ref))
    )
    .Migrate
```

例えば以下の実装はコンパイルエラーを吐き出します．

```go
// NG
err := mgorm.AlterTable(db, "employees").
    Rename("people").
    DropColumn("id").Migrate()

// NG
err := mgorm.AlterTable(db, "employees").
    AddCons("UC_id").Migrate()
```


## Rename
`Rename`はRENAME TO句を呼び出します．

引数にはテーブル名を指定します．

#### 例
```go
err := mgorm.AlterTable(db, "employees").
    Rename("people").Migrate()
// ALTER TABLE clients
//      RENAME TO people;
```


## AddColumn
`AddColumn`はADD COLUMN句を呼び出します．

引数にはカラム名，型名を指定します．

#### 例
```go
err := mgorm.AlterTable(db, "employees").
    AddColumn("nickname", "VARCHAR(64)").Migrate()
// ALTER TABLE clients
//      ADD COLUMN nickname VARCHAR(64);
```

`AddColumn`に続けて`NotNull`や`Default`を呼び出すことができます．
また，`NotNull`と`Default`は併用することができます．

### AddColumn.NotNull
`NotNull`はNOT NULL句を呼び出します．

#### 例
```go
err := mgorm.AlterTable(db, "clients").
    AddColumn("nickname", "VARCHAR(64)").
    NotNull().Migrate()
// ALTER TABLE clients
//      ADD COLUMN nickanme VARCHAR(64) NOT NULL;

err := mgorm.AlterTable(db, "clients").
    AddColumn("nickname", "VARCHAR(64)").
    NotNull().
    Default("none").Migrate()
// ALTER TABLE clients
//      ADD COLUMN nickanme VARCHAR(64) NOT NULL DEFAULT 'none';
```

### AddColumn.Default
`Default`はDEFAULT句を呼び出します．

引数には値を指定します．

#### 例
```go
err := mgorm.AlterTable(db, "clients").
    AddColumn("nickname", "VARCHAR(64)").
    Default("none").Migrate()
// ALTER TABLE clients
//      ADD COLUMN nickanme VARCHAR(64) DEFAULT 'none';

err := mgorm.AlterTable(db, "clients").
    AddColumn("nickname", "VARCHAR(64)").
    Default("none").
    NotNull().Migrate()
// ALTER TABLE clients
//      ADD COLUMN nickanme VARCHAR(64) DEFAULT 'none'  NOT NULL;
```


## DropColumn
`DropColumn`はDROP COLUMN句を呼び出します．

引数にはカラム名を指定します．

#### 例
```go
err := mgorm.AlterTable(db, "clients").
    DropColumn("nickname").Migrate()
// ALTER TABLE clients
//      DROP COLUMN nickname;
```


## RenameColumn
`RenameColumn`はRENAME COLUMN句を呼び出します．

引数には古いカラム名と新しいカラム名を指定します．

#### 例
```go
err := mgorm.AlterTable(db, "clients").
    RenameColumn("emp_no", "id").Migrate()
// ALTER TABLE clients
//      RENAME COLUMN emp_no TO id;
```


## AddCons
`AddCons`はADD CONSTRAINT句を呼び出します．

引数には制約名を指定します．

`AddCons`に続けて`Unique`，`Primary`，`Foreign`のいずれかを呼び出す必要があります．

### AddCons.Unique
`Unique`はUNIQUE句を呼び出します．

引数には複数カラム名を指定します．

#### 例
```go
err := mgorm.AlterTable(db, "employees").
    AddCons("UC_nickname").Unique("nickname").Migrate()
// ALTER TABLE employees
//      ADD CONSTRAINT UC_nickname UNIQUE (nickname);

err := mgorm.AlterTable(db, "employees").
    AddCons("UC_nickname").Unique("nickname", "first_name").Migrate()
// ALTER TABLE employees
//      ADD CONSTRAINT UC_nickname UNIQUE (nickname, first_name);
```

### AddCons.Primary
`Primary`はPRIMARY KEY句を呼び出します．

引数には複数カラム名を指定します．

#### 例
```go
err := mgorm.AlterTable(db, "employees").
    AddCons("PK_emp_no").Primary("emp_no").Migrate()
// ALTER TABLE employees
//      ADD CONSTRAINT PK_emp_no PRIMARY KEY (emp_no);

err := mgorm.AlterTable(db, "employees").
    AddCons("PK_emp_no").Primary("emp_no", "first_name").Migrate()
// ALTER TABLE employees
//      ADD CONSTRAINT PK_emp_no PRIMARY KEY (emp_no, first_name);
```

### AddCons.Foreign
`Foreign`はFOREIGN KEY句を呼び出します．

引数には複数カラム名を指定します．

`Foreign`に続けて`Ref`を呼び出す必要があります．

### AddCons.Foreign.Ref
`Ref`はREFERENCES句を呼び出します．

第一引数にテーブル名，第二引数以降に複数カラム名を指定します．

#### 例
```go
err := mgorm.AlterTable(db, "dept_emp").
    AddCons("FK_emp_no").Foreign("emp_no").Ref("employees", "emp_no").Migrate()
// ALTER TABLE dept_emp
//      ADD CONSTRAINT FK_emp_no FOREIGN KEY (emp_no) REFERENCES employees(emp_no);

err := mgorm.AlterTable(db, "dept_emp").
    AddCons("FK_emp_no").Foreign("emp_no", "from_date").Ref("employees", "emp_no", "hire_date").Migrate()
// ALTER TABLE dept_emp
//      ADD CONSTRAINT FK_emp_no FOREIGN KEY (emp_no, from_date) REFERENCES employees(emp_no, hire_date);
```
