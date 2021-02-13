package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
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
			`AND("lhs = rhs")`,
		},
		{
			&clause.And{Expr: "lhs = ?", Values: []interface{}{10}},
			`AND("lhs = ?", 10)`,
		},
		{
			&clause.And{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			`AND("lhs1 = ? AND lhs2 = ?", 10, "str")`,
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
		Result *syntax.StmtSet
	}{
		{
			&clause.And{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Keyword: "AND", Value: "lhs = 10", Parens: true},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.And.Build()
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
