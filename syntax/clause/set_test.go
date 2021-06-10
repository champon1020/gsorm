package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSet_String(t *testing.T) {
	testCases := []struct {
		Set    *clause.Set
		Result string
	}{
		{
			&clause.Set{Column: "col", Value: 10},
			`Set(col, 10)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Set.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestSet_Build(t *testing.T) {
	testCases := []struct {
		Set    *clause.Set
		Result *syntax.ClauseSet
	}{
		{
			&clause.Set{Column: "lhs", Value: "rhs"},
			&syntax.ClauseSet{Keyword: "SET", Value: `lhs = 'rhs'`},
		},
		{
			&clause.Set{Column: "lhs1", Value: 10},
			&syntax.ClauseSet{Keyword: "SET", Value: `lhs1 = 10`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Set.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
