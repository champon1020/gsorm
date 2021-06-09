package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestAnd_String(t *testing.T) {
	testCases := []struct {
		And    *clause.And
		Result string
	}{
		{
			&clause.And{Expr: "lhs = rhs"},
			`And("lhs = rhs")`,
		},
		{
			&clause.And{Expr: "lhs = ?", Values: []interface{}{10}},
			`And("lhs = ?", 10)`,
		},
		{
			&clause.And{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			`And("lhs1 = ? AND lhs2 = ?", 10, "str")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.And.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestAnd_Build(t *testing.T) {
	testCases := []struct {
		And    *clause.And
		Result *syntax.ClauseSet
	}{
		{
			&clause.And{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.ClauseSet{Keyword: "AND", Value: "lhs = 10", Parens: true},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.And.Build()
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestAnd_Build_Fail(t *testing.T) {
	a := &clause.And{Expr: "column = ?"}
	_, err := a.Build()
	if err == nil {
		t.Errorf("Error was not occurred")
	}
}
