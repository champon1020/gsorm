package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
)

func TestBuildStmtSetForExpression(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.StmtSet
	}{
		{
			`lhs = "rhs"`,
			[]interface{}{},
			&syntax.StmtSet{Value: `lhs = "rhs"`},
		},
		{
			"lhs = ?",
			[]interface{}{"rhs"},
			&syntax.StmtSet{Value: `lhs = "rhs"`},
		},
		{
			"NOT lhs = ?",
			[]interface{}{"rhs"},
			&syntax.StmtSet{Value: `NOT lhs = "rhs"`},
		},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{"rhs", 100},
			&syntax.StmtSet{Value: `lhs1 = "rhs" AND lhs2 = 100`},
		},
		{
			"IN lhs (?, ?, ?)",
			[]interface{}{"rhs", 100, true},
			&syntax.StmtSet{Value: `IN lhs ("rhs", 100, true)`},
		},
		{
			"lhs LIKE %%?%%",
			[]interface{}{"rhs"},
			&syntax.StmtSet{Value: `lhs LIKE %"rhs"%`},
		},
		{
			"lhs BETWEEN ? AND ?",
			[]interface{}{10, 100},
			&syntax.StmtSet{Value: "lhs BETWEEN 10 AND 100"},
		},
		{
			"IN ?",
			[]interface{}{mgorm.Select(nil, "*").
				From("table").
				Where("lhs = ?", "rhs").
				Sub()},
			&syntax.StmtSet{Value: `IN (SELECT * FROM table WHERE lhs = "rhs")`},
		},
	}

	for _, testCase := range testCases {
		res, _ := syntax.BuildStmtSetForExpression(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestBuildStmtSetForExpression_Fail(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Error  error
	}{
		{
			"lhs = ? AND rhs = ?",
			[]interface{}{10},
			errors.New("Length of values is not valid", errors.InvalidValueError),
		},
		{
			"lhs = ?",
			[]interface{}{[]string{}},
			errors.New("Type slice is not supported", errors.InvalidTypeError),
		},
	}

	for _, testCase := range testCases {
		_, err := syntax.BuildStmtSetForExpression(testCase.Expr, testCase.Values...)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}
		actualError, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}
		resultError := testCase.Error.(*errors.Error)
		if !resultError.Is(actualError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", resultError.Error(), resultError.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualError.Error(), actualError.Code)
		}
	}
}
