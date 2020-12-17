package syntax_test

import (
	"testing"

	"github.com/champon1020/minigorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSet_Build(t *testing.T) {
	testCases := []struct {
		Set    *syntax.Set
		Result *syntax.StmtSet
	}{
		{
			&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs", RHS: "rhs"}}},
			&syntax.StmtSet{Clause: "SET", Value: "lhs = rhs"},
		},
		{
			&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
			&syntax.StmtSet{Clause: "SET", Value: "lhs1 = rhs1, lhs2 = rhs2"},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Set.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestNewSet(t *testing.T) {
	testCases := []struct {
		LHS    []string
		RHS    []interface{}
		Result *syntax.Set
	}{
		{
			[]string{"lhs"},
			[]interface{}{10},
			&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs", RHS: 10}}},
		},
		{
			[]string{"lhs1", "lhs2"},
			[]interface{}{10, 100},
			&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs1", RHS: 10}, {LHS: "lhs2", RHS: 100}}},
		},
	}

	for _, testCase := range testCases {
		res, _ := syntax.NewSet(testCase.LHS, testCase.RHS)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestNewSet_Fail(t *testing.T) {
	testCases := []struct {
		LHS       []string
		RHS       []interface{}
		ErrorCode int
	}{
		{
			[]string{"lhs1", "lhs2"},
			[]interface{}{10},
			syntax.ErrInvalid,
		},
	}

	for _, testCase := range testCases {
		_, err := syntax.NewSet(testCase.LHS, testCase.RHS)
		assert.Equal(t, testCase.ErrorCode, err.(syntax.Error).Code)
	}
}
