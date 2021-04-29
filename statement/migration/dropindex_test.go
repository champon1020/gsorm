package migration_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/statement/migration"
	"github.com/stretchr/testify/assert"
)

func TestDropIndex_String(t *testing.T) {
	testCases := []struct {
		Stmt     *migration.DropIndexStmt
		Expected string
	}{
		{
			mgorm.DropIndex(migration.ExportedMySQLDB, "IDX_id").
				On("person").(*migration.DropIndexStmt),
			`DROP INDEX IDX_id ON person`,
		},
		{
			mgorm.DropIndex(migration.ExportedPSQLDB, "IDX_id").(*migration.DropIndexStmt),
			`DROP INDEX IDX_id`,
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
