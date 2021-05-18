package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestRawClause_String(t *testing.T) {
	testCases := []struct {
		RawClause *syntax.RawClause
		Expected  string
	}{
		{
			&syntax.RawClause{RawStr: "AUTO_INCREMENT"},
			`RAW CLAUSE("AUTO_INCREMENT")`,
		},
		{
			&syntax.RawClause{RawStr: "WHERE lhs = ?", Values: []interface{}{10}},
			`RAW CLAUSE("WHERE lhs = ?", 10)`,
		},
	}

	for _, testCase := range testCases {
		res := testCase.RawClause.String()
		assert.Equal(t, testCase.Expected, res)
	}
}

func TestRawClause_Build(t *testing.T) {
	testCases := []struct {
		RawClause *syntax.RawClause
		Expected  *syntax.StmtSet
	}{
		{
			&syntax.RawClause{RawStr: "AUTO_INCREMENT"},
			&syntax.StmtSet{Keyword: "AUTO_INCREMENT"},
		},
		{
			&syntax.RawClause{RawStr: "WHERE lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Keyword: "WHERE lhs = 10"},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.RawClause.Build()
		if diff := cmp.Diff(testCase.Expected, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}