package syntax

import (
	"fmt"
	"strings"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
)

// BuildForExpression makes StmtSet with expr and values.
func BuildForExpression(expr string, vals ...interface{}) (string, error) {
	if strings.Count(expr, "?") != len(vals) {
		return "", errors.New("Length of values is not valid", errors.InvalidValueError)
	}

	values := []interface{}{}
	for _, v := range vals {
		if stmt, ok := v.(Stmt); ok {
			values = append(values, fmt.Sprintf("(%s)", stmt.String()))
			continue
		}

		vStr, err := internal.ToString(v, true)
		if err != nil {
			return "", err
		}
		values = append(values, vStr)
	}

	return fmt.Sprintf(strings.ReplaceAll(expr, "?", "%s"), values...), nil
}
