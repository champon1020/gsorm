# Fundamental
In gsorm, the SQL clauses are called by calling methods, and the SQL is executed by calling Query, Exec or Migrate methods.

Example implementation is as follows.

```go
// SELECT id FROM people;
err := gsorm.Select(db, "id").From("people").Query(&person)

// INSERT INTO id, name VALUES (1, 'Taro');
err := gsorm.Insert(db, "people", "id", "name").Values(1, "Taro").Exec()

// CREATE TABLE teams (id INT NOT NULL, name VARCHAR(64) NOT NULL);
err := gsorm.CreateTable(db, "teams").
    Column("id", "INT").NotNull().
    Column("name", "VARCHAR(64)").NotNull().Migrate()
```

The order of calling methods is restricted because gsorm provides the `RawClause` method which allows the flexible implementation.

Also, the restriction makes the implementation using gsorm understandable.

If the order of calling methods is wrong, a compile error will occurr instead of a runtime error because the restriction on method calls is controlled by interface.


## Query
`Query` executes the SELECT statement with fetching rows.

#### Example
```go
type Person struct {
    ID        int
	FirstName string    `gsorm:"name"`
	BirthDate time.Time
}

person := Person{}

// SELECT id FROM people;
err := gsorm.Select(db, "id").From("people").Query(&person)
```

You can determine the correspondence between field name of Go structure and column name of database.


## Exec
`Exec` executes the INSERT, UPDATE or DELETE statement without fetching rows.

#### Example
```go
// INSERT INTO id, name VALUES (1, 'Taro');
err := gsorm.Insert(db, "people", "id", "name").Values(1, "Taro").Exec()
```

Especially, `InsertStmt.Model` and `UpdateStmt.Model` methods map the Go structure into the SQL.


## Migarte
`Migrate` executes the database migrations.

#### Example1
```go
// CREATE TABLE teams (id INT NOT NULL, name VARCHAR(64) NOT NULL);
err := gsorm.CreateTable(db, "teams").
    Column("id", "INT").NotNull().
    Column("name", "VARCHAR(64)").NotNull().Migrate()
```

Expecially, `CreateTableStmt.Model` method maps the Go structure into the SQL.

#### Example2
```go
type Person struct {
    ID        int    `gsorm:"notnull=t"`
    FirstName string `gsorm:"name typ=VARCHAR(64) notnull=t"`
}

// CREATE TABLE teams (id INT NOT NULL, name VARCHAR(64) NOT NULL);
err := gsorm.CreateTable(db, "teams").Model(&person).Migrate()
```

In this case, the properties of database columns are determined by field tag of Go structure.


## Mock
Providing the own mock structure is one of the gsorm features.

#### Example
```go
type Employee struct {
    EmpNo     int
    Firstname string
}

func TestWithMock(t *testing.T) {
    // Open the connection to mock database.
    mock := gsorm.OpenMock()

    // Expect statements that will be executed.
    mock.Expect(gsorm.Insert(nil, "employees", "emp_no", "first_name").
        Values(1001, "Taro").
        Values(1002, "Jiro"))
    mock.ExpectWithReturn(gsorm.Select(nil, "emp_no", "first_name").From("employees"), []Employee{
        {ID: 1001, FirstName: "Taro"},
        {ID: 1002, FirstName: "Jiro"},
    })

    actualProcess := func(db gsorm.DB) error {
        model := []Employee{}

        if err := gsorm.Insert(nil, "employees", "emp_no", "first_name").
            Values(1001, "Taro").
            Values(1002, "Jiro").Exec(); err != nil {
            return err
        }

        if err := gsorm.Select(nil, "emp_no", "first_name").From("employees").Query(&model); err != nil {
            return err
        }

        return nil
    }

    if err := actualProcess(mock); err != nil {
        t.Error(err)
    }

    // Check if all expected statements was executed.
    if err := mock.Complete(); err != nil {
        t.Error(err)
    }
}
```

Unlike [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock), it is not required to write SQL with raw string.

Details are given in [Mock](https://github.com/champon1020/gsorm/tree/main/docs/mock.md).
