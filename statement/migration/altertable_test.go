package migration_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/statement/migration"
	"github.com/stretchr/testify/assert"
)

func TestAlterTableStmt_Rename(t *testing.T) {
	testCases := []struct {
		Stmt     *migration.AlterTableStmt
		Expected string
	}{
		{
			mgorm.AlterTable(nil, "employees").Rename("people").(*migration.AlterTableStmt),
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
		Stmt     *migration.AlterTableStmt
		Expected string
	}{
		{
			mgorm.AlterTable(nil, "employees").AddColumn("nickname", "VARCHAR(64)").(*migration.AlterTableStmt),
			`ALTER TABLE employees ADD COLUMN nickname VARCHAR(64)`,
		},
		{
			mgorm.AlterTable(nil, "employees").
				AddColumn("nickname", "VARCHAR(64)").
				NotNull().(*migration.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD COLUMN nickname VARCHAR(64) NOT NULL`,
		},
		{
			mgorm.AlterTable(nil, "employees").
				AddColumn("nickname", "VARCHAR(64)").
				NotNull().
				Default("none").(*migration.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD COLUMN nickname VARCHAR(64) NOT NULL DEFAULT 'none'`,
		},
		{
			mgorm.AlterTable(nil, "employees").
				AddColumn("nickname", "VARCHAR(64)").
				Default("none").(*migration.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD COLUMN nickname VARCHAR(64) DEFAULT 'none'`,
		},
		{
			mgorm.AlterTable(nil, "employees").
				AddColumn("nickname", "VARCHAR(64)").
				Default("none").
				NotNull().(*migration.AlterTableStmt),
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
		Stmt     *migration.AlterTableStmt
		Expected string
	}{
		{
			mgorm.AlterTable(nil, "employees").DropColumn("nickname").(*migration.AlterTableStmt),
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
		Stmt     *migration.AlterTableStmt
		Expected string
	}{
		{
			mgorm.AlterTable(nil, "employees").RenameColumn("emp_no", "id").(*migration.AlterTableStmt),
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
		Stmt     *migration.AlterTableStmt
		Expected string
	}{
		{
			mgorm.AlterTable(nil, "employees").
				AddCons("UC_nickname").Unique("nickname").(*migration.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD CONSTRAINT UC_nickname UNIQUE (nickname)`,
		},
		{
			mgorm.AlterTable(nil, "employees").
				AddCons("UC_nickname").Unique("nickname", "first_name").(*migration.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD CONSTRAINT UC_nickname UNIQUE (nickname, first_name)`,
		},
		{
			mgorm.AlterTable(nil, "employees").
				AddCons("PK_emp_no").Primary("emp_no").(*migration.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD CONSTRAINT PK_emp_no PRIMARY KEY (emp_no)`,
		},
		{
			mgorm.AlterTable(nil, "employees").
				AddCons("PK_emp_no").Primary("emp_no", "first_name").(*migration.AlterTableStmt),
			`ALTER TABLE employees ` +
				`ADD CONSTRAINT PK_emp_no PRIMARY KEY (emp_no, first_name)`,
		},
		{
			mgorm.AlterTable(nil, "dept_emp").
				AddCons("FK_emp_no").Foreign("emp_no").Ref("employees", "emp_no").(*migration.AlterTableStmt),
			`ALTER TABLE dept_emp ` +
				`ADD CONSTRAINT FK_emp_no ` +
				`FOREIGN KEY (emp_no) REFERENCES employees (emp_no)`,
		},
		{
			mgorm.AlterTable(nil, "dept_emp").
				AddCons("FK_emp_no").
				Foreign("emp_no", "from_date").
				Ref("employees", "emp_no", "hire_date").(*migration.AlterTableStmt),
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
