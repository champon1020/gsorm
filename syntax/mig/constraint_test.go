package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestConstraint_Build(t *testing.T) {
	testCases := []struct {
		Constraint *mig.Constraint
		Expected   *syntax.StmtSet
	}{
		{
			&mig.Constraint{Key: "key_name"},
			&syntax.StmtSet{Keyword: "CONSTRAINT", Value: "key_name"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Constraint.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
