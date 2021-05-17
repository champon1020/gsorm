package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestCreateIndex_Build(t *testing.T) {
	testCases := []struct {
		CreateIndex *mig.CreateIndex
		Expected    *syntax.StmtSet
	}{
		{
			&mig.CreateIndex{IdxName: "idx"},
			&syntax.StmtSet{Keyword: "CREATE INDEX", Value: "idx"},
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
