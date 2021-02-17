package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestRef_Build(t *testing.T) {
	testCases := []struct {
		Ref      *mig.Ref
		Expected *syntax.StmtSet
	}{
		{
			&mig.Ref{Table: "table", Column: "column"},
			&syntax.StmtSet{Keyword: "REFERENCES", Value: "table(column)"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Ref.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
