package syntax_test

import (
	"testing"

	"github.com/champon1020/minigorm/syntax"
	"github.com/google/go-cmp/cmp"
)

func TestSelect_Build(t *testing.T) {
	testCases := []struct {
		Select *syntax.Select
		Result *syntax.StmtSet
	}{
		{
			&syntax.Select{Columns: []syntax.Column{{Name: "column"}}},
			&syntax.StmtSet{Clause: "SELECT", Value: "column"},
		},
		{
			&syntax.Select{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			&syntax.StmtSet{Clause: "SELECT", Value: "column AS c"},
		},
		{
			&syntax.Select{Columns: []syntax.Column{{Name: "column1"}, {Name: "column2"}}},
			&syntax.StmtSet{Clause: "SELECT", Value: "column1, column2"},
		},
		{
			&syntax.Select{Columns: []syntax.Column{{Name: "column1", Alias: "c1"}, {Name: "column2"}}},
			&syntax.StmtSet{Clause: "SELECT", Value: "column1 AS c1, column2"},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Select.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestNewSelect(t *testing.T) {
	testCases := []struct {
		Cols   []string
		Result *syntax.Select
	}{
		{
			[]string{"column1"},
			&syntax.Select{Columns: []syntax.Column{{Name: "column1"}}},
		},
		{
			[]string{"column1 AS c1"},
			&syntax.Select{Columns: []syntax.Column{{Name: "column1", Alias: "c1"}}},
		},
		{
			[]string{"column1 AS c1", "column2"},
			&syntax.Select{Columns: []syntax.Column{{Name: "column1", Alias: "c1"}, {Name: "column2"}}},
		},
		{
			[]string{"column1", "column2"},
			&syntax.Select{Columns: []syntax.Column{{Name: "column1"}, {Name: "column2"}}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewSelect(testCase.Cols)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}
