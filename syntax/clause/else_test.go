package clause_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestElse_String(t *testing.T) {
	testCases := []struct {
		Else   *clause.Else
		Result string
	}{
		{
			&clause.Else{Value: 10},
			`ELSE(10)`,
		},
		{
			&clause.Else{Value: "str"},
			`ELSE("str")`,
		},
		{
			&clause.Else{Value: true},
			`ELSE(true)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.Else.String()
		assert.Equal(t, testCase.Result, res)
	}
}

func TestElse_Build(t *testing.T) {
	testCases := []struct {
		Else   *clause.Else
		Result *syntax.StmtSet
	}{
		{
			&clause.Else{Value: 10},
			&syntax.StmtSet{Keyword: "ELSE", Value: "10"},
		},
		{
			&clause.Else{Value: "str"},
			&syntax.StmtSet{Keyword: "ELSE", Value: `"str"`},
		},
		{
			&clause.Else{Value: "str", IsColumn: true},
			&syntax.StmtSet{Keyword: "ELSE", Value: "str"},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Else.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestElse_Build_Fail(t *testing.T) {
	a := &clause.Else{Value: []int{10}}
	_, err := a.Build()
	if err == nil {
		t.Errorf("Error was not occurred")
	}
}
