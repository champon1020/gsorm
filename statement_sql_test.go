package mgorm_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/stretchr/testify/assert"
)

func TestDeleteStmt_From(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").(*mgorm.DeleteStmt),
			`DELETE FROM employees`,
		},
		{
			mgorm.Delete(nil).From("employees").(*mgorm.DeleteStmt),
			`DELETE FROM employees`,
		},
		{
			mgorm.Delete(nil).From("employees AS e").(*mgorm.DeleteStmt),
			`DELETE FROM employees AS e`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestDeleteStmt_Where(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = 1001").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("first_name = ?", "Taro").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE first_name = 'Taro'`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("first_name LIKE ?", "%Taro").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE first_name LIKE '%Taro'`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no BETWEEN ? AND ?", 1001, 1003).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no BETWEEN 1001 AND 1003`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", []int{1001, 1002}).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", [2]int{1001, 1002}).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (1001, 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestDeleteStmt_And(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = 1002").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name = ? OR first_name = ?", "Taro", "Jiro").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).
				And("emp_no = ?", 1003).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no = 1002) AND (emp_no = 1003)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name = ?", "Taro").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name = 'Taro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("first_name LIKE ?", "%Taro").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (first_name LIKE '%Taro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no BETWEEN ? AND ?", 1001, 1003).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", []int{1001, 1002}).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", [2]int{1001, 1002}).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 AND (emp_no IN (SELECT emp_no FROM dept_manager))`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestDeleteStmt_Or(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DeleteStmt
		Expected string
	}{
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = 1002").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ? AND first_name = ?", 1002, "Taro").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002 AND first_name = 'Taro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).
				Or("emp_no = ?", 1003).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no = 1002) OR (emp_no = 1003)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("first_name = ?", "Taro").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (first_name = 'Taro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("first_name LIKE ?", "%Taro").(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (first_name LIKE '%Taro')`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no BETWEEN ? AND ?", 1001, 1003).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", []int{1001, 1002}).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", [2]int{1001, 1002}).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (1001, 1002))`,
		},
		{
			mgorm.Delete(nil).From("employees").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", mgorm.Select(nil, "emp_no").From("dept_manager")).(*mgorm.DeleteStmt),
			`DELETE FROM employees WHERE emp_no = 1001 OR (emp_no IN (SELECT emp_no FROM dept_manager))`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestInsertStmt_Values(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").
				Values(1001, "Taro").(*mgorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").
				Values(1001, "Taro").
				Values(1002, "Jiro").(*mgorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestInsertStmt_Select(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "dept_manager").Select(mgorm.Select(nil).From("dept_emp")).(*mgorm.InsertStmt),
			`INSERT INTO dept_manager SELECT * FROM dept_emp`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestInsertStmt_Model(t *testing.T) {
	type Employee struct {
		ID        int `mgorm:"emp_no"`
		FirstName string
	}
	structModel := Employee{ID: 1001, FirstName: "Taro"}
	structSlice := []Employee{{ID: 1001, FirstName: "Taro"}, {ID: 1002, FirstName: "Jiro"}}
	mapModel := map[string]interface{}{"emp_no": 1001, "first_name": "Taro"}
	mapSlice := []map[string]interface{}{
		{"emp_no": 1001, "first_name": "Taro"},
		{"emp_no": 1002, "first_name": "Jiro"},
	}
	varSlice := []string{"Taro", "Jiro"}

	testCases := []struct {
		Stmt     *mgorm.InsertStmt
		Expected string
	}{
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").Model(&structModel).(*mgorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").Model(&structSlice).(*mgorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`,
		},
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").Model(&mapModel).(*mgorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			mgorm.Insert(nil, "employees", "emp_no", "first_name").Model(&mapSlice).(*mgorm.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`,
		},
		{
			mgorm.Insert(nil, "employees", "first_name").Model(&varSlice).(*mgorm.InsertStmt),
			`INSERT INTO employees (first_name) VALUES ('Taro'), ('Jiro')`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_From(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_Join(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
func TestSelectStmt_LeftJoin(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
func TestSelectStmt_RightJoin(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
func TestSelectStmt_Where(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_And(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_Or(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_GroupBy(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_Having(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_Union(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_UnionAll(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_OrderBy(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestSelectStmt_Limit(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.SelectStmt
		Expected string
	}{}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %+v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
