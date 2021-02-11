package expr_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm/errors"
	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSet_String(t *testing.T) {
	testCases := []struct {
		Set    *expr.Set
		Result string
	}{
		{
			&expr.Set{Eqs: []syntax.Eq{{LHS: "lhs", RHS: "rhs"}}},
			`SET("rhs")`,
		},
		{
			&expr.Set{Eqs: []syntax.Eq{
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
		Set    *expr.Set
		Result *syntax.StmtSet
	}{
		{
			&expr.Set{Eqs: []syntax.Eq{{LHS: "lhs", RHS: "rhs"}}},
			&syntax.StmtSet{Clause: "SET", Value: `lhs = "rhs"`},
		},
		{
			&expr.Set{Eqs: []syntax.Eq{
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
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestSet_Build_Fail(t *testing.T) {
	testCases := []struct {
		Set *expr.Set
	}{
		{&expr.Set{Eqs: []syntax.Eq{{LHS: "lhs", RHS: time.Now()}}}},
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
		Result *expr.Set
	}{
		{
			[]string{"lhs"},
			[]interface{}{10},
			&expr.Set{Eqs: []syntax.Eq{{LHS: "lhs", RHS: 10}}},
		},
		{
			[]string{"lhs1", "lhs2"},
			[]interface{}{10, "rhs2"},
			&expr.Set{Eqs: []syntax.Eq{
				{LHS: "lhs1", RHS: 10},
				{LHS: "lhs2", RHS: "rhs2"},
			}},
		},
	}

	for _, testCase := range testCases {
		res, err := expr.NewSet(testCase.LHS, testCase.RHS)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
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
			errors.New("Length is different between lhs and rhs", errors.InvalidValueError),
		},
	}

	for _, testCase := range testCases {
		_, err := expr.NewSet(testCase.LHS, testCase.RHS)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}
		actualError, ok := err.(*errors.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}
		resultError := testCase.Error.(*errors.Error)
		if !resultError.Is(actualError) {
			t.Errorf("Different error was occurred")
			t.Errorf("  Expected: %s, Code: %d", resultError.Error(), resultError.Code)
			t.Errorf("  Actual:   %s, Code: %d", actualError.Error(), actualError.Code)
		}
	}
}
