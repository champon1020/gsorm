package expr_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestGroupBy_String(t *testing.T) {
	testCases := []struct {
		GroupBy *expr.GroupBy
		Result  string
	}{
		{
			&expr.GroupBy{Columns: []syntax.Column{{Name: "column"}}},
			`GROUP BY("column")`,
		},
		{
			&expr.GroupBy{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			`GROUP BY("column AS c")`,
		},
		{
			&expr.GroupBy{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"},
			}},
			`GROUP BY("column1 AS c1", "column2 AS c2")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.GroupBy.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestGroupBy_Build(t *testing.T) {
	testCases := []struct {
		GroupBy *expr.GroupBy
		Result  *syntax.StmtSet
	}{
		{
			&expr.GroupBy{Columns: []syntax.Column{{Name: "column"}}},
			&syntax.StmtSet{Clause: "GROUP BY", Value: "column"},
		},
		{
			&expr.GroupBy{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			&syntax.StmtSet{Clause: "GROUP BY", Value: "column AS c"},
		},
		{
			&expr.GroupBy{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"},
			}},
			&syntax.StmtSet{Clause: "GROUP BY", Value: "column1 AS c1, column2 AS c2"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.GroupBy.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewGroupBy(t *testing.T) {
	testCases := []struct {
		Columns []string
		Result  *expr.GroupBy
	}{
		{
			[]string{"column"},
			&expr.GroupBy{Columns: []syntax.Column{{Name: "column"}}},
		},
		{
			[]string{"column AS c"},
			&expr.GroupBy{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
		},
		{
			[]string{"column1 AS c1", "column2 AS c2"},
			&expr.GroupBy{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"}},
			},
		},
	}

	for _, testCase := range testCases {
		res := expr.NewGroupBy(testCase.Columns)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
