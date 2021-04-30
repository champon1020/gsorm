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
			mgorm.CreateIndex(nil, "IDX_id").
				On("person", "id").(*migration.CreateIndexStmt),
			`CREATE INDEX IDX_id ON person (id)`,
		},
		{
			mgorm.CreateIndex(nil, "IDX_id").
				On("person", "id", "name").(*migration.CreateIndexStmt),
			`CREATE INDEX IDX_id ON person (id, name)`,
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
