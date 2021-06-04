package parser_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/stretchr/testify/assert"
)

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
