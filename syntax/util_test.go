package syntax_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/syntax"
	"github.com/stretchr/testify/assert"
)

func TestBuildExpr(t *testing.T) {
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
			"IN (?)",
			[]interface{}{gsorm.Select(nil, "*").
				From("table").
				Where("lhs = ?", "rhs")},
			`IN (SELECT * FROM table WHERE lhs = 'rhs')`,
		},
	}

	for _, testCase := range testCases {
		actual, _ := syntax.BuildExpr(testCase.Expr, testCase.Values...)
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestBuildExpr_Fail(t *testing.T) {
	expectedErr := "number of values doesn't match the number of '?'"
	expr := "lhs = ? AND rhs = ?"
	vals := []interface{}{10}

	// Actual process.
	_, err := syntax.BuildExpr(expr, vals)

	// Validate.
	assert.EqualError(t, err, expectedErr)
}

func TestBuildExprWithoutQuotes(t *testing.T) {
	testCases := []struct {
		Expr     string
		Values   []interface{}
		Expected string
	}{
		{
			"lhs = ?",
			[]interface{}{"rhs"},
			`lhs = rhs`,
		},
	}

	for _, testCase := range testCases {
		actual, _ := syntax.BuildExprWithoutQuotes(testCase.Expr, testCase.Values...)
		assert.Equal(t, testCase.Expected, actual)
	}
}
