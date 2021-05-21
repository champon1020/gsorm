package statement_test

import (
	"testing"
	"time"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/statement"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/morikuni/failure"
	"github.com/stretchr/testify/assert"
)

func TestUpdateStmt_BuildSQLWithClauses_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := gsorm.Update(nil, "table").(*statement.UpdateStmt)
				s.ExportedSetCalled(&clause.Join{})

				// Actual build.
				var sql internal.SQL
				err := statement.UpdateStmtBuildSQL(s, &sql)
				return err
			},
		},
	}

	for _, testCase := range testCases {
		err := testCase.Build()
		if !failure.Is(err, testCase.ExpectedError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", testCase.ExpectedError)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestUpdateStmt_BuildSQLWithModel_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := map[string]interface{}{
					"id":   1000,
					"name": "Taro",
				}
				s := gsorm.Update(nil, "sample").Model(&model, "id", "first_name").(*statement.UpdateStmt)

				// Actual build.
				var sql internal.SQL
				err := statement.UpdateStmtBuildSQL(s, &sql)
				return err
			},
		},
		{
			statement.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := []int{1000}
				s := gsorm.Update(nil, "sample").Model(&model, "id", "first_name").(*statement.UpdateStmt)

				// Actual build.
				var sql internal.SQL
				err := statement.UpdateStmtBuildSQL(s, &sql)
				return err
			},
		},
	}

	for _, testCase := range testCases {
		err := testCase.Build()
		if !failure.Is(err, testCase.ExpectedError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", testCase.ExpectedError)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestUpdateStmt_CompareStmts(t *testing.T) {
	testCases := []struct {
		ExpectedStmt  *statement.UpdateStmt
		ActualStmt    *statement.UpdateStmt
		ExpectedError failure.StringCode
	}{
		{
			gsorm.Update(nil, "table").Set("col1", 10).(*statement.UpdateStmt),
			gsorm.Update(nil, "table").Set("col1", 10).Set("col2", 100).(*statement.UpdateStmt),
			statement.ErrInvalidValue,
		},
		{
			gsorm.Update(nil, "table").Set("col1", 10).(*statement.UpdateStmt),
			gsorm.Update(nil, "table").Set("col1", 100).(*statement.UpdateStmt),
			statement.ErrInvalidValue,
		},
	}

	for _, testCase := range testCases {
		err := testCase.ExpectedStmt.CompareWith(testCase.ActualStmt)

		// Validate if the expected error was occurred.
		if !failure.Is(err, testCase.ExpectedError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %+v", testCase.ExpectedError)
			t.Errorf("  Actual:   %+v", err)
		}
	}
}

func TestUpdateStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "table").
				RawClause("RAW").
				Set("column", "value").(*statement.UpdateStmt),
			`UPDATE table RAW SET column = 'value'`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				RawClause("RAW").
				Set("column2", "value2").(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' RAW, column2 = 'value2'`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				RawClause("RAW").
				Where("column = ?", 10).(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`RAW WHERE column = 10`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				RawClause("RAW").
				And("column = ?", 100).(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`RAW AND (column = 100)`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				RawClause("RAW").
				Or("column = ?", 100).(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`RAW OR (column = 100)`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				And("column = ?", 100).
				RawClause("RAW").(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`AND (column = 100) RAW`,
		},
		{
			gsorm.Update(nil, "table").
				Set("column", "value").
				Where("column = ?", 10).
				Or("column = ?", 100).
				RawClause("RAW").(*statement.UpdateStmt),
			`UPDATE table SET column = 'value' ` +
				`WHERE column = 10 ` +
				`OR (column = 100) RAW`,
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

func TestUpdateStmt_Set(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako'`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Set("last_name", "Suzuki").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako', last_name = 'Suzuki'`,
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

func TestUpdateStmt_Where(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = 1001").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("first_name = ?", "Taro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE first_name = 'Taro'`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE birth_date = '2006-01-02 00:00:00'`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("first_name LIKE ?", "%Taro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE first_name LIKE '%Taro'`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no BETWEEN 1001 AND 1003`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no IN (?)", []int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no IN (?)", [2]int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no IN (1001, 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no IN (SELECT emp_no FROM dept_manager)`,
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

func TestUpdateStmt_And(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no = 1002").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no = 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no = 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("first_name = ? OR first_name = ?", "Taro", "Jiro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no = ?", 1002).
				And("emp_no = ?", 1003).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no = 1002) ` +
				`AND (emp_no = 1003)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("first_name LIKE ?", "%Taro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (first_name LIKE '%Taro')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", []int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", [2]int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				And("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`AND (emp_no IN (SELECT emp_no FROM dept_manager))`,
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

func TestUpdateStmt_Or(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no = 1002").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no = 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no = 1002)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("first_name = ? OR first_name = ?", "Taro", "Jiro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (first_name = 'Taro' OR first_name = 'Jiro')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no = ?", 1002).
				Or("emp_no = ?", 1003).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no = 1002) ` +
				`OR (emp_no = 1003)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("birth_date = ?", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (birth_date = '2006-01-02 00:00:00')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("first_name LIKE ?", "%Taro").(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (first_name LIKE '%Taro')`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no BETWEEN ? AND ?", 1001, 1003).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no BETWEEN 1001 AND 1003)`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", []int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", [2]int{1001, 1002}).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no IN (1001, 1002))`,
		},
		{
			gsorm.Update(nil, "employees").
				Set("first_name", "Hanako").
				Where("emp_no = ?", 1001).
				Or("emp_no IN (?)", gsorm.Select(nil, "emp_no").From("dept_manager")).(*statement.UpdateStmt),
			`UPDATE employees SET first_name = 'Hanako' ` +
				`WHERE emp_no = 1001 ` +
				`OR (emp_no IN (SELECT emp_no FROM dept_manager))`,
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

func TestUpdateStmt_Model(t *testing.T) {
	type Employee struct {
		ID        int `gsorm:"emp_no"`
		FirstName string
	}
	structModel := Employee{ID: 1001, FirstName: "Taro"}
	mapModel := map[string]interface{}{"emp_no": 1001, "first_name": "Taro"}

	testCases := []struct {
		Stmt     *statement.UpdateStmt
		Expected string
	}{
		{
			gsorm.Update(nil, "employees").Model(&structModel, "emp_no", "first_name").(*statement.UpdateStmt),
			`UPDATE employees SET emp_no = 1001, first_name = 'Taro'`,
		},
		{
			gsorm.Update(nil, "employees").Model(&mapModel, "emp_no", "first_name").(*statement.UpdateStmt),
			`UPDATE employees SET emp_no = 1001, first_name = 'Taro'`,
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
