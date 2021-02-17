package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestAutoInc_Build(t *testing.T) {
	testCases := []struct {
		AutoInc  *mig.AutoInc
		Expected *syntax.StmtSet
	}{
		{
			&mig.AutoInc{},
			&syntax.StmtSet{Keyword: "AUTO_INCREMENT"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.AutoInc.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
