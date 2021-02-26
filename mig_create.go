package mgorm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"

	provider "github.com/champon1020/mgorm/provider/create"
)

// CreateDBStmt is CREATE DATABASE statement.
type CreateDBStmt struct {
	migStmt
	cmd *mig.CreateDB
}

func (s *CreateDBStmt) String() string {
	return s.string(s.buildSQL)
}

// Migration executes database migration.
func (s *CreateDBStmt) Migration() error {
	return s.migration(s.buildSQL)
}

func (s *CreateDBStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())
	return nil
}

// CreateIndexStmt is CREATE INDEX statement.
type CreateIndexStmt struct {
	migStmt
	cmd *mig.CreateIndex
}

func (s *CreateIndexStmt) String() string {
	return s.string(s.buildSQL)
}

// Migration executes database migration.
func (s *CreateIndexStmt) Migration() error {
	return s.migration(s.buildSQL)
}

func (s *CreateIndexStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	for len(s.called) > 0 {
		e := s.headClause()
		if e == nil {
			break
		}

		switch e := e.(type) {
		case *mig.On:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
		default:
			msg := fmt.Sprintf("%v is not supported for CREATE INDEX statement", reflect.TypeOf(e).String())
			return errors.New(msg, errors.InvalidTypeError)
		}
	}

	return nil
}

// On calls ON clause.
func (s *CreateIndexStmt) On(table string, cols ...string) provider.OnMP {
	s.call(&mig.On{Table: table, Columns: cols})
	return s
}

// CreateTableStmt is CREATE TABLE statement.
type CreateTableStmt struct {
	migStmt
	cmd *mig.CreateTable
}

func (s *CreateTableStmt) String() string {
	return s.string(s.buildSQL)
}

// Migration executes database migration.
func (s *CreateTableStmt) Migration() error {
	return s.migration(s.buildSQL)
}

func (s *CreateTableStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	sql.Write("(")
	for len(s.called) > 0 {
		e := s.headClause()
		if e == nil {
			msg := "Called clauses have already been processed but SQL is not completed."
			return errors.New(msg, errors.InvalidSyntaxError)
		}

		switch e := e.(type) {
		case *mig.Column:
			if !strings.HasSuffix(sql.String(), "(") {
				sql.Write(",")
			}
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
			if err := s.buildColumnOptSQL(sql); err != nil {
				return err
			}
		case *mig.Cons:
			if !strings.HasSuffix(sql.String(), "(") {
				sql.Write(",")
			}
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
			if err := s.buildConstraintSQL(sql); err != nil {
				return err
			}
		default:
			msg := fmt.Sprintf("%v is not supported for CREATE TABLE", reflect.TypeOf(e).String())
			return errors.New(msg, errors.InvalidTypeError)
		}
	}
	sql.Write(")")

	return nil
}

// Column calls table column definition.
func (s *CreateTableStmt) Column(col, typ string) provider.ColumnMP {
	s.call(&mig.Column{Col: col, Type: typ})
	return s
}

// NotNull calls NOT NULL option.
func (s *CreateTableStmt) NotNull() provider.NotNullMP {
	s.call(&mig.NotNull{})
	return s
}

// AutoInc calls AUTO_INCREMENT option (only MySQL).
func (s *CreateTableStmt) AutoInc() provider.AutoIncMP {
	s.call(&mig.AutoInc{})
	return s
}

// Default calls DEFAULT option.
func (s *CreateTableStmt) Default(val interface{}) provider.DefaultMP {
	s.call(&mig.Default{Value: val})
	return s
}

// Cons calls CONSTRAINT option.
func (s *CreateTableStmt) Cons(key string) provider.ConsMP {
	s.call(&mig.Cons{Key: key})
	return s
}

// Unique calls UNIQUE keyword.
func (s *CreateTableStmt) Unique(cols ...string) provider.UniqueMP {
	s.call(&mig.Unique{Columns: cols})
	return s
}

// Primary calls PRIMARY KEY keyword.
func (s *CreateTableStmt) Primary(cols ...string) provider.PrimaryMP {
	s.call(&mig.Primary{Columns: cols})
	return s
}

// Foreign calls FOREIGN KEY keyword.
func (s *CreateTableStmt) Foreign(cols ...string) provider.ForeignMP {
	s.call(&mig.Foreign{Columns: cols})
	return s
}

// Ref calls REFERENCES keyword.
func (s *CreateTableStmt) Ref(table, col string) provider.RefMP {
	s.call(&mig.Ref{Table: table, Column: col})
	return s
}
