package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestInsert_Query(t *testing.T) {
	i := &syntax.Insert{}
	assert.Equal(t, "INSERT INTO", syntax.InsertQuery(i))
}

func TestInsert_AddTable(t *testing.T) {
	testCases := []struct {
		Table  string
		Insert *syntax.Insert
		Result *syntax.Insert
	}{
		{
			"table",
			&syntax.Insert{},
			&syntax.Insert{Table: syntax.Table{Name: "table"}},
		},
		{
			"table AS t",
			&syntax.Insert{},
			&syntax.Insert{Table: syntax.Table{Name: "table", Alias: "t"}},
		},
		{
			"table2",
			&syntax.Insert{Table: syntax.Table{Name: "table1", Alias: "t1"}},
			&syntax.Insert{Table: syntax.Table{Name: "table2"}},
		},
	}

	for _, testCase := range testCases {
		syntax.InsertAddTable(testCase.Insert, testCase.Table)
		if diff := cmp.Diff(testCase.Insert, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestInsert_AddColumn(t *testing.T) {
	testCases := []struct {
		Column string
		Insert *syntax.Insert
		Result *syntax.Insert
	}{
		{
			"column",
			&syntax.Insert{},
			&syntax.Insert{Columns: []syntax.Column{{Name: "column"}}},
		},
		{
			"column AS c",
			&syntax.Insert{},
			&syntax.Insert{Columns: []syntax.Column{{Name: "column", Alias: "c"}}},
		},
		{
			"column2",
			&syntax.Insert{Columns: []syntax.Column{{Name: "column1"}}},
			&syntax.Insert{Columns: []syntax.Column{{Name: "column1"}, {Name: "column2"}}},
		},
	}

	for _, testCase := range testCases {
		syntax.InsertAddColumn(testCase.Insert, testCase.Column)
		if diff := cmp.Diff(testCase.Insert, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

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
			internal.PrintTestDiff(t, diff)
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
			internal.PrintTestDiff(t, diff)
		}
	}
}
