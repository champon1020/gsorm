package mgorm_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/stretchr/testify/assert"
)

func TestMigStmt_String(t *testing.T) {
	testCases := []struct {
		MigStmt  *mgorm.MigStmt
		Expected string
	}{
		{
			mgorm.CreateDB(nil, "sample").(*mgorm.MigStmt),
			`CREATE DATABASE sample`,
		},
		{
			mgorm.CreateTable(nil, "sample").
				Column("id", "INT").NotNull().AutoInc().
				Column("name", "VARCHAR(64)").NotNull().Default("champon").
				Cons("PK_id").PK("id").
				Cons("FK_category_id").FK("category_id").Ref("category", "id").(*mgorm.MigStmt),
			`CREATE TABLE sample (` +
				`id INT NOT NULL AUTO_INCREMENT, ` +
				`name VARCHAR(64) NOT NULL DEFAULT "champon", ` +
				`CONSTRAINT PK_id PRIMARY KEY (id), ` +
				`CONSTRAINT FK_category_id FOREIGN KEY (category_id) REFERENCES category(id)` +
				`)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.MigStmt.String()
		errs := testCase.MigStmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %v", errs[0])
			return
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
