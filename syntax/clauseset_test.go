package syntax_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestClauseSet_WriteKeyword(t *testing.T) {
	testCases := []struct {
		ClauseSet *syntax.ClauseSet
		Keyword   string
		Result    *syntax.ClauseSet
	}{
		{
			&syntax.ClauseSet{Keyword: ""},
			"clause",
			&syntax.ClauseSet{Keyword: "clause"},
		},
		{
			&syntax.ClauseSet{Keyword: "clause1"},
			"clause2",
			&syntax.ClauseSet{Keyword: "clause1 clause2"},
		},
	}

	for _, testCase := range testCases {
		testCase.ClauseSet.WriteKeyword(testCase.Keyword)
		if diff := cmp.Diff(testCase.Result, testCase.ClauseSet); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestClauseSet_WriteValue(t *testing.T) {
	testCases := []struct {
		ClauseSet *syntax.ClauseSet
		Value     string
		Result    *syntax.ClauseSet
	}{
		{
			&syntax.ClauseSet{Value: ""},
			"value",
			&syntax.ClauseSet{Value: "value"},
		},
		{
			&syntax.ClauseSet{Value: "("},
			"value",
			&syntax.ClauseSet{Value: "(value"},
		},
		{
			&syntax.ClauseSet{Value: "(value"},
			")",
			&syntax.ClauseSet{Value: "(value)"},
		},
		{
			&syntax.ClauseSet{Value: "value1"},
			"value2",
			&syntax.ClauseSet{Value: "value1 value2"},
		},
		{
			&syntax.ClauseSet{Value: "value1"},
			",",
			&syntax.ClauseSet{Value: "value1,"},
		},
		{
			&syntax.ClauseSet{Value: "value1,"},
			"value2",
			&syntax.ClauseSet{Value: "value1, value2"},
		},
	}

	for _, testCase := range testCases {
		testCase.ClauseSet.WriteValue(testCase.Value)
		if diff := cmp.Diff(testCase.Result, testCase.ClauseSet); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}

}

func TestClauseSet_Build(t *testing.T) {
	testCases := []struct {
		ClauseSet *syntax.ClauseSet
		Result    string
	}{
		{
			&syntax.ClauseSet{Keyword: "clause", Value: "value"},
			"clause value",
		},
		{
			&syntax.ClauseSet{Keyword: "clause", Value: "value", Parens: true},
			"clause (value)",
		},
		{
			&syntax.ClauseSet{Keyword: "clause", Value: ""},
			"clause",
		},
		{
			&syntax.ClauseSet{Value: "value"},
			"value",
		},
		{
			&syntax.ClauseSet{Value: "value", Parens: true},
			"(value)",
		},
	}

	for _, testCase := range testCases {
		res := testCase.ClauseSet.Build()
		assert.Equal(t, testCase.Result, res)
	}
}
