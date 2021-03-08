# Fundamental

mgormでは，実行したいSQLに含まれる句をメソッドとして呼び出し，Query，Exec，Migrateのいずれかのメソッドを用いてSQLを実行します．
例えば，以下のようになります．
```
err := mgorm.Select(db, "*").From("persons").Query(&person)

err := mgorm.Insert(db, "id", "name").Values(1, "Taro").Exec()

err := mgorm.CreateTable(db, "")
```

しかし，実行できるメソッドには制限を設けてあります．なぜなら，SQLの文法では句の順序が決まっており，SQL-likeなORMライブラリであるmgormもこの性質を受け継いでいるからです．
一見，制約が強くて使いづらいように思えますが，SQLは特にSELECT文において複雑になりやすく，ORMライブラリをシンプルにするほど実際にどのようなSQLが実行されているのかが分かりにくくなります．
また，どのようなSQLが実行されているかわからないと，ORMライブラリに詳しくない人が使い方を誤ってしまいランタイムエラーが引き起こされてしまう，というような事態もありえます．
<例>

mgormではこのようなことを防ぐため，できるだけコンパイルエラーとして安全に解決する方法として以上のような制約を設けています．
また，基本的にSQLと同じような構文で実行できるため，学習コストも抑えることができます．


## Query
さて，まずはQueryの使い方から簡単に紹介いたします．
SELECT文などのQueryを実行する際は，Queryというメソッドを実行することで実際にSQLが実行されます．

<例>

このとき，Queryの引数にはmodelを渡すことができます．

modelには構造体のスライス，構造体，map，事前定義された型のスライス，事前定義された型の5種類の型を使用することができます．
ここでは例として，以下のような構造体を用いて説明いたします．
```
type Person struct {
	ID        int
	Name      string    `mgorm:"first_name typ=VARCHAR(64)"`
	BirthDate time.Time `mgorm:"layout=time.RFC3339"`
}
```
注目していただきたいのがフィールドタグです．
`Name`において`first_name`が指定されていますが，これはDBのテーブルにおけるカラム名を表しています．つまり，`Name`は`first_name`というカラムにマッピングされます．
また，`BirthDate`には`layout=time.RFC3339`が指定されていますが，このように日付のフォーマットを指定することができます．
これらについての詳細はは[model]()に記載されています．
