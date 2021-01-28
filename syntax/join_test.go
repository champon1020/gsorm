package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestJoin_String(t *testing.T) {
	testCases := []struct {
		Join   *syntax.Join
		Result string
	}{
		{
			&syntax.Join{Table: syntax.Table{Name: "table"}, Type: syntax.InnerJoin},
			`INNER JOIN("table")`,
		},
		{
			&syntax.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: syntax.LeftJoin},
			`LEFT JOIN("table AS t")`,
		},
		{
			&syntax.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: syntax.RightJoin},
			`RIGHT JOIN("table AS t")`,
		},
		{
			&syntax.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: syntax.FullJoin},
			`FULL OUTER JOIN("table AS t")`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Join.String()
		assert.Equal(t, testCase.Result, res)
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
			&syntax.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: syntax.LeftJoin},
			&syntax.StmtSet{Clause: "LEFT JOIN", Value: "table AS t"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Join.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
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
			syntax.RightJoin,
			&syntax.Join{Table: syntax.Table{Name: "table"}, Type: "RIGHT JOIN"},
		},
		{
			"table AS t",
			syntax.FullJoin,
			&syntax.Join{Table: syntax.Table{Name: "table", Alias: "t"}, Type: "FULL OUTER JOIN"},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewJoin(testCase.Table, testCase.Type)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
