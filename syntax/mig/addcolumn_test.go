package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
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

func TestAddColumn_String(t *testing.T) {
	testCases := []struct {
		AddColumn *mig.AddColumn
		Expected  string
	}{
		{
			&mig.AddColumn{Column: "column", Type: "type"},
			`ADD COLUMN(column, type)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.AddColumn.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
