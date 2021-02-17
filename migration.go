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
	cmd    syntax.MigCmd
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

// nextCalled returns pops head of called.
func (m *MigStmt) nextCalled() syntax.MigClause {
	if len(m.called) == 0 {
		return nil
	}
	defer m.popCalled()
	return m.called[0]
}

// popCalled removes head of called.
func (m *MigStmt) popCalled() {
	m.called = m.called[1:]
}

// String returns query with string.
func (m *MigStmt) String() string {
	sql, err := m.processMigrationSQL()
	if err != nil {
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

	return nil
}

func (m *MigStmt) processMigrationSQL() (internal.SQL, error) {
	var sql internal.SQL

	switch cmd := m.cmd.(type) {
	case *mig.CreateDB:
		sql.Write(cmd.Build().Build())
	case *mig.CreateTable:
		sql.Write(cmd.Build().Build())
		sql.Write("(")
		for len(m.called) > 0 {
			m.processCreateTableSQL(&sql)
		}
		sql.Write(")")
	default:
		msg := fmt.Sprintf("Type %v is not supported for migration", reflect.TypeOf(cmd).String())
		return "", errors.New(msg, errors.InvalidTypeError)
	}

	return sql, nil
}

func (m *MigStmt) processCreateTableSQL(sql *internal.SQL) error {
	e := m.nextCalled()
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
	case *mig.NotNull,
		*mig.AutoInc,
		*mig.Default:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
	case *mig.Constraint:
		if sql.Len() > 0 {
			sql.Write(",")
		}

		// Write CONSTRAINT.
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())

		// Write PK/FK.
		m.processKeySQL(sql)
	default:
		msg := fmt.Sprintf("Type %v is not supported for migration", reflect.TypeOf(e).String())
		return errors.New(msg, errors.InvalidTypeError)
	}
	return nil
}

func (m *MigStmt) processKeySQL(sql *internal.SQL) error {
	e := m.nextCalled()
	if e == nil {
		msg := "Called claues have already been processed but SQL is not completed."
		return errors.New(msg, errors.InvalidSyntaxError)
	}

	switch e := e.(type) {
	case *mig.PK:
		// Write PRIMARY KEY.
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
	case *mig.FK:
		// Write FOREIGN KEY.
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())

		// Write REFERENCES.
		m.processRefSQL(sql)
	default:
		msg := fmt.Sprintf("Type %v is not supported for migration", reflect.TypeOf(e).String())
		return errors.New(msg, errors.InvalidTypeError)
	}
	return nil
}

func (m *MigStmt) processRefSQL(sql *internal.SQL) error {
	e := m.nextCalled()
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
	default:
		msg := fmt.Sprintf("Type %v is not supported for migration", reflect.TypeOf(e).String())
		return errors.New(msg, errors.InvalidTypeError)
	}
	return nil
}

// Column calls table column definition.
func (m *MigStmt) Column(col string, typ string) ColumnMig {
	m.call(&mig.Column{Col: col, Type: typ})
	return m
}

// Rename calls RENAME TO clause.
func (m *MigStmt) Rename(table string) RenameMig {
	m.call(&mig.Rename{Table: table})
	return m
}

// Add calls ADD clause.
func (m *MigStmt) Add(col string, typ string) AddMig {
	m.call(&mig.Add{Column: col, Type: typ})
	return m
}

// Chnage calls CHANGE clause.
func (m *MigStmt) Change(col string, dest string, typ string) ChangeMig {
	m.call(&mig.Change{Column: col, Dest: dest, Type: typ})
	return m
}

// Modify calls MODIFY clause.
func (m *MigStmt) Modify(col string, typ string) ModifyMig {
	m.call(&mig.Modify{Column: col, Type: typ})
	return m
}

// Drop calls DROP clause.
func (m *MigStmt) Drop(col string) DropMig {
	m.call(&mig.Drop{Column: col})
	return m
}

// Charset calls CHARSET clause.
func (m *MigStmt) Charset(format string) CharsetMig {
	m.call(&mig.Charset{Format: format})
	return m
}

// NotNull calls NOT NULL option.
func (m *MigStmt) NotNull() NotNullMig {
	m.call(&mig.NotNull{})
	return m
}

// AutoInc calls AUTO_INCREMENT option.
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
	m.call(&mig.Constraint{Key: key})
	return m
}

// PK calls PRIMARY KEY keyword.
func (m *MigStmt) PK(col string) PKMig {
	p := new(mig.PK)
	p.AddColumns(col)
	m.call(p)
	return m
}

// FK calls FOREIGN KEY keyword.
func (m *MigStmt) FK(col string) FKMig {
	m.call(&mig.FK{Column: col})
	return m
}

// Ref calls REFERENCES keyword.
func (m *MigStmt) Ref(table string, col string) RefMig {
	m.call(&mig.Ref{Table: table, Column: col})
	return m
}
