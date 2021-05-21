package statement

import (
	"reflect"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/interfaces/iselect"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/morikuni/failure"
)

// SelectStmt is SELECT statement.
type SelectStmt struct {
	stmt
	cmd *clause.Select
}

func NewSelectStmt(conn domain.Conn, cols ...string) *SelectStmt {
	sel := new(clause.Select)
	if len(cols) == 0 {
		sel.AddColumns("*")
	} else {
		sel.AddColumns(cols...)
	}
	stmt := &SelectStmt{cmd: sel}
	stmt.conn = conn
	return stmt
}

// String returns SQL statement with string.
func (s *SelectStmt) String() string {
	return s.string(s.buildSQL)
}

// FuncString returns function call as string.
func (s *SelectStmt) FuncString() string {
	return s.funcString(s.cmd)
}

// Cmd returns cmd clause.
func (s *SelectStmt) Cmd() domain.Clause {
	return s.cmd
}

// CompareWith compares the statements and returns error if the statements is not same.
// In this case, same means that stmt.cmd and stmt.called is corresponding.
func (s *SelectStmt) CompareWith(targetStmt domain.Stmt) error {
	return s.compareWith(s.Cmd(), targetStmt)
}

// Query executes SQL statement with mapping to model.
// If type of (*SelectStmt).conn is gsorm.MockDB, compare statements between called and expected.
// Then, it maps expected values to model.
func (s *SelectStmt) Query(model interface{}) error {
	return s.query(s.buildSQL, s, model)
}

// buildSQL builds SQL statement from called clauses.
func (s *SelectStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	for _, e := range s.called {
		switch e := e.(type) {
		case *syntax.RawClause,
			*clause.From,
			*clause.Join,
			*clause.On,
			*clause.Where,
			*clause.And,
			*clause.Or,
			*clause.GroupBy,
			*clause.Having,
			*clause.OrderBy,
			*clause.Limit,
			*clause.Offset,
			*clause.Union:
			s, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(s.Build())
		default:
			return failure.New(errInvalidClause,
				failure.Context{"clause": reflect.TypeOf(e).Elem().String()},
				failure.Message("invalid clause for SELECT"))
		}
	}

	return nil
}

// RawClause calls the raw string clause.
func (s *SelectStmt) RawClause(rs string, v ...interface{}) iselect.RawClause {
	s.call(&syntax.RawClause{RawStr: rs, Values: v})
	return s
}

// From calls FROM clause.
func (s *SelectStmt) From(tables ...string) iselect.From {
	f := new(clause.From)
	for _, t := range tables {
		f.AddTable(t)
	}
	s.call(f)
	return s
}

// Where calls WHERE clause.
func (s *SelectStmt) Where(expr string, vals ...interface{}) iselect.Where {
	s.call(&clause.Where{Expr: expr, Values: vals})
	return s
}

// And calls AND clause.
func (s *SelectStmt) And(expr string, vals ...interface{}) iselect.And {
	s.call(&clause.And{Expr: expr, Values: vals})
	return s
}

// Or calls OR clause.
func (s *SelectStmt) Or(expr string, vals ...interface{}) iselect.Or {
	s.call(&clause.Or{Expr: expr, Values: vals})
	return s
}

// Limit calls LIMIT clause.
func (s *SelectStmt) Limit(num int) iselect.Limit {
	s.call(&clause.Limit{Num: num})
	return s
}

// Offset calls OFFSET clause.
func (s *SelectStmt) Offset(num int) iselect.Offset {
	s.call(&clause.Offset{Num: num})
	return s
}

// OrderBy calls ORDER BY clause.
func (s *SelectStmt) OrderBy(cols ...string) iselect.OrderBy {
	s.call(&clause.OrderBy{Columns: cols})
	return s
}

// Join calls (INNER) JOIN clause.
func (s *SelectStmt) Join(table string) iselect.Join {
	j := &clause.Join{Type: clause.InnerJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// LeftJoin calls (INNER) JOIN clause.
func (s *SelectStmt) LeftJoin(table string) iselect.Join {
	j := &clause.Join{Type: clause.LeftJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// RightJoin calls (INNER) JOIN clause.
func (s *SelectStmt) RightJoin(table string) iselect.Join {
	j := &clause.Join{Type: clause.RightJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// On calls ON clause.
func (s *SelectStmt) On(expr string, vals ...interface{}) iselect.On {
	s.call(&clause.On{Expr: expr, Values: vals})
	return s
}

// Union calls UNION clause.
func (s *SelectStmt) Union(stmt domain.Stmt) iselect.Union {
	s.call(&clause.Union{Stmt: stmt, All: false})
	return s
}

// UnionAll calls UNION ALL clause.
func (s *SelectStmt) UnionAll(stmt domain.Stmt) iselect.Union {
	s.call(&clause.Union{Stmt: stmt, All: true})
	return s
}

// GroupBy calls GROUP BY clause.
func (s *SelectStmt) GroupBy(cols ...string) iselect.GroupBy {
	g := new(clause.GroupBy)
	for _, c := range cols {
		g.AddColumn(c)
	}
	s.call(g)
	return s
}

// Having calls HAVING clause.
func (s *SelectStmt) Having(expr string, vals ...interface{}) iselect.Having {
	s.call(&clause.Having{Expr: expr, Values: vals})
	return s
}
