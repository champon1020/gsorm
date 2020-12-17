package syntax

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestFrom_Name(t *testing.T) {
	f := From{}
	assert.Equal(t, "FROM", f.name())
}

func TestFrom_AddTable(t *testing.T) {
	testCases := []struct {
		Table  string
		From   *From
		Result *From
	}{
		{
			"column",
			&From{},
			&From{Tables: []Table{{Name: "column"}}},
		},
		{
			"column AS c",
			&From{},
			&From{Tables: []Table{{Name: "column", Alias: "c"}}},
		},
		{
			"column2",
			&From{Tables: []Table{{Name: "column1"}}},
			&From{Tables: []Table{{Name: "column1"}, {Name: "column2"}}},
		},
	}

	for _, testCase := range testCases {
		testCase.From.addTable(testCase.Table)
		if diff := cmp.Diff(testCase.From, testCase.Result); diff != "" {
			PrintTestDiff(t, diff)
		}
	}
}
