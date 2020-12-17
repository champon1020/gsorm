package syntax

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestValues_Name(t *testing.T) {
	v := new(Values)
	assert.Equal(t, "VALUES", v.name())
}

func TestValues_AddColumn(t *testing.T) {
	testCases := []struct {
		Value  interface{}
		Values *Values
		Result *Values
	}{
		{
			"val",
			&Values{},
			&Values{Columns: []interface{}{"val"}},
		},
	}

	for _, testCase := range testCases {
		testCase.Values.addColumn(testCase.Value)
		if diff := cmp.Diff(testCase.Values, testCase.Result); diff != "" {
			PrintTestDiff(t, diff)
		}
	}
}
