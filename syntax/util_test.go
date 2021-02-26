package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/syntax"
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
			`IN lhs ('rhs', 100, true)`,
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
		ExpectedError error
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
		_, err := syntax.BuildForExpression(testCase.Expr, testCase.Values...)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}
		actualError, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}
		resultError := testCase.ExpectedError.(*errors.Error)
		if !resultError.Is(actualError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", resultError.Error(), resultError.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualError.Error(), actualError.Code)
		}
	}
}
