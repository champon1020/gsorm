package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestGroupBy_String(t *testing.T) {
	testCases := []struct {
		GroupBy *syntax.GroupBy
		Result  string
	}{
		{
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column"}}},
			`GROUP BY("column")`,
		},
		{
			&syntax.GroupBy{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			`GROUP BY("column AS c")`,
		},
		{
			&syntax.GroupBy{Columns: []syntax.Column{
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
			&syntax.GroupBy{Columns: []syntax.Column{
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
			[]string{"column1 AS c1", "column2 AS c2"},
			&syntax.GroupBy{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"}},
			},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewGroupBy(testCase.Columns)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
