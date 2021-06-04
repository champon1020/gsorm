package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
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
			`Values("column")`,
		},
		{
			&clause.Values{Values: []interface{}{"column", 2, true}},
			`Values("column", 2, 1)`,
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
			&syntax.StmtSet{Keyword: "VALUES", Value: `('column', 2, 1)`},
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

func TestValues_AddValue(t *testing.T) {
	v := &clause.Values{}
	val := 100
	v.AddValue(val)
	assert.Equal(t, v.Values, []interface{}{val})
}
