package syntax

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/champon1020/mgorm/internal"
	"github.com/morikuni/failure"
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
		err := failure.New(errInvalidArgument,
			failure.Context{"expr": fmt.Sprintf("'%s'", expr), "length of values": strconv.Itoa(len(vals))},
			failure.Message("length of values is not valid"))
		return "", err
	}

	values := []interface{}{}
	for _, v := range vals {
		if stmt, ok := v.(Stmt); ok {
			values = append(values, fmt.Sprintf("(%s)", stmt.String()))
			continue
		}
		opt := &internal.ToStringOpt{Quotes: option.quotes}
		values = append(values, internal.ToString(v, opt))
	}

	return fmt.Sprintf(strings.ReplaceAll(expr, "?", "%s"), values...), nil
}
