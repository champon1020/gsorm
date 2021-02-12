package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOrderBy_String(t *testing.T) {
	testCases := []struct {
		OrderBy *clause.OrderBy
		Result  string
	}{
		{
			&clause.OrderBy{Column: "column"},
			`ORDER BY("column", false)`,
		},
		{
			&clause.OrderBy{Column: "column", Desc: true},
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
		OrderBy *clause.OrderBy
		Result  *syntax.StmtSet
	}{
		{
			&clause.OrderBy{Column: "column"},
			&syntax.StmtSet{Keyword: "ORDER BY", Value: "column"},
		},
		{
			&clause.OrderBy{Column: "column", Desc: true},
			&syntax.StmtSet{Keyword: "ORDER BY", Value: "column DESC"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.OrderBy.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestNewOrderBy(t *testing.T) {
	testCases := []struct {
		Column string
		Desc   bool
		Result *clause.OrderBy
	}{
		{
			"column",
			false,
			&clause.OrderBy{Column: "column", Desc: false},
		},
		{
			"column",
			true,
			&clause.OrderBy{Column: "column", Desc: true},
		},
	}

	for _, testCase := range testCases {
		res := clause.NewOrderBy(testCase.Column, testCase.Desc)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
