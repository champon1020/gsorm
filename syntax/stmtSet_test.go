package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestStmtSet_WriteKeyword(t *testing.T) {
	testCases := []struct {
		StmtSet *syntax.StmtSet
		Keyword string
		Result  *syntax.StmtSet
	}{
		{
			&syntax.StmtSet{Keyword: ""},
			"clause",
			&syntax.StmtSet{Keyword: "clause"},
		},
		{
			&syntax.StmtSet{Keyword: "clause1"},
			"clause2",
			&syntax.StmtSet{Keyword: "clause1 clause2"},
		},
	}

	for _, testCase := range testCases {
		testCase.StmtSet.WriteKeyword(testCase.Keyword)
		if diff := cmp.Diff(testCase.Result, testCase.StmtSet); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestStmtSet_WriteValue(t *testing.T) {
	testCases := []struct {
		StmtSet *syntax.StmtSet
		Value   string
		Result  *syntax.StmtSet
	}{
		{
			&syntax.StmtSet{Value: ""},
			"value",
			&syntax.StmtSet{Value: "value"},
		},
		{
			&syntax.StmtSet{Value: "("},
			"value",
			&syntax.StmtSet{Value: "(value"},
		},
		{
			&syntax.StmtSet{Value: "(value"},
			")",
			&syntax.StmtSet{Value: "(value)"},
		},
		{
			&syntax.StmtSet{Value: "value1"},
			"value2",
			&syntax.StmtSet{Value: "value1 value2"},
		},
		{
			&syntax.StmtSet{Value: "value1"},
			",",
			&syntax.StmtSet{Value: "value1,"},
		},
		{
			&syntax.StmtSet{Value: "value1,"},
			"value2",
			&syntax.StmtSet{Value: "value1, value2"},
		},
	}

	for _, testCase := range testCases {
		testCase.StmtSet.WriteValue(testCase.Value)
		if diff := cmp.Diff(testCase.Result, testCase.StmtSet); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}

}

func TestStmtSet_Build(t *testing.T) {
	testCases := []struct {
		StmtSet *syntax.StmtSet
		Result  string
	}{
		{
			&syntax.StmtSet{Keyword: "clause", Value: "value"},
			"clause value",
		},
		{
			&syntax.StmtSet{Keyword: "clause", Value: "value", Parens: true},
			"clause (value)",
		},
		{
			&syntax.StmtSet{Keyword: "clause", Value: ""},
			"clause",
		},
	}

	for _, testCase := range testCases {
		res := testCase.StmtSet.Build()
		assert.Equal(t, testCase.Result, res)
	}
}
