package mig_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestColumn_String(t *testing.T) {
	testCases := []struct {
		Column   *mig.Column
		Expected string
	}{
		{
			&mig.Column{Col: "column", Type: "type"},
			`Column(column, type)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Column.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestColumn_Build(t *testing.T) {
	testCases := []struct {
		Column   *mig.Column
		Expected *syntax.ClauseSet
	}{
		{
			&mig.Column{Col: "id", Type: "INT"},
			&syntax.ClauseSet{Keyword: "id", Value: "INT"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Column.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
