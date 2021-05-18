package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestAddCons_Build(t *testing.T) {
	testCases := []struct {
		AddCons  *mig.AddCons
		Expected *syntax.StmtSet
	}{
		{
			&mig.AddCons{Key: "key"},
			&syntax.StmtSet{Keyword: "ADD CONSTRAINT", Value: "key"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.AddCons.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestAddCons_String(t *testing.T) {
	testCases := []struct {
		AddCons  *mig.AddCons
		Expected string
	}{
		{
			&mig.AddCons{Key: "key"},
			`ADD CONSTRAINT(key)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.AddCons.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
