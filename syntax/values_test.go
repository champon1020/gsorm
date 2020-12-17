package syntax_test

import (
	"testing"

	"github.com/champon1020/minigorm/syntax"
	"github.com/google/go-cmp/cmp"
)

func TestValues_Build(t *testing.T) {
	testCases := []struct {
		Values *syntax.Values
		Result *syntax.StmtSet
	}{
		{
			&syntax.Values{Columns: []interface{}{"column", 2, true}},
			&syntax.StmtSet{Clause: "VALUES", Value: "(column, 2, true)"},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Values.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestNewValues(t *testing.T) {
	testCases := []struct {
		Columns []interface{}
		Result  *syntax.Values
	}{
		{
			[]interface{}{"column", 2, true},
			&syntax.Values{Columns: []interface{}{"column", 2, true}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewValues(testCase.Columns)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}
