package gsorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/interfaces/idelete"
	"github.com/champon1020/gsorm/interfaces/iinsert"
	"github.com/champon1020/gsorm/interfaces/iselect"
	"github.com/champon1020/gsorm/interfaces/iupdate"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/xerrors"
)

// stmt stores information about query.
type stmt struct {
	conn   conn
	cmd    interfaces.Clause
	called []interfaces.Clause
	errors []error
}

// call appends called clause.
func (s *stmt) call(e interfaces.Clause) {
	s.called = append(s.called, e)
}

// throw appends occurred error.
func (s *stmt) throw(err error) {
	s.errors = append(s.errors, err)
}

// Cmd returns the command clause.
func (s *stmt) Cmd() interfaces.Clause {
	return s.cmd
}

// Clauses returns the called clauses.
func (s *stmt) Clauses() []interfaces.Clause {
	return s.called
}

func (s *stmt) sql(buildSQL func(*internal.SQL) error) string {
	var sql internal.SQL
	if err := buildSQL(&sql); err != nil {
		s.throw(err)
		return err.Error()
	}
	return sql.String()
}

// String returns the method calls as string.
func (s *stmt) String() string {
	str := s.cmd.String()
	for _, e := range s.called {
		str += fmt.Sprintf(".%s", e.String())
	}
	return str
}

// CompareWith compares the statements and returns error if the statements is not same.
// In this case, same means that stmt.cmd and stmt.called is corresponding.
func (s *stmt) CompareWith(targetStmt interfaces.Stmt) error {
	if diff := cmp.Diff(s.cmd, targetStmt.Cmd()); diff != "" {
		return xerrors.Errorf("statements comparison was failed:\nexpected: %s\nactual:   %s\n",
			s.String(), targetStmt.String())
	}

	expected := s.called
	actual := targetStmt.Clauses()
	if len(expected) != len(actual) {
		return xerrors.Errorf("statements comparison was failed:\nexpected: %s\nactual:   %s\n",
			s.String(), targetStmt.String())
	}
	for i, e := range expected {
		if diff := cmp.Diff(actual[i], e); diff != "" {
			return xerrors.Errorf("statements comparison was failed:\nexpected: %s\nactual:   %s\n",
				s.String(), targetStmt.String())
		}
	}
	return nil
}

func (s *stmt) query(buildSQL func(*internal.SQL) error, stmt interfaces.Stmt, model interface{}) error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch conn := s.conn.(type) {
	case Mock:
		returned, err := conn.compareWith(stmt)
		if err != nil || returned == nil {
			return err
		}

		v := reflect.ValueOf(returned)
		if v.Kind() == reflect.Ptr {
			return xerrors.New("returned value should not be a pointer")
		}
		mv := reflect.ValueOf(model)
		if mv.Kind() != reflect.Ptr {
			return xerrors.New("model must be a pointer")
		}

		mv.Elem().Set(v)
		return nil
	case DB, Tx:
		var sql internal.SQL
		if err := buildSQL(&sql); err != nil {
			return err
		}

		rows, err := conn.Query(sql.String())
		if err != nil {
			return err
		}
		defer rows.Close()

		p, err := newRowsParser(rows, model)
		if err != nil {
			return err
		}

		v, err := p.Parse()
		if err != nil {
			return err
		}

		ref := reflect.ValueOf(model).Elem()
		ref.Set(*v)
		return nil
	}

	return xerrors.Errorf("database connection should not be %s", reflect.TypeOf(s.conn).String())
}

func (s *stmt) exec(buildSQL func(*internal.SQL) error, stmt interfaces.Stmt) error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch conn := s.conn.(type) {
	case Mock:
		_, err := conn.compareWith(stmt)
		if err != nil {
			return err
		}
		return nil
	case DB, Tx:
		var sql internal.SQL
		if err := buildSQL(&sql); err != nil {
			return err
		}
		if _, err := conn.Exec(sql.String()); err != nil {
			return err
		}
		return nil
	}

	return xerrors.Errorf("database connection should not be %s", reflect.TypeOf(s.conn).String())
}

// DeleteStmt is DELETE statement.
type DeleteStmt struct {
	stmt
}

// newDeleteStmt creates DeleteStmt instance.
func newDeleteStmt(conn conn) *DeleteStmt {
	s := &DeleteStmt{}
	s.conn = conn
	s.cmd = &clause.Delete{}
	return s
}

// SQL returns the built SQL string.
func (s *DeleteStmt) SQL() string {
	return s.sql(s.buildSQL)
}

// Exec executed SQL statement without mapping to model.
// If type of conn is gsorm.MockDB, compare statements between called and expected.
func (s *DeleteStmt) Exec() error {
	return s.exec(s.buildSQL, s)
}

// buildSQL builds SQL statement.
func (s *DeleteStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	for _, e := range s.called {
		switch e := e.(type) {
		case *syntax.RawClause,
			*clause.From,
			*clause.Where,
			*clause.And,
			*clause.Or:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
		default:
			return xerrors.Errorf("%s is invalid clause for DELETE", reflect.TypeOf(e).Elem().String())
		}
	}
	return nil
}

// RawClause calls the raw string clause.
func (s *DeleteStmt) RawClause(raw string, values ...interface{}) idelete.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: values})
	return s
}

// From calls FROM clause.
func (s *DeleteStmt) From(tables ...string) idelete.From {
	f := new(clause.From)
	for _, t := range tables {
		f.AddTable(t)
	}
	s.call(f)
	return s
}

// Where calls WHERE clause.
func (s *DeleteStmt) Where(expr string, values ...interface{}) idelete.Where {
	s.call(&clause.Where{Expr: expr, Values: values})
	return s
}

// And calls AND clause.
func (s *DeleteStmt) And(expr string, values ...interface{}) idelete.And {
	s.call(&clause.And{Expr: expr, Values: values})
	return s
}

// Or calls OR clause.
func (s *DeleteStmt) Or(expr string, values ...interface{}) idelete.Or {
	s.call(&clause.Or{Expr: expr, Values: values})
	return s
}

// InsertStmt is INSERT statement.
type InsertStmt struct {
	stmt
	model interface{}
	sel   interfaces.Stmt
}

// newInsertStmt creates InsertStmt instance.
func newInsertStmt(conn conn, table string, cols ...string) *InsertStmt {
	i := &clause.Insert{}
	i.AddTable(table)
	i.AddColumns(cols...)
	s := &InsertStmt{}
	s.conn = conn
	s.cmd = i
	return s
}

// SQL returns the built SQL string.
func (s *InsertStmt) SQL() string {
	return s.sql(s.buildSQL)
}

// Exec executed SQL statement without mapping to model.
// If type of conn is gsorm.MockDB, compare statements between called and expected.
func (s *InsertStmt) Exec() error {
	return s.exec(s.buildSQL, s)
}

// buildSQL builds SQL statement.
func (s *InsertStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())

	if s.model != nil {
		insertCmd, ok := s.cmd.(*clause.Insert)
		if !ok {
			return xerrors.New("command must be clause.Insert")
		}
		var cols []string
		for _, c := range insertCmd.Columns {
			if c.Alias != "" {
				cols = append(cols, c.Alias)
				continue
			}
			cols = append(cols, c.Name)
		}
		if err := s.buildSQLWithModel(cols, s.model, sql); err != nil {
			return err
		}
		return nil
	}

	if s.sel != nil {
		sql.Write(s.sel.SQL())
		return nil
	}

	if err := s.buildSQLWithClauses(sql); err != nil {
		return err
	}
	return nil
}

// buildSQLWithClauses builds SQL statement from called clauses.
func (s *InsertStmt) buildSQLWithClauses(sql *internal.SQL) error {
	valuesCalled := false
	for _, e := range s.called {
		switch e := e.(type) {
		case *clause.Values:
			s, err := e.Build()
			if err != nil {
				return err
			}
			if valuesCalled {
				sql.Write(",")
				sql.Write(s.BuildValue())
				continue
			}
			sql.Write(s.Build())
			valuesCalled = true
		case *syntax.RawClause:
			ss, err := e.Build()
			if err != nil {
				return err
			}
			sql.Write(ss.Build())
		default:
			return xerrors.Errorf("%s is invalid clause for INSERT", reflect.TypeOf(e).Elem().String())
		}
	}
	return nil
}

// buildSQLWithModel builds SQL statement from model.
func (s *InsertStmt) buildSQLWithModel(cols []string, model interface{}, sql *internal.SQL) error {
	sql.Write("VALUES")
	p, err := newInsertModelParser(cols, model)
	if err != nil {
		return err
	}

	modelSQL, err := p.Parse()
	if err != nil {
		return err
	}

	sql.Write(modelSQL.String())
	return nil
}

// RawClause calls the raw string clause.
func (s *InsertStmt) RawClause(raw string, values ...interface{}) iinsert.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: values})
	return s
}

// Model sets model to InsertStmt.
func (s *InsertStmt) Model(model interface{}) iinsert.Model {
	s.model = model
	return s
}

// Select calls SELECT statement.
func (s *InsertStmt) Select(stmt interfaces.Stmt) iinsert.Select {
	s.sel = stmt
	return s
}

// Values calls VALUES clause.
func (s *InsertStmt) Values(values ...interface{}) iinsert.Values {
	v := new(clause.Values)
	for _, val := range values {
		v.AddValue(val)
	}
	s.call(v)
	return s
}

// SelectStmt is SELECT statement.
type SelectStmt struct {
	stmt
}

// newSelectStmt creates SelectStmt instance.
func newSelectStmt(conn conn, cols ...string) *SelectStmt {
	sel := &clause.Select{}
	if len(cols) == 0 {
		sel.AddColumns("*")
	} else {
		sel.AddColumns(cols...)
	}
	s := &SelectStmt{}
	s.conn = conn
	s.cmd = sel
	return s
}

// SQL returns the built SQL string.
func (s *SelectStmt) SQL() string {
	return s.sql(s.buildSQL)
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
			return xerrors.Errorf("%s is invalid clause for SELECT", reflect.TypeOf(e).Elem().String())
		}
	}

	return nil
}

// RawClause calls the raw string clause.
func (s *SelectStmt) RawClause(raw string, values ...interface{}) iselect.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: values})
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
func (s *SelectStmt) Where(expr string, values ...interface{}) iselect.Where {
	s.call(&clause.Where{Expr: expr, Values: values})
	return s
}

// And calls AND clause.
func (s *SelectStmt) And(expr string, values ...interface{}) iselect.And {
	s.call(&clause.And{Expr: expr, Values: values})
	return s
}

// Or calls OR clause.
func (s *SelectStmt) Or(expr string, values ...interface{}) iselect.Or {
	s.call(&clause.Or{Expr: expr, Values: values})
	return s
}

// Limit calls LIMIT clause.
func (s *SelectStmt) Limit(limit int) iselect.Limit {
	s.call(&clause.Limit{Num: limit})
	return s
}

// Offset calls OFFSET clause.
func (s *SelectStmt) Offset(offset int) iselect.Offset {
	s.call(&clause.Offset{Num: offset})
	return s
}

// OrderBy calls ORDER BY clause.
func (s *SelectStmt) OrderBy(columns ...string) iselect.OrderBy {
	s.call(&clause.OrderBy{Columns: columns})
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
func (s *SelectStmt) On(expr string, values ...interface{}) iselect.On {
	s.call(&clause.On{Expr: expr, Values: values})
	return s
}

// Union calls UNION clause.
func (s *SelectStmt) Union(stmt interfaces.Stmt) iselect.Union {
	s.call(&clause.Union{Stmt: stmt, All: false})
	return s
}

// UnionAll calls UNION ALL clause.
func (s *SelectStmt) UnionAll(stmt interfaces.Stmt) iselect.Union {
	s.call(&clause.Union{Stmt: stmt, All: true})
	return s
}

// GroupBy calls GROUP BY clause.
func (s *SelectStmt) GroupBy(columns ...string) iselect.GroupBy {
	g := new(clause.GroupBy)
	for _, c := range columns {
		g.AddColumn(c)
	}
	s.call(g)
	return s
}

// Having calls HAVING clause.
func (s *SelectStmt) Having(expr string, values ...interface{}) iselect.Having {
	s.call(&clause.Having{Expr: expr, Values: values})
	return s
}

// UpdateStmt is UPDATE statement..
type UpdateStmt struct {
	stmt
	model     interface{}
	modelCols []string
}

// newUpdateStmt creates UpdateStmt instance.
func newUpdateStmt(conn conn, table string) *UpdateStmt {
	u := &clause.Update{}
	u.AddTable(table)
	s := &UpdateStmt{}
	s.conn = conn
	s.cmd = u
	return s
}

// SQL returns the built SQL string.
func (s *UpdateStmt) SQL() string {
	return s.sql(s.buildSQL)
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
			return xerrors.Errorf("%s is invalid clause for UPDATE", reflect.TypeOf(e).Elem().String())
		}
	}
	return nil
}

// buildSQLWithModel builds SQL statement from model.
func (s *UpdateStmt) buildSQLWithModel(cols []string, model interface{}, sql *internal.SQL) error {
	sql.Write("SET")
	p, err := newUpdateModelParser(cols, model)
	if err != nil {
		return err
	}

	modelSQL, err := p.Parse()
	if err != nil {
		return err
	}

	sql.Write(modelSQL.String())
	return nil
}

// RawClause calls the raw string clause.
func (s *UpdateStmt) RawClause(raw string, values ...interface{}) iupdate.RawClause {
	s.call(&syntax.RawClause{RawStr: raw, Values: values})
	return s
}

// Model sets model to UpdateStmt.
func (s *UpdateStmt) Model(model interface{}, columns ...string) iupdate.Model {
	s.model = model
	s.modelCols = columns
	return s
}

// Set calls SET clause.
func (s *UpdateStmt) Set(column string, value interface{}) iupdate.Set {
	s.call(&clause.Set{Column: column, Value: value})
	return s
}

// Where calls WHERE clause.
func (s *UpdateStmt) Where(expr string, values ...interface{}) iupdate.Where {
	s.call(&clause.Where{Expr: expr, Values: values})
	return s
}

// And calls AND clause.
func (s *UpdateStmt) And(expr string, values ...interface{}) iupdate.And {
	s.call(&clause.And{Expr: expr, Values: values})
	return s
}

// Or calls OR clause.
func (s *UpdateStmt) Or(expr string, values ...interface{}) iupdate.Or {
	s.call(&clause.Or{Expr: expr, Values: values})
	return s
}

// rawStmt is raw string statement.
type rawStmt struct {
	stmt
}

// newRawStmt creates rawStmt instance.
func newRawStmt(conn conn, rs string, v ...interface{}) *rawStmt {
	s := &rawStmt{}
	s.conn = conn
	s.cmd = &syntax.RawClause{RawStr: rs, Values: v}
	return s
}

// SQL returns the built SQL string.
func (s *rawStmt) SQL() string {
	return s.sql(s.buildSQL)
}

// Query executes SQL statement with mapping to model.
// If type of (*SelectStmt).conn is gsorm.MockDB, compare statements between called and expected.
// Then, it maps expected values to model.
func (s *rawStmt) Query(model interface{}) error {
	return s.query(s.buildSQL, s, model)
}

// Exec executed SQL statement without mapping to model.
// If type of conn is gsorm.MockDB, compare statements between called and expected.
func (s *rawStmt) Exec() error {
	return s.exec(s.buildSQL, s)
}

// Migrate executes database migration.
func (s *rawStmt) Migrate() error {
	if len(s.errors) > 0 {
		return s.errors[0]
	}

	switch conn := s.conn.(type) {
	case Mock:
		return nil
	case DB, Tx:
		var sql internal.SQL
		if err := s.buildSQL(&sql); err != nil {
			return err
		}
		if _, err := conn.Exec(sql.String()); err != nil {
			return err
		}
		return nil
	}

	return xerrors.Errorf("database connection should not be %s", reflect.TypeOf(s.conn).String())
}

// buildSQL builds SQL statement from called clauses.
func (s *rawStmt) buildSQL(sql *internal.SQL) error {
	ss, err := s.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())
	return nil
}
