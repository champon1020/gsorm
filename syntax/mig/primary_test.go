package mig_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestPrimary_String(t *testing.T) {
	testCases := []struct {
		Primary  *mig.Primary
		Expected string
	}{
		{
			&mig.Primary{Columns: []string{"column"}},
			`Primary([column])`,
		},
		{
			&mig.Primary{Columns: []string{"column1", "column2"}},
			`Primary([column1 column2])`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Primary.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestPrimary_Build(t *testing.T) {
	testCases := []struct {
		Primary  *mig.Primary
		Expected *syntax.ClauseSet
	}{
		{
			&mig.Primary{Columns: []string{"id"}},
			&syntax.ClauseSet{Keyword: "PRIMARY KEY", Value: "(id)"},
		},
		{
			&mig.Primary{Columns: []string{"id", "name"}},
			&syntax.ClauseSet{Keyword: "PRIMARY KEY", Value: "(id, name)"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Primary.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
