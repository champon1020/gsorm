# Model
`Model` maps the struct contents to the SQL.

This method is available on `Insert`, `Update` and `CreateTable`.

#### Example
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
`Model` can handle the pointer of these types.

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

Type of struct field or map element should be below.

- int, int8, int16, in32, in64
- uint, uint8, uint16, uint32, uint64
- float32, float64
- bool
- string
- time.Time

In case of using struct, only exported fields are used.


## Tag
Tagging the struct field with `gsorm` defines the properties of database column corresponding to the struct field.

Multiple tags should be separated by comma.

#### Example
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
Defines the database column name.

`json` tag is also used as the database column name.

If both `gsorm` and `json` are tagged, `gsorm` will be applied.

#### Example
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
Defines the type of database column.

`typ` is required for `gsorm.CreateTable.Model`.

#### Example
```go
type Employee struct {
    EmpNo      int     `gsorm:"typ=INT"`
    FirstName  string  `gsorm:"typ=VARCHAR(14)"`
}
```


### notnull
Adds `NOT NULL` property.

#### Example
```go
type Employee struct {
    EmpNo int `gsorm:"notnull=t"`
}
```


### default
Adds `DEFAULT` property.

#### Example
```go
type Employee struct {
    EmpNo      int     `gsorm:"default=60000"`
    FirstName  string  `gsorm:"default='Taro'"`
}
```


### uc
Sets `UNIQUE`.

#### Example
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


### pk
Sets `PRIMARY KEY`.

#### Example
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
Sets `FOREIGN KEY`.

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
