package syntax_test

import (
	"testing"

	"github.com/champon1020/minigorm/syntax"
	"github.com/google/go-cmp/cmp"
)

func TestStmtSet_WriteClause(t *testing.T) {
	testCases := []struct {
		StmtSet *syntax.StmtSet
		Clause  string
		Result  *syntax.StmtSet
	}{
		{&syntax.StmtSet{Clause: ""}, "clause", &syntax.StmtSet{Clause: "clause"}},
		{&syntax.StmtSet{Clause: "clause1"}, "clause2", &syntax.StmtSet{Clause: "clause1 clause2"}},
	}

	for _, testCase := range testCases {
		testCase.StmtSet.WriteClause(testCase.Clause)
		if diff := cmp.Diff(testCase.StmtSet, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}

func TestStmtSet_WriteValue(t *testing.T) {
	testCases := []struct {
		StmtSet *syntax.StmtSet
		Value   string
		Result  *syntax.StmtSet
	}{
		{&syntax.StmtSet{Value: ""}, "value", &syntax.StmtSet{Value: "value"}},
		{&syntax.StmtSet{Value: "("}, "value", &syntax.StmtSet{Value: "(value"}},
		{&syntax.StmtSet{Value: "(value"}, ")", &syntax.StmtSet{Value: "(value)"}},
		{&syntax.StmtSet{Value: "value1"}, "value2", &syntax.StmtSet{Value: "value1 value2"}},
		{&syntax.StmtSet{Value: "value1"}, ",", &syntax.StmtSet{Value: "value1,"}},
		{&syntax.StmtSet{Value: "value1,"}, "value2", &syntax.StmtSet{Value: "value1, value2"}},
	}

	for _, testCase := range testCases {
		testCase.StmtSet.WriteValue(testCase.Value)
		if diff := cmp.Diff(testCase.StmtSet, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}

}
