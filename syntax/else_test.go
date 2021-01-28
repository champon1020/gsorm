package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestElse_String(t *testing.T) {
	testCases := []struct {
		Else   *syntax.Else
		Result string
	}{
		{
			&syntax.Else{Value: 10},
			`ELSE(10)`,
		},
		{
			&syntax.Else{Value: "str"},
			`ELSE("str")`,
		},
		{
			&syntax.Else{Value: true},
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
		Else   *syntax.Else
		Result *syntax.StmtSet
	}{
		{
			&syntax.Else{Value: 10},
			&syntax.StmtSet{Clause: "ELSE", Value: "10"},
		},
		{
			&syntax.Else{Value: "str"},
			&syntax.StmtSet{Clause: "ELSE", Value: `"str"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Else.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewElse(t *testing.T) {
	testCases := []struct {
		Value  interface{}
		Result *syntax.Else
	}{
		{10, &syntax.Else{Value: 10}},
		{"str", &syntax.Else{Value: "str"}},
	}

	for _, testCase := range testCases {
		res := syntax.NewElse(testCase.Value)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
