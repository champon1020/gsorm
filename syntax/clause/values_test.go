package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestValues_String(t *testing.T) {
	testCases := []struct {
		Values *clause.Values
		Result string
	}{
		{
			&clause.Values{Values: []interface{}{"column"}},
			`VALUES("column")`,
		},
		{
			&clause.Values{Values: []interface{}{"column", 2, true}},
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
		Values *clause.Values
		Result *syntax.StmtSet
	}{
		{
			&clause.Values{Values: []interface{}{"column"}},
			&syntax.StmtSet{Keyword: "VALUES", Value: `('column')`},
		},
		{
			&clause.Values{Values: []interface{}{"column", 2, true}},
			&syntax.StmtSet{Keyword: "VALUES", Value: `('column', 2, true)`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Values.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestValues_Build_Fail(t *testing.T) {
	testCases := []struct {
		Values *clause.Values
	}{
		{&clause.Values{Values: []interface{}{nil}}},
	}

	for _, testCase := range testCases {
		_, err := testCase.Values.Build()
		if err == nil {
			t.Errorf("Error was not occurred")
		}
	}
}
