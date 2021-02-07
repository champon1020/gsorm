package cmd_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/cmd"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestInsert_String(t *testing.T) {
	testCases := []struct {
		Insert *cmd.Insert
		Result string
	}{
		{
			&cmd.Insert{Table: syntax.Table{Name: "table"}},
			`INSERT INTO("table")`,
		},
		{
			&cmd.Insert{Table: syntax.Table{Name: "table", Alias: "t"}},
			`INSERT INTO("table AS t")`,
		},
		{
			&cmd.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column"}},
			},
			`INSERT INTO("table AS t", "column")`,
		},
		{
			&cmd.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column", Alias: "c"}},
			},
			`INSERT INTO("table AS t", "column AS c")`,
		},
		{
			&cmd.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column1", Alias: "c1"}, {Name: "column2", Alias: "c2"}},
			},
			`INSERT INTO("table AS t", "column1 AS c1", "column2 AS c2")`,
		},
	}
	for _, testCase := range testCases {
		res := testCase.Insert.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestInsert_Build(t *testing.T) {
	testCases := []struct {
		Insert *cmd.Insert
		Result *syntax.StmtSet
	}{
		{
			&cmd.Insert{Table: syntax.Table{Name: "table"}},
			&syntax.StmtSet{Clause: "INSERT INTO", Value: "table"},
		},
		{
			&cmd.Insert{Table: syntax.Table{Name: "table", Alias: "t"}},
			&syntax.StmtSet{Clause: "INSERT INTO", Value: "table AS t"},
		},
		{
			&cmd.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column"}},
			},
			&syntax.StmtSet{Clause: "INSERT INTO", Value: "table AS t (column)"},
		},
		{
			&cmd.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column", Alias: "c"}},
			},
			&syntax.StmtSet{Clause: "INSERT INTO", Value: "table AS t (column AS c)"},
		},
		{
			&cmd.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column1", Alias: "c1"}, {Name: "column2", Alias: "c2"}},
			},
			&syntax.StmtSet{Clause: "INSERT INTO", Value: "table AS t (column1 AS c1, column2 AS c2)"},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Insert.Build()
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewInsert(t *testing.T) {
	testCases := []struct {
		Table  string
		Cols   []string
		Result *cmd.Insert
	}{
		{
			"table",
			[]string{},
			&cmd.Insert{Table: syntax.Table{Name: "table"}},
		},
		{
			"table AS t",
			[]string{},
			&cmd.Insert{Table: syntax.Table{Name: "table", Alias: "t"}},
		},
		{
			"table AS t",
			[]string{"column"},
			&cmd.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column"}},
			},
		},
		{
			"table AS t",
			[]string{"column AS c"},
			&cmd.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column", Alias: "c"}},
			},
		},
		{
			"table AS t",
			[]string{"column1 AS c1", "column2 AS c2"},
			&cmd.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column1", Alias: "c1"}, {Name: "column2", Alias: "c2"}},
			},
		},
	}

	for _, testCase := range testCases {
		res := cmd.NewInsert(testCase.Table, testCase.Cols)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
