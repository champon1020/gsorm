package expr_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestElse_String(t *testing.T) {
	testCases := []struct {
		Else   *expr.Else
		Result string
	}{
		{
			&expr.Else{Value: 10},
			`ELSE(10)`,
		},
		{
			&expr.Else{Value: "str"},
			`ELSE("str")`,
		},
		{
			&expr.Else{Value: true},
			`ELSE(true)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Else.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestElse_Build(t *testing.T) {
	testCases := []struct {
		Else   *expr.Else
		Result *syntax.StmtSet
	}{
		{
			&expr.Else{Value: 10},
			&syntax.StmtSet{Clause: "ELSE", Value: "10"},
		},
		{
			&expr.Else{Value: "str", IsColumn: true},
			&syntax.StmtSet{Clause: "ELSE", Value: `"str"`},
		},
		{
			&expr.Else{Value: "str"},
			&syntax.StmtSet{Clause: "ELSE", Value: "str"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Else.Build(testCase.IsColumn)
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewElse(t *testing.T) {
	testCases := []struct {
		Value  interface{}
		Result *expr.Else
	}{
		{10, &expr.Else{Value: 10}},
		{"str", &expr.Else{Value: "str"}},
	}

	for _, testCase := range testCases {
		res := expr.NewElse(testCase.Value)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
