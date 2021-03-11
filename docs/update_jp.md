# Update
`mgorm.Update`を使用したとき，`Exec`を呼び出すことでテーブル上のカラムを更新することができます．

`mgorm.Update`の第1引数は`mgorm.Conn`の型，第2引数はテーブル名をstring型として，第3引数以降は複数のカラム名をstring型として受け取ることができます．
カラム名を指定しない場合は，全てのカラムとして適用されます．

`mgorm.Conn`を実装した型としては`*mgorm.DB`，`*mgorm.Tx`，`*mgorm.MockDB`，`*mgorm.MockTx`があります．

詳細は[Transaction]()，[Mock]()に記載されています．

#### 例
```go
// UPDATE employees
//  SET emp_no=1001,
//      birth_date='1995-07-07',
//      first_name='Hanako',
//      last_name='Suzuki',
//      gender='W',
//      hire_date='2019-09-01';
mgorm.Update(db).Set(10, "employees").
    Set(1001,
        "1995-07-07",
        "Hanako",
        "Suzuki",
        "W",
        "2019-09-01").Exec()

// UPDATE employees
//  SET emp_no=1001,
//      first_name='Hanako';
mgorm.Update(db, "employees", "emp_no", "first_name").
    Set(1001, "Hanako").Exec()
```


## Set
`mgorm.Update`を用いてカラムを更新するとき，`Set`を用いることで値を更新することができます．

`mgorm.Update`にカラム名を渡した場合，`Set`の引数の数は渡したカラム名の数と等しくなければいけません．

#### 例
```go
// UPDATE employees
//  SET emp_no=1001,
//      first_name='Hanako';
mgorm.Update(db, "employees", "emp_no", "first_name").
    Set(1001, "Hanako").Exec()
```


## Where
`Where`は引数に条件式を受け取ります．

詳しい使用方法は`mgorm.Select`における[Where]()と同様です．

#### 例
```go
// UPDATE employees
//  SET first_name='Jiro'
//      last_name='Kaneko'
//  WHERE emp_no = 1000;
mgorm.Update(db, "employees", "first_name").
    Set("Jiro").
    Where("emp_no = ?", 1000).Exec()
```


## And / Or
`And / Or`は引数に条件式を受け取ります．

詳しい使用方法は`mgorm.Select`における[And / Or]()と同様です．

#### 例
```go
// UPDATE employees
//  SET first_name='Jiro'
//      last_name='Kaneko'
//  WHERE emp_no = 1000
//  AND (first_name = 'Taro' OR last_name = 'Sato');
mgorm.Update(db, "employees", "first_name", "last_name").
    Set("Jiro", "Kaneko").
    Where("emp_no = ?", 1000).
    And("first_name = ? OR last_name = ?", "Taro", "Sato").Exec()

// UPDATE employees
//  SET first_name='Jiro'
//      last_name='Kaneko'
//  WHERE emp_no > 1000
//  OR (emp_no < 1000 AND first_name = 'Taro');
mgorm.Update(db, "employees", "first_name", "last_name").
    Set("Jiro", "Kaneko").
    Where("emp_no > ?", 1000).
    And("emp_no < ? AND first_name = ?", 1000, "Taro").Exec()
```
