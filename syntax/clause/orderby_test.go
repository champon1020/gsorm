package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOrderBy_String(t *testing.T) {
	testCases := []struct {
		OrderBy *clause.OrderBy
		Result  string
	}{
		{
			&clause.OrderBy{Columns: []string{"column"}},
			`ORDER BY(["column"])`,
		},
		{
			&clause.OrderBy{Columns: []string{"column1", "column2 DESC"}},
			`ORDER BY(["column1" "column2 DESC"])`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.OrderBy.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestOrderBy_Build(t *testing.T) {
	testCases := []struct {
		OrderBy *clause.OrderBy
		Result  *syntax.StmtSet
	}{
		{
			&clause.OrderBy{Columns: []string{"column"}},
			&syntax.StmtSet{Keyword: "ORDER BY", Value: "column"},
		},
		{
			&clause.OrderBy{Columns: []string{"column1", "column2 DESC"}},
			&syntax.StmtSet{Keyword: "ORDER BY", Value: "column1, column2 DESC"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.OrderBy.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
