package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestValues_Name(t *testing.T) {
	v := new(syntax.Values)
	assert.Equal(t, "VALUES", syntax.ValuesName(v))
}

func TestValues_AddColumn(t *testing.T) {
	testCases := []struct {
		Value  interface{}
		Values *syntax.Values
		Result *syntax.Values
	}{
		{
			"val",
			&syntax.Values{},
			&syntax.Values{Columns: []interface{}{"val"}},
		},
	}

	for _, testCase := range testCases {
		syntax.ValuesAddColumn(testCase.Values, testCase.Value)
		if diff := cmp.Diff(testCase.Values, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestValues_Build(t *testing.T) {
	testCases := []struct {
		Values *syntax.Values
		Result *syntax.StmtSet
	}{
		{
			&syntax.Values{Columns: []interface{}{"column", 2, true}},
			&syntax.StmtSet{Clause: "VALUES", Value: `("column", 2, true)`},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Values.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewValues(t *testing.T) {
	testCases := []struct {
		Columns []interface{}
		Result  *syntax.Values
	}{
		{
			[]interface{}{"column", 2, true},
			&syntax.Values{Columns: []interface{}{"column", 2, true}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewValues(testCase.Columns)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
