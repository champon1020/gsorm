# Update
`gsorm.Update`はUPDATE句を呼び出します．

引数にはデータベースのコネクション(`gsorm.Conn`)，テーブル名を指定します．

#### 例
```go
gsorm.Update(db).Set(10, "employees").
    Set("emp_no", 1001).
    Set("birth_date", "1995-07-07").
    Set("first_name", "Hanako").
    Set("last_name", "Suzuki").
    Set("gender", "W").
    Set("hire_date", time.Date(2019, time.September, 1, 0, 0, 0, 0, time.UTC)).Exec()
// UPDATE employees
//      SET emp_no = 1001,
//          birth_date = '1995-07-07',
//          first_name = 'Hanako',
//          last_name = 'Suzuki',
//          gender = 'W',
//          hire_date = '2019-09-01';
```


# Methods
`gsorm.Update`で使用できるメソッドは以下です．

- [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw_ja.md#rawclause)
- [Set](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#set)
- [Where](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#where)
- [And](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#and)
- [Or](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#or)
- [Model](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#model)

これらのメソッドは以下のEBNFに従って実行することができます．
但し，例外として`RawClause`は任意で呼び出すことができます．

```
| alternation
() grouping
[] option (0 to 1 times)
{} repetition (0 to n times)

gsorm.Update(DB, table, columns...)
    (.Set {.Set}) | .Model
    [.Where [{.And} | {.Or}]]
    .Exec
```

例えば以下の実装はコンパイルエラーを吐き出します．

```go
// NG
err := gsorm.Update(db, "employees", "emp_no", "first_name").Exec()

// NG
err := gsorm.Update(db, "employees", "emp_no", "first_name").
    Set("emp_no", 1001).
    Set("first_name", "Hanako").
    And("emp_no < ? AND first_name = ?", 1000, "Taro")
    Where("emp_no > ?", 1000).Exec()
```


## Set
`Set`はSET句を呼び出します．

引数にはカラム名と値を指定します．

`Set`は複数回呼び出すことが可能です．


#### 例
```go
err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").Exec()
// UPDATE employees
//      SET first_name = 'Hanako';

err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").
    Set("last_name", "Suzuki").Exec()
// UPDATE employees
//      SET first_name = 'Hanako',
//          last_name = 'Suzuki';
```


## Where
`Where`はWHERE句を呼び出します．

第1引数に条件式，第2引数以降に複数値を指定できます．
この際，条件式における`?`に値が代入されます．
また，代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が展開されます．
- 値が`*gsorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない値はそのまま展開される．

#### 例
```go
err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").
    Where("emp_no = 1001").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001;

err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001;

err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").
    Where("first_name = ?", "Taro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE first_name = 'Taro';

err := gsorm.Update(db, "employees").
    Set("first_name", "Hanako").
    Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE birth_date = '2006-01-02 00:00:00';

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("first_name LIKE ?", "%Taro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE first_name LIKE '%Taro';

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no BETWEEN 1001 AND 1003;

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no IN (?)", []int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no IN (1001, 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no IN (1001, 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no IN (SELECT emp_no FROM dept_manager);
```


## And
`And`はAND句を呼び出します．
このとき実行されるSQLは，条件式が`()`で括られた形となります．

第1引数に条件式，第2引数以降に複数値を指定できます．
この際，条件式における`?`に値が代入されます．
また，代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が展開されます．
- 値が`*gsorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`And`は複数回呼び出すことができます．

#### 例
```go
err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no = 1002").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no = ?", 1002).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("first_name = ? OR first_name = ?", "Taro", "Jiro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (first_name = 'Taro' OR first_name = 'Jiro');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no = ?", 1002).
    And("emp_no = ?", 1003).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no = 1002);
//      AND (emp_no = 1003);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (birth_date = '2006-01-02 00:00:00');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("first_name LIKE ?", "%Taro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (first_name LIKE '%Taro');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no BETWEEN 1001 AND 1003);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", []int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no IN (1001, 1002));

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no IN (1001, 1002));

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    And("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      AND (emp_no IN (SELECT emp_no FROM dept_manager));
```

## Or
`Or`はOR句を呼び出します．
このとき実行されるSQLは，条件式が`()`で括られた形となります．

第1引数に条件式，第2引数以降に複数値を指定できます．
この際，条件式における`?`に値が代入されます．
また，代入規則は以下に従います．

- 値が`string`型もしくは`time.Time`型の場合，値はシングルクオートで囲まれます．
- 値が事前定義型のスライスもしくは配列の場合，その要素が展開されます．
- 値が`*gsorm.SelectStmt`型の場合，SELECT文が展開されます．
- 以上の条件に該当しない値はそのまま展開される．

`Or`は複数回呼び出すことができます．

#### 例
```go
err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no = 1002").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no = ? AND first_name = ?", 1002, "Taro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no = 1002 OR first_name = 'Taro');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no = ?", 1002).
    Or("emp_no = ?", 1003).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no = 1002);
//      OR (emp_no = 1003);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (birth_date = '2006-01-02 00:00:00');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("first_name LIKE ?", "%Taro").Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (first_name LIKE '%Taro');

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no BETWEEN ? AND ?", 1001, 1003).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no BETWEEN 1001 AND 1003);

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", []int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no IN (1001, 1002));

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", [2]int{1001, 1002}).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no IN (1001, 1002));

err := gsorm.Update(db).From("employees").
    Set("first_name", "Hanako").
    Where("emp_no = ?", 1001).
    Or("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).Exec()
// UPDATE employees
//      SET first_name = 'Hanako'
//      WHERE emp_no = 1001
//      OR (emp_no IN (SELECT emp_no FROM dept_manager));
```


## Model
`Model`は構造体をマッピングします．

引数にはモデルのポインタ，複数カラム名をしてします．

モデルには構造体のポインタ，マップのポインタのいずれかを指定します．

構造体もしくは構造体のスライスをマッピングする際，対象のカラム名はフィールド名もしくはフィールドタグから推定されます．

Modelについての詳細は[Model](https://github.com/champon1020/gsorm/blob/main/docs/model_ja.md)に記載されています．

#### 例
```go
type Employee struct {
    ID        int       `gsorm:"emp_no"`
    BirthDate time.Time
    FirstName string
    LastName  string
    Gender    string
    HireDate  string
}

emp1 := Employee{ID: 1000, FirstName: "Taro"}

gsorm.Update(db, "employees").
    Model(&emp1, "emp_no", "first_name").Exec()
// UPDATE employees
//  SET emp_no = 1000,
//      first_name = 'Taro';

emp2 = Employee{
    EmpNo: 1000,
    BirthDate: time.Date(1965, time.April, 4, 0, 0, 0, 0, time.UTC),
    FirstName: "Taro",
    LastName: "Sato",
    Gender: "M",
    HireDate: "1988-04-01",
}

gsorm.Update(db, "employees").
    Model(&emp2).Exec()
// UPDATE employees
//  SET emp_no = 1000,
//      birth_date = '1965-04-04 00:00:00'
//      first_name = 'Taro',
//      last_name = 'Sato',
//      gender = 'M',
//      hire_date = '1988-04-01';
```
