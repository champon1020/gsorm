package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestGroupBy_Name(t *testing.T) {
	g := new(syntax.GroupBy)
	assert.Equal(t, "GROUP BY", syntax.GroupByName(g))
}

func TestGroupBy_AddColumn(t *testing.T) {
	testCases := []struct {
		GroupBy *syntax.GroupBy
		Column  string
		Result  *syntax.GroupBy
	}{
		{
			&syntax.GroupBy{},
			"column",
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column"}}},
		},
		{
			&syntax.GroupBy{},
			"column AS c",
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
		},
		{
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column1"}}},
			"column2 AS c2",
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column1"}, {Name: "column2", Alias: "c2"}}},
		},
	}

	for _, testCase := range testCases {
		syntax.GroupByAddColumn(testCase.GroupBy, testCase.Column)
		if diff := cmp.Diff(testCase.GroupBy, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestGroupBy_Build(t *testing.T) {
	testCases := []struct {
		GroupBy *syntax.GroupBy
		Result  *syntax.StmtSet
	}{
		{
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column"}}},
			&syntax.StmtSet{Clause: "GROUP BY", Value: "column"},
		},
		{
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			&syntax.StmtSet{Clause: "GROUP BY", Value: "column AS c"},
		},
		{
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column1"}, {Name: "column2", Alias: "c2"}}},
			&syntax.StmtSet{Clause: "GROUP BY", Value: "column1, column2 AS c2"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.GroupBy.Build()
		if err != nil {
			t.Errorf("%v\n", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewGroupBy(t *testing.T) {
	testCases := []struct {
		Columns []string
		Result  *syntax.GroupBy
	}{
		{
			[]string{"column"},
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column"}}},
		},
		{
			[]string{"column AS c"},
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
		},
		{
			[]string{"column1", "column2 AS c2"},
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column1"}, {Name: "column2", Alias: "c2"}}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewGroupBy(testCase.Columns)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
