package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestLimit_String(t *testing.T) {
	testCases := []struct {
		Limit  *syntax.Limit
		Result string
	}{
		{
			&syntax.Limit{Num: 10},
			`LIMIT(10)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Limit.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestLimit_Build(t *testing.T) {
	testCases := []struct {
		Limit  *syntax.Limit
		Result *syntax.StmtSet
	}{
		{
			&syntax.Limit{Num: 5},
			&syntax.StmtSet{Clause: "LIMIT", Value: "5"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Limit.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewLimit(t *testing.T) {
	testCases := []struct {
		Num    int
		Result *syntax.Limit
	}{
		{
			Num:    5,
			Result: &syntax.Limit{Num: 5},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewLimit(testCase.Num)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
