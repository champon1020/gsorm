package migration_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/statement/migration"
	"github.com/stretchr/testify/assert"
)

func TestDropTable_String(t *testing.T) {
	testCases := []struct {
		Stmt     *migration.DropTableStmt
		Expected string
	}{
		{
			gsorm.DropTable(nil, "employees").(*migration.DropTableStmt),
			`DROP TABLE employees`,
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

func TestDropTable_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *migration.DropTableStmt
		Expected string
	}{
		{
			gsorm.DropTable(nil, "table").RawClause("RAW").(*migration.DropTableStmt),
			`DROP TABLE table RAW`,
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
