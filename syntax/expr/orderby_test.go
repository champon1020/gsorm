package expr_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOrderBy_String(t *testing.T) {
	testCases := []struct {
		OrderBy *expr.OrderBy
		Result  string
	}{
		{
			&expr.OrderBy{Column: "column"},
			`ORDER BY("column", false)`,
		},
		{
			&expr.OrderBy{Column: "column", Desc: true},
			`ORDER BY("column", true)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.OrderBy.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestOrderBy_Build(t *testing.T) {
	testCases := []struct {
		OrderBy *expr.OrderBy
		Result  *syntax.StmtSet
	}{
		{
			&expr.OrderBy{Column: "column"},
			&syntax.StmtSet{Clause: "ORDER BY", Value: "column"},
		},
		{
			&expr.OrderBy{Column: "column", Desc: true},
			&syntax.StmtSet{Clause: "ORDER BY", Value: "column DESC"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.OrderBy.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewOrderBy(t *testing.T) {
	testCases := []struct {
		Column string
		Desc   bool
		Result *expr.OrderBy
	}{
		{
			"column",
			false,
			&expr.OrderBy{Column: "column", Desc: false},
		},
		{
			"column",
			true,
			&expr.OrderBy{Column: "column", Desc: true},
		},
	}

	for _, testCase := range testCases {
		res := expr.NewOrderBy(testCase.Column, testCase.Desc)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
