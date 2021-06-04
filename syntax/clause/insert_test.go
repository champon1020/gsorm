package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestInsert_String(t *testing.T) {
	testCases := []struct {
		Insert *clause.Insert
		Result string
	}{
		{
			&clause.Insert{Table: syntax.Table{Name: "table"}},
			`Insert("table")`,
		},
		{
			&clause.Insert{Table: syntax.Table{Name: "table", Alias: "t"}},
			`Insert("table AS t")`,
		},
		{
			&clause.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column"}},
			},
			`Insert("table AS t", "column")`,
		},
		{
			&clause.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column", Alias: "c"}},
			},
			`Insert("table AS t", "column AS c")`,
		},
		{
			&clause.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column1", Alias: "c1"}, {Name: "column2", Alias: "c2"}},
			},
			`Insert("table AS t", "column1 AS c1", "column2 AS c2")`,
		},
	}
	for _, testCase := range testCases {
		res := testCase.Insert.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestInsert_Build(t *testing.T) {
	testCases := []struct {
		Insert *clause.Insert
		Result *syntax.StmtSet
	}{
		{
			&clause.Insert{Table: syntax.Table{Name: "table"}},
			&syntax.StmtSet{Keyword: "INSERT INTO", Value: "table"},
		},
		{
			&clause.Insert{Table: syntax.Table{Name: "table", Alias: "t"}},
			&syntax.StmtSet{Keyword: "INSERT INTO", Value: "table AS t"},
		},
		{
			&clause.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column"}},
			},
			&syntax.StmtSet{Keyword: "INSERT INTO", Value: "table AS t (column)"},
		},
		{
			&clause.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column", Alias: "c"}},
			},
			&syntax.StmtSet{Keyword: "INSERT INTO", Value: "table AS t (column AS c)"},
		},
		{
			&clause.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column1", Alias: "c1"}, {Name: "column2", Alias: "c2"}},
			},
			&syntax.StmtSet{Keyword: "INSERT INTO", Value: "table AS t (column1 AS c1, column2 AS c2)"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Insert.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestInsert_AddTable(t *testing.T) {
	{
		j := &clause.Insert{}
		table := "table"
		j.AddTable(table)

		assert.Equal(t, j.Table, syntax.Table{Name: "table"})
	}
	{
		j := &clause.Insert{}
		table := "table as t"
		j.AddTable(table)

		assert.Equal(t, j.Table, syntax.Table{Name: "table", Alias: "t"})
	}
}

func TestInsert_AddColumn(t *testing.T) {
	g := &clause.Insert{}
	c := []string{"column1", "column2 as c"}
	g.AddColumns(c...)
	assert.Equal(t, []syntax.Column{{Name: "column1"}, {Name: "column2", Alias: "c"}}, g.Columns)
}
