# Introduction

gsormでは，実行したいSQLに含まれる句をメソッドとして呼び出し，Query，Exec，Migrateのいずれかのメソッドを用いてSQLを実行します．

例えば，以下のようになります．

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

gsormは柔軟な実装を可能にする`RawClause`メソッドを提供しているため，メソッドの呼び出し順序には制限を設けてあります．

また，メソッドの呼び出し順序を制限することで理解しやすい実装をすることができます．

もし間違った順序でメソッドを呼び出しても，メソッドの呼び出し順序はinterfaceによって制御されているため，ランタイムエラーではなくコンパイルエラーが起こります．


```
// コンパイルエラー：OFFSET句はLIMIT句が無いと使用できない．
err := gsorm.Select(db, "id").From("people").Offset(5).Query(&person)

// gsormなら複雑なSELECT文でもSQLのように実装することができる．
err := gsorm.Select(db, "id", "name").From("people").
    Join("others").On("people.id = others.id").
    Where("people.id > ?", 100).
    And("others.id = 10 OR others.id IN (?)",
        gsorm.Select(nil, "owner_id").From("teams).Where("name = 'Fighters'"))
```


## Query
`Query`はSELECT文を実行し，データを取得します．

#### 例
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

注目していただきたいのがフィールドタグです．

`FirstName`において`name`が指定されていますが，これはDBのテーブルにおけるカラム名を表しています．
つまり，`FirstName`は`name`というカラムにマッピングされます．


## Exec
`Exec`はINSERT文，UPDATE文，DELETE文を実行します．

#### 例
```go
// INSERT INTO id, name VALUES (1, 'Taro');
err := gsorm.Insert(db, "people", "id", "name").Values(1, "Taro").Exec()
```

特にINSERT文とUPDATE文では`Model`メソッドを使用することで，構造体をSQLへマッピングをすることができます．

#### 例
```go
person := Person{ID: 1, FirstName: "Taro"}

// INSERT INTO id, name VALUES (1, 'Taro');
err := gsorm.Insert(db, "people", "id", "name").Model(&person).Exec()
```


## Migrate
`Migrate`はCREATE TABLE文やALTER TABLE文などのマイグレーションを実行します．

#### 例
```go
// CREATE TABLE teams (id INT NOT NULL, name VARCHAR(64) NOT NULL);
err := gsorm.CreateTable(db, "teams").
    Column("id", "INT").NotNull().
    Column("name", "VARCHAR(64)").NotNull().Migrate()
```

特にCREATE TABLE文では`Model`メソッドを使用することで，構造体をSQLへマッピングすることができます．

#### 例
```go
type Person struct {
    ID        int    `gsorm:"notnull=t"`
    FirstName string `gsorm:"name typ=VARCHAR(64) notnull=t"`
}

// CREATE TABLE teams (id INT NOT NULL, name VARCHAR(64) NOT NULL);
err := gsorm.CreateTable(db, "teams").Model(&person).Migrate()
```

このとき，カラムの属性は構造体のフィールドタグによって指定することができます．


## Mock
gsormの特徴の1つとして，独自のmockを提供しているというところがあります．

#### 例
```go
mock := gsorm.NewMock()

// あらかじめ，実行が予期されるSQLと返り値を指定する．
mock.Expect(gsorm.Select(db, "id", "name").From("people")).
    Return(&[]Person{{ID: 1, Name: "Taro"}, {ID: 2, Name: "Jiro"}})

// 実際に実行される．
err := func(db gsorm.Conn) error {
    person := []Person{}

    err := db.Select(db, "id", "name").From("people").Query(&person)
    if err != nil {
        return err
    }

    return nil
}(mock)

// Expectされたものが全て実行されたかチェックする．
err := mock.Complete()
if err != nil{
    fmt.Fatal(err)
}
```

[go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)と異なり，生文字列でSQLを書かずに済みます．

詳細は[Mock](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md)に記載されています．
