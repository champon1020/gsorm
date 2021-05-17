# Model
`Model`を使用することで構造体を用いたマッピングを行うことができます．
このメソッドは[Insert](https://github.com/champon1020/mgorm/tree/main/docs/insert_ja.md)，[Update](https://github.com/champon1020/mgorm/tree/main/docs/update_ja.md)，[CreateTable](https://github.com/champon1020/mgorm/tree/main/docs/createtable_ja.md)において使用することができます．

`Model`の引数には，スライス，構造体，マップ，`int`や`string`などの事前定義型を渡すことができます．
この際，必ず参照を渡すようにしてください．

#### 例
```go
type Employee struct {
    EmpNo     int    `mgorm:"typ=INT"`
    FirstName string `mgorm:"typ=VARCHAR(14)"`
}

employee := Employee{EmpNo: 1000, FirstName: "Taro"}

err := mgorm.Insert(db, "employees", "emp_no", "first_name").
    Model(&employee).Exec()
```


## Type
`Model`の引数に渡せる値は以下の型の参照になります．

- int, int8, int16, in32, in64
- uint, uint8, uint16, uint32, uint64
- string
- time.Time
- map{}
- struct{} (エクスポートされたフィールドのみ適用)
- 上記の型によるスライス


## Tag
構造体のフィールドに`mgorm`タグを付与することでデータベースのスキーマを定義することができます．

複数のタグを使用したい場合は「,」で区切ることで指定できます．

#### 例
```go
type Employees struct {
    ID        int       `mgorm:"emp_no,typ=INT,notnull=t"`
    BirthDate time.Time `mgorm:"typ=DATE,notnull=t,layout=time.RFC3339"`
    FirstName string    `mgorm:"typ=VARCHAR(14),notnull=t"`
    LastName  string    `mgorm:"typ=VARCHAR(16),notnull=t"`
    Gender    string    `mgorm:"typ=CHAR(3),notnull=t"`
    HireDate  time.Time `mgorm:"typ=DATE,notnull=t,layout=2006-01-02"`
}
```


### column
データベースのカラム名を指定します．
`mgorm`タグではなく`json`をタグを用いた場合にも反映されます．
両方指定してある場合は`mgorm`タグが優先されます．

#### 例
```go
type Employee struct {
    // emp_no
    ID        int    `mgorm:"emp_no"`
    // first_name
    FirstName string `json:"first_name"`
    // last_name
    LastName  string `mgorm:"last_name" json:"lastname"`
}
```


### typ
データベース型を指定します．

#### 例
```go
type Employee struct {
    EmpNo      int     `mgorm:"typ=INT"`
    FirstName  string  `mgorm:"typ=VARCHAR(14)"`
}
```


### notnull
カラムに`NOT NULL`オプションを付与します．

#### 例
```go
type Employee struct {
    EmpNo int `mgorm:"notnull=t"`
}
```


### default
カラムに`DEFAULT`オプションを付与します．

#### 例
```go
type Employee struct {
    EmpNo      int     `mgorm:"default=60000"`
    FirstName  string  `mgorm:"default='Taro'"`
}
```


### pk
カラムに`PRIMARY KEY`を付与します．

#### 例
```go
type Employee struct {
    // CONSTRAINT PK_emp_no PRIAMRY KEY (emp_no)
    EmpNo int `mgorm:"pk=PK_emp_no"`
}

type DeptEmp struct {
    // CONSTRAINT PK_emp_no PRIMARY KEY (emp_no, dept_no)
    EmpNo  int `mgorm:"pk=PK_emp_no"`
    DeptNo int `mgorm:"pk=PK_emp_no"`
}
```


### fk
カラムに`FOREIGN KEY`を付与します．

#### 例
```go
type DeptEmp struct {
    // CONSTRAINT FK_emp_no FOREIGN KEY (emp_no) REFERENCES employees(emp_no)
    EmpNo int `mgorm:"fk=FK_emp_no:employees(emp_no)"`
}

type DeptEmp struct {
    // CONSTRAINT FK_emp_no FOREIGN KEY (emp_no) REFERENCES employees(emp_no, dept_no)
    EmpNo  int `mgorm:"fk=FK_emp_no:employees(emp_no)"`
    DeptNo int `mgorm:"fk=FK_emp_no:employees(emp_no)"`
}
```


### uc
カラムに`UNIQUE`を付与します．

#### 例
```go
type Employee struct {
    // CONSTRAINT PK_emp_no UNIQUE (emp_no)
    EmpNo int `mgorm:"uc=UC_emp_no"`
}

type DeptEmp struct {
    // CONSTRAINT UC_emp_no UNIQUE (emp_no, dept_no)
    EmpNo  int `mgorm:"uc=UC_emp_no"`
    DeptNo int `mgorm:"uc=UC_emp_no"`
}
```


### layout
カラムの型が`time.Time`の場合のみ，そのフォーマットを指定します．

#### 例
```go
type Employees struct {
    BirthDate   time.Time `mgorm:"layout=time.RFC3339"`
    HireDate    time.Time `mgorm:"layout=2006-01-02 15:05:06"`
}
```
