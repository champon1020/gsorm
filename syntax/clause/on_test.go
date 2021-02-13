package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOn_String(t *testing.T) {
	testCases := []struct {
		On     *clause.On
		Result string
	}{
		{
			&clause.On{Expr: "lhs = rhs"},
			`ON("lhs = rhs")`,
		},
		{
			&clause.On{Expr: "lhs = ?", Values: []interface{}{10}},
			`ON("lhs = ?", 10)`,
		},
		{
			&clause.On{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			`ON("lhs1 = ? AND lhs2 = ?", 10, "str")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.On.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestOn_Build(t *testing.T) {
	testCases := []struct {
		On     *clause.On
		Result *syntax.StmtSet
	}{
		{
			&clause.On{Expr: "lhs = rhs"},
			&syntax.StmtSet{Keyword: "ON", Value: "lhs = rhs"},
		},
		{
			&clause.On{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Keyword: "ON", Value: "lhs = 10"},
		},
		{
			&clause.On{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			&syntax.StmtSet{Keyword: "ON", Value: `lhs1 = 10 AND lhs2 = "str"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.On.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
