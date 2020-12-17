package syntax

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSelect_Query(t *testing.T) {
	s := Select{}
	assert.Equal(t, "SELECT", s.query())
}

func TestSelect_AddColumn(t *testing.T) {
	testCases := []struct {
		Col    string
		Select *Select
		Result *Select
	}{
		{
			Col:    "column1",
			Select: &Select{},
			Result: &Select{Columns: []Column{{Name: "column1"}}},
		},
		{
			Col:    "column1 AS c1",
			Select: &Select{},
			Result: &Select{Columns: []Column{{Name: "column1", Alias: "c1"}}},
		},
		{
			Col:    "column2",
			Select: &Select{Columns: []Column{{Name: "column1"}}},
			Result: &Select{Columns: []Column{{Name: "column1"}, {Name: "column2"}}},
		},
	}

	for _, testCase := range testCases {
		testCase.Select.addColumn(testCase.Col)
		if diff := cmp.Diff(testCase.Select, testCase.Result); diff != "" {
			PrintTestDiff(t, diff)
		}
	}
}
