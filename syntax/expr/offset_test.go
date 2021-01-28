package expr_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOffset_String(t *testing.T) {
	testCases := []struct {
		Offset *expr.Offset
		Result string
	}{
		{
			&expr.Offset{Num: 10},
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
		Offset *expr.Offset
		Result *syntax.StmtSet
	}{
		{
			&expr.Offset{Num: 5},
			&syntax.StmtSet{Clause: "OFFSET", Value: "5"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Offset.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewOffset(t *testing.T) {
	testCases := []struct {
		Num    int
		Result *expr.Offset
	}{
		{
			Num:    5,
			Result: &expr.Offset{Num: 5},
		},
	}

	for _, testCase := range testCases {
		res := expr.NewOffset(testCase.Num)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
