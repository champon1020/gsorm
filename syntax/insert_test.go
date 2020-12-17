package syntax_test

import (
	"testing"

	"github.com/champon1020/minigorm/syntax"
	"github.com/google/go-cmp/cmp"
)

func TestInsert_Build(t *testing.T) {
	testCases := []struct {
		Insert *syntax.Insert
		Result *syntax.StmtSet
	}{
		{
			&syntax.Insert{Table: syntax.Table{Name: "table"}, Columns: []syntax.Column{{Name: "column1"}}},
			&syntax.StmtSet{Clause: "INSERT INTO", Value: "table (column1)"},
		},
		{
			&syntax.Insert{
				Table:   syntax.Table{Name: "table"},
				Columns: []syntax.Column{{Name: "column1"}, {Name: "column2"}},
			},
			&syntax.StmtSet{Clause: "INSERT INTO", Value: "table (column1, column2)"},
		},
		{
			&syntax.Insert{
				Table:   syntax.Table{Name: "table"},
				Columns: []syntax.Column{{Name: "column1", Alias: "c1"}},
			},
			&syntax.StmtSet{Clause: "INSERT INTO", Value: "table (column1 AS c1)"},
		},
		{
			&syntax.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column1", Alias: "c1"}, {Name: "column2", Alias: "c2"}},
			},
			&syntax.StmtSet{Clause: "INSERT INTO", Value: "table AS t (column1 AS c1, column2 AS c2)"},
		},
	}

	for _, testCase := range testCases {
		res := testCase.Insert.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestNewInsert(t *testing.T) {
	testCases := []struct {
		Table  string
		Cols   []string
		Result *syntax.Insert
	}{
		{
			"table",
			[]string{"column1"},
			&syntax.Insert{Table: syntax.Table{Name: "table"}, Columns: []syntax.Column{{Name: "column1"}}},
		},
		{
			"table AS t",
			[]string{"column1 AS c1"},
			&syntax.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column1", Alias: "c1"}},
			},
		},
		{
			"table AS t",
			[]string{"column1 AS c1", "column2 AS c2"},
			&syntax.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column1", Alias: "c1"}, {Name: "column2", Alias: "c2"}},
			},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewInsert(testCase.Table, testCase.Cols)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}
