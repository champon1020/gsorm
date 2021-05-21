package migration

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/interfaces/icreatetable"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
	"github.com/morikuni/failure"
)

// CreateTableStmt is CREATE TABLE statement.
type CreateTableStmt struct {
	migStmt
	model interface{}
	cmd   *mig.CreateTable
}

// NewCreateTableStmt creates CreateTableStmt instance.
func NewCreateTableStmt(conn domain.Conn, table string) *CreateTableStmt {
	stmt := &CreateTableStmt{cmd: &mig.CreateTable{Table: table}}
	stmt.conn = conn
	return stmt
}

func (s *CreateTableStmt) String() string {
	return s.string(s.buildSQL)
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
			return failure.New(errInvalidSyntax,
				failure.Message("the SQL statement is not completed or the syntax is not supported"))
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
			return failure.New(errInvalidClause,
				failure.Context{"clause": reflect.TypeOf(e).String()},
				failure.Message("invalid clause for CREATE TABLE"))
		}
	}
	sql.Write(")")
	return nil
}

func (s *CreateTableStmt) buildSQLWithModel(sql *internal.SQL) error {
	typ := reflect.TypeOf(s.model)
	if typ.Kind() != reflect.Ptr {
		return failure.New(errInvalidValue, failure.Message("model must be pointer"))
	}

	typ = typ.Elem()
	if typ.Kind() != reflect.Struct {
		return failure.New(errInvalidValue, failure.Message("model must be pointer of struct"))
	}

	var (
		uc  = make(map[string][]string)
		pk  = make(map[string][]string)
		fk  = make(map[string][]string)
		ref = make(map[string]string)
	)
	sql.Write("(")
	for i := 0; i < typ.NumField(); i++ {
		if i > 0 {
			sql.Write(",")
		}
		f := typ.Field(i)
		tag := internal.ExtractTag(f)

		// Write column name.
		var name string
		if tag.Lookup("col") {
			name = tag.Column
		} else {
			name = internal.SnakeCase(f.Name)
		}
		sql.Write(name)

		// Write column type.
		if tag.Lookup("typ") {
			sql.Write(tag.Type)
		} else {
			if s.conn == nil {
				return failure.New(errFailedDBConnection, failure.Message("gsorm.db.conn is nil"))
			}
			dbtyp := s.conn.GetDriver().LookupDefaultType(f.Type)
			if dbtyp == "" {
				return failure.New(errInvalidType,
					failure.Context{"type": f.Type.String()},
					failure.Message("invalid type for database column"))
			}
			sql.Write(dbtyp)
		}

		// Write NOT NULL option if exist.
		if tag.Lookup("notnull") {
			sql.Write("NOT NULL")
		}
		// Write DEFAULT option if exist.
		if tag.Lookup("default") {
			sql.Write(fmt.Sprintf("DEFAULT %s", tag.Default))
		}
		// Store unique key if exist.
		if tag.Lookup("uc") {
			uc[tag.UC] = append(uc[tag.UC], name)
		}
		// Store primary key if exist.
		if tag.Lookup("pk") {
			pk[tag.PK] = append(pk[tag.PK], name)
		}
		// Store foreign key if exist.
		if tag.Lookup("fk") {
			fk[tag.FK] = append(fk[tag.FK], name)
			ref[tag.FK] = tag.Ref
		}
	}

	// Write unique key if exist.
	for k, v := range uc {
		sql.Write(",")
		sql.Write(fmt.Sprintf("CONSTRAINT %s UNIQUE (%s)", k, strings.Join(v, ", ")))
	}
	// Write primary key if exist.
	for k, v := range pk {
		sql.Write(",")
		sql.Write(fmt.Sprintf("CONSTRAINT %s PRIMARY KEY (%s)", k, strings.Join(v, ", ")))
	}
	// Write foreign key if exist.
	for k, v := range fk {
		sql.Write(",")
		sql.Write(fmt.Sprintf("CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s", k, strings.Join(v, ", "), ref[k]))
	}
	sql.Write(")")
	return nil
}

// RawClause calls the raw string clause.
func (s *CreateTableStmt) RawClause(rs string, v ...interface{}) icreatetable.RawClause {
	s.call(&syntax.RawClause{RawStr: rs, Values: v})
	return s
}

// Model sets model to CreateTableStmt.
func (s *CreateTableStmt) Model(model interface{}) icreatetable.Model {
	s.model = model
	return s
}

// Column calls table column definition.
func (s *CreateTableStmt) Column(col, typ string) icreatetable.Column {
	s.call(&mig.Column{Col: col, Type: typ})
	return s
}

// NotNull calls NOT NULL option.
func (s *CreateTableStmt) NotNull() icreatetable.NotNull {
	s.call(&mig.NotNull{})
	return s
}

// Default calls DEFAULT option.
func (s *CreateTableStmt) Default(val interface{}) icreatetable.Default {
	s.call(&mig.Default{Value: val})
	return s
}

// Cons calls CONSTRAINT option.
func (s *CreateTableStmt) Cons(key string) icreatetable.Cons {
	s.call(&mig.Cons{Key: key})
	return s
}

// Unique calls UNIQUE keyword.
func (s *CreateTableStmt) Unique(cols ...string) icreatetable.Unique {
	s.call(&mig.Unique{Columns: cols})
	return s
}

// Primary calls PRIMARY KEY keyword.
func (s *CreateTableStmt) Primary(cols ...string) icreatetable.Primary {
	s.call(&mig.Primary{Columns: cols})
	return s
}

// Foreign calls FOREIGN KEY keyword.
func (s *CreateTableStmt) Foreign(cols ...string) icreatetable.Foreign {
	s.call(&mig.Foreign{Columns: cols})
	return s
}

// Ref calls REFERENCES keyword.
func (s *CreateTableStmt) Ref(table string, cols ...string) icreatetable.Ref {
	s.call(&mig.Ref{Table: table, Columns: cols})
	return s
}
