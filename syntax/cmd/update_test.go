package cmd_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/cmd"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestUpdate_String(t *testing.T) {
	testCases := []struct {
		Update *cmd.Update
		Result string
	}{
		{
			&cmd.Update{Table: syntax.Table{Name: "table"}},
			`UPDATE("table")`,
		},
		{
			&cmd.Update{Table: syntax.Table{Name: "table", Alias: "t"}},
			`UPDATE("table AS t")`,
		},
		{
			&cmd.Update{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []string{"column"},
			},
			`UPDATE("table AS t", "column")`,
		},
		{
			&cmd.Update{
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
		Update *cmd.Update
		Result *syntax.StmtSet
	}{
		{
			&cmd.Update{Table: syntax.Table{Name: "table"}},
			&syntax.StmtSet{Clause: "UPDATE", Value: "table"},
		},
		{
			&cmd.Update{Table: syntax.Table{Name: "table", Alias: "t"}},
			&syntax.StmtSet{Clause: "UPDATE", Value: "table AS t"},
		},
		{
			&cmd.Update{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []string{"column"},
			},
			&syntax.StmtSet{Clause: "UPDATE", Value: "table AS t"},
		},
		{
			&cmd.Update{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []string{"column1", "column2"},
			},
			&syntax.StmtSet{Clause: "UPDATE", Value: "table AS t"},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Update.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewUpdate(t *testing.T) {
	testCases := []struct {
		Table   string
		Columns []string
		Result  *cmd.Update
	}{
		{
			"table",
			[]string{},
			&cmd.Update{Table: syntax.Table{Name: "table"}, Columns: []string{}},
		},
		{
			"table AS t",
			[]string{},
			&cmd.Update{Table: syntax.Table{Name: "table", Alias: "t"}, Columns: []string{}},
		},
		{
			"table AS t",
			[]string{"column"},
			&cmd.Update{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []string{"column"},
			},
		},
		{
			"table AS t",
			[]string{"column1", "column2"},
			&cmd.Update{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []string{"column1", "column2"},
			},
		},
	}

	for _, testCase := range testCases {
		res := cmd.NewUpdate(testCase.Table, testCase.Columns)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
