package migration_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/statement/migration"
	"github.com/stretchr/testify/assert"
)

func TestAlterTable_String(t *testing.T) {
	testCases := []struct {
		Stmt     *migration.AlterTableStmt
		Expected string
	}{
		// RENAME TO action.
		{
			mgorm.AlterTable(nil, "person").
				Rename("human").(*migration.AlterTableStmt),
			`ALTER TABLE person RENAME TO human`,
		},
		// ADD COLUMN action.
		{
			mgorm.AlterTable(nil, "person").
				AddColumn("id", "INT").NotNull().(*migration.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD COLUMN id INT NOT NULL`,
		},
		{
			mgorm.AlterTable(nil, "person").
				AddColumn("birth_date", "DATE").NotNull().
				Default(time.Date(2021, time.January, 2, 0, 0, 0, 0, time.UTC)).(*migration.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD COLUMN birth_date DATE NOT NULL DEFAULT '2021-01-02 00:00:00'`,
		},
		// DROP COLUMN action.
		{
			mgorm.AlterTable(nil, "person").
				DropColumn("nickname").(*migration.AlterTableStmt),
			`ALTER TABLE person ` +
				`DROP COLUMN nickname`,
		},
		// RENAME COLUMN action.
		{
			mgorm.AlterTable(nil, "person").
				RenameColumn("name", "first_name").(*migration.AlterTableStmt),
			`ALTER TABLE person ` +
				`RENAME COLUMN name TO first_name`,
		},
		// ADD CONSTRAINT action.
		{
			mgorm.AlterTable(nil, "person").
				AddCons("UC_name").Unique("name", "nickname").(*migration.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD CONSTRAINT UC_name UNIQUE (name, nickname)`,
		},
		{
			mgorm.AlterTable(nil, "person").
				AddCons("PK_id").Primary("id").(*migration.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD CONSTRAINT PK_id PRIMARY KEY (id)`,
		},
		{
			mgorm.AlterTable(nil, "person").
				AddCons("FK_country_code").Foreign("country_code").Ref("country", "code").(*migration.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD CONSTRAINT FK_country_code FOREIGN KEY (country_code) REFERENCES country(code)`,
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
