package mig_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestForeign_String(t *testing.T) {
	testCases := []struct {
		Foreign  *mig.Foreign
		Expected string
	}{
		{
			&mig.Foreign{Columns: []string{"column"}},
			`Foreign([column])`,
		},
		{
			&mig.Foreign{Columns: []string{"column1", "column2"}},
			`Foreign([column1 column2])`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Foreign.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestForeign_Build(t *testing.T) {
	testCases := []struct {
		Foreign  *mig.Foreign
		Expected *syntax.ClauseSet
	}{
		{
			&mig.Foreign{Columns: []string{"column"}},
			&syntax.ClauseSet{Keyword: "FOREIGN KEY", Value: "(column)"},
		},
		{
			&mig.Foreign{Columns: []string{"column1", "column2"}},
			&syntax.ClauseSet{Keyword: "FOREIGN KEY", Value: "(column1, column2)"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Foreign.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
