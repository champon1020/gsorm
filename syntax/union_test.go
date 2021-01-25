package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestUnion_Name(t *testing.T) {
	testCases := []struct {
		Union  *syntax.Union
		Result string
	}{
		{&syntax.Union{}, "UNION"},
		{&syntax.Union{All: true}, "UNION ALL"},
	}

	for _, testCase := range testCases {
		res := syntax.UnionName(testCase.Union)
		assert.Equal(t, testCase.Result, res)
	}
}

func TestUnion_Build(t *testing.T) {
	testCases := []struct {
		Union  *syntax.Union
		Result *syntax.StmtSet
	}{
		{
			&syntax.Union{Stmt: "SELECT * FROM dept_manager"},
			&syntax.StmtSet{Clause: "UNION", Value: "SELECT * FROM dept_manager"},
		},
		{
			&syntax.Union{Stmt: "SELECT * FROM dept_manager", All: true},
			&syntax.StmtSet{Clause: "UNION ALL", Value: "SELECT * FROM dept_manager"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Union.Build()
		if err != nil {
			t.Errorf("%v\n", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewUnion(t *testing.T) {
	testCases := []struct {
		Stmt   string
		All    bool
		Result *syntax.Union
	}{
		{
			"SELECT * FROM dept_manager",
			false,
			&syntax.Union{Stmt: "SELECT * FROM dept_manager", All: false},
		},
		{
			"SELECT * FROM dept_manager",
			true,
			&syntax.Union{Stmt: "SELECT * FROM dept_manager", All: true},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewUnion(testCase.Stmt, testCase.All)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
