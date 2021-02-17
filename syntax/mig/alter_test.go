package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestAlterTable_Build(t *testing.T) {
	testCases := []struct {
		AlterTable *mig.AlterTable
		Expected   *syntax.StmtSet
	}{
		{
			&mig.AlterTable{Table: "table"},
			&syntax.StmtSet{Keyword: "ALTER TABLE", Value: "table"},
		},
	}

	for _, testCase := range testCases {
		actual := testCase.AlterTable.Build()
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
