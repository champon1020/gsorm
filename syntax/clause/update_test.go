package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
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
			`UPDATE("table")`,
		},
		{
			&clause.Update{Table: syntax.Table{Name: "table", Alias: "t"}},
			`UPDATE("table AS t")`,
		},
		{
			&clause.Update{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []string{"column"},
			},
			`UPDATE("table AS t", "column")`,
		},
		{
			&clause.Update{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []string{"column1", "column2"},
			},
			`UPDATE("table AS t", "column1", "column2")`,
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
		{
			&clause.Update{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []string{"column"},
			},
			&syntax.StmtSet{Keyword: "UPDATE", Value: "table AS t"},
		},
		{
			&clause.Update{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []string{"column1", "column2"},
			},
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
