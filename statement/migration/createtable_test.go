package migration_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/database"
	"github.com/champon1020/mgorm/statement/migration"
	"github.com/stretchr/testify/assert"
)

func TestCreateTable_String(t *testing.T) {
	type Model struct {
		ID          int       `mgorm:"notnull=t,pk=PK_id"`
		CountryCode string    `mgorm:"typ=CHAR(3),notnull=t,default='0',fk=FK_country_code:country(code)"`
		Name        string    `mgorm:"typ=VARCHAR(16),notnull=t,default='anonymous',uc=UC_name"`
		Nickname    string    `mgorm:"typ=VARCHAR(32),uc=UC_name"`
		BirthDate   time.Time `mgorm:"notnull=t"`
	}
	model := new(Model)
	db := database.NewMockDB("mysql")

	testCases := []struct {
		Stmt     *migration.CreateTableStmt
		Expected string
	}{
		{
			mgorm.CreateTable(db, "person").
				Column("id", "INT").NotNull().
				Column("country_code", "CHAR(3)").NotNull().Default("0").
				Column("name", "VARCHAR(16)").NotNull().Default("anonymous").
				Column("nickname", "VARCHAR(32)").
				Column("birth_date", "DATE").NotNull().
				Cons("UC_name").Unique("name", "nickname").
				Cons("PK_id").Primary("id").
				Cons("FK_country_code").Foreign("country_code").Ref("country", "code").(*migration.CreateTableStmt),
			`CREATE TABLE person (` +
				`id INT NOT NULL, ` +
				`country_code CHAR(3) NOT NULL DEFAULT '0', ` +
				`name VARCHAR(16) NOT NULL DEFAULT 'anonymous', ` +
				`nickname VARCHAR(32), ` +
				`birth_date DATE NOT NULL, ` +
				`CONSTRAINT UC_name UNIQUE (name, nickname), ` +
				`CONSTRAINT PK_id PRIMARY KEY (id), ` +
				`CONSTRAINT FK_country_code FOREIGN KEY (country_code) REFERENCES country(code)` +
				`)`,
		},
		{
			mgorm.CreateTable(db, "person").
				Model(model).(*migration.CreateTableStmt),
			`CREATE TABLE person (` +
				`id INT NOT NULL, ` +
				`country_code CHAR(3) NOT NULL DEFAULT '0', ` +
				`name VARCHAR(16) NOT NULL DEFAULT 'anonymous', ` +
				`nickname VARCHAR(32), ` +
				`birth_date DATE NOT NULL, ` +
				`CONSTRAINT UC_name UNIQUE (name, nickname), ` +
				`CONSTRAINT PK_id PRIMARY KEY (id), ` +
				`CONSTRAINT FK_country_code FOREIGN KEY (country_code) REFERENCES country(code)` +
				`)`,
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
