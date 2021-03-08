# Fundamental

mgormでは，実行したいSQLに含まれる句をメソッドとして呼び出し，Query，Exec，Migrateのいずれかのメソッドを用いてSQLを実行します．
例えば，以下のようになります．

```go
// SELECT id FROM people;
err := mgorm.Select(db, "id").From("people").Query(&person)

// INSERT INTO id, name VALUES (1, 'Taro');
err := mgorm.Insert(db, "people", "id", "name").Values(1, "Taro").Exec()

// CREATE TABLE teams (id INT NOT NULL, name VARCHAR(64) NOT NULL);
err := mgorm.CreateTable(db, "teams").
    Column("id", "INT").NotNull().
    Column("name", "VARCHAR(64)").NotNull().Migrate()
```

しかし，実行できるメソッドには制限を設けてあります．
なぜなら，SQLの文法では句の順序が決まっており，SQL-likeなORMライブラリであるmgormもこの性質を受け継いでいるからです．
一見，制約が強くて使いづらいように思えますが，SQLは特にSELECT文において複雑になりやすく，ORMライブラリをシンプルにするほど実際にどのようなSQLが実行されているのかが分かりにくくなります．
また，どのようなSQLが実行されているかわからないと，ORMライブラリに詳しくない人が使い方を誤ってしまいランタイムエラーが引き起こされてしまう，というような事態もありえます．

```sql
-- OFFSET句はLIMIT句が無いと使用できない．--
SELECT * FROM people OFFSET 5;

-- SELECT文は複雑になりやすい．--
SELECT id, name FROM people
  JOIN others ON people.id = others.id
  WHERE people.id > 100
    AND (others.id = 10 OR others.id IN (SELECT owner_id FROM teams WHERE name = "Fighters"));
```

mgormではこのようなことを防ぐため，できるだけコンパイルエラーとして安全に解決する方法として以上のような制約を設けています．
また，基本的にSQLと同じような構文で実行できるため，学習コストも抑えることができます．


## Query
まずはQueryの使い方から簡単に紹介いたします．
SELECT文などのQueryを実行する際は，Queryというメソッドを実行することで実際にSQLが実行されます．

```go
// SELECT id FROM people;
err := mgorm.Select(db, "id").From("people").Query(&person)
```

このとき，Queryの引数にはmodelを渡すことができます．

modelには構造体のスライス，構造体，map，事前定義された型のスライス，事前定義された型の5種類の型を使用することができます．
ここでは例として，以下のような構造体を用いて説明いたします．

```go
type Person struct {
    ID        int
	FirstName string    `mgorm:"name typ=VARCHAR(64)"`
	BirthDate time.Time `mgorm:"layout=time.RFC3339"`
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
err := mgorm.Insert(db, "people", "id", "name").Values(1, "Taro").Exec()
```

特にINSERT文とUPDATE文ではModelをそのままマッピングすることができます．

```go
person := Person{ID: 1, FirstName: "Taro"}

// INSERT INTO id, name VALUES (1, 'Taro');
err := mgorm.Insert(db, "people", "id", "name").Model(&person).Exec()
```
