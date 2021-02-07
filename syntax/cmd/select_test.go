package cmd_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/cmd"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSelect_String(t *testing.T) {
	testCases := []struct {
		Select *cmd.Select
		Result string
	}{
		{
			&cmd.Select{Columns: []syntax.Column{{Name: "column"}}},
			`SELECT("column")`,
		},
		{
			&cmd.Select{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			`SELECT("column AS c")`,
		},
		{
			&cmd.Select{Columns: []syntax.Column{
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
		Select *cmd.Select
		Result *syntax.StmtSet
	}{
		{
			&cmd.Select{Columns: []syntax.Column{{Name: "column"}}},
			&syntax.StmtSet{Clause: "SELECT", Value: "column"},
		},
		{
			&cmd.Select{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
			&syntax.StmtSet{Clause: "SELECT", Value: "column AS c"},
		},
		{
			&cmd.Select{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"},
			}},
			&syntax.StmtSet{Clause: "SELECT", Value: "column1 AS c1, column2 AS c2"},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Select.Build()
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewSelect(t *testing.T) {
	testCases := []struct {
		Cols   []string
		Result *cmd.Select
	}{
		{
			[]string{"column1"},
			&cmd.Select{Columns: []syntax.Column{{Name: "column1"}}},
		},
		{
			[]string{"column AS c"},
			&cmd.Select{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
		},
		{
			[]string{"column1 AS c1", "column2 AS c2"},
			&cmd.Select{Columns: []syntax.Column{
				{Name: "column1", Alias: "c1"},
				{Name: "column2", Alias: "c2"},
			}},
		},
	}

	for _, testCase := range testCases {
		res := cmd.NewSelect(testCase.Cols)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
