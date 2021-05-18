package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestDropDB_Build(t *testing.T) {
	testCases := []struct {
		DropDB   *mig.DropDB
		Expected *syntax.StmtSet
	}{
		{
			&mig.DropDB{DBName: "dbname"},
			&syntax.StmtSet{Keyword: "DROP DATABASE", Value: "dbname"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropDB.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestDropDB_String(t *testing.T) {
	testCases := []struct {
		DropDB   *mig.DropDB
		Expected string
	}{
		{
			&mig.DropDB{DBName: "dbname"},
			`DROP DATABASE(dbname)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.DropDB.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
