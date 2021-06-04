package parser_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/stretchr/testify/assert"
)

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
