package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestAddColumn_Build(t *testing.T) {
	testCases := []struct {
		AddColumn *mig.AddColumn
		Expected  *syntax.StmtSet
	}{
		{
			&mig.AddColumn{Column: "column", Type: "type"},
			&syntax.StmtSet{Keyword: "ADD COLUMN", Value: "column type"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.AddColumn.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestAddCons_Build(t *testing.T) {
	testCases := []struct {
		AddCons  *mig.AddCons
		Expected *syntax.StmtSet
	}{
		{
			&mig.AddCons{Key: "key"},
			&syntax.StmtSet{Keyword: "ADD CONSTRAINT", Value: "key"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.AddCons.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
