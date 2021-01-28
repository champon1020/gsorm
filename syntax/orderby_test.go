package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOrderBy_String(t *testing.T) {
	testCases := []struct {
		OrderBy *syntax.OrderBy
		Result  string
	}{
		{
			&syntax.OrderBy{Column: "column"},
			`ORDER BY("column", false)`,
		},
		{
			&syntax.OrderBy{Column: "column", Desc: true},
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
		OrderBy *syntax.OrderBy
		Result  *syntax.StmtSet
	}{
		{
			&syntax.OrderBy{Column: "column"},
			&syntax.StmtSet{Clause: "ORDER BY", Value: "column"},
		},
		{
			&syntax.OrderBy{Column: "column", Desc: true},
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
		Result *syntax.OrderBy
	}{
		{
			"column",
			false,
			&syntax.OrderBy{Column: "column", Desc: false},
		},
		{
			"column",
			true,
			&syntax.OrderBy{Column: "column", Desc: true},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewOrderBy(testCase.Column, testCase.Desc)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
