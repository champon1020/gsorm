package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
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
