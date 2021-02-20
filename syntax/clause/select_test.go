package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSelect_String(t *testing.T) {
	testCases := []struct {
		Select *clause.Select
		Result string
	}{
		{
			&clause.Select{Columns: []syntax.Column{{Name: "column"}}},
			`SELECT("column")`,
		},
		{
			&clause.Select{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			`SELECT("column AS c")`,
		},
		{
			&clause.Select{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"},
			}},
			`SELECT("column1 AS c1", "column2 AS c2")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Select.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestSelect_Build(t *testing.T) {
	testCases := []struct {
		Select *clause.Select
		Result *syntax.StmtSet
	}{
		{
			&clause.Select{Columns: []syntax.Column{{Name: "column"}}},
			&syntax.StmtSet{Keyword: "SELECT", Value: "column"},
		},
		{
			&clause.Select{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			&syntax.StmtSet{Keyword: "SELECT", Value: "column AS c"},
		},
		{
			&clause.Select{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"},
			}},
			&syntax.StmtSet{Keyword: "SELECT", Value: "column1 AS c1, column2 AS c2"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Select.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestNewSelect(t *testing.T) {
	testCases := []struct {
		Cols   []string
		Result *clause.Select
	}{
		{
			[]string{"column1"},
			&clause.Select{Columns: []syntax.Column{{Name: "column1"}}},
		},
		{
			[]string{"column AS c"},
			&clause.Select{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
		},
		{
			[]string{"column1 AS c1", "column2 AS c2"},
			&clause.Select{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"},
			}},
		},
	}

	for _, testCase := range testCases {
		res := clause.NewSelect(testCase.Cols)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}