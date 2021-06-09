package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOr_String(t *testing.T) {
	testCases := []struct {
		Or     *clause.Or
		Result string
	}{
		{
			&clause.Or{Expr: "lhs = rhs"},
			`Or("lhs = rhs")`,
		},
		{
			&clause.Or{Expr: "lhs = ?", Values: []interface{}{10}},
			`Or("lhs = ?", 10)`,
		},
		{
			&clause.Or{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			`Or("lhs1 = ? AND lhs2 = ?", 10, "str")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Or.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestOr_Build(t *testing.T) {
	testCases := []struct {
		Or     *clause.Or
		Result *syntax.ClauseSet
	}{
		{
			&clause.Or{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.ClauseSet{Keyword: "OR", Value: "lhs = 10", Parens: true},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Or.Build()
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestOr_Build_Fail(t *testing.T) {
	a := &clause.Or{Expr: "column = ?"}
	_, err := a.Build()
	if err == nil {
		t.Errorf("Error was not occurred")
	}
}
