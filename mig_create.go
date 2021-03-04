package mgorm

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax/mig"

	prCreate "github.com/champon1020/mgorm/provider/create"
)

// CreateDBStmt is CREATE DATABASE statement.
type CreateDBStmt struct {
	migStmt
	cmd *mig.CreateDB
}

func (s *CreateDBStmt) String() string {
	return s.string(s.buildSQL)
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
func (s *CreateIndexStmt) On(table string, cols ...string) prCreate.OnMP {
	s.call(&mig.On{Table: table, Columns: cols})
	return s
}

// These are Column options which is used as field tag.
const (
	// Column name.
	colName = "mgorm"

	// Column type.
	colType = "type"

	// NOT NULL option.
	notnull = "notnull"

	// DEFAULT option.
	def = "def"

	// AUTO_INCREMENT option.
	autoinc = "autoinc"

	// Unique key.
	unique = "uc"

	// Primary key.
	primary = "pk"

	// Foreign key.
	foreign = "fk"
)

// CreateTableStmt is CREATE TABLE statement.
type CreateTableStmt struct {
	migStmt
	model interface{}
	cmd   *mig.CreateTable
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

func (s *CreateTableStmt) buildSQLWithModel(sql *internal.SQL) error {
	typ := reflect.TypeOf(s.model)
	if typ.Kind() != reflect.Ptr {
		return errors.New("Model must be pointer", errors.InvalidValueError)
	}

	typ = typ.Elem()
	if typ.Kind() != reflect.Struct {
		return errors.New("Type of model must be pointer of struct", errors.InvalidTypeError)
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

		// Write column name.
		var name string
		if v, ok := f.Tag.Lookup(colName); ok {
			name = v
		} else {
			name = internal.ConvertToSnakeCase(f.Name)
		}
		sql.Write(name)

		// Write column type.
		if v := f.Tag.Get(colType); v != "" {
			sql.Write(v)
		} else {
			dbtyp := convertToDBType(f.Type, s.conn.getDriver())
			if dbtyp == "" {
				msg := fmt.Sprintf("Type of %v is not supported for database column", f.Type)
				return errors.New(msg, errors.InvalidTypeError)
			}
			sql.Write(dbtyp)
		}

		// Write NOT NULL option if exist.
		if _, ok := f.Tag.Lookup(notnull); ok {
			sql.Write("NOT NULL")
		}
		// Write DEFAULT option if exist.
		if v, ok := f.Tag.Lookup(def); ok {
			sql.Write(fmt.Sprintf("DEFAULT %s", v))
		}
		// Write AUTO_INCREMENT option if exist.
		if _, ok := f.Tag.Lookup(autoinc); ok {
			sql.Write("AUTO_INCREMENT")
		}
		// Store unique key if exist.
		if v := f.Tag.Get(unique); v != "" {
			uc[v] = append(uc[v], name)
		}
		// Store primary key if exist.
		if v := f.Tag.Get(primary); v != "" {
			pk[v] = append(pk[v], name)
		}
		// Store foreign key if exist.
		if v := f.Tag.Get(foreign); v != "" {
			el := strings.Split(v, " ")
			if len(el) != 2 {
				msg := `Format of tag for foreign key must be fk:"<key> <ref_table>(<ref_column>)"`
				return errors.New(msg, errors.InvalidSyntaxError)
			}
			fk[el[0]] = append(fk[el[0]], name)
			if _, ok := ref[el[0]]; ok && ref[el[0]] != el[1] {
				msg := "Different reference is used with same key"
				return errors.New(msg, errors.InvalidSyntaxError)
			}
			ref[el[0]] = el[1]
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

// convertToDBType converts golang type to database type.
func convertToDBType(t reflect.Type, d internal.SQLDriver) string {
	switch t.Kind() {
	case reflect.String:
		return "VARCHAR(128)"
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return "INT"
	case reflect.Float32, reflect.Float64:
		if d == internal.MySQL {
			return "FLOAT"
		}
		return "NUMERIC"
	case reflect.Struct:
		if t == reflect.TypeOf(time.Time{}) {
			return "DATE"
		}
	case reflect.Bool:
		return "SMALLINT"
	}
	return ""
}

// Model sets model to CreateTableStmt.
func (s *CreateTableStmt) Model(model interface{}) prCreate.ModelMP {
	s.model = model
	return s
}

// Column calls table column definition.
func (s *CreateTableStmt) Column(col, typ string) prCreate.ColumnMP {
	s.call(&mig.Column{Col: col, Type: typ})
	return s
}

// NotNull calls NOT NULL option.
func (s *CreateTableStmt) NotNull() prCreate.NotNullMP {
	s.call(&mig.NotNull{})
	return s
}

// AutoInc calls AUTO_INCREMENT option (only MySQL).
func (s *CreateTableStmt) AutoInc() prCreate.AutoIncMP {
	s.call(&mig.AutoInc{})
	return s
}

// Default calls DEFAULT option.
func (s *CreateTableStmt) Default(val interface{}) prCreate.DefaultMP {
	s.call(&mig.Default{Value: val})
	return s
}

// Cons calls CONSTRAINT option.
func (s *CreateTableStmt) Cons(key string) prCreate.ConsMP {
	s.call(&mig.Cons{Key: key})
	return s
}

// Unique calls UNIQUE keyword.
func (s *CreateTableStmt) Unique(cols ...string) prCreate.UniqueMP {
	s.call(&mig.Unique{Columns: cols})
	return s
}

// Primary calls PRIMARY KEY keyword.
func (s *CreateTableStmt) Primary(cols ...string) prCreate.PrimaryMP {
	s.call(&mig.Primary{Columns: cols})
	return s
}

// Foreign calls FOREIGN KEY keyword.
func (s *CreateTableStmt) Foreign(cols ...string) prCreate.ForeignMP {
	s.call(&mig.Foreign{Columns: cols})
	return s
}

// Ref calls REFERENCES keyword.
func (s *CreateTableStmt) Ref(table, col string) prCreate.RefMP {
	s.call(&mig.Ref{Table: table, Column: col})
	return s
}
