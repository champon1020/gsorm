package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestJoin_Name(t *testing.T) {
	testCases := []struct {
		Join   *syntax.Join
		Result string
	}{
		{&syntax.Join{Type: syntax.InnerJoin}, "INNER JOIN"},
		{&syntax.Join{Type: syntax.LeftJoin}, "LEFT JOIN"},
		{&syntax.Join{Type: syntax.RightJoin}, "RIGHT JOIN"},
		{&syntax.Join{Type: syntax.FullJoin}, "FULL OUTER JOIN"},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.Result, syntax.JoinName(testCase.Join))
	}
}

func TestJoin_AddTable(t *testing.T) {
	testCases := []struct {
		Join   *syntax.Join
		Table  string
		Result *syntax.Join
	}{
		{
			&syntax.Join{},
			"table",
			&syntax.Join{Table: syntax.Table{Name: "table"}},
		},
		{
			&syntax.Join{},
			"table AS t",
			&syntax.Join{Table: syntax.Table{Name: "table", Alias: "t"}},
		},
	}

	for _, testCase := range testCases {
		syntax.JoinAddTable(testCase.Join, testCase.Table)
		if diff := cmp.Diff(testCase.Join, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestJoin_Build(t *testing.T) {
	testCases := []struct {
		Join   *syntax.Join
		Result *syntax.StmtSet
	}{
		{
			&syntax.Join{Table: syntax.Table{Name: "table"}, Type: syntax.InnerJoin},
			&syntax.StmtSet{Clause: "INNER JOIN", Value: "table"},
		},
		{
			&syntax.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: syntax.InnerJoin},
			&syntax.StmtSet{Clause: "INNER JOIN", Value: "table AS t"},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Join.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewJoin(t *testing.T) {
	testCases := []struct {
		Table  string
		Type   syntax.JoinType
		Result *syntax.Join
	}{
		{
			"table",
			syntax.InnerJoin,
			&syntax.Join{Table: syntax.Table{Name: "table"}, Type: "INNER JOIN"},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewJoin(testCase.Table, testCase.Type)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
