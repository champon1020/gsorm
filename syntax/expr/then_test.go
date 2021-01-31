package expr_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestThen_String(t *testing.T) {
	testCases := []struct {
		Then   *expr.Then
		Result string
	}{
		{
			&expr.Then{Value: 10},
			`THEN(10)`,
		},
		{
			&expr.Then{Value: "str"},
			`THEN("str")`,
		},
		{
			&expr.Then{Value: true},
			`THEN(true)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Then.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestThen_Build(t *testing.T) {
	testCases := []struct {
		Then   *expr.Then
		Result *syntax.StmtSet
	}{
		{
			&expr.Then{Value: 10},
			&syntax.StmtSet{Clause: "THEN", Value: "10"},
		},
		{
			&expr.Then{Value: "str"},
			&syntax.StmtSet{Clause: "THEN", Value: `"str"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Then.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewThen(t *testing.T) {
	testCases := []struct {
		Value  interface{}
		Result *expr.Then
	}{
		{
			10,
			&expr.Then{Value: 10},
		},
		{
			"str",
			&expr.Then{Value: "str"},
		},
	}

	for _, testCase := range testCases {
		res := expr.NewThen(testCase.Value)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
