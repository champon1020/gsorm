package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestPrimary_Build(t *testing.T) {
	testCases := []struct {
		Primary  *mig.Primary
		Expected *syntax.StmtSet
	}{
		{
			&mig.Primary{Columns: []string{"id"}},
			&syntax.StmtSet{Keyword: "PRIMARY KEY", Value: "(id)"},
		},
		{
			&mig.Primary{Columns: []string{"id", "name"}},
			&syntax.StmtSet{Keyword: "PRIMARY KEY", Value: "(id, name)"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Primary.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}