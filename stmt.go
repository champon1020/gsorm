package mgorm

import (
	"fmt"
	"reflect"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
)

// Stmt stores information about query.
type Stmt struct {
	db     Pool
	cmd    syntax.Clause
	called []syntax.Clause
	model  interface{}
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

// String returns query with string.
func (s *Stmt) String() string {
	switch s.cmd.(type) {
	case *clause.Select:
		sql, err := s.processQuerySQL()
		if err != nil {
			s.addError(err)
			return err.Error()
		}
		return sql.String()
	case *clause.Insert, *clause.Update, *clause.Delete:
		sql, err := s.processExecSQL()
		if err != nil {
			s.addError(err)
			return err.Error()
		}
		return sql.String()
	}

	return "Error was occurred"
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
		_, err := pool.CompareWith(s)
		if err != nil {
			return err
		}
	default:
		return errors.New("DB type must be *DB, *Tx, *MockDB or *MockTx", errors.InvalidValueError)
	}

	return nil
}

// ExpectQuery returns *Stmt. This function is used for mock test.
func (s *Stmt) ExpectQuery(model interface{}) *Stmt {
	return s
}

// ExpectExec returns *Stmt. This function is used for mock test.
func (s *Stmt) ExpectExec() *Stmt {
	return s
}

// processQuerySQL builds SQL with called clauses.
func (s *Stmt) processQuerySQL() (internal.SQL, error) {
	var sql internal.SQL

	sel, ok := s.cmd.(*clause.Select)
	if !ok {
		return "", errors.New("Command must be SELECT", errors.InvalidValueError)
	}
	ss, err := sel.Build()
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

// processQuerySQL builds SQL with called clauses.
func (s *Stmt) processExecSQL() (internal.SQL, error) {
	var sql internal.SQL

	switch cmd := s.cmd.(type) {
	case *clause.Insert:
		ss, err := cmd.Build()
		if err != nil {
			return "", err
		}
		sql.Write(ss.Build())

		if s.model != nil {
			cols := []string{}
			for _, c := range cmd.Columns {
				if c.Alias != "" {
					cols = append(cols, c.Alias)
					continue
				}
				cols = append(cols, c.Name)
			}
			if err := s.processInsertModelSQL(cols, s.model, &sql); err != nil {
				return "", err
			}
			return sql, nil
		}

		for _, e := range s.called {
			if err := s.processInsertSQL(e, &sql); err != nil {
				return "", err
			}
		}
		return sql, nil
	case *clause.Update:
		ss, err := cmd.Build()
		if err != nil {
			return "", err
		}
		sql.Write(ss.Build())

		if s.model != nil {
			cols := []string{}
			for _, c := range cmd.Columns {
				cols = append(cols, c)
			}
			if err := s.processUpdateModelSQL(cols, s.model, &sql); err != nil {
				return "", err
			}
		}

		for _, e := range s.called {
			if err := s.processUpdateSQL(e, &sql); err != nil {
				return "", err
			}
		}
		return sql, nil
	case *clause.Delete:
		ss, err := cmd.Build()
		if err != nil {
			return "", err
		}
		sql.Write(ss.Build())
		for _, e := range s.called {
			if err := s.processDeleteSQL(e, &sql); err != nil {
				return "", err
			}
		}
		return sql, nil
	}

	return "", errors.New("Command must be INSERT, UPDATE or DELETE", errors.InvalidValueError)
}

func (s *Stmt) processInsertModelSQL(cols []string, model interface{}, sql *internal.SQL) error {
	ref := reflect.ValueOf(model)
	if ref.Kind() != reflect.Ptr {
		return errors.New("Model must be pointer", errors.InvalidValueError)
	}
	ref = ref.Elem()

	sql.Write("VALUES")
	switch ref.Kind() {
	case reflect.Slice, reflect.Array:
		// Type of slice element.
		typ := reflect.TypeOf(ref.Interface()).Elem()

		// If undelying type of slice element is struct.
		if typ.Kind() == reflect.Struct {
			idxC2F := internal.MapOfColumnsToFields(cols, typ)
			for i := 0; i < ref.Len(); i++ {
				if i > 0 {
					sql.Write(",")
				}
				sql.Write("(")
				for j := 0; j < len(cols); j++ {
					if j > 0 {
						sql.Write(",")
					}
					vStr, err := internal.ToString(ref.Index(i).Field(idxC2F[j]).Interface(), true)
					if err != nil {
						return err
					}
					sql.Write(vStr)
				}
				sql.Write(")")
			}
			return nil
		}

		for i := 0; i < ref.Len(); i++ {
			if i > 0 {
				sql.Write(",")
			}
			vStr, err := internal.ToString(ref.Index(i).Interface(), true)
			if err != nil {
				return err
			}
			sql.Write(fmt.Sprintf("(%s)", vStr))
		}
		return nil
	case reflect.Struct:
		idxC2F := internal.MapOfColumnsToFields(cols, reflect.TypeOf(ref.Interface()))
		sql.Write("(")
		for j := 0; j < len(cols); j++ {
			if j > 0 {
				sql.Write(",")
			}
			vStr, err := internal.ToString(ref.Field(idxC2F[j]).Interface(), true)
			if err != nil {
				return err
			}
			sql.Write(vStr)
		}
		sql.Write(")")
		return nil
	case reflect.Map:
		r := ref.MapRange()
		fst := true
		for r.Next() {
			if !fst {
				sql.Write(",")
			}
			key, err := internal.ToString(r.Key().Interface(), true)
			if err != nil {
				return err
			}
			val, err := internal.ToString(r.Value().Interface(), true)
			if err != nil {
				return err
			}
			sql.Write(fmt.Sprintf("(%s, %s)", key, val))
			fst = false
		}
		return nil
	}

	msg := fmt.Sprintf("Type %s is not supported for Model with INSERT", reflect.TypeOf(model).String())
	return errors.New(msg, errors.InvalidTypeError)
}

func (s *Stmt) processUpdateModelSQL(cols []string, model interface{}, sql *internal.SQL) error {
	ref := reflect.ValueOf(model)
	if ref.Kind() != reflect.Ptr {
		return errors.New("Model must be pointer", errors.InvalidValueError)
	}
	ref = ref.Elem()

	sql.Write("SET")
	switch ref.Kind() {
	case reflect.Struct:
		idxC2F := internal.MapOfColumnsToFields(cols, reflect.TypeOf(ref.Interface()))
		for i, c := range cols {
			if i > 0 {
				sql.Write(",")
			}
			vStr, err := internal.ToString(ref.Field(idxC2F[i]).Interface(), true)
			if err != nil {
				return err
			}
			sql.Write(fmt.Sprintf("%s = %s", c, vStr))
		}
		return nil
	case reflect.Map:
		if len(cols) != 2 {
			msg := fmt.Sprintf("If you set map to Model, number of columns must be 2, not %d", len(cols))
			return errors.New(msg, errors.InvalidSyntaxError)
		}
		r := ref.MapRange()
		for r.Next() {
			key, err := internal.ToString(r.Key().Interface(), true)
			if err != nil {
				return err
			}
			val, err := internal.ToString(r.Value().Interface(), true)
			if err != nil {
				return err
			}
			sql.Write(fmt.Sprintf("%s = %s, %s = %s", cols[0], key, cols[1], val))
			break
		}
		return nil
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Bool,
		reflect.String:
		if len(cols) != 1 {
			msg := fmt.Sprintf("If you set variable to Model, number of columns must be 1, not %d", len(cols))
			return errors.New(msg, errors.InvalidSyntaxError)
		}
		vStr, err := internal.ToString(ref.Interface(), true)
		if err != nil {
			return err
		}
		sql.Write(fmt.Sprintf("%s = %s", cols[0], vStr))
		return nil
	}

	msg := fmt.Sprintf("Type %s is not supported for Model with UPDATE", reflect.TypeOf(model).String())
	return errors.New(msg, errors.InvalidTypeError)
}

func (s *Stmt) processInsertSQL(e syntax.Clause, sql *internal.SQL) error {
	switch e := e.(type) {
	case *clause.Values:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		return nil
	}

	msg := fmt.Sprintf("Type %s is not supported for INSERT", reflect.TypeOf(e).Elem().String())
	return errors.New(msg, errors.InvalidTypeError)
}

func (s *Stmt) processUpdateSQL(e syntax.Clause, sql *internal.SQL) error {
	switch e := e.(type) {
	case *clause.Set,
		*clause.Where,
		*clause.And,
		*clause.Or:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		return nil
	}

	msg := fmt.Sprintf("Type %s is not supported for UPDATE", reflect.TypeOf(e).Elem().String())
	return errors.New(msg, errors.InvalidTypeError)
}

func (s *Stmt) processDeleteSQL(e syntax.Clause, sql *internal.SQL) error {
	switch e := e.(type) {
	case *clause.From:
		s, err := e.Build()
		if err != nil {
			return err
		}
		sql.Write(s.Build())
		return nil
	}

	msg := fmt.Sprintf("Type %s is not supported for DELETE", reflect.TypeOf(e).Elem().String())
	return errors.New(msg, errors.InvalidTypeError)
}

// Model sets model to Stmt.
func (s *Stmt) Model(model interface{}) ModelStmt {
	s.model = model
	return s
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
	u, ok := s.cmd.(*clause.Update)
	if !ok {
		s.addError(errors.New("SET clause can be used with UPDATE command", errors.InvalidValueError))
		return s
	}
	if len(u.Columns) != len(vals) {
		s.addError(errors.New("Length is different between lhs and rhs", errors.InvalidValueError))
		return s
	}
	set := new(clause.Set)
	for i, c := range u.Columns {
		set.AddEq(c, vals[i])
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
func (s *Stmt) OrderBy(cols ...string) OrderByStmt {
	s.call(&clause.OrderBy{Columns: cols})
	return s
}

/*
// OrderByDesc calls ORDER BY ... DESC clause.
func (s *Stmt) OrderByDesc(col string) OrderByStmt {
	s.call(&clause.OrderBy{Column: col, Desc: true})
	return s
}*/

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
func (s *Stmt) Union(stmt syntax.Stmt) UnionStmt {
	s.call(&clause.Union{Stmt: stmt, All: false})
	return s
}

// UnionAll calls UNION ALL clause.
func (s *Stmt) UnionAll(stmt syntax.Stmt) UnionStmt {
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
