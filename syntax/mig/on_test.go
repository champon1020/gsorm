package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
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

func TestOn_String(t *testing.T) {
	testCases := []struct {
		On       *mig.On
		Expected string
	}{
		{
			&mig.On{Table: "table", Columns: []string{"column"}},
			`ON(table, [column])`,
		},
		{
			&mig.On{Table: "table", Columns: []string{"column1", "column2"}},
			`ON(table, [column1 column2])`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.On.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
