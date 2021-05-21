package migration_test

import (
	"testing"
	"time"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/database"
	"github.com/champon1020/gsorm/statement/migration"
	"github.com/stretchr/testify/assert"
)

func TestCreateTable_String(t *testing.T) {
	type Model struct {
		ID          int       `gsorm:"notnull=t,pk=PK_id"`
		CountryCode string    `gsorm:"typ=CHAR(3),notnull=t,default='0',fk=FK_country_code:country (code)"`
		Name        string    `gsorm:"typ=VARCHAR(16),notnull=t,default='anonymous',uc=UC_name"`
		Nickname    string    `gsorm:"typ=VARCHAR(32),uc=UC_name"`
		BirthDate   time.Time `gsorm:"notnull=t"`
	}
	model := new(Model)
	db := database.NewMockDB("mysql")

	testCases := []struct {
		Stmt     *migration.CreateTableStmt
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
				Cons("FK_country_code").Foreign("country_code").Ref("country", "code").(*migration.CreateTableStmt),
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
				Model(model).(*migration.CreateTableStmt),
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
		Stmt     *migration.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "table").
				RawClause("RAW").
				Column("column", "type").(*migration.CreateTableStmt),
			`CREATE TABLE table (RAW, column type)`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				RawClause("RAW").NotNull().(*migration.CreateTableStmt),
			`CREATE TABLE table (column type RAW NOT NULL)`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().RawClause("RAW").Default("value").(*migration.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL RAW DEFAULT 'value')`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				RawClause("RAW").
				Cons("key").Unique("column").(*migration.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value' RAW, ` +
				`CONSTRAINT key UNIQUE (column))`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").RawClause("RAW").Unique("column").(*migration.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key RAW UNIQUE (column))`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").Unique("column").RawClause("RAW").(*migration.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key UNIQUE (column) RAW)`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").RawClause("RAW").Primary("column").(*migration.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key RAW PRIMARY KEY (column))`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").Primary("column").RawClause("RAW").(*migration.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key PRIMARY KEY (column) RAW)`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").RawClause("RAW").
				Foreign("column").Ref("table2", "column2").(*migration.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key RAW FOREIGN KEY (column) REFERENCES table2 (column2))`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").Foreign("column").
				RawClause("RAW").Ref("table2", "column2").(*migration.CreateTableStmt),
			`CREATE TABLE table (column type NOT NULL DEFAULT 'value', ` +
				`CONSTRAINT key FOREIGN KEY (column) RAW REFERENCES table2 (column2))`,
		},
		{
			gsorm.CreateTable(nil, "table").
				Column("column", "type").
				NotNull().Default("value").
				Cons("key").Foreign("column").Ref("table2", "column2").
				RawClause("RAW").(*migration.CreateTableStmt),
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
		Stmt     *migration.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").(*migration.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT)`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").
				Column("birth_date", "DATE").(*migration.CreateTableStmt),
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
		Stmt     *migration.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().(*migration.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT NOT NULL)`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Column("birth_date", "DATE").(*migration.CreateTableStmt),
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
		Stmt     *migration.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").Default(1).(*migration.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT DEFAULT 1)`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().Default(1).(*migration.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT NOT NULL DEFAULT 1)`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").Default(1).NotNull().(*migration.CreateTableStmt),
			`CREATE TABLE employees (emp_no INT DEFAULT 1 NOT NULL)`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().Default(1).
				Column("birth_date", "DATE").(*migration.CreateTableStmt),
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
		Stmt     *migration.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Cons("UC_emp_no").Unique("emp_no").(*migration.CreateTableStmt),
			`CREATE TABLE employees (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT UC_emp_no UNIQUE (emp_no))`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Cons("PK_emp_no").Primary("emp_no").(*migration.CreateTableStmt),
			`CREATE TABLE employees (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT PK_emp_no PRIMARY KEY (emp_no))`,
		},
		{
			gsorm.CreateTable(nil, "dept_emp").
				Column("emp_no", "INT").NotNull().
				Cons("FK_dept_emp").Foreign("emp_no").
				Ref("employees", "emp_no").(*migration.CreateTableStmt),
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
		Stmt     *migration.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Cons("UC_emp_no").Unique("emp_no").(*migration.CreateTableStmt),
			`CREATE TABLE employees (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT UC_emp_no UNIQUE (emp_no))`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Column("first_name", "VARCHAR(14)").NotNull().
				Cons("UC_emp_no_first_name").Unique("emp_no", "first_name").(*migration.CreateTableStmt),
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
		Stmt     *migration.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Cons("PK_emp_no").Primary("emp_no").(*migration.CreateTableStmt),
			`CREATE TABLE employees (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT PK_emp_no PRIMARY KEY (emp_no))`,
		},
		{
			gsorm.CreateTable(nil, "employees").
				Column("emp_no", "INT").NotNull().
				Column("first_name", "VARCHAR(14)").NotNull().
				Cons("PK_emp_no_first_name").Primary("emp_no", "first_name").(*migration.CreateTableStmt),
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
		Stmt     *migration.CreateTableStmt
		Expected string
	}{
		{
			gsorm.CreateTable(nil, "dept_emp").
				Column("emp_no", "INT").NotNull().
				Cons("FK_dept_emp").Foreign("emp_no").
				Ref("employees", "emp_no").(*migration.CreateTableStmt),
			`CREATE TABLE dept_emp (` +
				`emp_no INT NOT NULL, ` +
				`CONSTRAINT FK_dept_emp FOREIGN KEY (emp_no) REFERENCES employees (emp_no))`,
		},
		{
			gsorm.CreateTable(nil, "dept_emp").
				Column("emp_no", "INT").NotNull().
				Column("first_name", "VARCHAR(14)").NotNull().
				Cons("FK_dept_emp").Foreign("emp_no", "first_name").
				Ref("employees", "emp_no", "first_name").(*migration.CreateTableStmt),
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
