package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestOn_Build(t *testing.T) {
	testCases := []struct {
		On       *mig.On
		Expected *syntax.StmtSet
	}{
		{
			&mig.On{Table: "table", Columns: []string{"column"}},
			&syntax.StmtSet{Keyword: "ON", Value: "table (column)"},
		},
		{
			&mig.On{Table: "table", Columns: []string{"column1", "column2"}},
			&syntax.StmtSet{Keyword: "ON", Value: "table (column1, column2)"},
		},
		{
			&mig.On{Table: "table"},
			&syntax.StmtSet{Keyword: "ON", Value: "table"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.On.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
