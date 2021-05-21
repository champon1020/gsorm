package mig_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestRef_Build(t *testing.T) {
	testCases := []struct {
		Ref      *mig.Ref
		Expected *syntax.StmtSet
	}{
		{
			&mig.Ref{Table: "table", Columns: []string{"column"}},
			&syntax.StmtSet{Keyword: "REFERENCES", Value: "table (column)"},
		},
		{
			&mig.Ref{Table: "table", Columns: []string{"column1", "column2"}},
			&syntax.StmtSet{Keyword: "REFERENCES", Value: "table (column1, column2)"},
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

func TestRef_String(t *testing.T) {
	testCases := []struct {
		Ref      *mig.Ref
		Expected string
	}{
		{
			&mig.Ref{Table: "table", Columns: []string{"column"}},
			`REFERENCES(table, [column])`,
		},
		{
			&mig.Ref{Table: "table", Columns: []string{"column1", "column2"}},
			`REFERENCES(table, [column1 column2])`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Ref.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
