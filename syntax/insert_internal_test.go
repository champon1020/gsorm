package syntax

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestInsert_Query(t *testing.T) {
	i := Insert{}
	assert.Equal(t, "INSERT INTO", i.query())
}

func TestInsert_AddTable(t *testing.T) {
	testCases := []struct {
		Table  string
		Insert *Insert
		Result *Insert
	}{
		{
			"table",
			&Insert{},
			&Insert{Table: Table{Name: "table"}},
		},
		{
			"table AS t",
			&Insert{},
			&Insert{Table: Table{Name: "table", Alias: "t"}},
		},
		{
			"table2",
			&Insert{Table: Table{Name: "table1", Alias: "t1"}},
			&Insert{Table: Table{Name: "table2"}},
		},
	}

	for _, testCase := range testCases {
		testCase.Insert.addTable(testCase.Table)
		if diff := cmp.Diff(testCase.Insert, testCase.Result); diff != "" {
			PrintTestDiff(t, diff)
		}
	}
}

func TestInsert_AddColumn(t *testing.T) {
	testCases := []struct {
		Column string
		Insert *Insert
		Result *Insert
	}{
		{
			"column",
			&Insert{},
			&Insert{Columns: []Column{{Name: "column"}}},
		},
		{
			"column AS c",
			&Insert{},
			&Insert{Columns: []Column{{Name: "column", Alias: "c"}}},
		},
		{
			"column2",
			&Insert{Columns: []Column{{Name: "column1"}}},
			&Insert{Columns: []Column{{Name: "column1"}, {Name: "column2"}}},
		},
	}

	for _, testCase := range testCases {
		testCase.Insert.addColumn(testCase.Column)
		if diff := cmp.Diff(testCase.Insert, testCase.Result); diff != "" {
			PrintTestDiff(t, diff)
		}
	}
}
