# Select

## メソッド
`mgorm.Select`を使用する際の文法を以下に示します．
各メソッドは上に行くほど呼び出しの優先度が高いです．

各メソッドについては順に説明いたします．

```
[]: optional,  |: or,  {}: one of them

mgorm.Select(DB, columns...).From(table)
    [{ .Join(tables) | .LeftJoin(table) | .RightJoin(table) }.On(expression)]
    [.Where(expression, values...)]
    [.Or(expression, values...) | .And(expression, values...)]
    [.GroupBy(columns...)]
    [.Having(expression, values...)]
    [.OrderBy(columns...)]
    [.Limit(number)]
    [.Offset(number)]
    [.Union(*mgorm.Stmt) | .UnionAll(*mgorm.Stmt)]
```


## Join
JOIN句で使用できるのは`Join`，`LeftJoin`，`RightJoin`の3種類です．
ここで，`Join`はINNER JOIN句としてSQLを実行します．

`Join`は結合したいテーブル名を引数としてstring型で受け取ります．
このとき受け取ることができるテーブル名は1つのみです．
複数テーブルを結合したい場合は，連続して`Join`のメソッドを呼び出してください．

また，これらJOIN句に関するメソッドの後には`On`を呼び出す必要があります．
`On`には結合条件となる式を引数としてstring型で受け取ります．

これらの使用方法は，`LeftJoin`や`RightJoin`を使用する場合も同様です．

#### 例
```go
// SELECT p.id, o.id FROM people AS p INNER JOIN others AS o ON p.id = o.id;
mgorm.Select(db, "p.id", "o.id").From("people AS p").Join("others AS o").On("p.id = o.id").Query(&model)

// SELECT p.id, o.id FROM people AS p
//  INNER JOIN others1 AS o1 ON p.id = o1.id
//  LEFT  JOIN others2 AS o2 ON p.id = o2.id;
mgorm.Select(db, "p.id", "o.id").From("people AS p").
    Join("others1 AS o1").On("p.id = o1.id").
    LeftJoin("others2 AS o2").On("p.id = o2.id").Query(&model)
```


## Where
`Where`は引数に条件式を受け取ります．

条件式自体はstring型で受け取りますが，式の中に`?`を書くことで値を置き換えることができます．
このとき，値は複数置き換えることができます．

#### 例
```go
// SELECT * FROM people WHERE id > 10;
mgorm.Select(db, "*").From("people").Where("id > ?", 10).Query(&model)

// SELECT * FROM people WHERE name LIKE '%Taro';
mgorm.Select(db, "*").From("people").Where("name LIKE ?", "%Taro").Query(&model)

// SELECT * FROM people WHERE id > 10 AND name Like '%Taro';
mgorm.Select(db, "*").From("people").Where("id > ? AND name LIKE ?", 10, "%Taro").Query(&model)
```


`Where`を用いることで副問合せを実行することもできます．
これは，値として`mgorm.Select`による文を渡すことで実現できます．

#### 例
```go
// SELECT * FROM people WHERE id IN (SELECT personal_id FROM companies WHERE name = 'ABC Company');
mgorm.Select(db, "*").From("people").Where("id IN (?)",
    mgorm.Select(nil, "personal_id").From("companies").Where("name = ?", "ABC Company")).Query(&model)
```
