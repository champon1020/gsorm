package gsorm_test

import (
	"testing"
	"time"

	"github.com/champon1020/gsorm"
	"github.com/stretchr/testify/assert"
)

func TestAlterTableStmt_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.AlterTableStmt
		Expected string
	}{
		{
			gsorm.AlterTable(nil, "table").
				RawClause("RAW").
				Rename("table").(*gsorm.AlterTableStmt),
			`ALTER TABLE table RAW RENAME TO table`,
		},
		{
			gsorm.AlterTable(nil, "table").
				Rename("table").
				RawClause("RAW").(*gsorm.AlterTableStmt),
			`ALTER TABLE table RENAME TO table RAW`,
		},
		{
			gsorm.AlterTable(nil, "table").
				RawClause("RAW").
				AddColumn("column", "type").(*gsorm.AlterTableStmt),
			`ALTER TABLE table RAW ADD COLUMN column type`,
		},
		{
			gsorm.AlterTable(nil, "table").
				AddColumn("column", "type").
				RawClause("RAW").NotNull().(*gsorm.AlterTableStmt),
			`ALTER TABLE table ADD COLUMN column type RAW NOT NULL`,
		},
		{
			gsorm.AlterTable(nil, "table").
				AddColumn("column", "type").
				NotNull().RawClause("RAW").Default("value").(*gsorm.AlterTableStmt),
			`ALTER TABLE table ADD COLUMN column type NOT NULL RAW DEFAULT 'value'`,
		},
		{
			gsorm.AlterTable(nil, "table").
				AddColumn("column", "type").
				NotNull().Default("value").RawClause("RAW").(*gsorm.AlterTableStmt),
			`ALTER TABLE table ADD COLUMN column type NOT NULL DEFAULT 'value' RAW`,
		},
		{
			gsorm.AlterTable(nil, "table").
				RawClause("RAW").
				DropColumn("column").(*gsorm.AlterTableStmt),
			`ALTER TABLE table RAW DROP COLUMN column`,
		},
		{
			gsorm.AlterTable(nil, "table").
				DropColumn("column").
				RawClause("RAW").(*gsorm.AlterTableStmt),
			`ALTER TABLE table DROP COLUMN column RAW`,
		},
		{
			gsorm.AlterTable(nil, "table").
				RawClause("RAW").
				RenameColumn("column", "dest").(*gsorm.AlterTableStmt),
			`ALTER TABLE table RAW RENAME COLUMN column TO dest`,
		},
		{
			gsorm.AlterTable(nil, "table").
				RenameColumn("column", "dest").
				RawClause("RAW").(*gsorm.AlterTableStmt),
			`ALTER TABLE table RENAME COLUMN column TO dest RAW`,
		},
		{
			gsorm.AlterTable(nil, "table").
				RawClause("RAW").
				AddCons("key").Unique("column").(*gsorm.AlterTableStmt),
			`ALTER TABLE table RAW ADD CONSTRAINT key UNIQUE (column)`,
		},
		{
			gsorm.AlterTable(nil, "table").
				AddCons("key").RawClause("RAW").Unique("column").(*gsorm.AlterTableStmt),
			`ALTER TABLE table ADD CONSTRAINT key RAW UNIQUE (column)`,
		},
		{
			gsorm.AlterTable(nil, "table").
				AddCons("key").Unique("column").RawClause("RAW").(*gsorm.AlterTableStmt),
			`ALTER TABLE table ADD CONSTRAINT key UNIQUE (column) RAW`,
		},
		{
			gsorm.AlterTable(nil, "table").
				AddCons("key").RawClause("RAW").Primary("column").(*gsorm.AlterTableStmt),
			`ALTER TABLE table ADD CONSTRAINT key RAW PRIMARY KEY (column)`,
		},
		{
			gsorm.AlterTable(nil, "table").
				AddCons("key").Primary("column").RawClause("RAW").(*gsorm.AlterTableStmt),
			`ALTER TABLE table ADD CONSTRAINT key PRIMARY KEY (column) RAW`,
		},
		{
			gsorm.AlterTable(nil, "table").
				AddCons("key").
				RawClause("RAW").
				Foreign("column").Ref("table2", "column2").(*gsorm.AlterTableStmt),
			`ALTER TABLE table ADD CONSTRAINT key RAW FOREIGN KEY (column) REFERENCES table2 (column2)`,
		},
		{
			gsorm.AlterTable(nil, "table").
				AddCons("key").Foreign("column").
				RawClause("RAW").
				Ref("table2", "column2").(*gsorm.AlterTableStmt),
			`ALTER TABLE table ADD CONSTRAINT key FOREIGN KEY (column) RAW REFERENCES table2 (column2)`,
		},
		{
			gsorm.AlterTable(nil, "table").
				AddCons("key").Foreign("column").
				Ref("table2", "column2").
				RawClause("RAW").(*gsorm.AlterTableStmt),
			`ALTER TABLE table ADD CONSTRAINT key FOREIGN KEY (column) REFERENCES table2 (column2) RAW`,
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

func TestAlterTableStmt_Rename(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.AlterTableStmt
		Expected string
	}{
		{
			gsorm.AlterTable(nil, "employees").Rename("people").(*gsorm.AlterTableStmt),
			`ALTER TABLE employees RENAME TO people`,
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

func TestAlterTableStmt_AddColumn(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.AlterTableStmt
		Expected string
	}{
		{
			gsorm.AlterTable(nil, "employees").AddColumn("nickname", "VARCHAR(64)").(*gsorm.AlterTableStmt),
			`ALTER TABLE employees ADD COLUMN nickname VARCHAR(64)`,
		},
		{
			gsorm.AlterTable(nil, "employees").
				AddColumn("nickname", "VARCHAR(64)").
				NotNull().(*gsorm.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD COLUMN nickname VARCHAR(64) NOT NULL`,
		},
		{
			gsorm.AlterTable(nil, "employees").
				AddColumn("nickname", "VARCHAR(64)").
				NotNull().
				Default("none").(*gsorm.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD COLUMN nickname VARCHAR(64) NOT NULL DEFAULT 'none'`,
		},
		{
			gsorm.AlterTable(nil, "employees").
				AddColumn("nickname", "VARCHAR(64)").
				Default("none").(*gsorm.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD COLUMN nickname VARCHAR(64) DEFAULT 'none'`,
		},
		{
			gsorm.AlterTable(nil, "employees").
				AddColumn("nickname", "VARCHAR(64)").
				Default("none").
				NotNull().(*gsorm.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD COLUMN nickname VARCHAR(64) DEFAULT 'none' NOT NULL`,
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

func TestAlterTableStmt_DropColumn(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.AlterTableStmt
		Expected string
	}{
		{
			gsorm.AlterTable(nil, "employees").DropColumn("nickname").(*gsorm.AlterTableStmt),
			`ALTER TABLE employees DROP COLUMN nickname`,
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

func TestAlterTableStmt_RenameColumn(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.AlterTableStmt
		Expected string
	}{
		{
			gsorm.AlterTable(nil, "employees").RenameColumn("emp_no", "id").(*gsorm.AlterTableStmt),
			`ALTER TABLE employees RENAME COLUMN emp_no TO id`,
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

func TestAlterTableStmt_AddCons(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.AlterTableStmt
		Expected string
	}{
		{
			gsorm.AlterTable(nil, "employees").
				AddCons("UC_nickname").Unique("nickname").(*gsorm.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD CONSTRAINT UC_nickname UNIQUE (nickname)`,
		},
		{
			gsorm.AlterTable(nil, "employees").
				AddCons("UC_nickname").Unique("nickname", "first_name").(*gsorm.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD CONSTRAINT UC_nickname UNIQUE (nickname, first_name)`,
		},
		{
			gsorm.AlterTable(nil, "employees").
				AddCons("PK_emp_no").Primary("emp_no").(*gsorm.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD CONSTRAINT PK_emp_no PRIMARY KEY (emp_no)`,
		},
		{
			gsorm.AlterTable(nil, "employees").
				AddCons("PK_emp_no").Primary("emp_no", "first_name").(*gsorm.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD CONSTRAINT PK_emp_no PRIMARY KEY (emp_no, first_name)`,
		},
		{
			gsorm.AlterTable(nil, "dept_emp").
				AddCons("FK_emp_no").Foreign("emp_no").Ref("employees", "emp_no").(*gsorm.AlterTableStmt),
			`ALTER TABLE dept_emp ` +
				`ADD CONSTRAINT FK_emp_no ` +
				`FOREIGN KEY (emp_no) REFERENCES employees (emp_no)`,
		},
		{
			gsorm.AlterTable(nil, "dept_emp").
				AddCons("FK_emp_no").
				Foreign("emp_no", "from_date").
				Ref("employees", "emp_no", "hire_date").(*gsorm.AlterTableStmt),
			`ALTER TABLE dept_emp ` +
				`ADD CONSTRAINT FK_emp_no ` +
				`FOREIGN KEY (emp_no, from_date) REFERENCES employees (emp_no, hire_date)`,
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

func TestCreateDB_String(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateDBStmt
		Expected string
	}{
		{
			gsorm.CreateDB(nil, "employees").(*gsorm.CreateDBStmt),
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
		Stmt     *gsorm.CreateDBStmt
		Expected string
	}{
		{
			gsorm.CreateDB(nil, "database").RawClause("RAW").(*gsorm.CreateDBStmt),
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

func TestCreateIndex_String(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateIndexStmt
		Expected string
	}{
		{
			gsorm.CreateIndex(nil, "IDX_emp").
				On("employees", "emp_no").(*gsorm.CreateIndexStmt),
			`CREATE INDEX IDX_emp ON employees (emp_no)`,
		},
		{
			gsorm.CreateIndex(nil, "IDX_emp").
				On("employees", "emp_no", "first_name").(*gsorm.CreateIndexStmt),
			`CREATE INDEX IDX_emp ON employees (emp_no, first_name)`,
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

func TestCreateIndex_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateIndexStmt
		Expected string
	}{
		{
			gsorm.CreateIndex(nil, "idx").RawClause("RAW").(*gsorm.CreateIndexStmt),
			`CREATE INDEX idx RAW`,
		},
		{
			gsorm.CreateIndex(nil, "idx").RawClause("RAW").On("table", "column").(*gsorm.CreateIndexStmt),
			`CREATE INDEX idx RAW ON table (column)`,
		},
		{
			gsorm.CreateIndex(nil, "idx").On("table", "column").RawClause("RAW").(*gsorm.CreateIndexStmt),
			`CREATE INDEX idx ON table (column) RAW`,
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

func TestCreateTable_String(t *testing.T) {
	type Model struct {
		ID          int       `gsorm:"typ=INT,notnull=t,pk=PK_id"`
		CountryCode string    `gsorm:"typ=CHAR(3),notnull=t,default='0',fk=FK_country_code:country (code)"`
		Name        string    `gsorm:"typ=VARCHAR(16),notnull=t,default='anonymous',uc=UC_name"`
		Nickname    string    `gsorm:"typ=VARCHAR(32),uc=UC_name"`
		BirthDate   time.Time `gsorm:"typ=DATE,notnull=t"`
	}
	model := new(Model)
	db := gsorm.OpenMock()

	testCases := []struct {
		Stmt     *gsorm.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(db, "person").
				Column("id", "INT").NotNull().
				Column("country_code", "CHAR(3)").NotNull().Default("0").
				Column("name", "VARCHAR(16)").NotNull().Default("anonymous").
				Column("nickname", "VARCHAR(32)").
				Column("birth_date", "DATE").NotNull().
				Cons("UC_name").Unique("name", "nickname").
				Cons("PK_id").Primary("id").
				Cons("FK_country_code").Foreign("country_code").Ref("country", "code").(*gsorm.CreateTableStmt),
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
		{
			gsorm.CreateTable(db, "person").
				Model(model).(*gsorm.CreateTableStmt),
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
		errs := testCase.Stmt.ExportedGetErrors()
		if len(errs) > 0 {
			t.Errorf("Error was occurred: %v", errs[0])
			continue
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestCreateTable_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "table").
				RawClause("RAW").
				Column("column", "type").(*gsorm.CreateTableStmt),
			`CREATE TABLE table (RAW, column type)`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				RawClause("RAW").NotNull().(*gsorm.CreateTableStmt),
			`CREATE TABLE table (column type RAW NOT NULL)`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().RawClause("RAW").Default("value").(*gsorm.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL RAW DEFAULT 'value')`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				RawClause("RAW").
				Cons("key").Unique("column").(*gsorm.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value' RAW, ` +
				`CONSTRAINT key UNIQUE (column))`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").RawClause("RAW").Unique("column").(*gsorm.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key RAW UNIQUE (column))`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").Unique("column").RawClause("RAW").(*gsorm.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key UNIQUE (column) RAW)`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").RawClause("RAW").Primary("column").(*gsorm.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key RAW PRIMARY KEY (column))`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").Primary("column").RawClause("RAW").(*gsorm.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key PRIMARY KEY (column) RAW)`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").RawClause("RAW").
				Foreign("column").Ref("table2", "column2").(*gsorm.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key RAW FOREIGN KEY (column) REFERENCES table2 (column2))`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").Foreign("column").
				RawClause("RAW").Ref("table2", "column2").(*gsorm.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key FOREIGN KEY (column) RAW REFERENCES table2 (column2))`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").Foreign("column").Ref("table2", "column2").
				RawClause("RAW").(*gsorm.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key FOREIGN KEY (column) REFERENCES table2 (column2) RAW)`,
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

func TestCreateTable_Column(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT)`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").
				Column("birth_date", "DATE").(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT, birth_date DATE)`,
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

func TestCreateTable_NotNull(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT NOT NULL)`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Column("birth_date", "DATE").(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT NOT NULL, birth_date DATE)`,
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

func TestCreateTable_Default(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").Default(1).(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT DEFAULT 1)`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().Default(1).(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT NOT NULL DEFAULT 1)`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").Default(1).NotNull().(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT DEFAULT 1 NOT NULL)`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().Default(1).
				Column("birth_date", "DATE").(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT NOT NULL DEFAULT 1, birth_date DATE)`,
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

func TestCreateTable_Cons(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Cons("UC_emp_no").Unique("emp_no").(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT UC_emp_no UNIQUE (emp_no))`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Cons("PK_emp_no").Primary("emp_no").(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT PK_emp_no PRIMARY KEY (emp_no))`,
		},
		{
			gsorm.CreateTable(nil, "dept_emp").
				Column("emp_no", "INT").NotNull().
				Cons("FK_dept_emp").Foreign("emp_no").
				Ref("employees", "emp_no").(*gsorm.CreateTableStmt),
			`CREATE TABLE dept_emp (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no) REFERENCES employees (emp_no))`,
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

func TestCreateTable_Unique(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Cons("UC_emp_no").Unique("emp_no").(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT UC_emp_no UNIQUE (emp_no))`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Column("first_name", "VARCHAR(14)").NotNull().
				Cons("UC_emp_no_first_name").Unique("emp_no", "first_name").(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (` +
				`emp_no INT NOT NULL, ` +
				`first_name VARCHAR(14) NOT NULL, ` +
				`CONSTRAINT UC_emp_no_first_name UNIQUE (emp_no, first_name))`,
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

func TestCreateTable_Primary(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Cons("PK_emp_no").Primary("emp_no").(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT PK_emp_no PRIMARY KEY (emp_no))`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Column("first_name", "VARCHAR(14)").NotNull().
				Cons("PK_emp_no_first_name").Primary("emp_no", "first_name").(*gsorm.CreateTableStmt),
			`CREATE TABLE employees (` +
				`emp_no INT NOT NULL, ` +
				`first_name VARCHAR(14) NOT NULL, ` +
				`CONSTRAINT PK_emp_no_first_name PRIMARY KEY (emp_no, first_name))`,
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

func TestCreateTable_Foreign(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "dept_emp").
				Column("emp_no", "INT").NotNull().
				Cons("FK_dept_emp").Foreign("emp_no").
				Ref("employees", "emp_no").(*gsorm.CreateTableStmt),
			`CREATE TABLE dept_emp (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no) REFERENCES employees (emp_no))`,
		},
		{
			gsorm.CreateTable(nil, "dept_emp").
				Column("emp_no", "INT").NotNull().
				Column("first_name", "VARCHAR(14)").NotNull().
				Cons("FK_dept_emp").Foreign("emp_no", "first_name").
				Ref("employees", "emp_no", "first_name").(*gsorm.CreateTableStmt),
			`CREATE TABLE dept_emp (` +
				`emp_no INT NOT NULL, ` +
				`first_name VARCHAR(14) NOT NULL, ` +
				`CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no, first_name) REFERENCES employees (emp_no, first_name))`,
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
		Stmt     *gsorm.DropDBStmt
		Expected string
	}{
		{
			gsorm.DropDB(nil, "employees").(*gsorm.DropDBStmt),
			`DROP DATABASE employees`,
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

func TestDropDB_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.DropDBStmt
		Expected string
	}{
		{
			gsorm.DropDB(nil, "database").RawClause("RAW").(*gsorm.DropDBStmt),
			`DROP DATABASE database RAW`,
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

func TestDropTable_String(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.DropTableStmt
		Expected string
	}{
		{
			gsorm.DropTable(nil, "employees").(*gsorm.DropTableStmt),
			`DROP TABLE employees`,
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

func TestDropTable_RawClause(t *testing.T) {
	testCases := []struct {
		Stmt     *gsorm.DropTableStmt
		Expected string
	}{
		{
			gsorm.DropTable(nil, "table").RawClause("RAW").(*gsorm.DropTableStmt),
			`DROP TABLE table RAW`,
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
