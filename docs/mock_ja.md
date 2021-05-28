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


# Methods
`gsorm.MockDB`が実装するメソッドは以下です．

また，`gsorm.MockDB`は`gsorm.DB`が実装するメソッドも使用することができます．

- [Expect](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#expect)
- [ExpectWithReturn](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#expectwithreturn)
- [Complete](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#complete)
- [ExpectBegin](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#expectbegin)


### Expect
`Expect`は返り値がない文が実行されることを予期します．

#### 例
```go
gsorm.Expect(gsorm.Insert(nil, "employees", "emp_no", "first_name").Values(1001, "Taro"))
```


### ExpectWithReturn
`ExpectWithReturn`は返り値がある文が実行されることを予期します．

#### 例
```go
gsorm.ExpectWithReturn(
	gsorm.Select(nil, "id", "first_name").From("employees"),
	[]Employee{{EmpNo: 1001, FirstName: "Taro"}, {EmpNo: 1002, FirstName: "Jiro"}})
```


### Complete
`Complete`は予期した文が全て実行されたかどうがを確認するメソッドです．

もし実行されていない文が存在する場合，エラーを返します．

mockを用いたテストを行う場合は，必ずこちらの関数を呼び出す必要があります．

#### 例
```go
if err := gsorm.Complete(); err != nil {
	log.Fatal(err)
}
```


### ExpectBegin
`ExpectBegin`はトランザクションが作成されることを予期します．

このとき，`gsorm.MockTx`を実装するインスタンスが返されます．

`gsorm.MockTx`についての詳細は[MockTx]()に記載されています．

#### 例
```go
mocktx := mock.ExpectBegin()
```
