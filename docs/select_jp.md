# Select
`mgorm.Select`を使用したとき，`Query`を呼び出すことでデータベースからの検索結果をマッピングすることができます．

`mgorm.Select`の第1引数は`mgorm.Conn`の型，第2引数以降は複数のカラム名をstring型として受け取ることができます．
`mgorm.Conn`を実装した型としては`*mgorm.DB`，`*mgorm.Tx`，`*mgorm.MockDB`，`*mgorm.MockTx`があります．

詳細は[Transaction]()，[Mock]()に記載されています．

#### 例
```go
// SELECT id FROM people;
mgorm.Select(db, "id").From("people").Query(&model)
```


## メソッド
`mgorm.Select`を使用する際の文法を以下に示します．
各メソッドは上に行くほど呼び出しの優先度が高いです．

各メソッドについては順に説明いたします．

```
[]: optional,  |: or,  {}: one of them

mgorm.Select(DB, columns...).From(tables...)
    [{ .Join(tables) | .LeftJoin(table) | .RightJoin(table) }.On(expression)]
    [.Where(expression, values...)]
    [.Or(expression, values...) | .And(expression, values...)]
    [.GroupBy(columns...)]
    [.Having(expression, values...)]
    [.Union(*mgorm.SelectStmt) | .UnionAll(*mgorm.SelectStmt)]
    [.OrderBy(columns...)]
    [.Limit(number)]
    [.Offset(number)]
    .Query(*model)
```


## From
`From`は複数のテーブル名をstring方で受け取ります．
必要であれば，テーブル名にエイリアスを含めることができます．

#### 例
```go
// SELECT id FROM people;
mgorm.Select(db, "id").From("people").Query(&model)

// SELECT p.id, t.id FROM "people AS p", "teams AS t" LIMIT 10;
mgorm.Select(db, "p1.id", "p2.id").From("people AS p1", "people AS p2").Limit(10).Query(&model)
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
err := mgorm.Select(db, "p.id", "o.id").From("people AS p").Join("others AS o").On("p.id = o.id").Query(&model)

// SELECT p.id, o.id FROM people AS p
//  INNER JOIN others1 AS o1 ON p.id = o1.id
//  LEFT  JOIN others2 AS o2 ON p.id = o2.id;
err := mgorm.Select(db, "p.id", "o.id").From("people AS p").
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
err := mgorm.Select(db, "*").From("people").Where("id > ?", 10).Query(&model)

// SELECT * FROM people WHERE name LIKE '%Taro';
err := mgorm.Select(db, "*").From("people").Where("name LIKE ?", "%Taro").Query(&model)

// SELECT * FROM people WHERE id > 10 AND name Like '%Taro';
err := mgorm.Select(db, "*").From("people").Where("id > ? AND name LIKE ?", 10, "%Taro").Query(&model)
```


また，`Where`を用いることで副問合せを実行することもできます．
これは，値として`mgorm.Select`による文を渡すことで実現できます．

#### 例
```go
// SELECT * FROM people WHERE id IN (SELECT personal_id FROM companies WHERE name = 'ABC Company');
err := mgorm.Select(db, "*").From("people").Where("id IN (?)",
    mgorm.Select(nil, "personal_id").From("companies").Where("name = ?", "ABC Company")).Query(&model)
```


## Or / And
`Or`，`And`は`Where`とほぼ同様の使用方法になります．
ただし，これらを使用した場合の条件式は`(`と`)`で囲まれたものになります．

#### 例
```go
// SELECT * FROM people WHERE id > 10 AND (name = 'Taro' OR name = 'Jiro');
err := mgorm.Select(db, "*").From("people").Where("id > ?" 10).And("name = ? OR name = ?", "Taro", "Jiro").Query(&model)

// SELECT * FROM people WHERE id > 10 OR (id = 5 AND name = 'Saburo');
err := mgorm.Select(db, "*").From("people").Where("id > ?", 10).Or("id = ? AND name = ?", 5, "Saburo").Query(&model)
```

また，`Where`と同様に副問合せも用いることができます．

#### 例
```go
// SELECT * FROM people WHERE id > 10
//  OR (name = 'Saburo' AND id IN (SELECT personal_id FROM companies));
err := mgorm.Select(db, "*").From("people").Where("id > ?", 10).
    Or("name = ? AND id IN (?)", "Saburo", mgorm.Select(nil, "*").From("companies")).Query(&model)
```


## GroupBy
`GroupBy`は引数に複数のカラム名をstring型で受け取ります．


#### 例
```go
// SELECT COUNT(birth_date) FROM people GROUP BY birth_date;
err := mgorm.Select(db, "COUNT(birth_date)").From("people").GroupBy("birth_date").Query(&model)
```


## Having
`Having`は引数に条件式を受け取ります．
条件式の受け取り方は`Where`や`And`，`Or`と同様です．

基本的には`GroupBy`と共に使用しますが，`Having`のみで使用することも可能です．

#### 例
```go
// SELECT id, COUNT(birth_date) FROM people GROUP BY birth_date HAVING COUNT(birth_date) > 10;
err := mgorm.Select(db, "id, COUNT(birth_date)").From("people").
    GroupBy("birth_date").Having("COUNT(birth_date) > ?", 10).Query(&model)

// SELECT SUM(salary) FROM people HAVING SUM(salary) > 100000;
err := mgorm.Sum(db, "salary").From("people").Having("SUM(salary) > ?", 10000).Query(&model);
```


## Union
`Union`は引数に`*mgorm.SelectStmt`の型を受け取ります．
つまり，`mgorm.Select`による文を受け取ることができます．
ただし，引数に渡す`*mgorm.SelectStmt`では`OrderBy`，`Limit`，`Offset`を呼び出してはいけません．

また，`Union`と同様に`UnionAll`も使用することができます．

#### 例
```go
// SELECT id FROM people UNION (SELECT id FROM teams);
mgorm.Select(db, "id").From("people").Union(mgorm.Select(db, "id").From("teams")).Query(&model)

// SELECT id FROM people UNION ALL (SELECT id FROM teams);
mgorm.Select(db, "id").From("people").UnionAll(mgorm.Select(db, "id").From("teams")).Query(&model)
```


## OrderBy
`OrderBy`は引数に複数のカラム名をstring型で受け取ります．

必要であれば，カラム名に`DESC`や`ASC`などの順序の方向も含めてください．
`DESC`や`ASC`は小文字でも問題ありません．

#### 例
```go
// SELECT id FROM people ORDER BY id;
err := mgorm.Select(db, "id").From("people").OrderBy("id").Query(&model)

// SELECT id FROM people ORDER BY name DESC;
err := mgorm.Select(db, "id").From("people").OrderBy("name DESC").Query(&model)

// SELECT id FROM people ORDER BY id ORDER BY name DESC;
err := mgorm.Select(db, "id").From("people").OrderBy("id", "name DESC")
```


## Limit
`Limit`は引数にint型を受け取ります．

#### 例
```go
// SELECT id FROM people LIMIT 10;
err := mgorm.Select(db, "id").From("people").Limit(10).Query(&model)
```


## Offset
`Offset`は引数にint型を受け取ります．
`Offset`は`Limit`の直後のみ呼び出すことができます．

#### 例
```go
// SELECT id FROM people LIMIT 10 OFFSET 5;
err := mgorm.Select(db, "id").From("people").Limit(10).Offset(5).Query(&model)
```
