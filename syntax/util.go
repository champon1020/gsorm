package syntax

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/champon1020/mgorm/internal"
	"github.com/morikuni/failure"
)

// BuildForExpression makes StmtSet with expr and values.
func BuildForExpression(expr string, vals ...interface{}) (string, error) {
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
		values = append(values, internal.ToString(v, nil))
	}

	return fmt.Sprintf(strings.ReplaceAll(expr, "?", "%s"), values...), nil
}
