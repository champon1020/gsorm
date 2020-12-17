package syntax

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSet_Name(t *testing.T) {
	s := Set{}
	assert.Equal(t, "SET", s.name())
}

func TestSet_AddEq(t *testing.T) {
	testCases := []struct {
		LHS    string
		RHS    interface{}
		Set    *Set
		Result *Set
	}{
		{"lhs", "rhs", &Set{}, &Set{Eqs: []Eq{{LHS: "lhs", RHS: "rhs"}}}},
		{
			"lhs2",
			"rhs2",
			&Set{Eqs: []Eq{{LHS: "lhs1", RHS: "rhs1"}}},
			&Set{Eqs: []Eq{{LHS: "lhs1", RHS: "rhs1"}, {LHS: "lhs2", RHS: "rhs2"}}},
		},
	}

	for _, testCase := range testCases {
		testCase.Set.addEq(testCase.LHS, testCase.RHS)
		if diff := cmp.Diff(testCase.Set, testCase.Result); diff != "" {
			PrintTestDiff(t, diff)
		}
	}
}
