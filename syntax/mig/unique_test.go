package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestUnique_Build(t *testing.T) {
	testCases := []struct {
		Unique   *mig.Unique
		Expected *syntax.StmtSet
	}{
		{
			&mig.Unique{Columns: []string{"column"}},
			&syntax.StmtSet{Keyword: "UNIQUE", Value: "(column)"},
		},
		{
			&mig.Unique{Columns: []string{"column1", "column2"}},
			&syntax.StmtSet{Keyword: "UNIQUE", Value: "(column1, column2)"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Unique.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
