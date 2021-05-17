package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestCreateTable_Build(t *testing.T) {
	testCases := []struct {
		CreateTable *mig.CreateTable
		Expected    *syntax.StmtSet
	}{
		{
			&mig.CreateTable{Table: "table"},
			&syntax.StmtSet{Keyword: "CREATE TABLE", Value: "table"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.CreateTable.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
