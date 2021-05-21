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

しかし，実行可能メソッドやメソッドの実行順序に制限を設けてあります．
なぜなら，SQLの文法では句の順序が決まっており，SQL-likeなORMライブラリであるgsormもこの性質を受け継いでいるからです．

一見，制約が強くて使いづらいように思えますが，SQLは特にSELECT文において複雑になりやすく，ORMライブラリをシンプルにするほど実際にどのようなSQLが実行されているのかが分かりにくくなります．

また，SQLの文法をある程度知っている人なら，ライブラリの学習コストを抑えることができます．

さらに，メソッドの制約はinterfaceによって制御されているため，誤った方法でメソッドを呼び出してもコンパイルエラーとして検出されます．

```
// OFFSET句はLIMIT句が無いと使用できない．
SELECT * FROM people OFFSET 5;

// コンパイルエラー．
err := gsorm.Select(db, "id").From("people").Offset(5).Query(&person)

// SELECT文は複雑になりやすい．
SELECT id, name FROM people
  JOIN others ON people.id = others.id
  WHERE people.id > 100
  AND (others.id = 10 OR others.id IN (SELECT owner_id FROM teams WHERE name = "Fighters"));

// gsormならSQL-likeに実装できる．
err := gsorm.Select(db, "id", "name").From("people").
    Join("others").On("people.id = others.id").
    Where("people.id > ?", 100).
    And("others.id = 10 OR others.id IN (?)",
        gsorm.Select(nil, "owner_id").From("teams).Where("name = 'Fighters'"))
```


## Query
まずはQueryの使い方から簡単に紹介いたします．

SELECT文などのQueryを実行する際は，Queryというメソッドを実行することで実際にSQLが実行されます．

```go
// SELECT id FROM people;
err := gsorm.Select(db, "id").From("people").Query(&person)
```

詳細は[Select]()に記載されています．

また，このときQueryの引数にはmodelを渡すことができます．

modelには構造体のスライス，構造体，map，事前定義された型のスライス，事前定義された型の5種類の型を使用することができます．

ここでは例として，以下のような構造体を用いて説明いたします．

```go
type Person struct {
    ID        int
	FirstName string    `gsorm:"name typ=VARCHAR(64)"`
	BirthDate time.Time `gsorm:"layout=time.RFC3339"`
}
```

注目していただきたいのがフィールドタグです．

`FirstName`において`name`が指定されていますが，これはDBのテーブルにおけるカラム名を表しています．
つまり，`FirstName`は`name`というカラムにマッピングされます．

また，`BirthDate`には`layout=time.RFC3339`が指定されていますが，このように日付のフォーマットを指定することができます．

これらについての詳細は[Model]()に記載されています．


## Exec
ExecはINSERT文，UPDATE文，DELETE文などのSQLを実行する際に使用されるメソッドです．

```go
// INSERT INTO id, name VALUES (1, 'Taro');
err := gsorm.Insert(db, "people", "id", "name").Values(1, "Taro").Exec()
```

特にINSERT文とUPDATE文ではModelをそのままマッピングすることができます．

```go
person := Person{ID: 1, FirstName: "Taro"}

// INSERT INTO id, name VALUES (1, 'Taro');
err := gsorm.Insert(db, "people", "id", "name").Model(&person).Exec()
```

詳細は[Insert]()，[Update]()，[Delete]()に記載されています．


## Migrate
MigrateはCREATE TABLE文やALTER TABLE文など，データベース自体を変更するようなSQLを実行する際に使用されるメソッドです．

```go
// CREATE TABLE teams (id INT NOT NULL, name VARCHAR(64) NOT NULL);
err := gsorm.CreateTable(db, "teams").
    Column("id", "INT").NotNull().
    Column("name", "VARCHAR(64)").NotNull().Migrate()
```

また，CREATE TABLE文ではModelをそのまま使用することができます．

```go
type Person struct {
    ID        int    `gsorm:"notnull=t"`
    FirstName string `gsorm:"name typ=VARCHAR(64) notnull=t"`
}

// CREATE TABLE teams (id INT NOT NULL, name VARCHAR(64) NOT NULL);
err := gsorm.CreateTable(db, "teams").Model(&person).Migrate()
```

このとき，カラムの属性は構造体のフィールドタグによって指定することができます．

こちらについての詳細は[Model]()に記載されています．

Migrateに関連するメソッドの使用方法は[CreateDB]()，[CreateTable]()，[AlterTable]()に記載されています．


## Mock
gsormの特徴の1つとして，独自のmockを提供しているというところがあります．

```go
mock := gsorm.NewMock()

// あらかじめ，実行が予期されるSQLと返り値を指定する．
mock.Expect(gsorm.Select(db, "id", "name").From("people")).
    Return(&[]Person{{ID: 1, Name: "Taro"}, {ID: 2, Name: "Jiro"}})

// 実際に実行される．
err := func(db gsorm.Conn) {
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

[go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)と異なり生SQLを書かずに済むので使いやすいです．

詳細は[Mock]()に記載されています．

<br>

以上でIntroductionは終了となります！

詳しい使用方法についてはそれぞれ記載してありますのでぜひ読んでみてください！

<br>

**次の項目ヘ進む -> [Select](https://github.com/champon1020/gsorm/tree/main/docs/select_jp.md)**
