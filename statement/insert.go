package statement

import (
	"reflect"

	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/internal/parser"
	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/morikuni/failure"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/interfaces/iinsert"
)

// InsertStmt is INSERT statement.
type InsertStmt struct {
	stmt
	model interface{}
	cmd   *clause.Insert
	sel   domain.Stmt
}

// NewInsertStmt creates InsertStmt instance.
func NewInsertStmt(conn domain.Conn, table string, cols ...string) *InsertStmt {
	i := new(clause.Insert)
	i.AddTable(table)
	i.AddColumns(cols...)
	stmt := &InsertStmt{cmd: i}
	stmt.conn = conn
	return stmt
}

// String returns SQL statement with string.
func (s *InsertStmt) String() string {
	return s.string(s.buildSQL)
}

// FuncString returns function call as string.
func (s *InsertStmt) FuncString() string {
	return s.funcString(s.cmd)
}

// Cmd returns cmd clause.
func (s *InsertStmt) Cmd() domain.Clause {
	return s.cmd
}

// CompareWith compares the statements and returns error if the statements is not same.
// In this case, same means that stmt.cmd and stmt.called is corresponding.
func (s *InsertStmt) CompareWith(targetStmt domain.Stmt) error {
	return s.compareWith(s.Cmd(), targetStmt)
}

// Exec executed SQL statement without mapping to model.
// If type of conn is gsorm.MockDB, compare statements between called and expected.
func (s *InsertStmt) Exec() error {
	return s.exec(s.buildSQL, s)
}

// buildSQL builds SQL statement.
func (s *InsertStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	if s.model != nil {
		cols := []string{}
		for _, c := range s.cmd.Columns {
			if c.Alias != "" {
				cols = append(cols, c.Alias)
				continue
			}
			cols = append(cols, c.Name)
		}
		if err := s.buildSQLWithModel(cols, s.model, sql); err != nil {
			return err
		}
		return nil
	}

	if s.sel != nil {
		sql.Write(s.sel.String())
		return nil
	}

	if err := s.buildSQLWithClauses(sql); err != nil {
		return err
	}
	return nil
}

// buildSQLWithClauses builds SQL statement from called clauses.
func (s *InsertStmt) buildSQLWithClauses(sql *internal.SQL) error {
	valuesCalled := false
	for _, e := range s.called {
		switch e := e.(type) {
		case *clause.Values:
			s, err := e.Build()
			if err != nil {
				return err
			}
			if valuesCalled {
				sql.Write(",")
				sql.Write(s.BuildValue())
				continue
			}
			sql.Write(s.Build())
			valuesCalled = true
		case *syntax.RawClause:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
		default:
			return failure.New(errInvalidClause,
				failure.Context{"clause": reflect.TypeOf(e).Elem().String()},
				failure.Message("invalid clause for INSERT"))
		}
	}
	return nil
}

// buildSQLWithModel builds SQL statement from model.
func (s *InsertStmt) buildSQLWithModel(cols []string, model interface{}, sql *internal.SQL) error {
	sql.Write("VALUES")
	p, err := parser.NewInsertModelParser(cols, model)
	if err != nil {
		return failure.Translate(err, errFailedParse)
	}

	modelSQL, err := p.Parse()
	if err != nil {
		return failure.Translate(err, errFailedParse)
	}

	sql.Write(modelSQL.String())
	return nil
}

// RawClause calls the raw string clause.
func (s *InsertStmt) RawClause(raw string, values ...interface{}) iinsert.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: values})
	return s
}

// Model sets model to InsertStmt.
func (s *InsertStmt) Model(model interface{}) iinsert.Model {
	s.model = model
	return s
}

// Select calls SELECT statement.
func (s *InsertStmt) Select(stmt domain.Stmt) iinsert.Select {
	s.sel = stmt
	return s
}

// Values calls VALUES clause.
func (s *InsertStmt) Values(values ...interface{}) iinsert.Values {
	v := new(clause.Values)
	for _, val := range values {
		v.AddValue(val)
	}
	s.call(v)
	return s
}
