package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestDropColumn_Build(t *testing.T) {
	testCases := []struct {
		DropColumn *mig.DropColumn
		Expected   *syntax.StmtSet
	}{
		{
			&mig.DropColumn{Column: "column"},
			&syntax.StmtSet{Keyword: "DROP COLUMN", Value: "column"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropColumn.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestDropColumn_String(t *testing.T) {
	testCases := []struct {
		DropColumn *mig.DropColumn
		Expected   string
	}{
		{
			&mig.DropColumn{Column: "column"},
			`DROP COLUMN(column)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.DropColumn.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
