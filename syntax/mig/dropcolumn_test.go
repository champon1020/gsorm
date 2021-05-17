package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestDropColumn_Build(t *testing.T) {
	testCases := []struct {
		DropColumn *mig.DropColumn
		Expected   *syntax.StmtSet
	}{
		{
			&mig.DropColumn{Column: "column"},
			&syntax.StmtSet{Keyword: "DROP COLUMN", Value: "column"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropColumn.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
