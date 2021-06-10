package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestHaving_String(t *testing.T) {
	testCases := []struct {
		Having *clause.Having
		Result string
	}{
		{
			&clause.Having{Expr: "lhs = rhs"},
			`Having("lhs = rhs")`,
		},
		{
			&clause.Having{Expr: "lhs = ?", Values: []interface{}{10}},
			`Having("lhs = ?", 10)`,
		},
		{
			&clause.Having{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			`Having("lhs1 = ? AND lhs2 = ?", 10, "str")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Having.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestHaving_Build(t *testing.T) {
	testCases := []struct {
		Having *clause.Having
		Result *syntax.ClauseSet
	}{
		{
			&clause.Having{Expr: "lhs = rhs"},
			&syntax.ClauseSet{Keyword: "HAVING", Value: "lhs = rhs"},
		},
		{
			&clause.Having{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.ClauseSet{Keyword: "HAVING", Value: "lhs = 10"},
		},
		{
			&clause.Having{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			&syntax.ClauseSet{Keyword: "HAVING", Value: `lhs1 = 10 AND lhs2 = 'str'`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Having.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestHaving_Build_Fail(t *testing.T) {
	a := &clause.Having{Expr: "column = ?"}
	_, err := a.Build()
	if err == nil {
		t.Errorf("Error was not occurred")
	}
}
