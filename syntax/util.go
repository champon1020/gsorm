package syntax

import (
	"errors"
	"fmt"
	"strings"

	"github.com/champon1020/mgorm/internal"
)

// Op values for error handling.
const (
	OpBuildStmtSet internal.Op = "syntax.BuildStmtSet"
)

// BuildStmtSetForExpression make StmtSet with expr and values.
func BuildStmtSetForExpression(expr string, vals ...interface{}) (*StmtSet, error) {
	if strings.Count(expr, "?") != len(vals) {
		err := errors.New("Length of values is not valid")
		return nil, internal.NewError(OpBuildStmtSet, internal.KindBasic, err)
	}

	ss := new(StmtSet)
	values := []interface{}{}
	for _, v := range vals {
		if sel, ok := v.(Var); ok {
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
