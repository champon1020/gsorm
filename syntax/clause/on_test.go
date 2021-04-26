package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOn_String(t *testing.T) {
	testCases := []struct {
		On     *clause.On
		Result string
	}{
		{
			&clause.On{Expr: "table1.column = table2.column"},
			`ON("table1.column = table2.column")`,
		},
		{
			&clause.On{Expr: "table1.column = ?", Values: []interface{}{"table2.column"}},
			`ON("table1.column = ?", 'table2.column')`,
		},
		{
			&clause.On{Expr: "table1.column = ? AND table2.column = ?", Values: []interface{}{"table3.column", "table4.column"}},
			`ON("table1.column = ? AND table2.column = ?", 'table3.column', 'table4.column')`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.On.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestOn_Build(t *testing.T) {
	testCases := []struct {
		On     *clause.On
		Result *syntax.StmtSet
	}{
		{
			&clause.On{Expr: "table1.column = table2.column"},
			&syntax.StmtSet{Keyword: "ON", Value: "table1.column = table2.column"},
		},
		{
			&clause.On{Expr: "table1.column = ?", Values: []interface{}{"table2.column"}},
			&syntax.StmtSet{Keyword: "ON", Value: `table1.column = table2.column`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.On.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
