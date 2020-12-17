package syntax_test

import (
	"testing"

	"github.com/champon1020/minigorm/syntax"
	"github.com/google/go-cmp/cmp"
)

func TestDelete_Build(t *testing.T) {
	testCases := []struct {
		Delete *syntax.Delete
		Result *syntax.StmtSet
	}{
		{&syntax.Delete{}, &syntax.StmtSet{Clause: "DELETE"}},
	}

	for _, testCase := range testCases {
		res := testCase.Delete.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			syntax.PrintTestDiff(t, diff)
		}
	}
}
