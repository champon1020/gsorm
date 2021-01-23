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
	opQuery               internal.Op = "mgorm.Stmt.OpQuery"
	opExec                internal.Op = "mgorm.Stmt.OpExec"
	opFrom                internal.Op = "mgorm.Stmt.From"
	opValues              internal.Op = "mgorm.Stmt.Values"
	opSet                 internal.Op = "mgorm.Stmt.Set"
	opWhere               internal.Op = "mgorm.Stmt.Where"
	opAnd                 internal.Op = "mgorm.Stmt.And"
	opOr                  internal.Op = "mgorm.Stmt.Or"
	opLimit               internal.Op = "mgorm.Stmt.Limit"
)

// Stmt keeps the sql statement.
type Stmt struct {
	db         sqlDB
	cmd        syntax.Cmd
	fromExpr   syntax.Expr
	valuesExpr syntax.Expr
	setExpr    syntax.Expr
	whereExpr  syntax.Expr
	andOr      []syntax.Expr
	limitExpr  syntax.Expr
	errors     []error

	// Used for test.
	called []*opArgs
}

func (s *Stmt) call(op internal.Op, args ...interface{}) {
	s.called = append(s.called, &opArgs{op: op, args: args})
}

func (s *Stmt) addError(err error) {
	s.errors = append(s.errors, err)
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

	// Build WHERE.
	if s.whereExpr != nil {
		w, err := s.whereExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(w.Build())
	}

	// Build AND or OR.
	if len(s.andOr) > 0 {
		for _, e := range s.andOr {
			ao, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.write(ao.Build())
		}
	}

	if s.limitExpr != nil {
		l, err := s.limitExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(l.Build())
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
func (s *Stmt) ExpectExec(model interface{}) *Stmt {
	s.call(opExec, model)
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

	// Build AND or OR.
	if len(s.andOr) > 0 {
		for _, e := range s.andOr {
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
		/* handle error */
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
	s.andOr = append(s.andOr, syntax.NewAnd(expr, vals...))
	s.call(opAnd, expr, vals)
	return s
}

// Or calls OR statement.
func (s *Stmt) Or(expr string, vals ...interface{}) *Stmt {
	s.andOr = append(s.andOr, syntax.NewOr(expr, vals...))
	s.call(opOr, expr, vals)
	return s
}

// Limit calls LIMIT statement.
func (s *Stmt) Limit(num int) *Stmt {
	s.limitExpr = syntax.NewLimit(num)
	s.call(opLimit, num)
	return s
}
