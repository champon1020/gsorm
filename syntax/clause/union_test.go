package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm"
	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestUnion_String(t *testing.T) {
	testCases := []struct {
		Union  *clause.Union
		Result string
	}{
		{
			&clause.Union{Stmt: gsorm.Select(nil, "*").From("table")},
			`UNION("SELECT * FROM table")`,
		},
		{
			&clause.Union{Stmt: gsorm.Select(nil, "*").From("table"), All: true},
			`UNION ALL("SELECT * FROM table")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Union.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestUnion_Build(t *testing.T) {
	testCases := []struct {
		Union  *clause.Union
		Result *syntax.StmtSet
	}{
		{
			&clause.Union{Stmt: gsorm.Select(nil, "*").From("table")},
			&syntax.StmtSet{Keyword: "UNION", Value: "(SELECT * FROM table)"},
		},
		{
			&clause.Union{Stmt: gsorm.Select(nil, "*").From("table"), All: true},
			&syntax.StmtSet{Keyword: "UNION ALL", Value: "(SELECT * FROM table)"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Union.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
