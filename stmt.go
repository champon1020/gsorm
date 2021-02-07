package mgorm

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/cmd"
	"github.com/champon1020/mgorm/syntax/expr"
)

// Op values.
const (
	opStmtProcessQuerySQL internal.Op = "mgorm.Stmt.processQuerySQL"
	opStmtProcessCaseSQL  internal.Op = "mgorm.Stmt.processCaseSQL"
	opStmtProcessExecSQL  internal.Op = "mgorm.Stmt.processExecSQL"
	opColumn              internal.Op = "mgorm.Stmt.Column"
	opVar                 internal.Op = "mgorm.Stmt.Var"
	opString              internal.Op = "mgorm.Stmt.String"
	opQuery               internal.Op = "mgorm.Stmt.Query"
	opExec                internal.Op = "mgorm.Stmt.Exec"
	opSet                 internal.Op = "mgorm.Stmt.Set"
)

// Stmt keeps the sql statement.
type Stmt struct {
	db       sqlDB
	cmd      syntax.Cmd
	called   []syntax.Expr
	executed *opArgs
	errors   []error
}

func (s *Stmt) call(e syntax.Expr) {
	s.called = append(s.called, e)
}

func (s *Stmt) execute(op internal.Op, args ...interface{}) {
	s.executed = &opArgs{op: op, args: args}
}

func (s *Stmt) addError(err error) {
	s.errors = append(s.errors, err)
}

// Column returns string with column format.
func (s *Stmt) Column() string {
	if _, ok := s.db.(*MockDB); ok {
		s.execute(opColumn)
	}
	sql, err := s.processCaseSQL(true)
	if err != nil {
		s.addError(err)
		return ""
	}
	return sql.String()
}

// Var returns Stmt.String with syntax.Var type.
func (s *Stmt) Var() syntax.Var {
	if _, ok := s.db.(*MockDB); ok {
		s.execute(opVar)
	}
	if _, ok := s.called[0].(*expr.When); ok {
		sql, err := s.processCaseSQL(false)
		if err != nil {
			s.addError(err)
			return ""
		}
		return syntax.Var(sql.String())
	}
	return syntax.Var(s.String())
}

// String returns query string.
func (s *Stmt) String() string {
	if _, ok := s.db.(*MockDB); ok {
		s.execute(opString)
	}
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

// Query executes a query that returns some results.
func (s *Stmt) Query(model interface{}) error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch db := s.db.(type) {
	case *DB:
		sql, err := s.processQuerySQL()
		if err != nil {
			return err
		}
		if err := internal.Query(db.db, &sql, model); err != nil {
			return err
		}
	case *MockDB:
		s.execute(opQuery, model)
		db.addExecuted(s)
	}
	return nil
}

// Exec executes a query without returning any results.
func (s *Stmt) Exec() error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch db := s.db.(type) {
	case *DB:
		sql, err := s.processExecSQL()
		if err != nil {
			return err
		}
		if err := internal.Exec(db.db, &sql); err != nil {
			return err
		}
	case *MockDB:
		s.execute(opExec)
		db.addExecuted(s)
	}
	return nil
}

// ExpectQuery executes a query as mock database.
func (s *Stmt) ExpectQuery(model interface{}) *Stmt {
	s.execute(opQuery, model)
	return s
}

// ExpectExec executes a query as mock database.
func (s *Stmt) ExpectExec() *Stmt {
	s.execute(opExec)
	return s
}

// processQuerySQL builds SQL with called expressions.
func (s *Stmt) processQuerySQL() (internal.SQL, error) {
	var sql internal.SQL

	sel, ok := s.cmd.(*cmd.Select)
	if !ok {
		err := errors.New("command must be SELECT")
		return "", internal.NewError(opStmtProcessQuerySQL, internal.KindRuntime, err)
	}
	sql.Write(sel.Build().Build())

	for _, e := range s.called {
		switch e := e.(type) {
		case *expr.From,
			*expr.Join,
			*expr.On,
			*expr.Where,
			*expr.And,
			*expr.Or,
			*expr.GroupBy,
			*expr.Having,
			*expr.OrderBy,
			*expr.Limit,
			*expr.Offset,
			*expr.Union:
			s, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.Write(s.Build())
		default:
			err := fmt.Errorf("%s is not supported", reflect.TypeOf(e).Elem().String())
			return "", internal.NewError(opStmtProcessQuerySQL, internal.KindRuntime, err)
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
		case *expr.When:
			s, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.Write(s.Build())
		case *expr.Then:
			e.IsColumn = isColumn
			s, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.Write(s.Build())
		case *expr.Else:
			e.IsColumn = isColumn
			s, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.Write(s.Build())
		default:
			err := fmt.Errorf("%s is not supported", reflect.TypeOf(e).Elem().String())
			return "", internal.NewError(opStmtProcessCaseSQL, internal.KindRuntime, err)
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
		err := errors.New("command must be INSERT, UPDATE or DELETE")
		return "", internal.NewError(opStmtProcessExecSQL, internal.KindRuntime, err)

	}

	for _, e := range s.called {
		switch e := e.(type) {
		case *expr.Values, *expr.Set, *expr.From, *expr.Where, *expr.And, *expr.Or:
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
	s.call(expr.NewFrom(tables))
	return s
}

// Values calls VALUES statement.
func (s *Stmt) Values(vals ...interface{}) ValuesStmt {
	s.call(expr.NewValues(vals))
	return s
}

// Set calls SET statement.
func (s *Stmt) Set(vals ...interface{}) SetStmt {
	u, ok := s.cmd.(*cmd.Update)
	if !ok {
		err := errors.New("SET statement can be used with UPDATE command")
		s.addError(internal.NewError(opSet, internal.KindRuntime, err))
		return s
	}
	set, err := expr.NewSet(u.Columns, vals)
	if err != nil {
		s.addError(err)
		return s
	}
	s.call(set)
	return s
}

// Where calls WHERE statement.
func (s *Stmt) Where(e string, vals ...interface{}) WhereStmt {
	s.call(expr.NewWhere(e, vals...))
	return s
}

// And calls AND statement.
func (s *Stmt) And(e string, vals ...interface{}) AndStmt {
	s.call(expr.NewAnd(e, vals...))
	return s
}

// Or calls OR statement.
func (s *Stmt) Or(e string, vals ...interface{}) OrStmt {
	s.call(expr.NewOr(e, vals...))
	return s
}

// Limit calls LIMIT statement.
func (s *Stmt) Limit(num int) LimitStmt {
	s.call(expr.NewLimit(num))
	return s
}

// Offset calls OFFSET statement.
func (s *Stmt) Offset(num int) OffsetStmt {
	s.call(expr.NewOffset(num))
	return s
}

// OrderBy calls ORDER BY statement.
func (s *Stmt) OrderBy(col string) OrderByStmt {
	s.call(expr.NewOrderBy(col, false))
	return s
}

// OrderByDesc calls ORDER BY ... DESC statement.
func (s *Stmt) OrderByDesc(col string) OrderByStmt {
	s.call(expr.NewOrderBy(col, true))
	return s
}

// Join calls (INNER) JOIN statement.
func (s *Stmt) Join(table string) JoinStmt {
	s.call(expr.NewJoin(table, expr.InnerJoin))
	return s
}

// LeftJoin calls (INNER) JOIN statement.
func (s *Stmt) LeftJoin(table string) JoinStmt {
	s.call(expr.NewJoin(table, expr.LeftJoin))
	return s
}

// RightJoin calls (INNER) JOIN statement.
func (s *Stmt) RightJoin(table string) JoinStmt {
	s.call(expr.NewJoin(table, expr.RightJoin))
	return s
}

// FullJoin calls (INNER) JOIN statement.
func (s *Stmt) FullJoin(table string) JoinStmt {
	s.call(expr.NewJoin(table, expr.FullJoin))
	return s
}

// On calls ON statement.
func (s *Stmt) On(e string, vals ...interface{}) OnStmt {
	s.call(expr.NewOn(e, vals...))
	return s
}

// Union calls UNION statement.
func (s *Stmt) Union(stmt syntax.Var) UnionStmt {
	s.call(expr.NewUnion(stmt, false))
	return s
}

// UnionAll calls UNION ALL statement.
func (s *Stmt) UnionAll(stmt syntax.Var) UnionStmt {
	s.call(expr.NewUnion(stmt, true))
	return s
}

// GroupBy calls GROUP BY statement.
func (s *Stmt) GroupBy(cols ...string) GroupByStmt {
	s.call(expr.NewGroupBy(cols))
	return s
}

// Having calls HAVING statement.
func (s *Stmt) Having(e string, vals ...interface{}) HavingStmt {
	s.call(expr.NewHaving(e, vals...))
	return s
}

// When calls WHEN statement.
func (s *Stmt) When(e string, vals ...interface{}) WhenStmt {
	s.call(expr.NewWhen(e, vals...))
	return s
}

// Then calls THEN statement.
func (s *Stmt) Then(val interface{}) ThenStmt {
	s.call(expr.NewThen(val))
	return s
}

// Else calls ELSE statement.
func (s *Stmt) Else(val interface{}) ElseStmt {
	s.call(expr.NewElse(val))
	return s
}
