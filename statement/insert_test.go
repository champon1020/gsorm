package statement_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/statement"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/morikuni/failure"
	"github.com/stretchr/testify/assert"
)

func TestInsertStmt_BuildSQLWithClauses_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrInvalidClause,
			func() error {
				// Prepare for test.
				s := gsorm.Insert(nil, "", "").(*statement.InsertStmt)
				s.ExportedSetCalled(&clause.Set{})

				// Actual build.
				var sql internal.SQL
				err := statement.InsertStmtBuildSQL(s, &sql)
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

func TestInsertStmt_BuildSQLWithModel_Fail(t *testing.T) {
	testCases := []struct {
		ExpectedError failure.StringCode
		Build         func() error
	}{
		{
			statement.ErrFailedParse,
			func() error {
				// Prepare for test.
				s := gsorm.Insert(nil, "", "").Model(1000).(*statement.InsertStmt)

				// Actual build.
				var sql internal.SQL
				err := statement.InsertStmtBuildSQL(s, &sql)
				return err
			},
		},
		{
			statement.ErrFailedParse,
			func() error {
				// Prepare for test.
				model := make(map[string]interface{})
				s := gsorm.Insert(nil, "table", "column").Model(&model).(*statement.InsertStmt)

				// Actual build.
				var sql internal.SQL
				err := statement.InsertStmtBuildSQL(s, &sql)
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

func TestInsertStmt_CompareStmts(t *testing.T) {
	testCases := []struct {
		ExpectedStmt  *statement.InsertStmt
		ActualStmt    *statement.InsertStmt
		ExpectedError failure.StringCode
	}{
		{
			gsorm.Insert(nil, "table").Values(10).(*statement.InsertStmt),
			gsorm.Insert(nil, "table").Values(10).Values(100).(*statement.InsertStmt),
			statement.ErrInvalidValue,
		},
		{
			gsorm.Insert(nil, "table").Values(10).(*statement.InsertStmt),
			gsorm.Insert(nil, "table").Values(10, 100).(*statement.InsertStmt),
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

func TestInsertStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *statement.InsertStmt
		Expected string
	}{
		{
			gsorm.Insert(nil, "table").
				RawClause("RAW").
				Values("value").(*statement.InsertStmt),
			`INSERT INTO table RAW VALUES ('value')`,
		},
		{
			gsorm.Insert(nil, "table").
				Values("value1").
				RawClause("RAW").
				Values("value2").(*statement.InsertStmt),
			`INSERT INTO table VALUES ('value1') RAW, ('value2')`,
		},
		{
			gsorm.Insert(nil, "table").
				Values("value").
				RawClause("RAW").(*statement.InsertStmt),
			`INSERT INTO table VALUES ('value') RAW`,
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
		Stmt     *statement.InsertStmt
		Expected string
	}{
		{
			gsorm.Insert(nil, "employees").
				Values(1001, "1996-03-09", "Taro", "Sato", "M", "2020-04-01").(*statement.InsertStmt),
			`INSERT INTO employees VALUES (1001, '1996-03-09', 'Taro', 'Sato', 'M', '2020-04-01')`,
		},
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").
				Values(1001, "Taro").(*statement.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").
				Values(1001, "Taro").
				Values(1002, "Jiro").(*statement.InsertStmt),
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
		Stmt     *statement.InsertStmt
		Expected string
	}{
		{
			gsorm.Insert(nil, "dept_manager").Select(gsorm.Select(nil).From("dept_emp")).(*statement.InsertStmt),
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
		ID        int `gsorm:"emp_no"`
		FirstName string
	}
	structModel := Employee{ID: 1001, FirstName: "Taro"}
	structSlice := []Employee{{ID: 1001, FirstName: "Taro"}, {ID: 1002, FirstName: "Jiro"}}
	mapModel := map[string]interface{}{"emp_no": 1001, "first_name": "Taro"}
	mapSlice := []map[string]interface{}{
		{"emp_no": 1001, "first_name": "Taro"},
		{"emp_no": 1002, "first_name": "Jiro"},
	}

	testCases := []struct {
		Stmt     *statement.InsertStmt
		Expected string
	}{
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&structModel).(*statement.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&structSlice).(*statement.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro'), (1002, 'Jiro')`,
		},
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&mapModel).(*statement.InsertStmt),
			`INSERT INTO employees (emp_no, first_name) VALUES (1001, 'Taro')`,
		},
		{
			gsorm.Insert(nil, "employees", "emp_no", "first_name").Model(&mapSlice).(*statement.InsertStmt),
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
