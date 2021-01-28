package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestUnion_String(t *testing.T) {
	testCases := []struct {
		Union  *syntax.Union
		Result string
	}{
		{
			&syntax.Union{Stmt: "SELECT * FROM table"},
			`UNION("SELECT * FROM table")`,
		},
		{
			&syntax.Union{Stmt: "SELECT * FROM table", All: true},
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
		Union  *syntax.Union
		Result *syntax.StmtSet
	}{
		{
			&syntax.Union{Stmt: "SELECT * FROM table"},
			&syntax.StmtSet{Clause: "UNION", Value: "SELECT * FROM table"},
		},
		{
			&syntax.Union{Stmt: "SELECT * FROM table", All: true},
			&syntax.StmtSet{Clause: "UNION ALL", Value: "SELECT * FROM table"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Union.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewUnion(t *testing.T) {
	testCases := []struct {
		Stmt   syntax.Var
		All    bool
		Result *syntax.Union
	}{
		{
			"SELECT * FROM table",
			false,
			&syntax.Union{Stmt: "SELECT * FROM table", All: false},
		},
		{
			"SELECT * FROM table",
			true,
			&syntax.Union{Stmt: "SELECT * FROM table", All: true},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewUnion(testCase.Stmt, testCase.All)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
