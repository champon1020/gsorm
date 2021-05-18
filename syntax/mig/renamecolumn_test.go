package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestRenameColumn_Build(t *testing.T) {
	testCases := []struct {
		RenameColumn *mig.RenameColumn
		Expected     *syntax.StmtSet
	}{
		{
			&mig.RenameColumn{Column: "column", Dest: "dest"},
			&syntax.StmtSet{Keyword: "RENAME COLUMN", Value: "column TO dest"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.RenameColumn.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestRenameColumn_String(t *testing.T) {
	testCases := []struct {
		RenameColumn *mig.RenameColumn
		Expected     string
	}{
		{
			&mig.RenameColumn{Column: "column", Dest: "dest"},
			`RENAME COLUMN(column, dest)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.RenameColumn.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
