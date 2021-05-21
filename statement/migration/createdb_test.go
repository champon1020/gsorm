package migration_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/statement/migration"
	"github.com/stretchr/testify/assert"
)

func TestCreateDB_String(t *testing.T) {
	testCases := []struct {
		Stmt     *migration.CreateDBStmt
		Expected string
	}{
		{
			gsorm.CreateDB(nil, "employees").(*migration.CreateDBStmt),
			`CREATE DATABASE employees`,
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

func TestCreateDB_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *migration.CreateDBStmt
		Expected string
	}{
		{
			gsorm.CreateDB(nil, "database").RawClause("RAW").(*migration.CreateDBStmt),
			`CREATE DATABASE database RAW`,
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
