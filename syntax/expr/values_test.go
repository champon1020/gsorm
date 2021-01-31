package expr_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/expr"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestValues_String(t *testing.T) {
	testCases := []struct {
		Values *expr.Values
		Result string
	}{
		{
			&expr.Values{Columns: []interface{}{"column"}},
			`VALUES("column")`,
		},
		{
			&expr.Values{Columns: []interface{}{"column", 2, true}},
			`VALUES("column", 2, true)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Values.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestValues_Build(t *testing.T) {
	testCases := []struct {
		Values *expr.Values
		Result *syntax.StmtSet
	}{
		{
			&expr.Values{Columns: []interface{}{"column"}},
			&syntax.StmtSet{Clause: "VALUES", Value: `("column")`},
		},
		{
			&expr.Values{Columns: []interface{}{"column", 2, true}},
			&syntax.StmtSet{Clause: "VALUES", Value: `("column", 2, true)`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Values.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestValues_Build_Fail(t *testing.T) {
	testCases := []struct {
		Values *expr.Values
	}{
		{&expr.Values{Columns: []interface{}{time.Now()}}},
	}

	for _, testCase := range testCases {
		_, err := testCase.Values.Build()
		if err == nil {
			t.Errorf("Error was not occurred")
		}
	}
}

func TestNewValues(t *testing.T) {
	testCases := []struct {
		Columns []interface{}
		Result  *expr.Values
	}{
		{
			[]interface{}{"column"},
			&expr.Values{Columns: []interface{}{"column"}},
		},
		{
			[]interface{}{"column", 2, true},
			&expr.Values{Columns: []interface{}{"column", 2, true}},
		},
	}

	for _, testCase := range testCases {
		res := expr.NewValues(testCase.Columns)
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
