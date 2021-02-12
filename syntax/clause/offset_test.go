package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOffset_String(t *testing.T) {
	testCases := []struct {
		Offset *clause.Offset
		Result string
	}{
		{
			&clause.Offset{Num: 10},
			`OFFSET(10)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Offset.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestOffset_Build(t *testing.T) {
	testCases := []struct {
		Offset *clause.Offset
		Result *syntax.StmtSet
	}{
		{
			&clause.Offset{Num: 5},
			&syntax.StmtSet{Keyword: "OFFSET", Value: "5"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Offset.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewOffset(t *testing.T) {
	testCases := []struct {
		Num    int
		Result *clause.Offset
	}{
		{
			Num:    5,
			Result: &clause.Offset{Num: 5},
		},
	}

	for _, testCase := range testCases {
		res := clause.NewOffset(testCase.Num)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
