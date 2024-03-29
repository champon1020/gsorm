package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
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
			`Offset(10)`,
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
		Result *syntax.ClauseSet
	}{
		{
			&clause.Offset{Num: 5},
			&syntax.ClauseSet{Keyword: "OFFSET", Value: "5"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Offset.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
