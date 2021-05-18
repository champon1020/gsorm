package statement

import "github.com/champon1020/mgorm/domain"

const (
	ErrInvalidValue  = errInvalidValue
	ErrInvalidClause = errInvalidClause
	ErrInvalidSyntax = errInvalidSyntax
	ErrInvalidType   = errInvalidType
	ErrFailedParse   = errFailedParse
)

// Exported values which is declared in stmt.go.
var (
	SelectStmtBuildSQL = (*SelectStmt).buildSQL
	InsertStmtBuildSQL = (*InsertStmt).buildSQL
	UpdateStmtBuildSQL = (*UpdateStmt).buildSQL
	DeleteStmtBuildSQL = (*DeleteStmt).buildSQL
)

func (s *stmt) ExportedGetErrors() []error {
	return s.errors
}

func (s *DeleteStmt) ExportedSetCalled(c ...domain.Clause) {
	s.called = append(s.called, c...)
}

func (s *InsertStmt) ExportedSetCalled(c ...domain.Clause) {
	s.called = append(s.called, c...)
}

func (s *SelectStmt) ExportedSetCalled(c ...domain.Clause) {
	s.called = append(s.called, c...)
}

func (s *UpdateStmt) ExportedSetCalled(c ...domain.Clause) {
	s.called = append(s.called, c...)
}
