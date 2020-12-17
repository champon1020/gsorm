package syntax

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestUpdate_Query(t *testing.T) {
	u := Update{}
	assert.Equal(t, "UPDATE", u.query())
}

func TestUpdate_AddTable(t *testing.T) {
	testCases := []struct {
		Table  string
		Update *Update
		Result *Update
	}{
		{
			"table",
			&Update{},
			&Update{Table: Table{Name: "table"}},
		},
		{
			"table AS t",
			&Update{},
			&Update{Table: Table{Name: "table", Alias: "t"}},
		},
		{
			"table2",
			&Update{Table: Table{Name: "table1", Alias: "t1"}},
			&Update{Table: Table{Name: "table2"}},
		},
	}

	for _, testCase := range testCases {
		testCase.Update.addTable(testCase.Table)
		if diff := cmp.Diff(testCase.Update, testCase.Result); diff != "" {
			PrintTestDiff(t, diff)
		}
	}
}
