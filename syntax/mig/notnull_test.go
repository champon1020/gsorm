package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestNotNull_Build(t *testing.T) {
	testCases := []struct {
		NotNull  *mig.NotNull
		Expected *syntax.StmtSet
	}{
		{
			&mig.NotNull{},
			&syntax.StmtSet{Keyword: "NOT NULL"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.NotNull.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
