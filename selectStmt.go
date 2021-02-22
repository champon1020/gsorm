package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
)

type MgormSelect interface {
	From(...string) SelectFrom
}

type SelectFrom interface {
	Join(string) SelectJoin
	LeftJoin(string) SelectJoin
	RightJoin(string) SelectJoin
	FullJoin(string) SelectJoin
	Where(string, ...interface{}) SelectWhere
	GroupBy(...string) SelectGroupBy
	OrderBy(...string) SelectOrderBy
	Limit(int) SelectLimit
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// JoinStmt is returned after (*Stmt).Join, (*Stmt).LeftJoin, (*Stmt).RightJoin or (*Stmt).FullJoin is called.
type SelectJoin interface {
	On(string, ...interface{}) SelectOn
}

// OnStmt is returned after (*Stmt).On is called.
type SelectOn interface {
	Where(string, ...interface{}) SelectWhere
	GroupBy(...string) SelectGroupBy
	OrderBy(...string) SelectOrderBy
	Limit(int) SelectLimit
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// WhereStmt is returned after (*Stmt).Where is called.
type SelectWhere interface {
	And(string, ...interface{}) SelectAnd
	Or(string, ...interface{}) SelectOr
	GroupBy(...string) SelectGroupBy
	OrderBy(...string) SelectOrderBy
	Limit(int) SelectLimit
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// AndStmt is returned after (*Stmt).And is called.
type SelectAnd interface {
	GroupBy(...string) SelectGroupBy
	OrderBy(...string) SelectOrderBy
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// OrStmt is returned after (*Stmt).Or is called.
type SelectOr interface {
	GroupBy(...string) SelectGroupBy
	OrderBy(...string) SelectOrderBy
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// GroupByStmt is returned after (*Stmt).GroupBy is called.
type SelectGroupBy interface {
	Having(string, ...interface{}) SelectHaving
	OrderBy(...string) SelectOrderBy
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// HavingStmt is returned after (*Stmt).Having is called.
type SelectHaving interface {
	OrderBy(...string) SelectOrderBy
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// OrderByStmt is returned after (*Stmt).OrderBy is called.
type SelectOrderBy interface {
	Limit(int) SelectLimit
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// LimitStmt is returned after (*Stmt).Limit is called.
type SelectLimit interface {
	Offset(int) SelectOffset
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// OffsetStmt is returned after (*Stmt).Offset is called.
type SelectOffset interface {
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// UnionStmt is returned after (*Stmt).Union is called.
type SelectUnion interface {
	OrderBy(...string) SelectOrderBy
	Limit(int) SelectLimit
	Union(syntax.Stmt) SelectUnion
	UnionAll(syntax.Stmt) SelectUnion
	QueryCallable
}

// SelectStmt is SELECT statement.
type SelectStmt struct {
	stmt
	cmd *clause.Select
}

// String returns statement with string.
func (s *SelectStmt) String() string {
	sql, err := s.processSQL()
	if err != nil {
		s.throw(err)
		return err.Error()
	}
	return sql.String()
}

func (s *SelectStmt) funcString() string {
	str := s.cmd.String()
	for _, e := range s.called {
		str += fmt.Sprintf(".%s", e.String())
	}
	return str
}

func (s *SelectStmt) Query(model interface{}) error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch pool := s.db.(type) {
	case *DB, *Tx:
		sql, err := s.processSQL()
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
		returned, err := pool.CompareWith(s)
		if err != nil || returned == nil {
			return err
		}

		v := reflect.ValueOf(returned)
		if v.Kind() == reflect.Ptr {
			return errors.New("Returned value must not be pointer", errors.InvalidValueError)
		}
		mv := reflect.ValueOf(model)
		if mv.Kind() != reflect.Ptr {
			return errors.New("Model must be pointer", errors.InvalidPointerError)
		}

		mv.Elem().Set(v)
	default:
		return errors.New("DB type must be *DB, *Tx, *MockDB or *MockTx", errors.InvalidValueError)
	}

	return nil
}

func (s *SelectStmt) processSQL() (internal.SQL, error) {
	var sql internal.SQL

	ss, err := s.cmd.Build()
	if err != nil {
		return "", err
	}
	sql.Write(ss.Build())

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
			msg := fmt.Sprintf("Type %s is not supported for SELECT", reflect.TypeOf(e).Elem().String())
			return "", errors.New(msg, errors.InvalidTypeError)
		}
	}

	return sql, nil
}

// ExpectQuery returns *Stmt. This function is used for mock test.
func (s *SelectStmt) ExpectQuery(model interface{}) *SelectStmt {
	return s
}

// From calls FROM clause.
func (s *SelectStmt) From(tables ...string) SelectFrom {
	f := new(clause.From)
	for _, t := range tables {
		f.AddTable(t)
	}
	s.call(f)
	return s
}

// Where calls WHERE clause.
func (s *SelectStmt) Where(expr string, vals ...interface{}) SelectWhere {
	s.call(&clause.Where{Expr: expr, Values: vals})
	return s
}

// And calls AND clause.
func (s *SelectStmt) And(expr string, vals ...interface{}) SelectAnd {
	s.call(&clause.And{Expr: expr, Values: vals})
	return s
}

// Or calls OR clause.
func (s *SelectStmt) Or(expr string, vals ...interface{}) SelectOr {
	s.call(&clause.Or{Expr: expr, Values: vals})
	return s
}

// Limit calls LIMIT clause.
func (s *SelectStmt) Limit(num int) SelectLimit {
	s.call(&clause.Limit{Num: num})
	return s
}

// Offset calls OFFSET clause.
func (s *SelectStmt) Offset(num int) SelectOffset {
	s.call(&clause.Offset{Num: num})
	return s
}

// OrderBy calls ORDER BY clause.
func (s *SelectStmt) OrderBy(cols ...string) SelectOrderBy {
	s.call(&clause.OrderBy{Columns: cols})
	return s
}

// Join calls (INNER) JOIN clause.
func (s *SelectStmt) Join(table string) SelectJoin {
	j := &clause.Join{Type: clause.InnerJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// LeftJoin calls (INNER) JOIN clause.
func (s *SelectStmt) LeftJoin(table string) SelectJoin {
	j := &clause.Join{Type: clause.LeftJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// RightJoin calls (INNER) JOIN clause.
func (s *SelectStmt) RightJoin(table string) SelectJoin {
	j := &clause.Join{Type: clause.RightJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// FullJoin calls (INNER) JOIN clause.
func (s *SelectStmt) FullJoin(table string) SelectJoin {
	j := &clause.Join{Type: clause.FullJoin}
	j.AddTable(table)
	s.call(j)
	return s
}

// On calls ON clause.
func (s *SelectStmt) On(expr string, vals ...interface{}) SelectOn {
	s.call(&clause.On{Expr: expr, Values: vals})
	return s
}

// Union calls UNION clause.
func (s *SelectStmt) Union(stmt syntax.Stmt) SelectUnion {
	s.call(&clause.Union{Stmt: stmt, All: false})
	return s
}

// UnionAll calls UNION ALL clause.
func (s *SelectStmt) UnionAll(stmt syntax.Stmt) SelectUnion {
	s.call(&clause.Union{Stmt: stmt, All: true})
	return s
}

// GroupBy calls GROUP BY clause.
func (s *SelectStmt) GroupBy(cols ...string) SelectGroupBy {
	g := new(clause.GroupBy)
	for _, c := range cols {
		g.AddColumn(c)
	}
	s.call(g)
	return s
}

// Having calls HAVING clause.
func (s *SelectStmt) Having(expr string, vals ...interface{}) SelectHaving {
	s.call(&clause.Having{Expr: expr, Values: vals})
	return s
}
