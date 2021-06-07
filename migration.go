package gsorm

import (
	"reflect"
	"strings"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/interfaces/ialtertable"
	"github.com/champon1020/gsorm/interfaces/icreatedb"
	"github.com/champon1020/gsorm/interfaces/icreateindex"
	"github.com/champon1020/gsorm/interfaces/icreatetable"
	"github.com/champon1020/gsorm/interfaces/idropdb"
	"github.com/champon1020/gsorm/interfaces/idroptable"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
	"github.com/morikuni/failure"
	"golang.org/x/xerrors"
)

// migStmt stores information about database migration query.
type migStmt struct {
	conn   conn
	called []domain.Clause
	errors []error
}

// call appends called clause.
func (s *migStmt) call(e domain.Clause) {
	s.called = append(s.called, e)
}

// throw appends occurred error.
func (s *migStmt) throw(e error) {
	s.errors = append(s.errors, e)
}

// headClause returns first element of called.
func (s *migStmt) headClause() domain.Clause {
	if len(s.called) == 0 {
		return nil
	}
	return s.called[0]
}

// advanceClause slides slice of called.
func (s *migStmt) advanceClause() {
	s.called = s.called[1:]
}

func (s *migStmt) sql(buildSQL func(*internal.SQL) error) string {
	var sql internal.SQL
	if err := buildSQL(&sql); err != nil {
		s.throw(err)
		return ""
	}
	return sql.String()
}

func (s *migStmt) migration(buildSQL func(*internal.SQL) error) error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch conn := s.conn.(type) {
	case Mock:
		return nil
	case DB, Tx:
		var sql internal.SQL
		if err := buildSQL(&sql); err != nil {
			return err
		}
		if _, err := conn.Exec(sql.String()); err != nil {
			return failure.Wrap(err)
		}
		return nil
	}

	return xerrors.Errorf("database connection should not be %s", reflect.TypeOf(s.conn).String())
}

func (s *migStmt) buildColumnOptSQL(sql *internal.SQL) error {
	for len(s.called) > 0 {
		e := s.headClause()
		if e == nil {
			return nil
		}

		switch e := e.(type) {
		case *syntax.RawClause,
			*mig.NotNull,
			*mig.Default:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
		default:
			return nil
		}

		s.advanceClause()
	}

	return nil
}

func (s *migStmt) buildConstraintSQL(sql *internal.SQL) error {
	e := s.headClause()
	if e == nil {
		return xerrors.New("syntax is invalid")
	}

	if rc, ok := e.(*syntax.RawClause); ok {
		ss, err := rc.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())
		s.advanceClause()
		e = s.headClause()
	}

	switch e := e.(type) {
	case *mig.Primary, *mig.Unique:
		ss, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())
		s.advanceClause()
		return nil
	case *mig.Foreign:
		ss, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())
		s.advanceClause()
		return s.buildRefSQL(sql)
	}

	return xerrors.Errorf("%s is invalid clause for CONSTRAINT", reflect.TypeOf(e).String())
}

func (s *migStmt) buildRefSQL(sql *internal.SQL) error {
	e := s.headClause()
	if e == nil {
		return xerrors.New("syntax is invalid")
	}

	if rc, ok := e.(*syntax.RawClause); ok {
		ss, err := rc.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())
		s.advanceClause()
		e = s.headClause()
	}

	switch e := e.(type) {
	case *mig.Ref:
		ss, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(ss.Build())
		s.advanceClause()
		return nil
	}

	return xerrors.Errorf("%s is invalid clause for FOREIGN KEY", reflect.TypeOf(e).String())
}

// AlterTableStmt is ALTER TABLE statement.
type AlterTableStmt struct {
	migStmt
	cmd *mig.AlterTable
}

// newAlterTableStmt creates AlterTableStmt instance.
func newAlterTableStmt(conn conn, table string) *AlterTableStmt {
	stmt := &AlterTableStmt{cmd: &mig.AlterTable{Table: table}}
	stmt.conn = conn
	return stmt
}

// SQL returns the built SQL string.
func (s *AlterTableStmt) SQL() string {
	return s.sql(s.buildSQL)
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
			return xerrors.New("syntax is invalid")
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
			return xerrors.Errorf("%s is invalid clause for ALTER TABLE", reflect.TypeOf(e).String())
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

// CreateDBStmt is CREATE DATABASE statement.
type CreateDBStmt struct {
	migStmt
	cmd *mig.CreateDB
}

// newCreateDBStmt creates CreateDBStmt instance.
func newCreateDBStmt(conn conn, dbName string) *CreateDBStmt {
	stmt := &CreateDBStmt{cmd: &mig.CreateDB{DBName: dbName}}
	stmt.conn = conn
	return stmt
}

// SQL returns the built SQL string.
func (s *CreateDBStmt) SQL() string {
	return s.sql(s.buildSQL)
}

// Migrate executes database migration.
func (s *CreateDBStmt) Migrate() error {
	return s.migration(s.buildSQL)
}

func (s *CreateDBStmt) buildSQL(sql *internal.SQL) error {
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
		case *syntax.RawClause:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
		default:
			return xerrors.Errorf("%s is invalid clause for CREATE DATABASE", reflect.TypeOf(e).String())
		}
	}

	return nil
}

// RawClause calls the raw string clause.
func (s *CreateDBStmt) RawClause(raw string, values ...interface{}) icreatedb.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: values})
	return s
}

// CreateIndexStmt is CREATE INDEX statement.
type CreateIndexStmt struct {
	migStmt
	cmd *mig.CreateIndex
}

// newCreateIndexStmt creates CreateIndexStmt instance.
func newCreateIndexStmt(conn conn, idx string) *CreateIndexStmt {
	stmt := &CreateIndexStmt{cmd: &mig.CreateIndex{IdxName: idx}}
	stmt.conn = conn
	return stmt
}

// SQL returns the built SQL string.
func (s *CreateIndexStmt) SQL() string {
	return s.sql(s.buildSQL)
}

// Migrate executes database migration.
func (s *CreateIndexStmt) Migrate() error {
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
		case *syntax.RawClause,
			*mig.On:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
		default:
			return xerrors.Errorf("%s is invalid clause for CREATE INDEX", reflect.TypeOf(e).String())
		}
	}

	return nil
}

// On calls ON clause.
func (s *CreateIndexStmt) On(table string, columns ...string) icreateindex.On {
	s.call(&mig.On{Table: table, Columns: columns})
	return s
}

// RawClause calls the raw string clause.
func (s *CreateIndexStmt) RawClause(raw string, values ...interface{}) icreateindex.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: values})
	return s
}

// CreateTableStmt is CREATE TABLE statement.
type CreateTableStmt struct {
	migStmt
	model interface{}
	cmd   *mig.CreateTable
}

// newCreateTableStmt creates CreateTableStmt instance.
func newCreateTableStmt(conn conn, table string) *CreateTableStmt {
	stmt := &CreateTableStmt{cmd: &mig.CreateTable{Table: table}}
	stmt.conn = conn
	return stmt
}

// SQL returns the built SQL string.
func (s *CreateTableStmt) SQL() string {
	return s.sql(s.buildSQL)
}

// Migrate executes database migration.
func (s *CreateTableStmt) Migrate() error {
	return s.migration(s.buildSQL)
}

func (s *CreateTableStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	if s.model != nil {
		return s.buildSQLWithModel(sql)
	}

	return s.buildSQLWithClauses(sql)
}

func (s *CreateTableStmt) buildSQLWithClauses(sql *internal.SQL) error {
	sql.Write("(")
	for len(s.called) > 0 {
		e := s.headClause()
		if e == nil {
			return xerrors.New("syntax is invalid")
		}

		switch e := e.(type) {
		case *syntax.RawClause:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
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
			return xerrors.Errorf("%s is invalid clause for CREATE TABLE", reflect.TypeOf(e).String())
		}
	}
	sql.Write(")")
	return nil
}

func (s *CreateTableStmt) buildSQLWithModel(sql *internal.SQL) error {
	p, err := newCreateTableModelParser(s.model)
	if err != nil {
		return err
	}

	modelSQL, err := p.Parse()
	if err != nil {
		return err
	}

	sql.Write(modelSQL.String())
	return nil
}

// RawClause calls the raw string clause.
func (s *CreateTableStmt) RawClause(raw string, values ...interface{}) icreatetable.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: values})
	return s
}

// Model sets model to CreateTableStmt.
func (s *CreateTableStmt) Model(model interface{}) icreatetable.Model {
	s.model = model
	return s
}

// Column calls table column definition.
func (s *CreateTableStmt) Column(column, typ string) icreatetable.Column {
	s.call(&mig.Column{Col: column, Type: typ})
	return s
}

// NotNull calls NOT NULL option.
func (s *CreateTableStmt) NotNull() icreatetable.NotNull {
	s.call(&mig.NotNull{})
	return s
}

// Default calls DEFAULT option.
func (s *CreateTableStmt) Default(value interface{}) icreatetable.Default {
	s.call(&mig.Default{Value: value})
	return s
}

// Cons calls CONSTRAINT option.
func (s *CreateTableStmt) Cons(key string) icreatetable.Cons {
	s.call(&mig.Cons{Key: key})
	return s
}

// Unique calls UNIQUE keyword.
func (s *CreateTableStmt) Unique(columns ...string) icreatetable.Unique {
	s.call(&mig.Unique{Columns: columns})
	return s
}

// Primary calls PRIMARY KEY keyword.
func (s *CreateTableStmt) Primary(columns ...string) icreatetable.Primary {
	s.call(&mig.Primary{Columns: columns})
	return s
}

// Foreign calls FOREIGN KEY keyword.
func (s *CreateTableStmt) Foreign(columns ...string) icreatetable.Foreign {
	s.call(&mig.Foreign{Columns: columns})
	return s
}

// Ref calls REFERENCES keyword.
func (s *CreateTableStmt) Ref(table string, columns ...string) icreatetable.Ref {
	s.call(&mig.Ref{Table: table, Columns: columns})
	return s
}

// DropDBStmt is DROP DATABASE statement.
type DropDBStmt struct {
	migStmt
	cmd *mig.DropDB
}

// newDropDBStmt creates DropDBStmt instance.
func newDropDBStmt(conn conn, dbName string) *DropDBStmt {
	stmt := &DropDBStmt{cmd: &mig.DropDB{DBName: dbName}}
	stmt.conn = conn
	return stmt
}

// SQL returns the built SQL string.
func (s *DropDBStmt) SQL() string {
	return s.sql(s.buildSQL)
}

// Migrate executes database migration.
func (s *DropDBStmt) Migrate() error {
	return s.migration(s.buildSQL)
}

func (s *DropDBStmt) buildSQL(sql *internal.SQL) error {
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
		case *syntax.RawClause:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
		default:
			return xerrors.Errorf("%s is invalid clause for DROP DATABASE", reflect.TypeOf(e).String())
		}
	}

	return nil
}

// RawClause calls the raw string clause.
func (s *DropDBStmt) RawClause(raw string, value ...interface{}) idropdb.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: value})
	return s
}

// DropTableStmt is DROP TABLE statement.
type DropTableStmt struct {
	migStmt
	cmd *mig.DropTable
}

// newDropTableStmt creates DropTableStmt instance.
func newDropTableStmt(conn conn, table string) *DropTableStmt {
	stmt := &DropTableStmt{cmd: &mig.DropTable{Table: table}}
	stmt.conn = conn
	return stmt
}

// SQL returns the built SQL string.
func (s *DropTableStmt) SQL() string {
	return s.sql(s.buildSQL)
}

// Migrate executes database migration.
func (s *DropTableStmt) Migrate() error {
	return s.migration(s.buildSQL)
}

func (s *DropTableStmt) buildSQL(sql *internal.SQL) error {
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
		case *syntax.RawClause:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
			s.advanceClause()
		default:
			return xerrors.Errorf("%s is invalid clause for DROP TABLE", reflect.TypeOf(e).String())
		}
	}

	return nil
}

// RawClause calls the raw string clause.
func (s *DropTableStmt) RawClause(raw string, value ...interface{}) idroptable.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: value})
	return s
}
