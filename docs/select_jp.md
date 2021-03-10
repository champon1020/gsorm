# Select

## メソッド
`mgorm.Select`を使用する際の文法を以下に示します．
また，各メソッドについては順に説明いたします．

```
mgorm.Select(DB, columns...).From(table)
    [{ .Join(tables) | .LeftJoin(table) | .RightJoin(table) }.On(expression)]
    [.Where(expression, values...)]
    [.Or(expression, values...) | .And(expression, values...)]
    [.GroupBy(columns...)] [.Having(expression, values...)]
    [.OrderBy(columns...)]
    [.Limit(number)] [.Offset(number)]
    [.Union(*mgorm.Stmt) | .UnionAll(*mgorm.Stmt)]
```


## Join
JOIN句で使用できるのは`Join`，`LeftJoin`，`RightJoin`の3種類です．
ここで，`Join`はINNER JOIN句とSQLを実行します．

`Join`は結合したいテーブル名を引数としてstring型で受け取ります．

また，これらJOIN句に関するメソッドの後には`On`を呼び出す必要があります．

`On`には結合条件となる式を引数としてstring型で受け取ります．

```go
mgorm.Select(db, "p.id", "o.id").From("people AS p").Join("others AS o").On("p.id = o.id").Query(&model)
```

これらの使用方法は，`LeftJoin`や`RightJoin`を使用する場合も同様です．
