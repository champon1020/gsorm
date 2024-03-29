package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/interfaces"
	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestFrom_String(t *testing.T) {
	testCases := []struct {
		From   *clause.From
		Result string
	}{
		{
			&clause.From{Tables: []syntax.Table{{Name: "table"}}},
			`From("table")`,
		},
		{
			&clause.From{Tables: []syntax.Table{{Name: "table", Alias: "t"}}},
			`From("table AS t")`,
		},
		{
			&clause.From{Tables: []syntax.Table{{Name: "table1", Alias: "t1"}, {Name: "table2", Alias: "t2"}}},
			`From("table1 AS t1", "table2 AS t2")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.From.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestFrom_Build(t *testing.T) {
	testCases := []struct {
		From   *clause.From
		Result interfaces.ClauseSet
	}{
		{
			&clause.From{Tables: []syntax.Table{{Name: "table"}}},
			&syntax.ClauseSet{Keyword: "FROM", Value: "table"},
		},
		{
			&clause.From{Tables: []syntax.Table{{Name: "table", Alias: "t"}}},
			&syntax.ClauseSet{Keyword: "FROM", Value: "table AS t"},
		},
		{
			&clause.From{Tables: []syntax.Table{{Name: "table1", Alias: "t1"}, {Name: "table2", Alias: "t2"}}},
			&syntax.ClauseSet{Keyword: "FROM", Value: "table1 AS t1, table2 AS t2"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.From.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
