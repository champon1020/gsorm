package expr_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestFrom_String(t *testing.T) {
	testCases := []struct {
		From   *expr.From
		Result string
	}{
		{
			&expr.From{Tables: []syntax.Table{{Name: "table"}}},
			`FROM("table")`,
		},
		{
			&expr.From{Tables: []syntax.Table{{Name: "table", Alias: "t"}}},
			`FROM("table AS t")`,
		},
		{
			&expr.From{Tables: []syntax.Table{{Name: "table1", Alias: "t1"}, {Name: "table2", Alias: "t2"}}},
			`FROM("table1 AS t1", "table2 AS t2")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.From.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestFrom_Build(t *testing.T) {
	testCases := []struct {
		From   *expr.From
		Result *syntax.StmtSet
	}{
		{
			&expr.From{Tables: []syntax.Table{{Name: "table"}}},
			&syntax.StmtSet{Clause: "FROM", Value: "table"},
		},
		{
			&expr.From{Tables: []syntax.Table{{Name: "table", Alias: "t"}}},
			&syntax.StmtSet{Clause: "FROM", Value: "table AS t"},
		},
		{
			&expr.From{Tables: []syntax.Table{{Name: "table1", Alias: "t1"}, {Name: "table2", Alias: "t2"}}},
			&syntax.StmtSet{Clause: "FROM", Value: "table1 AS t1, table2 AS t2"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.From.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewFrom(t *testing.T) {
	testCases := []struct {
		Tables []string
		Result *expr.From
	}{
		{
			[]string{"table"},
			&expr.From{Tables: []syntax.Table{{Name: "table"}}},
		},
		{
			[]string{"table AS t"},
			&expr.From{Tables: []syntax.Table{{Name: "table", Alias: "t"}}},
		},
		{
			[]string{"table1 AS t1", "table2 AS t2"},
			&expr.From{Tables: []syntax.Table{{Name: "table1", Alias: "t1"}, {Name: "table2", Alias: "t2"}}},
		},
	}

	for _, testCase := range testCases {
		res := expr.NewFrom(testCase.Tables)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
