# Update
`mgorm.Update`を使用したとき，`Exec`を呼び出すことでテーブル上のカラムを更新することができます．

#### 例
```go
mgorm.Update(db).Set(10, "employees").
    Set("emp_no", 1001).
    Set("birth_date", "1995-07-07").
    Set("first_name", "Hanako").
    Set("last_name", "Suzuki").
    Set("gender", "W").
    Set("hire_date", time.Date(2019, time.September, 1, 0, 0, 0, 0, time.UTC)).Exec()
// UPDATE employees
//  SET emp_no=1001,
//      birth_date='1995-07-07',
//      first_name='Hanako',
//      last_name='Suzuki',
//      gender='W',
//      hire_date='2019-09-01';
```


# Methods
`mgorm.Update`で使用できるメソッドを以下に示します．

- [Set](https://github.com/champon1020/mgorm/tree/main/docs/update_jp.md#set)
- [Where](https://github.com/champon1020/mgorm/tree/main/docs/update_jp.md#where)
- [And / Or](https://github.com/champon1020/mgorm/tree/main/docs/update_jp.md#and--or)
- [Model](https://github.com/champon1020/mgorm/tree/main/docs/update_jp.md#model)

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

mgorm.Update(DB, table, columns...)
    ( .Set {.Set} [.Where {.And | .Or}] ) | ( .Model )
    .Exec()
```

例えば以下に反した場合，コンパイルエラーを吐き出します．

```go
// NG
err := mgorm.Update(db, "employees", "emp_no", "first_name").
    Set("emp_no", 1001).
    Set("first_name", "Hanako").
    And("emp_no < ? AND first_name = ?", 1000, "Taro")
    Where("emp_no > ?", 1000).Exec()
```


## Set
`Set`はSET句を呼び出します．

`Set`は複数回呼び出すことが可能です．


#### 例
```go
mgorm.Update(db, "employees").
    Set("emp_no", 1001).
    Set("first_name", "Hanako").Exec()
// UPDATE employees
//  SET emp_no=1001,
//      first_name='Hanako';
```


## Where
`Where`は引数に条件式を受け取ります．

詳しい使用方法は`mgorm.Select`における[Where]()と同様です．

#### 例
```go
mgorm.Update(db, "employees").
    Set("first_name", "Jiro").
    Where("emp_no = ?", 1000).Exec()
// UPDATE employees
//  SET first_name='Jiro'
//      last_name='Kaneko'
//  WHERE emp_no = 1000;
```


## And / Or
`And / Or`は引数に条件式を受け取ります．

詳しい使用方法は`mgorm.Select`における[And / Or]()と同様です．

#### 例
```go
mgorm.Update(db, "employees", "first_name", "last_name").
    Set("first_name", "Jiro").
    Set("last_name", "Kaneko").
    Where("emp_no = ?", 1000).
    And("first_name = ? OR last_name = ?", "Taro", "Sato").Exec()
// UPDATE employees
//  SET first_name='Jiro'
//      last_name='Kaneko'
//  WHERE emp_no = 1000
//  AND (first_name = 'Taro' OR last_name = 'Sato');

mgorm.Update(db, "employees", "first_name", "last_name").
    Set("first_name", "Kaneko").
    Set("last_name", "Kaneko").
    Where("emp_no > ?", 1000).
    And("emp_no < ? AND first_name = ?", 1000, "Taro").Exec()
// UPDATE employees
//  SET first_name='Jiro'
//      last_name='Kaneko'
//  WHERE emp_no > 1000
//  OR (emp_no < 1000 AND first_name = 'Taro');
```


## Model
`mgorm.Update`を使用すとき，`Model`を使用することで構造体をマッピングしてカラムを更新することができます．

`Model`は引数として構造体のポインタ，構造体スライスのポインタ，マップ型のポインタなどを受け取ることができます．

また，フィールドタグを変更することで対応するカラム名を変更することができます．
指定しない場合は，フィールド名のスネークケースとなります．

Modelの型やタグについての詳細は[Model]()に記載されています．

#### 例
```go
type Employee struct {
    ID        int       `mgorm:"emp_no"`
    BirthDate time.Time
    FirstName string
    LastName  string
    Gender    string
    HireDate  string
}

emp1 := Employee{ID: 1000, FirstName: "Taro"}

mgorm.Update(db, "employees").
    Model(&emp1, "emp_no", "first_name").Exec()
// UPDATE employees
//  SET emp_no=1000,
//      first_name='Taro';

emp2 = Employee{
    EmpNo: 1000,
    BirthDate: time.Date(1965, time.April, 4, 0, 0, 0, 0, time.UTC),
    FirstName: "Taro",
    LastName: "Sato",
    Gender: "M",
    HireDate: "1988-04-01",
}

mgorm.Update(db, "employees").
    Model(&emp2).Exec()
// UPDATE employees
//  SET emp_no=1000,
//      birth_date='1965-04-04 00:00:00'
//      first_name='Taro',
//      last_name='Sato',
//      gender='M',
//      hire_date='1988-04-01';
```
