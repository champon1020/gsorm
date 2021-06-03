# Mock
gsorm provides the mock.

gsorm's mock can be used without using raw SQL.

#### Example
```go
mock := gsorm.OpenMock()

mock.Expect(gsorm.Insert(nil, "employees", "emp_no", "first_name").Values(1001, "Taro"))

mock.ExpectWithReturn(
	gsorm.Select(nil, "emp_no", "first_name").From("employees"),
	[]Employee{{EmpNo: 1001, FirstName: "Taro"}, {EmpNo: 1002, FirstName: "Jiro"}})

process := func(db gsorm.Conn) error {
	if err := gsorm.Insert(db, "employees", "emp_no", "first_name").Values(1001, "Taro").Exec(); err != nil {
		return err
	}

	emp := []Employee{}
	if err := gsorm.Select(nil, "emp_no", "first_name").From("employees").Query(&emp); err != nil {
		return err
	}

	return nil
}

if err := process(mock); err != nil {
	log.Fatal(err)
}

if err := mock.Complete(); err != nil {
	log.Fatal(err)
}
```


# MockDB
## Methods
- [Expect](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mockdbexpect)
- [ExpectWithReturn](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mockdbexpectwithreturn)
- [Complete](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mockdbcomplete)
- [ExpectBegin](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mockdbexpectbegin)


## (MockDB).Expect
`Expect` expects the SQL statement without return value.

#### Example
```go
mock.Expect(gsorm.Insert(nil, "employees", "emp_no", "first_name").Values(1001, "Taro"))
```


## (MockDB).ExpectWithReturn
`ExpectWithReturn` expects the SQL statement with specifing return value.

#### Exmample
```go
mock.ExpectWithReturn(
	gsorm.Select(nil, "id", "first_name").From("employees"),
	[]Employee{{EmpNo: 1001, FirstName: "Taro"}, {EmpNo: 1002, FirstName: "Jiro"}})
```


## (MockDB).Complete
`Complete` validates whether all expected statements are executed.

If there are the statements that are not executed but expected, `Complete` returns error.

`Complete` should be executed at last when the testing with mock.

#### Example
```go
if err := gsorm.Complete(); err != nil {
	log.Fatal(err)
}
```


## (MockDB).ExpectBegin
`ExpectBegin` expects the execution of `Begin` method.

`ExpectBegin` returns the instance that implements `gsorm.MockTx`.

Details of `gsorm.MockTx` is in [MockTx](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktx).

#### Example
```go
mocktx := mock.ExpectBegin()
```


# MockTx
## Methods
- [Expect](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktxexpect)
- [ExpectWithReturn](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktxexpectwithreturn)
- [ExpectCommit](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktxexpectcommit)
- [ExpectRollback](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktxexpectrollback)


## (MockTx).Expect
`Expect` expects the SQL statement without return value.

#### Example
```go
mocktx.Expect(gsorm.Insert(nil, "employees", "emp_no", "first_name").Values(1001, "Taro"))
```


## (MockTx).ExpectWithReturn
`ExpectWithReturn` expects the SQL statement with specifing return value.

#### Example
```go
mocktx.ExpectWithReturn(
	gsorm.Select(nil, "id", "first_name").From("employees"),
	[]Employee{{EmpNo: 1001, FirstName: "Taro"}, {EmpNo: 1002, FirstName: "Jiro"}})
```


## (MockTx).ExpectCommit
`ExpectCommit` expects the transaction commit.

#### Example
```go
mocktx.ExpectCommit()
```


## (MockTx).ExpectRollback
`ExpectRollback` expects the transaction rollback.

#### Example
```go
mocktx.ExpectRollback()
```
