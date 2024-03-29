package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestLimit_String(t *testing.T) {
	testCases := []struct {
		Limit  *clause.Limit
		Result string
	}{
		{
			&clause.Limit{Num: 10},
			`Limit(10)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Limit.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestLimit_Build(t *testing.T) {
	testCases := []struct {
		Limit  *clause.Limit
		Result *syntax.ClauseSet
	}{
		{
			&clause.Limit{Num: 5},
			&syntax.ClauseSet{Keyword: "LIMIT", Value: "5"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Limit.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
