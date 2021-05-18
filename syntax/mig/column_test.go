package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestColumn_Build(t *testing.T) {
	testCases := []struct {
		Column   *mig.Column
		Expected *syntax.StmtSet
	}{
		{
			&mig.Column{Col: "id", Type: "INT"},
			&syntax.StmtSet{Keyword: "id", Value: "INT"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Column.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestColumn_String(t *testing.T) {
	testCases := []struct {
		Column   *mig.Column
		Expected string
	}{
		{
			&mig.Column{Col: "column", Type: "type"},
			`COLUMN(column, type)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Column.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
