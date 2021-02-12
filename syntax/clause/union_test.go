package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestUnion_String(t *testing.T) {
	testCases := []struct {
		Union  *clause.Union
		Result string
	}{
		{
			&clause.Union{Stmt: "SELECT * FROM table"},
			`UNION("SELECT * FROM table")`,
		},
		{
			&clause.Union{Stmt: "SELECT * FROM table", All: true},
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
			&clause.Union{Stmt: "SELECT * FROM table"},
			&syntax.StmtSet{Keyword: "UNION", Value: "SELECT * FROM table"},
		},
		{
			&clause.Union{Stmt: "SELECT * FROM table", All: true},
			&syntax.StmtSet{Keyword: "UNION ALL", Value: "SELECT * FROM table"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Union.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewUnion(t *testing.T) {
	testCases := []struct {
		Stmt   syntax.Sub
		All    bool
		Result *clause.Union
	}{
		{
			"SELECT * FROM table",
			false,
			&clause.Union{Stmt: "SELECT * FROM table", All: false},
		},
		{
			"SELECT * FROM table",
			true,
			&clause.Union{Stmt: "SELECT * FROM table", All: true},
		},
	}

	for _, testCase := range testCases {
		res := clause.NewUnion(testCase.Stmt, testCase.All)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
