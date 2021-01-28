package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestThen_String(t *testing.T) {
	testCases := []struct {
		Then   *syntax.Then
		Result string
	}{
		{
			&syntax.Then{Value: 10},
			`THEN(10)`,
		},
		{
			&syntax.Then{Value: "str"},
			`THEN("str")`,
		},
		{
			&syntax.Then{Value: true},
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
		Then   *syntax.Then
		Result *syntax.StmtSet
	}{
		{
			&syntax.Then{Value: 10},
			&syntax.StmtSet{Clause: "THEN", Value: "10"},
		},
		{
			&syntax.Then{Value: "str"},
			&syntax.StmtSet{Clause: "THEN", Value: `"str"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Then.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewThen(t *testing.T) {
	testCases := []struct {
		Value  interface{}
		Result *syntax.Then
	}{
		{
			10,
			&syntax.Then{Value: 10},
		},
		{
			"str",
			&syntax.Then{Value: "str"},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewThen(testCase.Value)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
