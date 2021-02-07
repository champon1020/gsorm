package syntax_test

import (
	"errors"
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
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
			internal.PrintTestDiff(t, diff)
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
			internal.NewError(
				syntax.OpBuildStmtSetForExpression,
				internal.KindBasic,
				errors.New("Length of values is not valid"),
			),
		},
		{
			"lhs = ?",
			[]interface{}{[]string{}},
			internal.NewError(
				internal.OpToString,
				internal.KindType,
				errors.New("type is invalid"),
			),
		},
	}

	for _, testCase := range testCases {
		_, err := syntax.BuildStmtSetForExpression(testCase.Expr, testCase.Values...)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}

		e, ok := err.(*internal.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}

		if diff := internal.CmpError(testCase.Error.(*internal.Error), e); diff != "" {
			t.Errorf(diff)
		}
	}
}
