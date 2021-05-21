package mig_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestDropTable_Build(t *testing.T) {
	testCases := []struct {
		DropTable *mig.DropTable
		Expected  *syntax.StmtSet
	}{
		{
			&mig.DropTable{Table: "table"},
			&syntax.StmtSet{Keyword: "DROP TABLE", Value: "table"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropTable.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestDropTable_String(t *testing.T) {
	testCases := []struct {
		DropTable *mig.DropTable
		Expected  string
	}{
		{
			&mig.DropTable{Table: "table"},
			`DROP TABLE(table)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.DropTable.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
