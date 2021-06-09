package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestGroupBy_AddColumn(t *testing.T) {
	g := &clause.GroupBy{}
	c := "column as c"
	g.AddColumn(c)
	assert.Equal(t, []syntax.Column{{Name: "column", Alias: "c"}}, g.Columns)
}

func TestGroupBy_String(t *testing.T) {
	testCases := []struct {
		GroupBy *clause.GroupBy
		Result  string
	}{
		{
			&clause.GroupBy{Columns: []syntax.Column{{Name: "column"}}},
			`GroupBy("column")`,
		},
		{
			&clause.GroupBy{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			`GroupBy("column AS c")`,
		},
		{
			&clause.GroupBy{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"},
			}},
			`GroupBy("column1 AS c1", "column2 AS c2")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.GroupBy.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestGroupBy_Build(t *testing.T) {
	testCases := []struct {
		GroupBy *clause.GroupBy
		Result  *syntax.ClauseSet
	}{
		{
			&clause.GroupBy{Columns: []syntax.Column{{Name: "column"}}},
			&syntax.ClauseSet{Keyword: "GROUP BY", Value: "column"},
		},
		{
			&clause.GroupBy{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			&syntax.ClauseSet{Keyword: "GROUP BY", Value: "column AS c"},
		},
		{
			&clause.GroupBy{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"},
			}},
			&syntax.ClauseSet{Keyword: "GROUP BY", Value: "column1 AS c1, column2 AS c2"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.GroupBy.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
