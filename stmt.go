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

// Stmt keeps the sql statement.
type Stmt struct {
	db     Pool
	cmd    syntax.Cmd
	called []syntax.Clause
	errors []error
}

func (s *Stmt) call(e syntax.Clause) {
	s.called = append(s.called, e)
}

func (s *Stmt) addError(err error) {
	s.errors = append(s.errors, err)
}

// CaseColumn returns string without double quotes.
// This is used for CASE WHEN ... statement
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

// CaseValue returns string with double quotes.
// This is used for CASE WHEN ... statement
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

// Sub returns Stmt.String with syntax.Sub type.
// This is used for UNION or WHERE with SELECT clause.
func (s *Stmt) Sub() syntax.Sub {
	return syntax.Sub(s.String())
}

// String returns query string.
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

// stmtFuncString returns called function like following example.
//   exmaple: SELECT(...).FROM(...).WHERE(...).QUERY(...).
func (s *Stmt) funcString() string {
	str := s.cmd.String()
	for _, e := range s.called {
		str += fmt.Sprintf(".%s", e.String())
	}
	return str
}

// Query executes a query that returns some results.
func (s *Stmt) Query(model interface{}) error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch pool := s.db.(type) {
	case *DB:
		sql, err := s.processQuerySQL()
		if err != nil {
			return err
		}

		rows, err := pool.db.Query(sql.String())
		if err != nil {
			return errors.New(err.Error(), errors.DBQueryError)
		}

		defer rows.Close()
		if err := internal.MapRowsToModel(rows, model); err != nil {
			return err
		}
	case *Tx:
		sql, err := s.processQuerySQL()
		if err != nil {
			return err
		}

		rows, err := pool.tx.Query(sql.String())
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

// Exec executes a query without returning any results.
func (s *Stmt) Exec() error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch pool := s.db.(type) {
	case *DB:
		sql, err := s.processExecSQL()
		if err != nil {
			return err
		}
		if _, err := pool.db.Exec(sql.String()); err != nil {
			return errors.New(err.Error(), errors.DBQueryError)
		}
	case *Tx:
		sql, err := s.processExecSQL()
		if err != nil {
			return err
		}
		if _, err := pool.tx.Exec(sql.String()); err != nil {
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

// processQuerySQL builds SQL with called expressions.
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

// processCaseSQL builds SQL with called expressions.
// isColumn flag indicates whether this is called from Stmt.Column() or not.
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

// processQuerySQL builds SQL with called expressions.
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

// From calls FROM statement.
func (s *Stmt) From(tables ...string) FromStmt {
	s.call(clause.NewFrom(tables))
	return s
}

// Values calls VALUES statement.
func (s *Stmt) Values(vals ...interface{}) ValuesStmt {
	s.call(clause.NewValues(vals))
	return s
}

// Set calls SET statement.
func (s *Stmt) Set(vals ...interface{}) SetStmt {
	u, ok := s.cmd.(*cmd.Update)
	if !ok {
		s.addError(errors.New("SET statement can be used with UPDATE command", errors.InvalidValueError))
		return s
	}
	set, err := clause.NewSet(u.Columns, vals)
	if err != nil {
		s.addError(err)
		return s
	}
	s.call(set)
	return s
}

// Where calls WHERE statement.
func (s *Stmt) Where(e string, vals ...interface{}) WhereStmt {
	s.call(clause.NewWhere(e, vals...))
	return s
}

// And calls AND statement.
func (s *Stmt) And(e string, vals ...interface{}) AndStmt {
	s.call(clause.NewAnd(e, vals...))
	return s
}

// Or calls OR statement.
func (s *Stmt) Or(e string, vals ...interface{}) OrStmt {
	s.call(clause.NewOr(e, vals...))
	return s
}

// Limit calls LIMIT statement.
func (s *Stmt) Limit(num int) LimitStmt {
	s.call(clause.NewLimit(num))
	return s
}

// Offset calls OFFSET statement.
func (s *Stmt) Offset(num int) OffsetStmt {
	s.call(clause.NewOffset(num))
	return s
}

// OrderBy calls ORDER BY statement.
func (s *Stmt) OrderBy(col string) OrderByStmt {
	s.call(clause.NewOrderBy(col, false))
	return s
}

// OrderByDesc calls ORDER BY ... DESC statement.
func (s *Stmt) OrderByDesc(col string) OrderByStmt {
	s.call(clause.NewOrderBy(col, true))
	return s
}

// Join calls (INNER) JOIN statement.
func (s *Stmt) Join(table string) JoinStmt {
	s.call(clause.NewJoin(table, clause.InnerJoin))
	return s
}

// LeftJoin calls (INNER) JOIN statement.
func (s *Stmt) LeftJoin(table string) JoinStmt {
	s.call(clause.NewJoin(table, clause.LeftJoin))
	return s
}

// RightJoin calls (INNER) JOIN statement.
func (s *Stmt) RightJoin(table string) JoinStmt {
	s.call(clause.NewJoin(table, clause.RightJoin))
	return s
}

// FullJoin calls (INNER) JOIN statement.
func (s *Stmt) FullJoin(table string) JoinStmt {
	s.call(clause.NewJoin(table, clause.FullJoin))
	return s
}

// On calls ON statement.
func (s *Stmt) On(e string, vals ...interface{}) OnStmt {
	s.call(clause.NewOn(e, vals...))
	return s
}

// Union calls UNION statement.
func (s *Stmt) Union(stmt syntax.Sub) UnionStmt {
	s.call(clause.NewUnion(stmt, false))
	return s
}

// UnionAll calls UNION ALL statement.
func (s *Stmt) UnionAll(stmt syntax.Sub) UnionStmt {
	s.call(clause.NewUnion(stmt, true))
	return s
}

// GroupBy calls GROUP BY statement.
func (s *Stmt) GroupBy(cols ...string) GroupByStmt {
	s.call(clause.NewGroupBy(cols))
	return s
}

// Having calls HAVING statement.
func (s *Stmt) Having(e string, vals ...interface{}) HavingStmt {
	s.call(clause.NewHaving(e, vals...))
	return s
}

// When calls WHEN statement.
func (s *Stmt) When(e string, vals ...interface{}) WhenStmt {
	s.call(clause.NewWhen(e, vals...))
	return s
}

// Then calls THEN statement.
func (s *Stmt) Then(val interface{}) ThenStmt {
	s.call(clause.NewThen(val))
	return s
}

// Else calls ELSE statement.
func (s *Stmt) Else(val interface{}) ElseStmt {
	s.call(clause.NewElse(val))
	return s
}
