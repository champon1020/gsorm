package mgorm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
)

// MigStmt stores information about database migration query.
type MigStmt struct {
	pool   Pool
	driver internal.SQLDriver
	cmd    syntax.MigClause
	called []syntax.MigClause
	errors []error
}

// call appends called clause.
func (m *MigStmt) call(e syntax.MigClause) {
	m.called = append(m.called, e)
}

// throw appends occurred error.
func (m *MigStmt) throw(e error) {
	m.errors = append(m.errors, e)
}

// headClause returns first element of called.
func (m *MigStmt) headClause() syntax.MigClause {
	if len(m.called) == 0 {
		return nil
	}
	return m.called[0]
}

// advanceClause slides slice of called.
func (m *MigStmt) advanceClause() {
	m.called = m.called[1:]
}

// String returns query with string.
func (m *MigStmt) String() string {
	sql, err := m.processMigrationSQL()
	if err != nil {
		m.throw(err)
		return ""
	}
	return sql.String()
}

// Migration executes database migration.
func (m *MigStmt) Migration() error {
	if len(m.errors) > 0 {
		return m.errors[0]
	}

	_, err := m.processMigrationSQL()
	if err != nil {
		return err
	}

	/* process */

	return nil
}

func (m *MigStmt) processMigrationSQL() (internal.SQL, error) {
	var sql internal.SQL

	switch cmd := m.cmd.(type) {
	case *mig.CreateDB,
		*mig.DropDB,
		*mig.DropTable,
		*mig.DropIndex:
		s, err := cmd.Build()
		if err != nil {
			return "", err
		}
		sql.Write(s.Build())
		return sql, nil
	case *mig.CreateTable:
		s, err := cmd.Build()
		if err != nil {
			return "", err
		}
		sql.Write(s.Build())
		sql.Write("(")
		for len(m.called) > 0 {
			err := m.processCreateTableSQL(&sql)
			if err != nil {
				return "", err
			}
		}
		sql.Write(")")
		return sql, nil
	case *mig.AlterTable:
		s, err := cmd.Build()
		if err != nil {
			return "", err
		}
		sql.Write(s.Build())
		for len(m.called) > 0 {
			err := m.processAlterTableSQL(&sql)
			if err != nil {
				return "", err
			}
		}
		return sql, nil
	case *mig.CreateIndex:
		s, err := cmd.Build()
		if err != nil {
			return "", err
		}
		sql.Write(s.Build())
		for len(m.called) > 0 {
			err := m.processCreateIndexSQL(&sql)
			if err != nil {
				return "", err
			}
		}
		return sql, nil
	}

	msg := fmt.Sprintf("Type %v is not supported for migration", reflect.TypeOf(m.cmd).String())
	return "", errors.New(msg, errors.InvalidTypeError)
}

func (m *MigStmt) processCreateTableSQL(sql *internal.SQL) error {
	e := m.headClause()
	if e == nil {
		msg := "Called claues have already been processed but SQL is not completed."
		return errors.New(msg, errors.InvalidSyntaxError)
	}

	switch e := e.(type) {
	case *mig.Column:
		if !strings.HasSuffix(sql.String(), "(") {
			sql.Write(",")
		}
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return nil
	case *mig.NotNull,
		*mig.AutoInc,
		*mig.Default:
		return m.processColumnOptSQL(sql)
	case *mig.Cons:
		if !strings.HasSuffix(sql.String(), "(") {
			sql.Write(",")
		}
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return m.processConstraintSQL(sql)
	}

	msg := fmt.Sprintf("Type %v is not supported for CREATE TABLE", reflect.TypeOf(e).String())
	return errors.New(msg, errors.InvalidTypeError)
}

func (m *MigStmt) processAlterTableSQL(sql *internal.SQL) error {
	e := m.headClause()
	if e == nil {
		msg := "Called claues have already been processed but SQL is not completed."
		return errors.New(msg, errors.InvalidSyntaxError)
	}

	switch e := e.(type) {
	case *mig.Rename,
		*mig.RenameColumn,
		*mig.DropColumn,
		*mig.DropPK,
		*mig.DropFK,
		*mig.DropUC,
		*mig.DropIndex:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return nil
	case *mig.AddCons:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return m.processConstraintSQL(sql)
	case *mig.AddColumn:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return m.processColumnOptSQL(sql)
	case *mig.NotNull,
		*mig.AutoInc,
		*mig.Default,
		*mig.Cons:
		return m.processColumnOptSQL(sql)
	}

	msg := fmt.Sprintf("Type %v is not supported for ALTER TABLE", reflect.TypeOf(e).String())
	return errors.New(msg, errors.InvalidTypeError)
}

func (m *MigStmt) processCreateIndexSQL(sql *internal.SQL) error {
	e := m.headClause()
	if e == nil {
		msg := "Called claues have already been processed but SQL is not completed."
		return errors.New(msg, errors.InvalidSyntaxError)
	}

	switch e := e.(type) {
	case *mig.On:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return nil
	}

	msg := fmt.Sprintf("Type %v is not supported for CREATE INDEX", reflect.TypeOf(e).String())
	return errors.New(msg, errors.InvalidTypeError)
}

func (m *MigStmt) processColumnOptSQL(sql *internal.SQL) error {
	e := m.headClause()
	if e == nil {
		msg := "Called claues have already been processed but SQL is not completed."
		return errors.New(msg, errors.InvalidSyntaxError)
	}

	switch e := e.(type) {
	case *mig.NotNull,
		*mig.AutoInc,
		*mig.Default:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return nil
	case *mig.Cons:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return m.processConstraintSQL(sql)
	}

	msg := fmt.Sprintf("Type %v is not supported for column option", reflect.TypeOf(e).String())
	return errors.New(msg, errors.InvalidTypeError)
}

func (m *MigStmt) processConstraintSQL(sql *internal.SQL) error {
	e := m.headClause()
	if e == nil {
		msg := "Called claues have already been processed but SQL is not completed."
		return errors.New(msg, errors.InvalidSyntaxError)
	}

	switch e := e.(type) {
	case *mig.PK, *mig.UC:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return nil
	case *mig.FK:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return m.processRefSQL(sql)
	}

	msg := fmt.Sprintf("Type %v is not supported for CONSTRAINT", reflect.TypeOf(e).String())
	return errors.New(msg, errors.InvalidTypeError)
}

func (m *MigStmt) processRefSQL(sql *internal.SQL) error {
	e := m.headClause()
	if e == nil {
		msg := "Called claues have already been processed but SQL is not completed."
		return errors.New(msg, errors.InvalidSyntaxError)
	}

	switch e := e.(type) {
	case *mig.Ref:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		m.advanceClause()
		return nil
	}

	msg := fmt.Sprintf("Type %v is not supported for CONSTRAINT KEY", reflect.TypeOf(e).String())
	return errors.New(msg, errors.InvalidTypeError)
}

// On calls ON clause.
func (m *MigStmt) On(table string, cols ...string) OnMig {
	m.call(&mig.On{Table: table, Columns: cols})
	return m
}

// Rename calls RENAME TO clause.
func (m *MigStmt) Rename(table string) RenameMig {
	m.call(&mig.Rename{Table: table})
	return m
}

// AddColumn calls ADD COLUMN clause.
func (m *MigStmt) AddColumn(col, typ string) AddColumnMig {
	m.call(&mig.AddColumn{Column: col, Type: typ})
	return m
}

// DropColumn calls DROP COLUMN clause.
func (m *MigStmt) DropColumn(col string) DropColumnMig {
	m.call(&mig.DropColumn{Column: col})
	return m
}

// RenameColumn calls RENAME COLUMN clause.
func (m *MigStmt) RenameColumn(col, dest string) RenameColumnMig {
	m.call(&mig.RenameColumn{Column: col, Dest: dest})
	return m
}

// AddCons calls ADD CONSTRAINT clause.
func (m *MigStmt) AddCons(key string) AddConsMig {
	m.call(&mig.AddCons{Key: key})
	return m
}

// DropPK calls DROP PRIMARY KEY | DROP CONSTRAINT clause.
func (m *MigStmt) DropPK(key string) DropPKMig {
	m.call(&mig.DropPK{Driver: m.driver, Key: key})
	return m
}

// DropFK calls DROP FOREIGN KEY | DROP CONSTRAINT clause.
func (m *MigStmt) DropFK(key string) DropFKMig {
	m.call(&mig.DropFK{Driver: m.driver, Key: key})
	return m
}

// DropUC calls DROP INDEX | DROP CONSTRAINT clause.
func (m *MigStmt) DropUC(key string) DropUCMig {
	m.call(&mig.DropUC{Driver: m.driver, Key: key})
	return m
}

// Column calls table column definition.
func (m *MigStmt) Column(col, typ string) ColumnMig {
	m.call(&mig.Column{Col: col, Type: typ})
	return m
}

// NotNull calls NOT NULL option.
func (m *MigStmt) NotNull() NotNullMig {
	m.call(&mig.NotNull{})
	return m
}

// AutoInc calls AUTO_INCREMENT option. (only MySQL)
func (m *MigStmt) AutoInc() AutoIncMig {
	m.call(&mig.AutoInc{})
	return m
}

// Default calls DEFAULT option.
func (m *MigStmt) Default(val interface{}) DefaultMig {
	m.call(&mig.Default{Value: val})
	return m
}

// Cons calls CONSTRAINT option.
func (m *MigStmt) Cons(key string) ConsMig {
	m.call(&mig.Cons{Key: key})
	return m
}

// UC calls UNIQUE keyword.
func (m *MigStmt) UC(cols ...string) UCMig {
	m.call(&mig.UC{Columns: cols})
	return m
}

// PK calls PRIMARY KEY keyword.
func (m *MigStmt) PK(cols ...string) PKMig {
	m.call(&mig.PK{Columns: cols})
	return m
}

// FK calls FOREIGN KEY keyword.
func (m *MigStmt) FK(cols ...string) FKMig {
	m.call(&mig.FK{Columns: cols})
	return m
}

// Ref calls REFERENCES keyword.
func (m *MigStmt) Ref(table, col string) RefMig {
	m.call(&mig.Ref{Table: table, Column: col})
	return m
}
