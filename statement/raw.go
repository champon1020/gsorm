package statement

import (
	"reflect"

	"github.com/champon1020/gsorm/interfaces/domain"
	"github.com/champon1020/gsorm/internal"
	"github.com/champon1020/gsorm/syntax"
	"github.com/morikuni/failure"
)

// RawStmt is raw string statement.
type RawStmt struct {
	stmt
	cmd *syntax.RawClause
}

// NewRawStmt creates RawStmt instance.
func NewRawStmt(conn domain.Conn, rs string, v ...interface{}) *RawStmt {
	s := &RawStmt{cmd: &syntax.RawClause{RawStr: rs, Values: v}}
	s.conn = conn
	return s
}

func (r *RawStmt) String() string {
	return r.string(r.buildSQL)
}

// FuncString returns function call as string.
func (r *RawStmt) FuncString() string {
	return r.funcString(r.cmd)
}

// Cmd returns cmd clause.
func (r *RawStmt) Cmd() domain.Clause {
	return r.cmd
}

// CompareWith compares the statements and returns error if the statements is not same.
// In this case, same means that stmt.cmd and stmt.called is corresponding.
func (r *RawStmt) CompareWith(s domain.Stmt) error {
	return r.compareWith(r.Cmd(), s)
}

// Query executes SQL statement with mapping to model.
// If type of (*SelectStmt).conn is gsorm.MockDB, compare statements between called and expected.
// Then, it maps expected values to model.
func (r *RawStmt) Query(model interface{}) error {
	return r.query(r.buildSQL, r, model)
}

// Exec executed SQL statement without mapping to model.
// If type of conn is gsorm.MockDB, compare statements between called and expected.
func (r *RawStmt) Exec() error {
	return r.exec(r.buildSQL, r)
}

// Migrate executes database migration.
func (r *RawStmt) Migrate() error {
	if len(r.errors) > 0 {
		return r.errors[0]
	}

	switch conn := r.conn.(type) {
	case domain.Mock:
		return nil
	case domain.DB, domain.Tx:
		var sql internal.SQL
		if err := r.buildSQL(&sql); err != nil {
			return err
		}
		if _, err := conn.Exec(sql.String()); err != nil {
			return failure.Wrap(err)
		}
		return nil
	}

	return failure.New(errInvalidValue,
		failure.Context{"conn": reflect.TypeOf(r.conn).String()},
		failure.Message("conn can be *DB, *Tx, *MockDB or *MockTx"))
}

// buildSQL builds SQL statement from called clauses.
func (r *RawStmt) buildSQL(sql *internal.SQL) error {
	ss, err := r.cmd.Build()
	if err != nil {
		return err
	}
	sql.Write(ss.Build())
	return nil
}
