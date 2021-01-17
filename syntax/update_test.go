package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestUpdate_Query(t *testing.T) {
	u := &syntax.Update{}
	assert.Equal(t, "UPDATE", syntax.UpdateQuery(u))
}

func TestUpdate_AddTable(t *testing.T) {
	testCases := []struct {
		Table  string
		Update *syntax.Update
		Result *syntax.Update
	}{
		{
			"table",
			&syntax.Update{},
			&syntax.Update{Table: syntax.Table{Name: "table"}},
		},
		{
			"table AS t",
			&syntax.Update{},
			&syntax.Update{Table: syntax.Table{Name: "table", Alias: "t"}},
		},
		{
			"table2",
			&syntax.Update{Table: syntax.Table{Name: "table1", Alias: "t1"}},
			&syntax.Update{Table: syntax.Table{Name: "table2"}},
		},
	}

	for _, testCase := range testCases {
		syntax.UpdateAddTable(testCase.Update, testCase.Table)
		if diff := cmp.Diff(testCase.Update, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestUpdate_Build(t *testing.T) {
	testCases := []struct {
		Update *syntax.Update
		Result *syntax.StmtSet
	}{
		{
			&syntax.Update{Table: syntax.Table{Name: "table"}},
			&syntax.StmtSet{Clause: "UPDATE", Value: "table"},
		},
		{
			&syntax.Update{Table: syntax.Table{Name: "table", Alias: "t"}},
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
		Result  *syntax.Update
	}{
		{
			"table",
			[]string{},
			&syntax.Update{Table: syntax.Table{Name: "table"}, Columns: []string{}},
		},
		{
			"table AS t",
			[]string{"column1", "column2"},
			&syntax.Update{Table: syntax.Table{Name: "table", Alias: "t"}, Columns: []string{"column1", "column2"}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewUpdate(testCase.Table, testCase.Columns)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
