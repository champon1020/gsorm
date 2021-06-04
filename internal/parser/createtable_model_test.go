package parser_test

import (
	"testing"
	"time"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/statement/migration"
	"github.com/stretchr/testify/assert"
)

func TestCreateTableModelParser(t *testing.T) {
	type Model struct {
		ID          int       `gsorm:"typ=INT,notnull=t,pk=PK_id"`
		CountryCode string    `gsorm:"typ=CHAR(3),notnull=t,default='0',fk=FK_country_code:country (code)"`
		Name        string    `gsorm:"typ=VARCHAR(16),notnull=t,default='anonymous',uc=UC_name"`
		Nickname    string    `gsorm:"typ=VARCHAR(32),uc=UC_name"`
		BirthDate   time.Time `gsorm:"typ=DATE,notnull=t"`
	}
	model := Model{}
	db := gsorm.OpenMock()

	testCases := []struct {
		Stmt     *migration.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(db, "person").
				Model(&model).(*migration.CreateTableStmt),
			`CREATE TABLE person (` +
				`id INT NOT NULL, ` +
				`country_code CHAR(3) NOT NULL DEFAULT '0', ` +
				`name VARCHAR(16) NOT NULL DEFAULT 'anonymous', ` +
				`nickname VARCHAR(32), ` +
				`birth_date DATE NOT NULL, ` +
				`CONSTRAINT UC_name UNIQUE (name, nickname), ` +
				`CONSTRAINT PK_id PRIMARY KEY (id), ` +
				`CONSTRAINT FK_country_code FOREIGN KEY (country_code) REFERENCES country (code)` +
				`)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Stmt.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
