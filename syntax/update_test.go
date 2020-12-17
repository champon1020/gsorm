package syntax_test

import (
	"testing"

	"github.com/champon1020/minigorm/syntax"
	"github.com/google/go-cmp/cmp"
)

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
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestNewUpdate(t *testing.T) {
	testCases := []struct {
		Table  string
		Result *syntax.Update
	}{
		{
			"table",
			&syntax.Update{Table: syntax.Table{Name: "table"}},
		},
		{
			"table AS t",
			&syntax.Update{Table: syntax.Table{Name: "table", Alias: "t"}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewUpdate(testCase.Table)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}
