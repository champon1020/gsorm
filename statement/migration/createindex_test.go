package migration_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/statement/migration"
	"github.com/stretchr/testify/assert"
)

func TestCreateIndex_String(t *testing.T) {
	testCases := []struct {
		Stmt     *migration.CreateIndexStmt
		Expected string
	}{
		{
			mgorm.CreateIndex(nil, "IDX_emp").
				On("employees", "emp_no").(*migration.CreateIndexStmt),
			`CREATE INDEX IDX_emp ON employees (emp_no)`,
		},
		{
			mgorm.CreateIndex(nil, "IDX_emp").
				On("employees", "emp_no", "first_name").(*migration.CreateIndexStmt),
			`CREATE INDEX IDX_emp ON employees (emp_no, first_name)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
