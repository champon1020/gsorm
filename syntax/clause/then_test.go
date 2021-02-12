package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestThen_String(t *testing.T) {
	testCases := []struct {
		Then   *clause.Then
		Result string
	}{
		{
			&clause.Then{Value: 10},
			`THEN(10)`,
		},
		{
			&clause.Then{Value: "str"},
			`THEN("str")`,
		},
		{
			&clause.Then{Value: true},
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
		Then   *clause.Then
		Result *syntax.StmtSet
	}{
		{
			&clause.Then{Value: 10},
			&syntax.StmtSet{Keyword: "THEN", Value: "10"},
		},
		{
			&clause.Then{Value: "str"},
			&syntax.StmtSet{Keyword: "THEN", Value: `"str"`},
		},
		{
			&clause.Then{Value: "str", IsColumn: true},
			&syntax.StmtSet{Keyword: "THEN", Value: "str"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Then.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestNewThen(t *testing.T) {
	testCases := []struct {
		Value  interface{}
		Result *clause.Then
	}{
		{
			10,
			&clause.Then{Value: 10},
		},
		{
			"str",
			&clause.Then{Value: "str"},
		},
	}

	for _, testCase := range testCases {
		res := clause.NewThen(testCase.Value)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
