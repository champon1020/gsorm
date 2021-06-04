package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestUpdate_String(t *testing.T) {
	testCases := []struct {
		Update *clause.Update
		Result string
	}{
		{
			&clause.Update{Table: syntax.Table{Name: "table"}},
			`Update("table")`,
		},
		{
			&clause.Update{Table: syntax.Table{Name: "table", Alias: "t"}},
			`Update("table AS t")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Update.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestUpdate_Build(t *testing.T) {
	testCases := []struct {
		Update *clause.Update
		Result *syntax.StmtSet
	}{
		{
			&clause.Update{Table: syntax.Table{Name: "table"}},
			&syntax.StmtSet{Keyword: "UPDATE", Value: "table"},
		},
		{
			&clause.Update{Table: syntax.Table{Name: "table", Alias: "t"}},
			&syntax.StmtSet{Keyword: "UPDATE", Value: "table AS t"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Update.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestUpdate_AddTable(t *testing.T) {
	{
		j := &clause.Update{}
		table := "table"
		j.AddTable(table)

		assert.Equal(t, j.Table, syntax.Table{Name: "table"})
	}
	{
		j := &clause.Update{}
		table := "table as t"
		j.AddTable(table)

		assert.Equal(t, j.Table, syntax.Table{Name: "table", Alias: "t"})
	}
}
