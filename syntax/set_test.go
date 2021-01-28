package syntax_test

import (
	"errors"
	"testing"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSet_String(t *testing.T) {
	testCases := []struct {
		Set    *syntax.Set
		Result string
	}{
		{
			&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs", RHS: "rhs"}}},
			`SET("rhs")`,
		},
		{
			&syntax.Set{Eqs: []syntax.Eq{
				{LHS: "lhs", RHS: "rhs"},
				{LHS: "lhs", RHS: 10},
			}},
			`SET("rhs", 10)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Set.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestSet_Build(t *testing.T) {
	testCases := []struct {
		Set    *syntax.Set
		Result *syntax.StmtSet
	}{
		{
			&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs", RHS: "rhs"}}},
			&syntax.StmtSet{Clause: "SET", Value: `lhs = "rhs"`},
		},
		{
			&syntax.Set{Eqs: []syntax.Eq{
				{LHS: "lhs1", RHS: 10},
				{LHS: "lhs2", RHS: "rhs2"},
			}},
			&syntax.StmtSet{Clause: "SET", Value: `lhs1 = 10, lhs2 = "rhs2"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Set.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestSet_Build_Fail(t *testing.T) {
	testCases := []struct {
		Set *syntax.Set
	}{
		{&syntax.Set{Eqs: []syntax.Eq{{LHS: "lhs", RHS: time.Now()}}}},
	}

	for _, testCase := range testCases {
		_, err := testCase.Set.Build()
		if err == nil {
			t.Errorf("Error was not occurred")
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
			[]interface{}{10, "rhs2"},
			&syntax.Set{Eqs: []syntax.Eq{
				{LHS: "lhs1", RHS: 10},
				{LHS: "lhs2", RHS: "rhs2"},
			}},
		},
	}

	for _, testCase := range testCases {
		res, err := syntax.NewSet(testCase.LHS, testCase.RHS)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
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
		if diff := internal.CmpError(e, testCase.Error.(*internal.Error)); diff != "" {
			t.Errorf(diff)
		}
	}
}
