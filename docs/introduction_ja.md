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

実行可能メソッドやメソッドの呼び出し順序には制限を設けてあります．
なぜなら，SQLの文法では句の順序が決まっており，SQL-likeなORMライブラリであるgsormもこの性質を受け継いでいるからです．

一見，制約が強くて使いづらいように思えますが，SQLは特にSELECT文において複雑になりやすく，ORMライブラリをシンプルにするほど実際にどのようなSQLが実行されているのかが分かりにくくなります．

また，SQL-likeにすることで汎用性を高めることができます．

gsormでは`RawClause`というメソッドを設けているため，ユーザーが自由に句を追加することができます．

さらに，メソッドの制約はinterfaceによって制御されているため，間違った順序でメソッドを呼び出してもコンパイルエラーとして検出されます．

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
SELECT文を実行するとき，`Query`メソッドを実行することで実際にSQLが実行されます．

`Query`メソッドの引数にはmodelを渡すことができます．

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

詳細は[Select](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md)に記載されています．


## Exec
`Exec`メソッドは，INSERT文，UPDATE文，DELETE文などのSQLを実行するときに使用されるメソッドです．

#### 例
```go
// INSERT INTO id, name VALUES (1, 'Taro');
err := gsorm.Insert(db, "people", "id", "name").Values(1, "Taro").Exec()
```

特にINSERT文とUPDATE文では`Model`メソッドを使用することでマッピングをすることができます．

#### 例
```go
person := Person{ID: 1, FirstName: "Taro"}

// INSERT INTO id, name VALUES (1, 'Taro');
err := gsorm.Insert(db, "people", "id", "name").Model(&person).Exec()
```

詳細は[Insert](https://github.com/champon1020/gsorm/tree/main/docs/insert_ja.md)，[Update](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md)，[Delete](https://github.com/champon1020/gsorm/tree/main/docs/delete_ja.md)に記載されています．


## Migrate
`Migrate`メソッドは，CREATE TABLE文やALTER TABLE文など，データベーススキーマを変更するようなSQLを実行するときに使用されるメソッドです．

#### 例
```go
// CREATE TABLE teams (id INT NOT NULL, name VARCHAR(64) NOT NULL);
err := gsorm.CreateTable(db, "teams").
    Column("id", "INT").NotNull().
    Column("name", "VARCHAR(64)").NotNull().Migrate()
```

CREATE TABLE文では`Model`メソッドを使用することで，構造体をマッピングしてテーブルを作成することができます．

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

Modelについての詳細は[Model](https://github.com/champon1020/gsorm/tree/main/docs/model_ja.md)に記載されています．


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

<br>

以上でIntroductionは終了となります！

詳しい使用方法についてはそれぞれ記載してありますのでぜひ読んでみてください！

<br>

**次の項目ヘ進む -> [Select](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md)**
