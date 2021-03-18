# Model

## Tag

### typ
データベース型を指定します．

#### 例
```go
type Employee struct {
    emp_no      int     `mgorm:"typ=INT"`
    first_name  string  `mgorm:"typ=VARCHAR(64)"`
}
```


### notnull
カラムに`NOT NULL`オプションを付与します．

#### 例
```go
type Employee struct {
    emp_no int `mgorm:"notnull=t"`
}
```


### default
カラムに`DEFAULT`オプションを付与します．

#### 例
```go
type Employee struct {
    emp_no      int     `mgorm:"default=60000"`
    first_name  string  `mgorm:"default='Taro'"`
}
```


### pk
カラムに`PRIMARY KEY`を付与します．

#### 例
```go
type Employee struct {
    // CONSTRAINT PK_emp_no PRIAMRY KEY (emp_no)
    emp_no int `mgorm:"pk=PK_emp_no"`
}

type DeptEmp struct {
    // CONSTRAINT PK_emp_no PRIMARY KEY (emp_no, dept_no)
    emp_no  int `mgorm:"pk=PK_emp_no"`
    dept_no int `mgorm:"pk=PK_emp_no"`
}
```


### fk
カラムに`FOREIGN KEY`を付与します．

#### 例
```go
type DeptEmp struct {
    // CONSTRAINT FK_emp_no FOREIGN KEY (emp_no) REFERENCES employees(emp_no)
    emp_no int `mgorm:"fk=FK_emp_no:employees(emp_no)"`
}

type DeptEmp struct {
    // CONSTRAINT FK_emp_no FOREIGN KEY (emp_no) REFERENCES employees(emp_no, dept_no)
    emp_no  int `mgorm:"fk=FK_emp_no:employees(emp_no)"`
    dept_no int `mgorm:"fk=FK_emp_no:employees(emp_no)"`
}
```


### uc
カラムに`UNIQUE`を付与します．

#### 例
```go
type Employee struct {
    // CONSTRAINT PK_emp_no UNIQUE (emp_no)
    emp_no int `mgorm:"uc=UC_emp_no"`
}

type DeptEmp struct {
    // CONSTRAINT UC_emp_no UNIQUE (emp_no, dept_no)
    emp_no  int `mgorm:"uc=UC_emp_no"`
    dept_no int `mgorm:"uc=UC_emp_no"`
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
