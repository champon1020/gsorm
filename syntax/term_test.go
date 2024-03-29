package syntax_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestTable_Build(t *testing.T) {
	testCases := []struct {
		Table  *syntax.Table
		Result string
	}{
		{&syntax.Table{Name: "table"}, "table"},
		{&syntax.Table{Name: "table", Alias: "t"}, "table AS t"},
	}

	for _, testCase := range testCases {
		res := testCase.Table.Build()
		assert.Equal(t, res, testCase.Result)
	}
}

func TestNewTable(t *testing.T) {
	testCases := []struct {
		TableStr string
		Result   *syntax.Table
	}{
		{"table", &syntax.Table{Name: "table"}},
		{"table AS t", &syntax.Table{Name: "table", Alias: "t"}},
		{"table as t", &syntax.Table{Name: "table", Alias: "t"}},
	}

	for _, testCase := range testCases {
		res := syntax.NewTable(testCase.TableStr)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestColumn_Build(t *testing.T) {
	testCases := []struct {
		Column *syntax.Column
		Result string
	}{
		{&syntax.Column{Name: "column"}, "column"},
		{&syntax.Column{Name: "column", Alias: "c"}, "column AS c"},
	}

	for _, testCase := range testCases {
		res := testCase.Column.Build()
		assert.Equal(t, res, testCase.Result)
	}
}

func TestNewColumn(t *testing.T) {
	testCases := []struct {
		ColStr string
		Result *syntax.Column
	}{
		{"column", &syntax.Column{Name: "column"}},
		{"column AS c", &syntax.Column{Name: "column", Alias: "c"}},
		{"column as c", &syntax.Column{Name: "column", Alias: "c"}},
	}

	for _, testCase := range testCases {
		res := syntax.NewColumn(testCase.ColStr)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
