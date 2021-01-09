package syntax

// Stmt keeps the sql statement.
type Stmt struct {
	DB        DbIface
	Mode      uint
	Cmd       Cmd
	From      Expr
	Values    Expr
	Set       Expr
	WhereExpr Expr
	AndOr     []Expr
	Errors    []error
}

// Query executes a query that returns some results.
func (s *Stmt) Query(model interface{}) error {
	sql, err := s.processQuerySQL()
	if err != nil {
		return err
	}
	if err := sql.doQuery(s.DB, model); err != nil {
		return err
	}
	return nil
}

func (s *Stmt) processQuerySQL() (SQL, error) {
	var sql SQL

	// Build SELECT.
	sel, ok := s.Cmd.(*Select)
	if !ok {
		return "", newError(ErrInvalidType, "command must be SELECT")
	}
	sql.write(sel.Build().Build())

	// Build FROM.
	if s.From != nil {
		from, err := s.From.Build()
		if err != nil {
			return "", err
		}
		sql.write(from.Build())
	}

	// Build WHERE.
	if s.WhereExpr != nil {
		w, err := s.WhereExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(w.Build())
	}

	// Build AND or OR.
	if len(s.AndOr) > 0 {
		for _, e := range s.AndOr {
			ao, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.write(ao.Build())
		}
	}
	return sql, nil
}

// Exec executes a query without returning any results.
func (s *Stmt) Exec() error {
	sql, err := s.processExecSQL()
	if err != nil {
		return err
	}
	if err := sql.doExec(s.DB); err != nil {
		return err
	}
	return nil
}

func (s *Stmt) processExecSQL() (SQL, error) {
	var sql SQL
	switch cmd := s.Cmd.(type) {
	case *Insert:
		sql.write(cmd.Build().Build())
		if s.Values != nil {
			values, err := s.Values.Build()
			if err != nil {
				return "", err
			}
			sql.write(values.Build())
		}
	case *Update:
		sql.write(cmd.Build().Build())
		if s.Set != nil {
			set, err := s.Set.Build()
			if err != nil {
				return "", err
			}
			sql.write(set.Build())
		}
	case *Delete:
		sql.write(cmd.Build().Build())
		if s.From != nil {
			from, err := s.From.Build()
			if err != nil {
				return "", err
			}
			sql.write(from.Build())
		}
	default:
		return "", newError(ErrInvalidType, "command must be INSERT, UPDATE or DELETE")
	}

	// Build WHERE.
	if s.WhereExpr != nil {
		w, err := s.WhereExpr.Build()
		if err != nil {
			return "", err
		}
		sql.write(w.Build())
	}

	// Build AND or OR.
	if len(s.AndOr) > 0 {
		for _, e := range s.AndOr {
			ao, err := e.Build()
			if err != nil {
				return "", err
			}
			sql.write(ao.Build())
		}
	}
	return sql, nil
}

// AddError append error to stmt.
func (s *Stmt) addError(err error) {
	s.Errors = append(s.Errors, err)
}

// Where calls WHERE statement.
func (s *Stmt) Where(expr string, vals ...interface{}) *Stmt {
	w := NewWhere(expr, vals...)
	s.WhereExpr = w
	return s
}

// And calls AND statement.
func (s *Stmt) And(expr string, vals ...interface{}) *Stmt {
	w := NewAnd(expr, vals...)
	s.AndOr = append(s.AndOr, w)
	return s
}

// Or calls OR statement.
func (s *Stmt) Or(expr string, vals ...interface{}) *Stmt {
	w := NewOr(expr, vals...)
	s.AndOr = append(s.AndOr, w)
	return s
}
