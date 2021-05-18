package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestCons_Build(t *testing.T) {
	testCases := []struct {
		Cons     *mig.Cons
		Expected *syntax.StmtSet
	}{
		{
			&mig.Cons{Key: "key_name"},
			&syntax.StmtSet{Keyword: "CONSTRAINT", Value: "key_name"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Cons.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestCons_String(t *testing.T) {
	testCases := []struct {
		Cons     *mig.Cons
		Expected string
	}{
		{
			&mig.Cons{Key: "key"},
			`CONSTRAINT(key)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.Cons.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
