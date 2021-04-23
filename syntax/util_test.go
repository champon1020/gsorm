package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/syntax"
	"github.com/morikuni/failure"
	"github.com/stretchr/testify/assert"
)

func TestBuildForExpression(t *testing.T) {
	testCases := []struct {
		Expr     string
		Values   []interface{}
		Expected string
	}{
		{
			`lhs = 'rhs'`,
			[]interface{}{},
			`lhs = 'rhs'`,
		},
		{
			"lhs = ?",
			[]interface{}{"rhs"},
			`lhs = 'rhs'`,
		},
		{
			"NOT lhs = ?",
			[]interface{}{"rhs"},
			`NOT lhs = 'rhs'`,
		},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{"rhs", 100},
			`lhs1 = 'rhs' AND lhs2 = 100`,
		},
		{
			"IN lhs (?, ?, ?)",
			[]interface{}{"rhs", 100, true},
			`IN lhs ('rhs', 100, 1)`,
		},
		{
			"lhs LIKE %%?%%",
			[]interface{}{"rhs"},
			`lhs LIKE %'rhs'%`,
		},
		{
			"lhs BETWEEN ? AND ?",
			[]interface{}{10, 100},
			"lhs BETWEEN 10 AND 100",
		},
		{
			"IN ?",
			[]interface{}{mgorm.Select(nil, "*").
				From("table").
				Where("lhs = ?", "rhs")},
			`IN (SELECT * FROM table WHERE lhs = 'rhs')`,
		},
	}

	for _, testCase := range testCases {
		actual, _ := syntax.BuildForExpression(testCase.Expr, testCase.Values...)
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestBuildForExpression_Fail(t *testing.T) {
	testCases := []struct {
		Expr          string
		Values        []interface{}
		ExpectedError failure.StringCode
	}{
		{
			"lhs = ? AND rhs = ?",
			[]interface{}{10},
			syntax.ErrInvalidArgument,
		},
	}

	for _, testCase := range testCases {
		_, err := syntax.BuildForExpression(testCase.Expr, testCase.Values...)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}
		if !failure.Is(err, syntax.ErrInvalidArgument) {
			t.Errorf("Different error")
			continue
		}
	}
}
