package syntax

import (
	"fmt"
	"strings"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/internal"
	"golang.org/x/xerrors"
)

// BuildExpr assigns the values to '?' of the expression.
func BuildExpr(expr string, vals ...interface{}) (string, error) {
	return buildExprWithOpt(&buildExprOpt{quotes: true}, expr, vals...)
}

// BuildExprWithoutQuotes assigns the values to '?' of the expression with single quotes.
func BuildExprWithoutQuotes(expr string, vals ...interface{}) (string, error) {
	return buildExprWithOpt(&buildExprOpt{quotes: false}, expr, vals...)
}

type buildExprOpt struct {
	quotes bool
}

func buildExprWithOpt(option *buildExprOpt, expr string, vals ...interface{}) (string, error) {
	if strings.Count(expr, "?") != len(vals) {
		return "", xerrors.New("number of values doesn't match the number of '?'")
	}

	values := []interface{}{}
	for _, v := range vals {
		if stmt, ok := v.(interfaces.Stmt); ok {
			values = append(values, stmt.SQL())
			continue
		}
		opt := &internal.ToStringOpt{Quotes: option.quotes}
		values = append(values, internal.ToString(v, opt))
	}

	return fmt.Sprintf(strings.ReplaceAll(expr, "?", "%s"), values...), nil
}
