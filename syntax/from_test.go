package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestFrom_Name(t *testing.T) {
	f := &syntax.From{}
	assert.Equal(t, "FROM", syntax.FromName(f))
}

func TestFrom_AddTable(t *testing.T) {
	testCases := []struct {
		Table  string
		From   *syntax.From
		Result *syntax.From
	}{
		{
			"column",
			&syntax.From{},
			&syntax.From{Tables: []syntax.Table{{Name: "column"}}},
		},
		{
			"column AS c",
			&syntax.From{},
			&syntax.From{Tables: []syntax.Table{{Name: "column", Alias: "c"}}},
		},
		{
			"column2",
			&syntax.From{Tables: []syntax.Table{{Name: "column1"}}},
			&syntax.From{Tables: []syntax.Table{{Name: "column1"}, {Name: "column2"}}},
		},
	}

	for _, testCase := range testCases {
		syntax.FromAddTable(testCase.From, testCase.Table)
		if diff := cmp.Diff(testCase.From, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestFrom_Build(t *testing.T) {
	testCases := []struct {
		From   *syntax.From
		Result *syntax.StmtSet
	}{
		{
			&syntax.From{Tables: []syntax.Table{{Name: "table"}}},
			&syntax.StmtSet{Clause: "FROM", Value: "table"},
		},
		{
			&syntax.From{Tables: []syntax.Table{{Name: "table", Alias: "t"}}},
			&syntax.StmtSet{Clause: "FROM", Value: "table AS t"},
		},
		{
			&syntax.From{Tables: []syntax.Table{{Name: "table1"}, {Name: "table2"}}},
			&syntax.StmtSet{Clause: "FROM", Value: "table1, table2"},
		},
		{
			&syntax.From{Tables: []syntax.Table{{Name: "table1", Alias: "t1"}, {Name: "table2", Alias: "t2"}}},
			&syntax.StmtSet{Clause: "FROM", Value: "table1 AS t1, table2 AS t2"},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.From.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewFrom(t *testing.T) {
	testCases := []struct {
		Tables []string
		Result *syntax.From
	}{
		{
			[]string{"table"},
			&syntax.From{Tables: []syntax.Table{{Name: "table"}}},
		},
		{
			[]string{"table AS t"},
			&syntax.From{Tables: []syntax.Table{{Name: "table", Alias: "t"}}},
		},
		{
			[]string{"table1", "table2"},
			&syntax.From{Tables: []syntax.Table{{Name: "table1"}, {Name: "table2"}}},
		},
		{
			[]string{"table1 AS t1", "table2 AS t2"},
			&syntax.From{Tables: []syntax.Table{{Name: "table1", Alias: "t1"}, {Name: "table2", Alias: "t2"}}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewFrom(testCase.Tables)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
