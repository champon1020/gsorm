package statement

import (
	"reflect"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/interfaces/iupdate"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/internal/parser"
	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/morikuni/failure"
)

// UpdateStmt is UPDATE statement..
type UpdateStmt struct {
	stmt
	model     interface{}
	modelCols []string
	cmd       *clause.Update
}

func NewUpdateStmt(conn domain.Conn, table string) *UpdateStmt {
	u := new(clause.Update)
	u.AddTable(table)
	stmt := &UpdateStmt{cmd: u}
	stmt.conn = conn
	return stmt
}

// String returns SQL statement with string.
func (s *UpdateStmt) String() string {
	return s.string(s.buildSQL)
}

// FuncString returns function call as string.
func (s *UpdateStmt) FuncString() string {
	return s.funcString(s.cmd)
}

// Cmd returns cmd clause.
func (s *UpdateStmt) Cmd() domain.Clause {
	return s.cmd
}

// CompareWith compares the statements and returns error if the statements is not same.
// In this case, same means that stmt.cmd and stmt.called is corresponding.
func (s *UpdateStmt) CompareWith(targetStmt domain.Stmt) error {
	return s.compareWith(s.Cmd(), targetStmt)
}

// Exec executes SQL statement without mapping to model.
// If type of conn is gsorm.MockDB, compare statements between called and expected.
func (s *UpdateStmt) Exec() error {
	return s.exec(s.buildSQL, s)
}

// buildSQL builds SQL statement.
func (s *UpdateStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	if s.model != nil {
		cols := []string{}
		cols = append(cols, s.modelCols...)
		if err = s.buildSQLWithModel(cols, s.model, sql); err != nil {
			return err
		}
	}

	if err = s.buildSQLWithClauses(sql); err != nil {
		return err
	}
	return nil
}

// buildSQLWithClauses builds SQL statement from called clauses.
func (s *UpdateStmt) buildSQLWithClauses(sql *internal.SQL) error {
	setCalled := false
	for _, e := range s.called {
		switch e := e.(type) {
		case *syntax.RawClause,
			*clause.Where,
			*clause.And,
			*clause.Or:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
		case *clause.Set:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			if setCalled {
				sql.Write(",")
				sql.Write(ss.BuildValue())
				continue
			}
			sql.Write(ss.Build())
			setCalled = true
		default:
			return failure.New(errInvalidClause,
				failure.Context{"clause": reflect.TypeOf(e).Elem().String()},
				failure.Message("invalid clause for UPDATE"))
		}
	}
	return nil
}

// buildSQLWithModel builds SQL statement from model.
func (s *UpdateStmt) buildSQLWithModel(cols []string, model interface{}, sql *internal.SQL) error {
	sql.Write("SET")
	p, err := parser.NewUpdateModelParser(cols, model)
	if err != nil {
		return err
	}

	modelSQL, err := p.Parse()
	if err != nil {
		return failure.Translate(err, errFailedParse)
	}

	sql.Write(modelSQL.String())
	return nil
}

// RawClause calls the raw string clause.
func (s *UpdateStmt) RawClause(rs string, v ...interface{}) iupdate.RawClause {
	s.call(&syntax.RawClause{RawStr: rs, Values: v})
	return s
}

// Model sets model to UpdateStmt.
func (s *UpdateStmt) Model(model interface{}, cols ...string) iupdate.Model {
	s.model = model
	s.modelCols = cols
	return s
}

// Set calls SET clause.
func (s *UpdateStmt) Set(col string, val interface{}) iupdate.Set {
	s.call(&clause.Set{Column: col, Value: val})
	return s
}

// Where calls WHERE clause.
func (s *UpdateStmt) Where(expr string, vals ...interface{}) iupdate.Where {
	s.call(&clause.Where{Expr: expr, Values: vals})
	return s
}

// And calls AND clause.
func (s *UpdateStmt) And(expr string, vals ...interface{}) iupdate.And {
	s.call(&clause.And{Expr: expr, Values: vals})
	return s
}

// Or calls OR clause.
func (s *UpdateStmt) Or(expr string, vals ...interface{}) iupdate.Or {
	s.call(&clause.Or{Expr: expr, Values: vals})
	return s
}
