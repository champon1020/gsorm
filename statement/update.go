package statement

import (
	"reflect"

	"github.com/champon1020/mgorm/domain"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/morikuni/failure"

	ifc "github.com/champon1020/mgorm/interfaces/update"
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
func (s *UpdateStmt) Cmd() syntax.Clause {
	return s.cmd
}

// Exec executes SQL statement without mapping to model.
// If type of conn is mgorm.MockDB, compare statements between called and expected.
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
	for i, e := range s.called {
		switch e := e.(type) {
		case *clause.Where,
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
			if i == 0 {
				sql.Write(ss.Build())
			} else if _, ok := s.called[i-1].(*clause.Set); !ok {
				sql.Write(ss.Build())
			} else {
				sql.Write(",")
				sql.Write(ss.Value)
			}
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
	parser, err := internal.NewUpdateModelParser(cols, model)
	if err != nil {
		return err
	}

	modelSQL, err := parser.Parse()
	if err != nil {
		return failure.Translate(err, errFailedParse)
	}

	sql.Write(modelSQL.String())
	return nil
}

// Model sets model to UpdateStmt.
func (s *UpdateStmt) Model(model interface{}, cols ...string) ifc.Model {
	s.model = model
	s.modelCols = cols
	return s
}

// Set calls SET clause.
func (s *UpdateStmt) Set(col string, val interface{}) ifc.Set {
	s.call(&clause.Set{Column: col, Value: val})
	return s
}

// Where calls WHERE clause.
func (s *UpdateStmt) Where(expr string, vals ...interface{}) ifc.Where {
	s.call(&clause.Where{Expr: expr, Values: vals})
	return s
}

// And calls AND clause.
func (s *UpdateStmt) And(expr string, vals ...interface{}) ifc.And {
	s.call(&clause.And{Expr: expr, Values: vals})
	return s
}

// Or calls OR clause.
func (s *UpdateStmt) Or(expr string, vals ...interface{}) ifc.Or {
	s.call(&clause.Or{Expr: expr, Values: vals})
	return s
}
