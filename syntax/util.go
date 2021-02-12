package syntax

import (
	"fmt"
	"strings"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
)

// BuildStmtSetForExpression makes StmtSet with expr and values.
func BuildStmtSetForExpression(expr string, vals ...interface{}) (*StmtSet, error) {
	if strings.Count(expr, "?") != len(vals) {
		return nil, errors.New("Length of values is not valid", errors.InvalidValueError)
	}

	ss := new(StmtSet)
	values := []interface{}{}
	for _, v := range vals {
		if sel, ok := v.(Sub); ok {
			values = append(values, fmt.Sprintf("(%s)", sel))
			continue
		}
		vStr, err := internal.ToString(v, true)
		if err != nil {
			return nil, err
		}
		values = append(values, vStr)
	}

	ss.WriteValue(fmt.Sprintf(strings.ReplaceAll(expr, "?", "%s"), values...))
	return ss, nil
}
