package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
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
			`INSERT INTO("table")`,
		},
		{
			&clause.Insert{Table: syntax.Table{Name: "table", Alias: "t"}},
			`INSERT INTO("table AS t")`,
		},
		{
			&clause.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column"}},
			},
			`INSERT INTO("table AS t", "column")`,
		},
		{
			&clause.Insert{
				Table:   syntax.Table{Name: "table", Alias: "t"},
				Columns: []syntax.Column{{Name: "column", Alias: "c"}},
			},
			`INSERT INTO("table AS t", "column AS c")`,
		},
		{
			&clause.Insert{
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
