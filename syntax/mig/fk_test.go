package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestFK_Build(t *testing.T) {
	testCases := []struct {
		FK       *mig.FK
		Expected *syntax.StmtSet
	}{
		{
			&mig.FK{Columns: []string{"column"}},
			&syntax.StmtSet{Keyword: "FOREIGN KEY", Value: "(column)"},
		},
		{
			&mig.FK{Columns: []string{"column1", "column2"}},
			&syntax.StmtSet{Keyword: "FOREIGN KEY", Value: "(column1, column2)"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.FK.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
