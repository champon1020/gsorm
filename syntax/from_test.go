package syntax_test

import (
	"testing"

	"github.com/champon1020/minigorm/syntax"
	"github.com/google/go-cmp/cmp"
)

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
			syntax.PrintTestDiff(t, diff)
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
			syntax.PrintTestDiff(t, diff)
		}
	}
}
