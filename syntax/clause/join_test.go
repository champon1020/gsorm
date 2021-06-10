package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestJoin_AddTable(t *testing.T) {
	{
		j := &clause.Join{}
		table := "table"
		j.AddTable(table)

		assert.Equal(t, j.Table, syntax.Table{Name: "table"})
	}
	{
		j := &clause.Join{}
		table := "table as t"
		j.AddTable(table)

		assert.Equal(t, j.Table, syntax.Table{Name: "table", Alias: "t"})
	}
}

func TestJoin_String(t *testing.T) {
	testCases := []struct {
		Join   *clause.Join
		Result string
	}{
		{
			&clause.Join{Table: syntax.Table{Name: "table"}, Type: clause.InnerJoin},
			`Join("table")`,
		},
		{
			&clause.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: clause.LeftJoin},
			`LeftJoin("table AS t")`,
		},
		{
			&clause.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: clause.RightJoin},
			`RightJoin("table AS t")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Join.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestJoin_Build(t *testing.T) {
	testCases := []struct {
		Join   *clause.Join
		Result *syntax.ClauseSet
	}{
		{
			&clause.Join{Table: syntax.Table{Name: "table"}, Type: clause.InnerJoin},
			&syntax.ClauseSet{Keyword: "INNER JOIN", Value: "table"},
		},
		{
			&clause.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: clause.LeftJoin},
			&syntax.ClauseSet{Keyword: "LEFT JOIN", Value: "table AS t"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Join.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
