package mgorm

import (
	"errors"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
)

// Op values.
const (
	opStmtProcessQuerySQL internal.Op = "mgorm.Stmt.processQuerySQL"
	opStmtProcessExecSQL  internal.Op = "mgorm.Stmt.processExecSQL"
	opQuery               internal.Op = "mgorm.Stmt.Query"
	opExec                internal.Op = "mgorm.Stmt.Exec"
	opFrom                internal.Op = "mgorm.Stmt.From"
	opValues              internal.Op = "mgorm.Stmt.Values"
	opSet                 internal.Op = "mgorm.Stmt.Set"
	opWhere               internal.Op = "mgorm.Stmt.Where"
	opAnd                 internal.Op = "mgorm.Stmt.And"
	opOr                  internal.Op = "mgorm.Stmt.Or"
	opNot                 internal.Op = "mgorm.Stmt.Not"
	opLimit               internal.Op = "mgorm.Stmt.Limit"
	opOffset              internal.Op = "mgorm.Stmt.Offset"
	opOrderBy             internal.Op = "mgorm.Stmt.OrderBy"
	opJoin                internal.Op = "mgorm.Stmt.Join"
	opLeftJoin            internal.Op = "mgorm.Stmt.LeftJoin"
	opRightJoin           internal.Op = "mgorm.Stmt.RightJoin"
	opFullJoin            internal.Op = "mgorm.Stmt.FullJoin"
	opOn                  internal.Op = "mgorm.Stmt.On"
	opUnion               internal.Op = "mgorm.Stmt.Union"
	opUnionAll            internal.Op = "mgorm.Stmt.UnionAll"
	opGroupBy             internal.Op = "mgorm.Stmt.GroupBy"
	opHaving              internal.Op = "mgorm.Stmt.Having"
	opWhen                internal.Op = "mgorm.Stmt.When"
	opThen                internal.Op = "mgorm.Stmt.Then"
	opElse                internal.Op = "mgorm.Stmt.Else"
)

// Stmt keeps the sql statement.
type Stmt struct {
	db          sqlDB
	cmd         syntax.Cmd
	fromExpr    syntax.Expr
	valuesExpr  syntax.Expr
	setExpr     syntax.Expr
	whereExpr   syntax.Expr
	andOrNot    []syntax.Expr
	limitExpr   syntax.Expr
	offsetExpr  syntax.Expr
	orderByExpr []syntax.Expr
	joinExpr    []syntax.Expr
	onExpr      []syntax.Expr
	unionExpr   syntax.Expr
	groupByExpr syntax.Expr
	havingExpr  syntax.Expr
	whenExpr    []syntax.Expr
	thenExpr    []syntax.Expr
	elseExpr    syntax.Expr
	errors      []error

	// Used for test.
	called []*opArgs
}

func (s *Stmt) call(op internal.Op, args ...interface{}) {
	s.called = append(s.called, &opArgs{op: op, args: args})
}

func (s *Stmt) addError(err error) {
	s.errors = append(s.errors, err)
}

// Var returns Stmt.String with syntax.Var type.
func (s *Stmt) Var() syntax.Var {
	return syntax.Var(s.String())
}

// String returns query string.
func (s *Stmt) String() string {
	_, ok := s.cmd.(*syntax.Select)
	if ok {
		sql, _ := s.processQuerySQL()
		return sql.string()
	}
	sql, _ := s.processExecSQL()
	return sql.string()
}

// Query executes a query that returns some results.
func (s *Stmt) Query(model interface{}) error {
	if db, ok := s.db.(*MockDB); ok {
		s.call(opQuery, model)
		db.addExecuted(s.called)
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
	s.call(opQuery, model)
	return s
}

// processQuerySQL aggregates the values stored in Stmt structure and returns as SQL object.
func (s *Stmt) processQuerySQL() (SQL, error) {
	var sql SQL

	// Build SELECT.
	sel, ok := s.cmd.(*syntax.Select)
	if !ok {
		err := errors.New("command must be SELECT")
		return "", internal.NewError(opStmtProcessQuerySQL, internal.KindBasic, err)
	}
	sql.write(sel.Build().Build())

	// Build FROM.
	if s.fromExpr != nil {
		from, err := s.fromExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(from.Build())
	}

	// Build JOIN and ON.
	if len(s.joinExpr) > 0 {
		for i, e := range s.joinExpr {
			// If onExpr is not sufficient, return error.
			if len(s.onExpr) <= i {
				err := errors.New("JOIN was executed but ON is not called")
				return "", internal.NewError(opStmtProcessQuerySQL, internal.KindRuntime, err)
			}

			// Build JOIN.
			j, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.write(j.Build())

			// Build ON.
			o, err := s.onExpr[i].Build()
			if err != nil {
				return "", err
			}
			sql.write(o.Build())
		}
	}

	// Build WHERE.
	if s.whereExpr != nil {
		w, err := s.whereExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(w.Build())
	}

	// Build AND, OR or NOT.
	if len(s.andOrNot) > 0 {
		for _, e := range s.andOrNot {
			ao, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.write(ao.Build())
		}
	}

	// Build GROUP BY.
	if s.groupByExpr != nil {
		g, err := s.groupByExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(g.Build())
	}

	// Build HAVING.
	if s.havingExpr != nil {
		h, err := s.havingExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(h.Build())
	}

	// Build ORDER BY.
	if len(s.orderByExpr) > 0 {
		for _, e := range s.orderByExpr {
			ob, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.write(ob.Build())
		}
	}

	// Build CASE ... END.
	if len(s.whenExpr) > 0 {
		sql.write("CASE")
		for i, e := range s.whenExpr {
			// If thenExpr is not sufficient, return error.
			if len(s.thenExpr) <= i {
				err := errors.New("WHEN was executed but THEN is not called")
				return "", internal.NewError(opStmtProcessQuerySQL, internal.KindRuntime, err)
			}

			// Build WHEN.
			w, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.write(w.Build())

			// Build THEN.
			t, err := s.thenExpr[i].Build()
			if err != nil {
				return "", err
			}
			sql.write(t.Build())
		}
		if s.elseExpr != nil {
			e, err := s.elseExpr.Build()
			if err != nil {
				return "", err
			}
			sql.write(e.Build())
		}
		sql.write("END")
	}

	// Build LIMIT.
	if s.limitExpr != nil {
		l, err := s.limitExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(l.Build())
	}

	// Build OFFSET.
	if s.offsetExpr != nil {
		l, err := s.offsetExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(l.Build())
	}

	// Build UNION.
	if s.unionExpr != nil {
		u, err := s.unionExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(u.Build())
	}

	return sql, nil
}

// Exec executes a query without returning any results.
func (s *Stmt) Exec() error {
	if db, ok := s.db.(*MockDB); ok {
		s.call(opQuery)
		db.addExecuted(s.called)
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
	s.call(opExec)
	return s
}

// processExecSQL aggregates the values stored in Stmt structure and returns as SQL object.
func (s *Stmt) processExecSQL() (SQL, error) {
	var sql SQL
	switch cmd := s.cmd.(type) {
	case *syntax.Insert:
		sql.write(cmd.Build().Build())
		if s.valuesExpr != nil {
			values, err := s.valuesExpr.Build()
			if err != nil {
				return "", err
			}
			sql.write(values.Build())
		}
	case *syntax.Update:
		sql.write(cmd.Build().Build())
		if s.setExpr != nil {
			set, err := s.setExpr.Build()
			if err != nil {
				return "", err
			}
			sql.write(set.Build())
		}
	case *syntax.Delete:
		sql.write(cmd.Build().Build())
		if s.fromExpr != nil {
			from, err := s.fromExpr.Build()
			if err != nil {
				return "", err
			}
			sql.write(from.Build())
		}
	default:
		err := errors.New("command must be INSERT, UPDATE or DELETE")
		return "", internal.NewError(opStmtProcessExecSQL, internal.KindType, err)

	}

	// Build WHERE.
	if s.whereExpr != nil {
		w, err := s.whereExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(w.Build())
	}

	// Build AND, OR or NOT.
	if len(s.andOrNot) > 0 {
		for _, e := range s.andOrNot {
			ao, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.write(ao.Build())
		}
	}
	return sql, nil
}

// From calls FROM statement.
func (s *Stmt) From(tables ...string) *Stmt {
	s.fromExpr = syntax.NewFrom(tables)
	s.call(opFrom, tables)
	return s
}

// Values calls VALUES statement.
func (s *Stmt) Values(vals ...interface{}) *Stmt {
	s.valuesExpr = syntax.NewValues(vals)
	s.call(opValues, vals)
	return s
}

// Set calls SET statement.
func (s *Stmt) Set(vals ...interface{}) *Stmt {
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
	s.setExpr = set
	s.call(opSet, vals)
	return s
}

// Where calls WHERE statement.
func (s *Stmt) Where(expr string, vals ...interface{}) *Stmt {
	s.whereExpr = syntax.NewWhere(expr, vals...)
	s.call(opWhere, expr, vals)
	return s
}

// And calls AND statement.
func (s *Stmt) And(expr string, vals ...interface{}) *Stmt {
	s.andOrNot = append(s.andOrNot, syntax.NewAnd(expr, vals...))
	s.call(opAnd, expr, vals)
	return s
}

// Or calls OR statement.
func (s *Stmt) Or(expr string, vals ...interface{}) *Stmt {
	s.andOrNot = append(s.andOrNot, syntax.NewOr(expr, vals...))
	s.call(opOr, expr, vals)
	return s
}

// Limit calls LIMIT statement.
func (s *Stmt) Limit(num int) *Stmt {
	s.limitExpr = syntax.NewLimit(num)
	s.call(opLimit, num)
	return s
}

// Offset calls OFFSET statement.
func (s *Stmt) Offset(num int) *Stmt {
	s.offsetExpr = syntax.NewOffset(num)
	s.call(opOffset, num)
	return s
}

// OrderBy calls ORDER BY statement.
func (s *Stmt) OrderBy(col string, desc bool) *Stmt {
	s.orderByExpr = append(s.orderByExpr, syntax.NewOrderBy(col, desc))
	s.call(opOrderBy, col, desc)
	return s
}

// Join calls (INNER) JOIN statement.
func (s *Stmt) Join(table string) *Stmt {
	s.joinExpr = append(s.joinExpr, syntax.NewJoin(table, syntax.InnerJoin))
	s.call(opJoin, table)
	return s
}

// LeftJoin calls (INNER) JOIN statement.
func (s *Stmt) LeftJoin(table string) *Stmt {
	s.joinExpr = append(s.joinExpr, syntax.NewJoin(table, syntax.LeftJoin))
	s.call(opJoin, table)
	return s
}

// RightJoin calls (INNER) JOIN statement.
func (s *Stmt) RightJoin(table string) *Stmt {
	s.joinExpr = append(s.joinExpr, syntax.NewJoin(table, syntax.RightJoin))
	s.call(opJoin, table)
	return s
}

// FullJoin calls (INNER) JOIN statement.
func (s *Stmt) FullJoin(table string) *Stmt {
	s.joinExpr = append(s.joinExpr, syntax.NewJoin(table, syntax.FullJoin))
	s.call(opJoin, table)
	return s
}

// On calls ON statement.
func (s *Stmt) On(expr string, vals ...interface{}) *Stmt {
	s.onExpr = append(s.onExpr, syntax.NewOn(expr, vals...))
	s.call(opOn, expr, vals)
	return s
}

// Union calls UNION statement.
func (s *Stmt) Union(stmt syntax.Var) *Stmt {
	s.unionExpr = syntax.NewUnion(stmt, false)
	s.call(opUnion, stmt)
	return s
}

// UnionAll calls UNION ALL statement.
func (s *Stmt) UnionAll(stmt syntax.Var) *Stmt {
	s.unionExpr = syntax.NewUnion(stmt, true)
	s.call(opUnionAll, stmt)
	return s
}

// GroupBy calls GROUP BY statement.
func (s *Stmt) GroupBy(cols ...string) *Stmt {
	s.groupByExpr = syntax.NewGroupBy(cols)
	s.call(opGroupBy, cols)
	return s
}

// Having calls HAVING statement.
func (s *Stmt) Having(expr string, vals ...interface{}) *Stmt {
	s.havingExpr = syntax.NewHaving(expr, vals...)
	s.call(opHaving, expr, vals)
	return s
}

// When calls WHEN statement.
func (s *Stmt) When(expr string, vals ...interface{}) *Stmt {
	s.whenExpr = append(s.whenExpr, syntax.NewWhen(expr, vals...))
	s.call(opWhen, expr, vals)
	return s
}

// Then calls THEN statement.
func (s *Stmt) Then(val interface{}) *Stmt {
	s.thenExpr = append(s.thenExpr, syntax.NewThen(val))
	s.call(opThen, val)
	return s
}

// Else calls ELSE statement.
func (s *Stmt) Else(val interface{}) *Stmt {
	s.elseExpr = syntax.NewElse(val)
	s.call(opElse, val)
	return s
}
