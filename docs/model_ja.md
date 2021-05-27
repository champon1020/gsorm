# Model
`Model`を使用することで構造体を用いたマッピングを行うことができます．
このメソッドは[Insert](https://github.com/champon1020/gsorm/tree/main/docs/insert_ja.md)，[Update](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md)，[CreateTable](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md)において使用することができます．

#### 例
```go
type Employee struct {
    EmpNo     int    `gsorm:"typ=INT"`
    FirstName string `gsorm:"typ=VARCHAR(14)"`
}

employee := Employee{EmpNo: 1000, FirstName: "Taro"}

err := gsorm.Insert(db, "employees", "emp_no", "first_name").
    Model(&employee).Exec()
```


## Type
`Model`の引数に渡せる値は以下の型の**参照**になります．

Insert，Update，CreateTableでそれぞれ使用できる型が異なるので注意してください．

### Insert
- map[string]interface{}
- struct{}
- []map[string]interface{}
- []struct{}

### Update
- map[string]interface{}
- struct{}

### CreateTable
- struct{}

構造体のフィールド，マップの要素に使用できる型は以下です．

- int, int8, int16, in32, in64
- uint, uint8, uint16, uint32, uint64
- float32, float64
- bool
- string
- time.Time

構造体はエクスポートされている型のみ適用されます．


## Tag
構造体のフィールドに`gsorm`タグを付与することでデータベースのスキーマを定義することができます．

複数のタグを使用したい場合はカンマで区切ることで指定できます．

#### 例
```go
type Employees struct {
    ID        int       `gsorm:"emp_no,typ=INT,notnull=t"`
    BirthDate time.Time `gsorm:"typ=DATE,notnull=t"`
    FirstName string    `gsorm:"typ=VARCHAR(14),notnull=t"`
    LastName  string    `gsorm:"typ=VARCHAR(16),notnull=t"`
    Gender    string    `gsorm:"typ=CHAR(3),notnull=t"`
    HireDate  time.Time `gsorm:"typ=DATE,notnull=t"`
}
```


### column
データベースのカラム名を指定します．
`gsorm`タグではなく`json`をタグを用いた場合にも反映されます．
両方指定してある場合は`gsorm`タグが優先されます．

#### 例
```go
type Employee struct {
    // emp_no
    ID        int    `gsorm:"emp_no"`
    // first_name
    FirstName string `json:"first_name"`
    // last_name
    LastName  string `gsorm:"last_name" json:"lastname"`
}
```


### typ
データベース型を指定します．

#### 例
```go
type Employee struct {
    EmpNo      int     `gsorm:"typ=INT"`
    FirstName  string  `gsorm:"typ=VARCHAR(14)"`
}
```


### notnull
カラムに`NOT NULL`オプションを付与します．

#### 例
```go
type Employee struct {
    EmpNo int `gsorm:"notnull=t"`
}
```


### default
カラムに`DEFAULT`オプションを付与します．

#### 例
```go
type Employee struct {
    EmpNo      int     `gsorm:"default=60000"`
    FirstName  string  `gsorm:"default='Taro'"`
}
```


### pk
カラムに`PRIMARY KEY`を付与します．

#### 例
```go
type Employee struct {
    // CONSTRAINT PK_emp_no PRIAMRY KEY (emp_no)
    EmpNo int `gsorm:"pk=PK_emp_no"`
}

type DeptEmp struct {
    // CONSTRAINT PK_emp_no PRIMARY KEY (emp_no, dept_no)
    EmpNo  int `gsorm:"pk=PK_emp_no"`
    DeptNo int `gsorm:"pk=PK_emp_no"`
}
```


### fk
カラムに`FOREIGN KEY`を付与します．

#### 例
```go
type DeptEmp struct {
    // CONSTRAINT FK_emp_no FOREIGN KEY (emp_no) REFERENCES employees(emp_no)
    EmpNo int `gsorm:"fk=FK_emp_no:employees(emp_no)"`
}

type DeptEmp struct {
    // CONSTRAINT FK_emp_no FOREIGN KEY (emp_no) REFERENCES employees(emp_no, dept_no)
    EmpNo  int `gsorm:"fk=FK_emp_no:employees(emp_no)"`
    DeptNo int `gsorm:"fk=FK_emp_no:employees(emp_no)"`
}
```


### uc
カラムに`UNIQUE`を付与します．

#### 例
```go
type Employee struct {
    // CONSTRAINT PK_emp_no UNIQUE (emp_no)
    EmpNo int `gsorm:"uc=UC_emp_no"`
}

type DeptEmp struct {
    // CONSTRAINT UC_emp_no UNIQUE (emp_no, dept_no)
    EmpNo  int `gsorm:"uc=UC_emp_no"`
    DeptNo int `gsorm:"uc=UC_emp_no"`
}
```
