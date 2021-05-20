package migration

import (
	"reflect"

	"github.com/champon1020/mgorm/interfaces/domain"
	"github.com/champon1020/mgorm/interfaces/ialtertable"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/morikuni/failure"
)

// AlterTableStmt is ALTER TABLE statement.
type AlterTableStmt struct {
	migStmt
	cmd *mig.AlterTable
}

// NewAlterTableStmt creates AlterTableStmt instance.
func NewAlterTableStmt(conn domain.Conn, table string) *AlterTableStmt {
	stmt := &AlterTableStmt{cmd: &mig.AlterTable{Table: table}}
	stmt.conn = conn
	return stmt
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
		case *syntax.RawClause,
			*mig.Rename,
			*mig.RenameColumn,
			*mig.DropColumn:
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

// RawClause calls the raw string clause.
func (s *AlterTableStmt) RawClause(rs string, v ...interface{}) ialtertable.RawClause {
	s.call(&syntax.RawClause{RawStr: rs, Values: v})
	return s
}

// Rename calls RENAME TO clause.
func (s *AlterTableStmt) Rename(table string) ialtertable.Rename {
	s.call(&mig.Rename{Table: table})
	return s
}

// AddColumn calls ADD COLUMN clause.
func (s *AlterTableStmt) AddColumn(col, typ string) ialtertable.AddColumn {
	s.call(&mig.AddColumn{Column: col, Type: typ})
	return s
}

// DropColumn calls DROP COLUMN clause.
func (s *AlterTableStmt) DropColumn(col string) ialtertable.DropColumn {
	s.call(&mig.DropColumn{Column: col})
	return s
}

// RenameColumn calls RENAME COLUMN clause.
func (s *AlterTableStmt) RenameColumn(col, dest string) ialtertable.RenameColumn {
	s.call(&mig.RenameColumn{Column: col, Dest: dest})
	return s
}

// NotNull calls NOT NULL option.
func (s *AlterTableStmt) NotNull() ialtertable.NotNull {
	s.call(&mig.NotNull{})
	return s
}

// Default calls DEFAULT option.
func (s *AlterTableStmt) Default(val interface{}) ialtertable.Default {
	s.call(&mig.Default{Value: val})
	return s
}

// AddCons calls ADD CONSTRAINT clause.
func (s *AlterTableStmt) AddCons(key string) ialtertable.AddCons {
	s.call(&mig.AddCons{Key: key})
	return s
}

// Unique calls UNIQUE keyword.
func (s *AlterTableStmt) Unique(cols ...string) ialtertable.Unique {
	s.call(&mig.Unique{Columns: cols})
	return s
}

// Primary calls PRIMARY KEY keyword.
func (s *AlterTableStmt) Primary(cols ...string) ialtertable.Primary {
	s.call(&mig.Primary{Columns: cols})
	return s
}

// Foreign calls FOREIGN KEY keyword.
func (s *AlterTableStmt) Foreign(cols ...string) ialtertable.Foreign {
	s.call(&mig.Foreign{Columns: cols})
	return s
}

// Ref calls REFERENCES keyword.
func (s *AlterTableStmt) Ref(table string, cols ...string) ialtertable.Ref {
	s.call(&mig.Ref{Table: table, Columns: cols})
	return s
}
