package mgorm_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/stretchr/testify/assert"
)

func TestCreateDB_String(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.CreateDBStmt
		Expected string
	}{
		{
			mgorm.CreateDB(mgorm.ExportedMySQLDB, "sample").(*mgorm.CreateDBStmt),
			`CREATE DATABASE sample`,
		},
		{
			mgorm.CreateDB(mgorm.ExportedPSQLDB, "sample").(*mgorm.CreateDBStmt),
			`CREATE DATABASE sample`,
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

func TestCreateIndex_String(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.CreateIndexStmt
		Expected string
	}{
		{
			mgorm.CreateIndex(mgorm.ExportedMySQLDB, "IDX_id").
				On("person", "id").(*mgorm.CreateIndexStmt),
			`CREATE INDEX IDX_id ON person (id)`,
		},
		{
			mgorm.CreateIndex(mgorm.ExportedMySQLDB, "IDX_id").
				On("person", "id", "name").(*mgorm.CreateIndexStmt),
			`CREATE INDEX IDX_id ON person (id, name)`,
		},
		{
			mgorm.CreateIndex(mgorm.ExportedPSQLDB, "IDX_id").
				On("person", "id").(*mgorm.CreateIndexStmt),
			`CREATE INDEX IDX_id ON person (id)`,
		},
		{
			mgorm.CreateIndex(mgorm.ExportedPSQLDB, "IDX_id").
				On("person", "id", "name").(*mgorm.CreateIndexStmt),
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

func TestCreateTable_String(t *testing.T) {
	type Model struct {
		ID          int       `mgorm:"notnull=t,pk=PK_id"`
		CountryCode string    `mgorm:"typ=CHAR(3),notnull=t,default='0',fk=FK_country_code:country(code)"`
		Name        string    `mgorm:"typ=VARCHAR(16),notnull=t,default='anonymous',uc=UC_name"`
		Nickname    string    `mgorm:"typ=VARCHAR(32),uc=UC_name"`
		BirthDate   time.Time `mgorm:"notnull=t"`
	}
	type Model2 struct {
		ID          int       `mgorm:"notnull=t,pk=PK_id"`
		CountryCode string    `mgorm:"typ=CHAR(3),notnull=t,default='0',fk=FK_country_code:country(code)"`
		Name        string    `mgorm:"typ=VARCHAR(16),notnull=t,default='anonymous',uc=UC_name"`
		Nickname    string    `mgorm:"typ=VARCHAR(32),uc=UC_name"`
		BirthDate   time.Time `mgorm:"notnull=t"`
	}

	model := new(Model)
	model2 := new(Model2)

	testCases := []struct {
		Stmt     *mgorm.CreateTableStmt
		Expected string
	}{
		{
			mgorm.CreateTable(mgorm.ExportedMySQLDB, "person").
				Column("id", "INT").NotNull().
				Column("country_code", "CHAR(3)").NotNull().Default("0").
				Column("name", "VARCHAR(16)").NotNull().Default("anonymous").
				Column("nickname", "VARCHAR(32)").
				Column("birth_date", "DATE").NotNull().
				Cons("UC_name").Unique("name", "nickname").
				Cons("PK_id").Primary("id").
				Cons("FK_country_code").Foreign("country_code").Ref("country", "code").(*mgorm.CreateTableStmt),
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
			mgorm.CreateTable(mgorm.ExportedPSQLDB, "person").
				Column("id", "INT").NotNull().
				Column("country_code", "CHAR(3)").NotNull().Default("0").
				Column("name", "VARCHAR(16)").NotNull().Default("anonymous").
				Column("nickname", "VARCHAR(32)").
				Column("birth_date", "DATE").NotNull().
				Cons("UC_name").Unique("name", "nickname").
				Cons("PK_id").Primary("id").
				Cons("FK_country_code").Foreign("country_code").Ref("country", "code").(*mgorm.CreateTableStmt),
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
			mgorm.CreateTable(mgorm.ExportedMySQLDB, "person").
				Model(model).(*mgorm.CreateTableStmt),
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
			mgorm.CreateTable(mgorm.ExportedPSQLDB, "person").
				Model(model2).(*mgorm.CreateTableStmt),
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

func TestDropDB_String(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DropDBStmt
		Expected string
	}{
		{
			mgorm.DropDB(mgorm.ExportedMySQLDB, "person").(*mgorm.DropDBStmt),
			`DROP DATABASE person`,
		},
		{
			mgorm.DropDB(mgorm.ExportedPSQLDB, "person").(*mgorm.DropDBStmt),
			`DROP DATABASE person`,
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

func TestDropIndex_String(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DropIndexStmt
		Expected string
	}{
		{
			mgorm.DropIndex(mgorm.ExportedMySQLDB, "IDX_id").
				On("person").(*mgorm.DropIndexStmt),
			`DROP INDEX IDX_id ON person`,
		},
		{
			mgorm.DropIndex(mgorm.ExportedPSQLDB, "IDX_id").(*mgorm.DropIndexStmt),
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

func TestDropTable_String(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.DropTableStmt
		Expected string
	}{
		{
			mgorm.DropTable(mgorm.ExportedMySQLDB, "person").(*mgorm.DropTableStmt),
			`DROP TABLE person`,
		},
		{
			mgorm.DropTable(mgorm.ExportedPSQLDB, "person").(*mgorm.DropTableStmt),
			`DROP TABLE person`,
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

func TestAlterTable_String(t *testing.T) {
	testCases := []struct {
		Stmt     *mgorm.AlterTableStmt
		Expected string
	}{
		// RENAME TO action.
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				Rename("human").(*mgorm.AlterTableStmt),
			`ALTER TABLE person RENAME TO human`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedPSQLDB, "person").
				Rename("human").(*mgorm.AlterTableStmt),
			`ALTER TABLE person RENAME TO human`,
		},
		// ADD COLUMN action.
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				AddColumn("id", "SERIAL").NotNull().(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD COLUMN id SERIAL NOT NULL`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				AddColumn("birth_date", "DATE").NotNull().
				Default(time.Date(2021, time.January, 2, 0, 0, 0, 0, time.UTC)).(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD COLUMN birth_date DATE NOT NULL DEFAULT '2021-01-02 00:00:00'`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedPSQLDB, "person").
				AddColumn("birth_date", "DATE").NotNull().
				Default(time.Date(2021, time.January, 2, 0, 0, 0, 0, time.UTC)).(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD COLUMN birth_date DATE NOT NULL DEFAULT '2021-01-02 00:00:00'`,
		},
		// DROP COLUMN action.
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				DropColumn("nickname").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`DROP COLUMN nickname`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedPSQLDB, "person").
				DropColumn("nickname").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`DROP COLUMN nickname`,
		},
		// RENAME COLUMN action.
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				RenameColumn("name", "first_name").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`RENAME COLUMN name TO first_name`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedPSQLDB, "person").
				RenameColumn("name", "first_name").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`RENAME COLUMN name TO first_name`,
		},
		// ADD CONSTRAINT action.
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				AddCons("UC_name").Unique("name", "nickname").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD CONSTRAINT UC_name UNIQUE (name, nickname)`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedPSQLDB, "person").
				AddCons("UC_name").Unique("name", "nickname").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD CONSTRAINT UC_name UNIQUE (name, nickname)`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				AddCons("PK_id").Primary("id").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD CONSTRAINT PK_id PRIMARY KEY (id)`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedPSQLDB, "person").
				AddCons("PK_id").Primary("id").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD CONSTRAINT PK_id PRIMARY KEY (id)`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				AddCons("FK_country_code").Foreign("country_code").Ref("country", "code").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD CONSTRAINT FK_country_code FOREIGN KEY (country_code) REFERENCES country(code)`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedPSQLDB, "person").
				AddCons("FK_country_code").Foreign("country_code").Ref("country", "code").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`ADD CONSTRAINT FK_country_code FOREIGN KEY (country_code) REFERENCES country(code)`,
		},
		// DROP CONSTRAINT action.
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				DropUnique("UC_name").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`DROP INDEX UC_name`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedPSQLDB, "person").
				DropUnique("UC_name").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`DROP CONSTRAINT UC_name`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				DropPrimary("PK_id").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`DROP PRIMARY KEY`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedPSQLDB, "person").
				DropPrimary("PK_id").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`DROP CONSTRAINT PK_id`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedMySQLDB, "person").
				DropForeign("FK_country_code").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`DROP FOREIGN KEY FK_country_code`,
		},
		{
			mgorm.AlterTable(mgorm.ExportedPSQLDB, "person").
				DropForeign("FK_country_code").(*mgorm.AlterTableStmt),
			`ALTER TABLE person ` +
				`DROP CONSTRAINT FK_country_code`,
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
