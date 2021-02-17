package mig_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
)

func TestDefault_Build(t *testing.T) {
	testCases := []struct {
		Default  *mig.Default
		Expected *syntax.StmtSet
	}{
		{
			&mig.Default{Value: "value"},
			&syntax.StmtSet{Keyword: "DEFAULT", Value: `"value"`},
		},
		{
			&mig.Default{Value: 10},
			&syntax.StmtSet{Keyword: "DEFAULT", Value: "10"},
		},
		{
			&mig.Default{Value: 10.1},
			&syntax.StmtSet{Keyword: "DEFAULT", Value: "10.1"},
		},
		{
			&mig.Default{Value: true},
			&syntax.StmtSet{Keyword: "DEFAULT", Value: "true"},
		},
		{
			&mig.Default{Value: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)},
			&syntax.StmtSet{Keyword: "DEFAULT", Value: "2021-01-01 00:00:00"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Default.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
