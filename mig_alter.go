package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"

	provider "github.com/champon1020/mgorm/provider/alter"
)

// AlterTableStmt is ALTER TABLE statement.
type AlterTableStmt struct {
	migStmt
	cmd *mig.AlterTable
}

func (s *AlterTableStmt) String() string {
	return s.string(s.buildSQL)
}

// Migration executes database migration.
func (s *AlterTableStmt) Migration() error {
	return s.migration(s.buildSQL)
}

func (s *AlterTableStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	for len(s.called) > 0 {
		e := s.headClause()
		if e == nil {
			msg := "Called clauses have already been processed but SQL is not completed."
			return errors.New(msg, errors.InvalidSyntaxError)
		}

		switch e := e.(type) {
		case *mig.Rename,
			*mig.RenameColumn,
			*mig.DropColumn,
			*mig.DropUnique,
			*mig.DropPrimary,
			*mig.DropForeign,
			*mig.DropIndex:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
		case *mig.AddColumn:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
			if err := s.buildColumnOptSQL(sql); err != nil {
				return err
			}
		case *mig.AddCons:
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
			msg := fmt.Sprintf("%v is not supported for ALTER TABLE statement", reflect.TypeOf(e).String())
			return errors.New(msg, errors.InvalidTypeError)
		}
	}

	return nil
}

// Rename calls RENAME TO clause.
func (s *AlterTableStmt) Rename(table string) provider.RenameMP {
	s.call(&mig.Rename{Table: table})
	return s
}

// AddColumn calls ADD COLUMN clause.
func (s *AlterTableStmt) AddColumn(col, typ string) provider.AddColumnMP {
	s.call(&mig.AddColumn{Column: col, Type: typ})
	return s
}

// DropColumn calls DROP COLUMN clause.
func (s *AlterTableStmt) DropColumn(col string) provider.DropColumnMP {
	s.call(&mig.DropColumn{Column: col})
	return s
}

// RenameColumn calls RENAME COLUMN clause.
func (s *AlterTableStmt) RenameColumn(col, dest string) provider.RenameColumnMP {
	s.call(&mig.RenameColumn{Column: col, Dest: dest})
	return s
}

// NotNull calls NOT NULL option.
func (s *AlterTableStmt) NotNull() provider.NotNullMP {
	s.call(&mig.NotNull{})
	return s
}

// AutoInc calls AUTO_INCREMENT option. (only MySQL)
func (s *AlterTableStmt) AutoInc() provider.AutoIncMP {
	s.call(&mig.AutoInc{})
	return s
}

// Default calls DEFAULT option.
func (s *AlterTableStmt) Default(val interface{}) provider.DefaultMP {
	s.call(&mig.Default{Value: val})
	return s
}

// AddCons calls ADD CONSTRAINT clause.
func (s *AlterTableStmt) AddCons(key string) provider.AddConsMP {
	s.call(&mig.AddCons{Key: key})
	return s
}

// DropUnique calls DROP INDEX | DROP CONSTRAINT clause.
func (s *AlterTableStmt) DropUnique(key string) provider.DropUniqueMP {
	s.call(&mig.DropUnique{Driver: s.conn.getDriver(), Key: key})
	return s
}

// DropPrimary calls DROP PRIMARY KEY | DROP CONSTRAINT clause.
func (s *AlterTableStmt) DropPrimary(key string) provider.DropPrimaryMP {
	s.call(&mig.DropPrimary{Driver: s.conn.getDriver(), Key: key})
	return s
}

// DropForeign calls DROP FOREIGN KEY | DROP CONSTRAINT clause.
func (s *AlterTableStmt) DropForeign(key string) provider.DropForeignMP {
	s.call(&mig.DropForeign{Driver: s.conn.getDriver(), Key: key})
	return s
}

// Unique calls UNIQUE keyword.
func (s *AlterTableStmt) Unique(cols ...string) provider.UniqueMP {
	s.call(&mig.Unique{Columns: cols})
	return s
}

// Primary calls PRIMARY KEY keyword.
func (s *AlterTableStmt) Primary(cols ...string) provider.PrimaryMP {
	s.call(&mig.Primary{Columns: cols})
	return s
}

// Foreign calls FOREIGN KEY keyword.
func (s *AlterTableStmt) Foreign(cols ...string) provider.ForeignMP {
	s.call(&mig.Foreign{Columns: cols})
	return s
}

// Ref calls REFERENCES keyword.
func (s *AlterTableStmt) Ref(table, col string) provider.RefMP {
	s.call(&mig.Ref{Table: table, Column: col})
	return s
}
