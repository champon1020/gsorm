package migration

import (
	"reflect"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/interfaces/ialtertable"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
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
func (s *AlterTableStmt) RawClause(raw string, values ...interface{}) ialtertable.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: values})
	return s
}

// Rename calls RENAME TO clause.
func (s *AlterTableStmt) Rename(table string) ialtertable.Rename {
	s.call(&mig.Rename{Table: table})
	return s
}

// AddColumn calls ADD COLUMN clause.
func (s *AlterTableStmt) AddColumn(column, typ string) ialtertable.AddColumn {
	s.call(&mig.AddColumn{Column: column, Type: typ})
	return s
}

// DropColumn calls DROP COLUMN clause.
func (s *AlterTableStmt) DropColumn(column string) ialtertable.DropColumn {
	s.call(&mig.DropColumn{Column: column})
	return s
}

// RenameColumn calls RENAME COLUMN clause.
func (s *AlterTableStmt) RenameColumn(column, dest string) ialtertable.RenameColumn {
	s.call(&mig.RenameColumn{Column: column, Dest: dest})
	return s
}

// NotNull calls NOT NULL option.
func (s *AlterTableStmt) NotNull() ialtertable.NotNull {
	s.call(&mig.NotNull{})
	return s
}

// Default calls DEFAULT option.
func (s *AlterTableStmt) Default(value interface{}) ialtertable.Default {
	s.call(&mig.Default{Value: value})
	return s
}

// AddCons calls ADD CONSTRAINT clause.
func (s *AlterTableStmt) AddCons(key string) ialtertable.AddCons {
	s.call(&mig.AddCons{Key: key})
	return s
}

// Unique calls UNIQUE keyword.
func (s *AlterTableStmt) Unique(columns ...string) ialtertable.Unique {
	s.call(&mig.Unique{Columns: columns})
	return s
}

// Primary calls PRIMARY KEY keyword.
func (s *AlterTableStmt) Primary(columns ...string) ialtertable.Primary {
	s.call(&mig.Primary{Columns: columns})
	return s
}

// Foreign calls FOREIGN KEY keyword.
func (s *AlterTableStmt) Foreign(columns ...string) ialtertable.Foreign {
	s.call(&mig.Foreign{Columns: columns})
	return s
}

// Ref calls REFERENCES keyword.
func (s *AlterTableStmt) Ref(table string, columns ...string) ialtertable.Ref {
	s.call(&mig.Ref{Table: table, Columns: columns})
	return s
}
