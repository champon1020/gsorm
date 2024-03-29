package mig_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestUnique_String(t *testing.T) {
	testCases := []struct {
		Unique   *mig.Unique
		Expected string
	}{
		{
			&mig.Unique{Columns: []string{"column"}},
			`Unique([column])`,
		},
		{
			&mig.Unique{Columns: []string{"column1", "column2"}},
			`Unique([column1 column2])`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Unique.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestUnique_Build(t *testing.T) {
	testCases := []struct {
		Unique   *mig.Unique
		Expected *syntax.ClauseSet
	}{
		{
			&mig.Unique{Columns: []string{"column"}},
			&syntax.ClauseSet{Keyword: "UNIQUE", Value: "(column)"},
		},
		{
			&mig.Unique{Columns: []string{"column1", "column2"}},
			&syntax.ClauseSet{Keyword: "UNIQUE", Value: "(column1, column2)"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Unique.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
