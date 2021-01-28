package syntax_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestValues_String(t *testing.T) {
	testCases := []struct {
		Values *syntax.Values
		Result string
	}{
		{
			&syntax.Values{Columns: []interface{}{"column"}},
			`VALUES("column")`,
		},
		{
			&syntax.Values{Columns: []interface{}{"column", 2, true}},
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
		Values *syntax.Values
		Result *syntax.StmtSet
	}{
		{
			&syntax.Values{Columns: []interface{}{"column"}},
			&syntax.StmtSet{Clause: "VALUES", Value: `("column")`},
		},
		{
			&syntax.Values{Columns: []interface{}{"column", 2, true}},
			&syntax.StmtSet{Clause: "VALUES", Value: `("column", 2, true)`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Values.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestValues_Build_Fail(t *testing.T) {
	testCases := []struct {
		Values *syntax.Values
	}{
		{&syntax.Values{Columns: []interface{}{time.Now()}}},
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
		Result  *syntax.Values
	}{
		{
			[]interface{}{"column"},
			&syntax.Values{Columns: []interface{}{"column"}},
		},
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
