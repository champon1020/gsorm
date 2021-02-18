package mgorm_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/stretchr/testify/assert"
)

func TestMigStmt_String(t *testing.T) {
	testCases := []struct {
		MigStmt  *mgorm.MigStmt
		Expected string
	}{
		{
			mgorm.CreateDB(nil, "sampledb").(*mgorm.MigStmt),
			`CREATE DATABASE sampledb`,
		},
		{
			mgorm.DropDB(nil, "sampledb").(*mgorm.MigStmt),
			`DROP DATABASE sampledb`,
		},
		{
			mgorm.CreateTable(nil, "sample").
				Column("id", "INT").NotNull().AutoInc().
				Column("name", "VARCHAR(64)").NotNull().Default("champon").
				Cons("CHK_id").Check("id < ?", 10000).
				Cons("UC_id").Unique("id").
				Cons("PK_id").PK("id").
				Cons("FK_sample2_id").FK("sample2_id").Ref("sample2", "id").(*mgorm.MigStmt),
			`CREATE TABLE sample (` +
				`id INT NOT NULL AUTO_INCREMENT, ` +
				`name VARCHAR(64) NOT NULL DEFAULT "champon", ` +
				`CONSTRAINT CHK_id CHECK (id < 10000), ` +
				`CONSTRAINT UC_id UNIQUE (id), ` +
				`CONSTRAINT PK_id PRIMARY KEY (id), ` +
				`CONSTRAINT FK_sample2_id FOREIGN KEY (sample2_id) REFERENCES sample2(id)` +
				`)`,
		},
		{
			mgorm.DropTable(nil, "sample").(*mgorm.MigStmt),
			`DROP TABLE sample`,
		},
		{
			mgorm.AlterTable(nil, "sample").
				Rename("example").(*mgorm.MigStmt),
			`ALTER TABLE sample RENAME TO example`,
		},
		{
			mgorm.AlterTable(nil, "sample").
				Add("birth_date", "DATE").NotNull().
				Default(time.Date(2021, time.January, 2, 0, 0, 0, 0, time.UTC)).(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`ADD birth_date DATE NOT NULL DEFAULT 2021-01-02 00:00:00`,
		},
		{
			mgorm.AlterTable(nil, "sample").
				Change("id", "uid", "CHAR(8)").NotNull().(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`CHANGE id uid CHAR(8) NOT NULL`,
		},
		{
			mgorm.AlterTable(nil, "sample").
				Modify("id", "INT").AutoInc().(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`MODIFY id INT AUTO_INCREMENT`,
		},
		{
			mgorm.AlterTable(nil, "sample").
				AddCons("PK_id").PK("id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`ADD CONSTRAINT PK_id PRIMARY KEY (id)`,
		},
		{
			mgorm.AlterTable(nil, "sample").
				AddCons("FK_id").FK("category_id").Ref("category", "id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`ADD CONSTRAINT FK_id FOREIGN KEY (category_id) REFERENCES category(id)`,
		},
		{
			mgorm.AlterTable(nil, "sample").
				AddCons("CHK_id").Check("id < ?", 10000).(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`ADD CONSTRAINT CHK_id CHECK (id < 10000)`,
		},
		{
			mgorm.AlterTable(nil, "sample").
				AddCons("UC_id").Unique("id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`ADD CONSTRAINT UC_id UNIQUE (id)`,
		},
		{
			mgorm.AlterTable(nil, "sample").
				DropCons("PK_id").(*mgorm.MigStmt),
			`ALTER TABLE sample ` +
				`DROP CONSTRAINT PK_id`,
		},
		{
			mgorm.CreateIndex(nil, "idx_id").On("sample", "id").(*mgorm.MigStmt),
			`CREATE INDEX idx_id ON sample (id)`,
		},
		{
			mgorm.AlterTable(nil, "sample").DropIndex("idx_id").(*mgorm.MigStmt),
			`ALTER TABLE sample DROP INDEX idx_id`,
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
