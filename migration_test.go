package mgorm_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/stretchr/testify/assert"
)

func TestMigStmt_String(t *testing.T) {
	psql := new(mgorm.DB)
	psql.ExportedSetDriver(internal.PSQL)
	mysql := new(mgorm.DB)
	mysql.ExportedSetDriver(internal.MySQL)

	testCases := []struct {
		MigStmt  *mgorm.MigStmt
		Expected string
	}{
		{
			mgorm.CreateDB(mysql, "sampledb").(*mgorm.MigStmt),
			`CREATE DATABASE sampledb`,
		},
		{
			mgorm.DropDB(mysql, "sampledb").(*mgorm.MigStmt),
			`DROP DATABASE sampledb`,
		},
		{
			mgorm.CreateTable(mysql, "sample").
				Column("id", "INT").NotNull().AutoInc().
				Column("name", "VARCHAR(64)").NotNull().Default("champon").
				Cons("UC_id").UC("id").
				Cons("PK_id").PK("id").
				Cons("FK_sample2_id").FK("sample2_id").Ref("sample2", "id").(*mgorm.MigStmt),
			`CREATE TABLE sample (` +
				`id INT NOT NULL AUTO_INCREMENT, ` +
				`name VARCHAR(64) NOT NULL DEFAULT "champon", ` +
				`CONSTRAINT UC_id UNIQUE (id), ` +
				`CONSTRAINT PK_id PRIMARY KEY (id), ` +
				`CONSTRAINT FK_sample2_id FOREIGN KEY (sample2_id) REFERENCES sample2(id)` +
				`)`,
		},
		{
			mgorm.DropTable(mysql, "sample").(*mgorm.MigStmt),
			`DROP TABLE sample`,
		},
		{
			mgorm.AlterTable(mysql, "sample").
				Rename("example").(*mgorm.MigStmt),
			`ALTER TABLE sample RENAME TO example`,
		},
		{
			mgorm.AlterTable(mysql, "sample").
				AddColumn("birth_date", "DATE").NotNull().
				Default(time.Date(2021, time.January, 2, 0, 0, 0, 0, time.UTC)).(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`ADD COLUMN birth_date DATE NOT NULL DEFAULT 2021-01-02 00:00:00`,
		},
		{
			mgorm.AlterTable(mysql, "sample").
				RenameColumn("name", "first_name").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`RENAME COLUMN name TO first_name`,
		},
		{
			mgorm.AlterTable(mysql, "sample").
				AddCons("PK_id").PK("id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`ADD CONSTRAINT PK_id PRIMARY KEY (id)`,
		},
		{
			mgorm.AlterTable(mysql, "sample").
				AddCons("FK_id").FK("category_id").Ref("category", "id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`ADD CONSTRAINT FK_id FOREIGN KEY (category_id) REFERENCES category(id)`,
		},
		{
			mgorm.AlterTable(mysql, "sample").
				AddCons("UC_id").UC("id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`ADD CONSTRAINT UC_id UNIQUE (id)`,
		},
		{
			mgorm.AlterTable(mysql, "sample").
				DropPK("PK_id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`DROP PRIMARY KEY`,
		},
		{
			mgorm.AlterTable(psql, "sample").
				DropPK("PK_id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`DROP CONSTRAINT PK_id`,
		},
		{
			mgorm.AlterTable(mysql, "sample").
				DropFK("FK_id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`DROP FOREIGN KEY FK_id`,
		},
		{
			mgorm.AlterTable(psql, "sample").
				DropFK("FK_id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`DROP CONSTRAINT FK_id`,
		},
		{
			mgorm.AlterTable(mysql, "sample").
				DropUC("UC_id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`DROP INDEX UC_id`,
		},
		{
			mgorm.AlterTable(psql, "sample").
				DropUC("UC_id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`DROP CONSTRAINT UC_id`,
		},
		{
			mgorm.CreateIndex(mysql, "idx_id").On("sample", "id").(*mgorm.MigStmt),
			`CREATE INDEX idx_id ON sample (id)`,
		},
		{
			mgorm.DropIndex(mysql, "sample", "IDX_id").(*mgorm.MigStmt),
			`ALTER TABLE sample DROP INDEX IDX_id`,
		},
		{
			mgorm.DropIndex(psql, "sample", "IDX_id").(*mgorm.MigStmt),
			`DROP INDEX IDX_id`,
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
