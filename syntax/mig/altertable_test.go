package mig_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestAlterTable_String(t *testing.T) {
	testCases := []struct {
		AlterTable *mig.AlterTable
		Expected   string
	}{
		{
			&mig.AlterTable{Table: "table"},
			`AlterTable(table)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.AlterTable.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestAlterTable_Build(t *testing.T) {
	testCases := []struct {
		AlterTable *mig.AlterTable
		Expected   *syntax.ClauseSet
	}{
		{
			&mig.AlterTable{Table: "table"},
			&syntax.ClauseSet{Keyword: "ALTER TABLE", Value: "table"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.AlterTable.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
