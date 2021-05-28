# Mock
gsormは独自のモックを提供しています．

gsormのモックは生文字列を指定せずに使用することができます．

#### 例
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
`gsorm.MockDB`が実装するメソッドは以下です．

また，`gsorm.MockDB`は`gsorm.DB`が実装するメソッドも使用することができます．

- [Expect](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#expect)
- [ExpectWithReturn](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#expectwithreturn)
- [Complete](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#complete)
- [ExpectBegin](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#expectbegin)


## Expect
`Expect`は返り値がない文が実行されることを予期します．

#### 例
```go
mock.Expect(gsorm.Insert(nil, "employees", "emp_no", "first_name").Values(1001, "Taro"))
```


## ExpectWithReturn
`ExpectWithReturn`は返り値がある文が実行されることを予期します．

#### 例
```go
mock.ExpectWithReturn(
	gsorm.Select(nil, "id", "first_name").From("employees"),
	[]Employee{{EmpNo: 1001, FirstName: "Taro"}, {EmpNo: 1002, FirstName: "Jiro"}})
```


## Complete
`Complete`は予期した文が全て実行されたかどうがを確認するメソッドです．

もし実行されていない文が存在する場合，エラーを返します．

mockを用いたテストを行う場合は，必ずこちらの関数を呼び出す必要があります．

#### 例
```go
if err := gsorm.Complete(); err != nil {
	log.Fatal(err)
}
```


## ExpectBegin
`ExpectBegin`はトランザクションが作成されることを予期します．

このとき，`gsorm.MockTx`を実装するインスタンスが返されます．

`gsorm.MockTx`についての詳細は[MockTx](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mocktx)に記載されています．

#### 例
```go
mocktx := mock.ExpectBegin()
```


# MockTx
## Methods
`gsorm.MockTx`が実装するメソッドは以下です．

また，`gsorm.MockTx`は`gsorm.Tx`が実装するメソッドも使用することができます．

基本的には`gsorm.MockDB`と同一ですが，`ExpectCommit`や`ExpectRollback`を使用することができます．

- [Expect](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#(mocktx).expect)
- [ExpectWithReturn](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#(mocktx).expectwithreturn)
- [Complete](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#(mocktx).complete)
- [ExpectCommit](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#(mocktx).expectcommit)
- [ExpectRollback](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#(mocktx).expectrollback)

## (MockTx).Expect
`Expect`は返り値がない文が実行されることを予期します．

#### 例
```go
mocktx.Expect(gsorm.Insert(nil, "employees", "emp_no", "first_name").Values(1001, "Taro"))
```


## (MockTx).ExpectWithReturn
`ExpectWithReturn`は返り値がある文が実行されることを予期します．

#### 例
```go
mocktx.ExpectWithReturn(
	gsorm.Select(nil, "id", "first_name").From("employees"),
	[]Employee{{EmpNo: 1001, FirstName: "Taro"}, {EmpNo: 1002, FirstName: "Jiro"}})
```


## (MockTx).Complete
通常は`gsorm.MockDB`において`Complete`を実行すると，作成されたトランザクションの`Complete`も反映されます．
そのため，ユーザーはこのメソッドを意識する必要はありません．


## (MockTx).ExpectCommit
`ExpectCommit`はトランザクションのCommitを予期します．

#### 例
```go
mocktx.ExpectCommit()
```


## (MockTx).ExpectRollback
`ExpectRollback`はトランザクションのRollbackを予期します．

#### 例
```go
mocktx.ExpectRollback()
```
