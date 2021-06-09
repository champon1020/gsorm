package mig_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestCreateIndex_String(t *testing.T) {
	testCases := []struct {
		CreateIndex *mig.CreateIndex
		Expected    string
	}{
		{
			&mig.CreateIndex{IdxName: "idx"},
			`CreateIndex(idx)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.CreateIndex.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestCreateIndex_Build(t *testing.T) {
	testCases := []struct {
		CreateIndex *mig.CreateIndex
		Expected    *syntax.ClauseSet
	}{
		{
			&mig.CreateIndex{IdxName: "idx"},
			&syntax.ClauseSet{Keyword: "CREATE INDEX", Value: "idx"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.CreateIndex.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
