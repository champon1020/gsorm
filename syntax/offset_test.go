package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOffset_Name(t *testing.T) {
	l := new(syntax.Offset)
	assert.Equal(t, "LIMIT", syntax.OffsetName(l))
}

func TestOffset_Build(t *testing.T) {
	testCases := []struct {
		Offset *syntax.Offset
		Result *syntax.StmtSet
	}{
		{
			&syntax.Offset{Num: 5},
			&syntax.StmtSet{Clause: "LIMIT", Value: "5"},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Offset.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewOffset(t *testing.T) {
	testCases := []struct {
		Num    int
		Result *syntax.Offset
	}{
		{
			Num:    5,
			Result: &syntax.Offset{Num: 5},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewOffset(testCase.Num)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
