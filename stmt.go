package mgorm

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Op values.
const (
	opStmtProcessQuerySQL internal.Op = "mgorm.Stmt.processQuerySQL"
	opStmtProcessCaseSQL  internal.Op = "mgorm.Stmt.processCaseSQL"
	opStmtProcessExecSQL  internal.Op = "mgorm.Stmt.processExecSQL"
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

func (s *Stmt) call(expr syntax.Expr) {
	s.called = append(s.called, expr)
}

func (s *Stmt) execute(op internal.Op, args ...interface{}) {
	s.executed = &opArgs{op: op, args: args}
}

func (s *Stmt) addError(err error) {
	s.errors = append(s.errors, err)
}

// Var returns Stmt.String with syntax.Var type.
func (s *Stmt) Var() syntax.Var {
	s.execute(opVar)
	return syntax.Var(s.String())
}

// String returns query string.
func (s *Stmt) String() string {
	if _, ok := s.called[0].(*syntax.When); ok {
		sql, _ := s.processCaseSQL()
		return sql.string()
	}
	if _, ok := s.cmd.(*syntax.Select); ok {
		sql, _ := s.processQuerySQL()
		return sql.string()
	}
	s.execute(opString)
	sql, _ := s.processExecSQL()
	return sql.string()
}

// Query executes a query that returns some results.
func (s *Stmt) Query(model interface{}) error {
	if db, ok := s.db.(*MockDB); ok {
		s.execute(opQuery, model)
		db.addExecuted(s)
		return nil
	}

	sql, err := s.processQuerySQL()
	if err != nil {
		return err
	}
	if err := sql.doQuery(s.db, model); err != nil {
		return err
	}
	return nil
}

// ExpectQuery executes a query as mock database.
func (s *Stmt) ExpectQuery(model interface{}) *Stmt {
	s.execute(opQuery, model)
	return s
}

func (s *Stmt) processQuerySQL() (SQL, error) {
	var sql SQL

	sel, ok := s.cmd.(*syntax.Select)
	if !ok {
		err := errors.New("command must be SELECT")
		return "", internal.NewError(opStmtProcessQuerySQL, internal.KindRuntime, err)
	}
	sql.write(sel.Build().Build())

	for _, expr := range s.called {
		switch expr := expr.(type) {
		case *syntax.From,
			*syntax.Join,
			*syntax.On,
			*syntax.Where,
			*syntax.And,
			*syntax.Or,
			*syntax.GroupBy,
			*syntax.Having,
			*syntax.OrderBy,
			*syntax.Limit,
			*syntax.Offset,
			*syntax.Union:
			e, err := expr.Build()
			if err != nil {
				return "", err
			}
			sql.write(e.Build())
		default:
			err := fmt.Errorf("%s is not supported", reflect.TypeOf(expr).Elem().String())
			return "", internal.NewError(opStmtProcessQuerySQL, internal.KindRuntime, err)
		}
	}

	return sql, nil
}

func (s *Stmt) processCaseSQL() (SQL, error) {
	var sql SQL
	sql.write("CASE")
	for _, expr := range s.called {
		switch expr := expr.(type) {
		case *syntax.When,
			*syntax.Then,
			*syntax.Else:
			e, err := expr.Build()
			if err != nil {
				return "", err
			}
			sql.write(e.Build())
		default:
			err := fmt.Errorf("%s is not supported", reflect.TypeOf(expr).Elem().String())
			return "", internal.NewError(opStmtProcessCaseSQL, internal.KindRuntime, err)
		}
	}
	sql.write("END")
	return sql, nil
}

// Exec executes a query without returning any results.
func (s *Stmt) Exec() error {
	if db, ok := s.db.(*MockDB); ok {
		s.execute(opExec)
		db.addExecuted(s)
		return nil
	}

	sql, err := s.processExecSQL()
	if err != nil {
		return err
	}
	if err := sql.doExec(s.db); err != nil {
		return err
	}
	return nil
}

// ExpectExec executes a query as mock database.
func (s *Stmt) ExpectExec() *Stmt {
	s.execute(opExec)
	return s
}

// processExecSQL aggregates the values stored in Stmt structure and returns as SQL object.
func (s *Stmt) processExecSQL() (SQL, error) {
	var sql SQL

	switch s.cmd.(type) {
	case *syntax.Insert, *syntax.Update, *syntax.Delete:
		sql.write(s.cmd.Build().Build())
	default:
		err := errors.New("command must be INSERT, UPDATE or DELETE")
		return "", internal.NewError(opStmtProcessExecSQL, internal.KindRuntime, err)

	}

	for _, expr := range s.called {
		switch expr := expr.(type) {
		case *syntax.Values, *syntax.Set, *syntax.From, *syntax.Where, *syntax.And, *syntax.Or:
			e, err := expr.Build()
			if err != nil {
				return "", err
			}
			sql.write(e.Build())
		}
	}

	return sql, nil
}

// From calls FROM statement.
func (s *Stmt) From(tables ...string) FromStmt {
	s.call(syntax.NewFrom(tables))
	return s
}

// Values calls VALUES statement.
func (s *Stmt) Values(vals ...interface{}) ValuesStmt {
	s.call(syntax.NewValues(vals))
	return s
}

// Set calls SET statement.
func (s *Stmt) Set(vals ...interface{}) SetStmt {
	u, ok := s.cmd.(*syntax.Update)
	if !ok {
		err := errors.New("SET statement can be used with UPDATE command")
		s.addError(internal.NewError(opSet, internal.KindRuntime, err))
		return s
	}
	set, err := syntax.NewSet(u.Columns, vals)
	if err != nil {
		s.addError(err)
		return s
	}
	s.call(set)
	return s
}

// Where calls WHERE statement.
func (s *Stmt) Where(expr string, vals ...interface{}) WhereStmt {
	s.call(syntax.NewWhere(expr, vals...))
	return s
}

// And calls AND statement.
func (s *Stmt) And(expr string, vals ...interface{}) AndStmt {
	s.call(syntax.NewAnd(expr, vals...))
	return s
}

// Or calls OR statement.
func (s *Stmt) Or(expr string, vals ...interface{}) OrStmt {
	s.call(syntax.NewOr(expr, vals...))
	return s
}

// Limit calls LIMIT statement.
func (s *Stmt) Limit(num int) LimitStmt {
	s.call(syntax.NewLimit(num))
	return s
}

// Offset calls OFFSET statement.
func (s *Stmt) Offset(num int) OffsetStmt {
	s.call(syntax.NewOffset(num))
	return s
}

// OrderBy calls ORDER BY statement.
func (s *Stmt) OrderBy(col string, desc bool) OrderByStmt {
	s.call(syntax.NewOrderBy(col, desc))
	return s
}

// Join calls (INNER) JOIN statement.
func (s *Stmt) Join(table string) JoinStmt {
	s.call(syntax.NewJoin(table, syntax.InnerJoin))
	return s
}

// LeftJoin calls (INNER) JOIN statement.
func (s *Stmt) LeftJoin(table string) JoinStmt {
	s.call(syntax.NewJoin(table, syntax.LeftJoin))
	return s
}

// RightJoin calls (INNER) JOIN statement.
func (s *Stmt) RightJoin(table string) JoinStmt {
	s.call(syntax.NewJoin(table, syntax.RightJoin))
	return s
}

// FullJoin calls (INNER) JOIN statement.
func (s *Stmt) FullJoin(table string) JoinStmt {
	s.call(syntax.NewJoin(table, syntax.FullJoin))
	return s
}

// On calls ON statement.
func (s *Stmt) On(expr string, vals ...interface{}) OnStmt {
	s.call(syntax.NewOn(expr, vals...))
	return s
}

// Union calls UNION statement.
func (s *Stmt) Union(stmt syntax.Var) UnionStmt {
	s.call(syntax.NewUnion(stmt, false))
	return s
}

// UnionAll calls UNION ALL statement.
func (s *Stmt) UnionAll(stmt syntax.Var) UnionStmt {
	s.call(syntax.NewUnion(stmt, true))
	return s
}

// GroupBy calls GROUP BY statement.
func (s *Stmt) GroupBy(cols ...string) GroupByStmt {
	s.call(syntax.NewGroupBy(cols))
	return s
}

// Having calls HAVING statement.
func (s *Stmt) Having(expr string, vals ...interface{}) HavingStmt {
	s.call(syntax.NewHaving(expr, vals...))
	return s
}

// When calls WHEN statement.
func (s *Stmt) When(expr string, vals ...interface{}) WhenStmt {
	s.call(syntax.NewWhen(expr, vals...))
	return s
}

// Then calls THEN statement.
func (s *Stmt) Then(val interface{}) ThenStmt {
	s.call(syntax.NewThen(val))
	return s
}

// Else calls ELSE statement.
func (s *Stmt) Else(val interface{}) ElseStmt {
	s.call(syntax.NewElse(val))
	return s
}
