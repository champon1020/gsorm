package syntax_test

import (
	"errors"
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSet_Name(t *testing.T) {
	s := &syntax.Set{}
	assert.Equal(t, "SET", syntax.SetName(s))
}

func TestSet_AddEq(t *testing.T) {
	testCases := []struct {
		LHS    string
		RHS    interface{}
		Set    *syntax.Set
		Result *syntax.Set
	}{
		{"lhs", "rhs", &syntax.Set{}, &syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs", RHS: "rhs"}}}},
		{
			"lhs2",
			"rhs2",
			&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs1", RHS: "rhs1"}}},
			&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
		},
	}

	for _, testCase := range testCases {
		syntax.SetAddEq(testCase.Set, testCase.LHS, testCase.RHS)
		if diff := cmp.Diff(testCase.Set, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

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
			internal.PrintTestDiff(t, diff)
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
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewSet_Fail(t *testing.T) {
	testCases := []struct {
		LHS   []string
		RHS   []interface{}
		Error error
	}{
		{
			[]string{"lhs1", "lhs2"},
			[]interface{}{10},
			internal.NewError(
				syntax.OpNewSet,
				internal.KindBasic,
				errors.New("Length is different between lhs and rhs"),
			),
		},
	}

	for _, testCase := range testCases {
		_, err := syntax.NewSet(testCase.LHS, testCase.RHS)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}

		e, ok := err.(*internal.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}

		if diff := internal.CmpError(*e, *testCase.Error.(*internal.Error)); diff != "" {
			t.Errorf(diff)
		}
	}
}
