package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/champon1020/mgorm/syntax/cmd"
)

// Stmt stores information about query.
type Stmt struct {
	db     Pool
	cmd    syntax.Cmd
	called []syntax.Clause
	errors []error
}

// call appends called clause.
func (s *Stmt) call(e syntax.Clause) {
	s.called = append(s.called, e)
}

// addError appends occurred error
func (s *Stmt) addError(err error) {
	s.errors = append(s.errors, err)
}

// CaseColumn returns string without double quotes. This is used for CASE WHEN ... statement.
func (s *Stmt) CaseColumn() string {
	if _, ok := s.called[0].(*clause.When); ok {
		sql, err := s.processCaseSQL(true)
		if err != nil {
			s.addError(err)
			return ""
		}
		return sql.String()
	}
	return s.String()
}

// CaseValue returns string with double quotes. This is used for CASE WHEN ... statement.
func (s *Stmt) CaseValue() string {
	if _, ok := s.called[0].(*clause.When); ok {
		sql, err := s.processCaseSQL(false)
		if err != nil {
			s.addError(err)
			return ""
		}
		return sql.String()
	}
	return s.String()
}

// Sub returns Stmt.String with syntax.Sub type. This is used for UNION or WHERE with SELECT clause.
func (s *Stmt) Sub() syntax.Sub {
	return syntax.Sub(s.String())
}

// String returns query with string.
func (s *Stmt) String() string {
	if _, ok := s.cmd.(*cmd.Select); ok {
		sql, err := s.processQuerySQL()
		if err != nil {
			s.addError(err)
			return ""
		}
		return sql.String()
	}
	sql, err := s.processExecSQL()
	if err != nil {
		s.addError(err)
		return ""
	}
	return sql.String()
}

// stmtFuncString returns called function like "SELECT(...).FROM(...).WHERE(...).QUERY(...)".
func (s *Stmt) funcString() string {
	str := s.cmd.String()
	for _, e := range s.called {
		str += fmt.Sprintf(".%s", e.String())
	}
	return str
}

// Query executes a query that maps values to model.
func (s *Stmt) Query(model interface{}) error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch pool := s.db.(type) {
	case *DB, *Tx:
		sql, err := s.processQuerySQL()
		if err != nil {
			return err
		}

		rows, err := pool.Query(sql.String())
		if err != nil {
			return errors.New(err.Error(), errors.DBQueryError)
		}

		defer rows.Close()
		if err := internal.MapRowsToModel(rows, model); err != nil {
			return err
		}
	case Mock:
		returned, err := compareTo(pool, s)
		if err != nil {
			return err
		}
		if returned != nil {
			/* process */
		}
	default:
		return errors.New("DB type must be *DB, *Tx, *MockDB or *MockTx", errors.InvalidValueError)
	}

	return nil
}

// Exec executes a query without without mapping.
func (s *Stmt) Exec() error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch pool := s.db.(type) {
	case *DB, *Tx:
		sql, err := s.processExecSQL()
		if err != nil {
			return err
		}
		if _, err := pool.Exec(sql.String()); err != nil {
			return errors.New(err.Error(), errors.DBQueryError)
		}
	case Mock:
		_, err := compareTo(pool, s)
		if err != nil {
			return err
		}
	default:
		return errors.New("DB type must be *DB, *Tx, *MockDB or *MockTx", errors.InvalidValueError)
	}

	return nil
}

// processQuerySQL builds SQL with called clauses.
func (s *Stmt) processQuerySQL() (internal.SQL, error) {
	var sql internal.SQL

	sel, ok := s.cmd.(*cmd.Select)
	if !ok {
		return "", errors.New("Command must be SELECT", errors.InvalidValueError)
	}
	sql.Write(sel.Build().Build())

	for _, e := range s.called {
		switch e := e.(type) {
		case *clause.From,
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
				return "", err
			}
			sql.Write(s.Build())
		default:
			msg := fmt.Sprintf("Type %s is not supported", reflect.TypeOf(e).Elem().String())
			return "", errors.New(msg, errors.InvalidTypeError)
		}
	}

	return sql, nil
}

// processCaseSQL builds SQL with called clauses.
// isColumn flag indicates whether this is called from (*Stmt).CaseColumn() or not.
func (s *Stmt) processCaseSQL(isColumn bool) (internal.SQL, error) {
	var sql internal.SQL
	sql.Write("CASE")
	for _, e := range s.called {
		switch e := e.(type) {
		case *clause.When:
			s, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.Write(s.Build())
		case *clause.Then:
			e.IsColumn = isColumn
			s, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.Write(s.Build())
		case *clause.Else:
			e.IsColumn = isColumn
			s, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.Write(s.Build())
		default:
			msg := fmt.Sprintf("Type %s is not supported", reflect.TypeOf(e).Elem().String())
			return "", errors.New(msg, errors.InvalidTypeError)
		}
	}
	sql.Write("END")
	return sql, nil
}

// processQuerySQL builds SQL with called clauses.
func (s *Stmt) processExecSQL() (internal.SQL, error) {
	var sql internal.SQL

	switch s.cmd.(type) {
	case *cmd.Insert, *cmd.Update, *cmd.Delete:
		sql.Write(s.cmd.Build().Build())
	default:
		return "", errors.New("Command must be INSERT, UPDATE or DELETE", errors.InvalidValueError)

	}

	for _, e := range s.called {
		switch e := e.(type) {
		case *clause.Values, *clause.Set, *clause.From, *clause.Where, *clause.And, *clause.Or:
			s, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.Write(s.Build())
		}
	}

	return sql, nil
}

// From calls FROM clause.
func (s *Stmt) From(tables ...string) FromStmt {
	f := new(clause.From)
	for _, t := range tables {
		f.AddTable(t)
	}
	s.call(f)
	return s
}

// Values calls VALUES clause.
func (s *Stmt) Values(vals ...interface{}) ValuesStmt {
	v := new(clause.Values)
	for _, val := range vals {
		v.AddValue(val)
	}
	s.call(v)
	return s
}

// Set calls SET clause.
func (s *Stmt) Set(vals ...interface{}) SetStmt {
	if s.cmd == nil {
		s.addError(errors.New("Command is nil", errors.InvalidValueError))
		return s
	}
	u, ok := s.cmd.(*cmd.Update)
	if !ok {
		s.addError(errors.New("SET clause can be used with UPDATE command", errors.InvalidValueError))
		return s
	}
	if len(u.Columns) != len(vals) {
		s.addError(errors.New("Length is different between lhs and rhs", errors.InvalidValueError))
		return s
	}
	set := new(clause.Set)
	for i := 0; i < len(u.Columns); i++ {
		set.AddEq(u.Columns[i], vals[i])
	}
	s.call(set)
	return s
}

// Where calls WHERE clause.
func (s *Stmt) Where(expr string, vals ...interface{}) WhereStmt {
	s.call(&clause.Where{Expr: expr, Values: vals})
	return s
}

// And calls AND clause.
func (s *Stmt) And(expr string, vals ...interface{}) AndStmt {
	s.call(&clause.And{Expr: expr, Values: vals})
	return s
}

// Or calls OR clause.
func (s *Stmt) Or(expr string, vals ...interface{}) OrStmt {
	s.call(&clause.Or{Expr: expr, Values: vals})
	return s
}

// Limit calls LIMIT clause.
func (s *Stmt) Limit(num int) LimitStmt {
	s.call(&clause.Limit{Num: num})
	return s
}

// Offset calls OFFSET clause.
func (s *Stmt) Offset(num int) OffsetStmt {
	s.call(&clause.Offset{Num: num})
	return s
}

// OrderBy calls ORDER BY clause.
func (s *Stmt) OrderBy(col string) OrderByStmt {
	s.call(&clause.OrderBy{Column: col, Desc: false})
	return s
}

// OrderByDesc calls ORDER BY ... DESC clause.
func (s *Stmt) OrderByDesc(col string) OrderByStmt {
	s.call(&clause.OrderBy{Column: col, Desc: true})
	return s
}

// Join calls (INNER) JOIN clause.
func (s *Stmt) Join(table string) JoinStmt {
	j := &clause.Join{Type: clause.InnerJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// LeftJoin calls (INNER) JOIN clause.
func (s *Stmt) LeftJoin(table string) JoinStmt {
	j := &clause.Join{Type: clause.LeftJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// RightJoin calls (INNER) JOIN clause.
func (s *Stmt) RightJoin(table string) JoinStmt {
	j := &clause.Join{Type: clause.RightJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// FullJoin calls (INNER) JOIN clause.
func (s *Stmt) FullJoin(table string) JoinStmt {
	j := &clause.Join{Type: clause.FullJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// On calls ON clause.
func (s *Stmt) On(expr string, vals ...interface{}) OnStmt {
	s.call(&clause.On{Expr: expr, Values: vals})
	return s
}

// Union calls UNION clause.
func (s *Stmt) Union(stmt syntax.Sub) UnionStmt {
	s.call(&clause.Union{Stmt: stmt, All: false})
	return s
}

// UnionAll calls UNION ALL clause.
func (s *Stmt) UnionAll(stmt syntax.Sub) UnionStmt {
	s.call(&clause.Union{Stmt: stmt, All: true})
	return s
}

// GroupBy calls GROUP BY clause.
func (s *Stmt) GroupBy(cols ...string) GroupByStmt {
	g := new(clause.GroupBy)
	for _, c := range cols {
		g.AddColumn(c)
	}
	s.call(g)
	return s
}

// Having calls HAVING clause.
func (s *Stmt) Having(expr string, vals ...interface{}) HavingStmt {
	s.call(&clause.Having{Expr: expr, Values: vals})
	return s
}

// When calls WHEN clause.
func (s *Stmt) When(expr string, vals ...interface{}) WhenStmt {
	s.call(&clause.When{Expr: expr, Values: vals})
	return s
}

// Then calls THEN clause.
func (s *Stmt) Then(val interface{}) ThenStmt {
	s.call(&clause.Then{Value: val})
	return s
}

// Else calls ELSE clause.
func (s *Stmt) Else(val interface{}) ElseStmt {
	s.call(&clause.Else{Value: val})
	return s
}
