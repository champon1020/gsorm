package gsorm_test

import (
	"testing"
	"time"

	"github.com/champon1020/gsorm"
	"github.com/stretchr/testify/assert"
)

func TestCreateTableModelParser(t *testing.T) {
	type Model struct {
		ID          int       `gsorm:"typ=INT,notnull=t,pk=PK_id"`
		CountryCode string    `gsorm:"typ=CHAR(3),notnull=t,default='0',fk=FK_country_code:country (code)"`
		Name        string    `gsorm:"typ=VARCHAR(16),notnull=t,default='anonymous',uc=UC_name"`
		Nickname    string    `gsorm:"typ=VARCHAR(32),uc=UC_name"`
		BirthDate   time.Time `gsorm:"typ=DATE,notnull=t"`
	}
	model := Model{}
	db := gsorm.OpenMock()

	testCases := []struct {
		Stmt     *gsorm.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(db, "person").
				Model(&model).(*gsorm.CreateTableStmt),
			`CREATE TABLE person (` +
				`id INT NOT NULL, ` +
				`country_code CHAR(3) NOT NULL DEFAULT '0', ` +
				`name VARCHAR(16) NOT NULL DEFAULT 'anonymous', ` +
				`nickname VARCHAR(32), ` +
				`birth_date DATE NOT NULL, ` +
				`CONSTRAINT UC_name UNIQUE (name, nickname), ` +
				`CONSTRAINT PK_id PRIMARY KEY (id), ` +
				`CONSTRAINT FK_country_code FOREIGN KEY (country_code) REFERENCES country (code)` +
				`)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestInsertModelParser_ParseStruct2(t *testing.T) {
	type Employee struct {
		ID        int `gsorm:"emp_no"`
		FirstName string
	}

	model := Employee{ID: 1001, FirstName: "Taro"}

	actual := gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&model).String()
	expected := `INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`

	assert.Equal(t, expected, actual)
}

func TestInsertModelParser_ParseStructSlice(t *testing.T) {
	type Employee struct {
		ID        int `gsorm:"emp_no"`
		FirstName string
	}

	model := []Employee{{ID: 1001, FirstName: "Taro"}, {ID: 1002, FirstName: "Jiro"}}

	actual := gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&model).String()
	expected := `INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`

	assert.Equal(t, expected, actual)
}

func TestInsertModelParser_ParseMap(t *testing.T) {
	model := map[string]interface{}{"emp_no": 1001, "first_name": "Taro"}

	actual := gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&model).String()
	expected := `INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`

	assert.Equal(t, expected, actual)
}

func TestInsertModelParser_ParseMapSlice(t *testing.T) {
	model := []map[string]interface{}{
		{"emp_no": 1001, "first_name": "Taro"},
		{"emp_no": 1002, "first_name": "Jiro"},
	}

	actual := gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&model).String()
	expected := `INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`

	assert.Equal(t, expected, actual)
}

func TestUpdateModelParser_ParseStruct(t *testing.T) {
	type Employee struct {
		ID        int `gsorm:"emp_no"`
		FirstName string
	}

	model := Employee{ID: 1001, FirstName: "Taro"}

	actual := gsorm.Update(nil, "employees").Model(&model, "emp_no", "first_name").String()
	expected := `UPDATE employees SET emp_no = 1001, first_name = 'Taro'`

	assert.Equal(t, expected, actual)
}

func TestUpdateModelParser_ParseMap(t *testing.T) {
	model := map[string]interface{}{"emp_no": 1001, "first_name": "Taro"}

	actual := gsorm.Update(nil, "employees").Model(&model, "emp_no", "first_name").String()
	expected := `UPDATE employees SET emp_no = 1001, first_name = 'Taro'`

	assert.Equal(t, expected, actual)
}
