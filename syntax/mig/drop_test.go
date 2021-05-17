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

func TestDropTable_Build(t *testing.T) {
	testCases := []struct {
		DropTable *mig.DropTable
		Expected  *syntax.StmtSet
	}{
		{
			&mig.DropTable{Table: "table"},
			&syntax.StmtSet{Keyword: "DROP TABLE", Value: "table"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropTable.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestDropColumn_Build(t *testing.T) {
	testCases := []struct {
		DropColumn *mig.DropColumn
		Expected   *syntax.StmtSet
	}{
		{
			&mig.DropColumn{Column: "column"},
			&syntax.StmtSet{Keyword: "DROP COLUMN", Value: "column"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropColumn.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
