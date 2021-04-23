package mgorm

import (
	"reflect"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/morikuni/failure"

	ifc "github.com/champon1020/mgorm/interfaces/altertable"
)

// AlterTableStmt is ALTER TABLE statement.
type AlterTableStmt struct {
	migStmt
	cmd *mig.AlterTable
}

func (s *AlterTableStmt) String() string {
	return s.string(s.buildSQL)
}

// Migrate executes database migration.
func (s *AlterTableStmt) Migrate() error {
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
			return failure.New(errInvalidSyntax,
				failure.Message("the SQL statement is not completed or the syntax is not supported"))
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
			return failure.New(errInvalidClause,
				failure.Context{"clause": reflect.TypeOf(e).String()},
				failure.Message("invalid clause for ALTER TABLE"))
		}
	}

	return nil
}

// Rename calls RENAME TO clause.
func (s *AlterTableStmt) Rename(table string) ifc.Rename {
	s.call(&mig.Rename{Table: table})
	return s
}

// AddColumn calls ADD COLUMN clause.
func (s *AlterTableStmt) AddColumn(col, typ string) ifc.AddColumn {
	s.call(&mig.AddColumn{Column: col, Type: typ})
	return s
}

// DropColumn calls DROP COLUMN clause.
func (s *AlterTableStmt) DropColumn(col string) ifc.DropColumn {
	s.call(&mig.DropColumn{Column: col})
	return s
}

// RenameColumn calls RENAME COLUMN clause.
func (s *AlterTableStmt) RenameColumn(col, dest string) ifc.RenameColumn {
	s.call(&mig.RenameColumn{Column: col, Dest: dest})
	return s
}

// NotNull calls NOT NULL option.
func (s *AlterTableStmt) NotNull() ifc.NotNull {
	s.call(&mig.NotNull{})
	return s
}

// Default calls DEFAULT option.
func (s *AlterTableStmt) Default(val interface{}) ifc.Default {
	s.call(&mig.Default{Value: val})
	return s
}

// AddCons calls ADD CONSTRAINT clause.
func (s *AlterTableStmt) AddCons(key string) ifc.AddCons {
	s.call(&mig.AddCons{Key: key})
	return s
}

// DropUnique calls DROP INDEX | DROP CONSTRAINT clause.
func (s *AlterTableStmt) DropUnique(key string) ifc.DropUnique {
	s.call(&mig.DropUnique{Driver: s.conn.getDriver(), Key: key})
	return s
}

// DropPrimary calls DROP PRIMARY KEY | DROP CONSTRAINT clause.
func (s *AlterTableStmt) DropPrimary(key string) ifc.DropPrimary {
	s.call(&mig.DropPrimary{Driver: s.conn.getDriver(), Key: key})
	return s
}

// DropForeign calls DROP FOREIGN KEY | DROP CONSTRAINT clause.
func (s *AlterTableStmt) DropForeign(key string) ifc.DropForeign {
	s.call(&mig.DropForeign{Driver: s.conn.getDriver(), Key: key})
	return s
}

// Unique calls UNIQUE keyword.
func (s *AlterTableStmt) Unique(cols ...string) ifc.Unique {
	s.call(&mig.Unique{Columns: cols})
	return s
}

// Primary calls PRIMARY KEY keyword.
func (s *AlterTableStmt) Primary(cols ...string) ifc.Primary {
	s.call(&mig.Primary{Columns: cols})
	return s
}

// Foreign calls FOREIGN KEY keyword.
func (s *AlterTableStmt) Foreign(cols ...string) ifc.Foreign {
	s.call(&mig.Foreign{Columns: cols})
	return s
}

// Ref calls REFERENCES keyword.
func (s *AlterTableStmt) Ref(table, col string) ifc.Ref {
	s.call(&mig.Ref{Table: table, Column: col})
	return s
}
