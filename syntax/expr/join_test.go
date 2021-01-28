package expr_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestJoin_String(t *testing.T) {
	testCases := []struct {
		Join   *expr.Join
		Result string
	}{
		{
			&expr.Join{Table: syntax.Table{Name: "table"}, Type: expr.InnerJoin},
			`INNER JOIN("table")`,
		},
		{
			&expr.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: expr.LeftJoin},
			`LEFT JOIN("table AS t")`,
		},
		{
			&expr.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: expr.RightJoin},
			`RIGHT JOIN("table AS t")`,
		},
		{
			&expr.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: expr.FullJoin},
			`FULL OUTER JOIN("table AS t")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Join.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestJoin_Build(t *testing.T) {
	testCases := []struct {
		Join   *expr.Join
		Result *syntax.StmtSet
	}{
		{
			&expr.Join{Table: syntax.Table{Name: "table"}, Type: expr.InnerJoin},
			&syntax.StmtSet{Clause: "INNER JOIN", Value: "table"},
		},
		{
			&expr.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: expr.LeftJoin},
			&syntax.StmtSet{Clause: "LEFT JOIN", Value: "table AS t"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Join.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewJoin(t *testing.T) {
	testCases := []struct {
		Table  string
		Type   expr.JoinType
		Result *expr.Join
	}{
		{
			"table",
			expr.RightJoin,
			&expr.Join{Table: syntax.Table{Name: "table"}, Type: "RIGHT JOIN"},
		},
		{
			"table AS t",
			expr.FullJoin,
			&expr.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: "FULL OUTER JOIN"},
		},
	}

	for _, testCase := range testCases {
		res := expr.NewJoin(testCase.Table, testCase.Type)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
