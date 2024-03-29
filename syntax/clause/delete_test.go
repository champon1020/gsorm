package clause_test

import (
	"testing"

	"github.com/champon1020/gsorm/syntax"
	"github.com/champon1020/gsorm/syntax/clause"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestDelete_String(t *testing.T) {
	d := new(clause.Delete)
	assert.Equal(t, "Delete()", d.String())
}

func TestDelete_Build(t *testing.T) {
	testCases := []struct {
		Delete *clause.Delete
		Result *syntax.ClauseSet
	}{
		{&clause.Delete{}, &syntax.ClauseSet{Keyword: "DELETE"}},
	}

	for _, testCase := range testCases {
		actual, err := testCase.Delete.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
