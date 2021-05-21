package migration_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/statement/migration"
	"github.com/stretchr/testify/assert"
)

func TestDropDB_String(t *testing.T) {
	testCases := []struct {
		Stmt     *migration.DropDBStmt
		Expected string
	}{
		{
			gsorm.DropDB(nil, "employees").(*migration.DropDBStmt),
			`DROP DATABASE employees`,
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

func TestDropDB_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *migration.DropDBStmt
		Expected string
	}{
		{
			gsorm.DropDB(nil, "database").RawClause("RAW").(*migration.DropDBStmt),
			`DROP DATABASE database RAW`,
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
