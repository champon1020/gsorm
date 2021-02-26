package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
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

func TestDropIndex_Build(t *testing.T) {
	testCases := []struct {
		DropIndex *mig.DropIndex
		Expected  *syntax.StmtSet
	}{
		{
			&mig.DropIndex{IdxName: "idx"},
			&syntax.StmtSet{Keyword: "DROP INDEX", Value: "idx"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropIndex.Build()
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

func TestDropPrimary_Build(t *testing.T) {
	testCases := []struct {
		DropPrimary *mig.DropPrimary
		Expected    *syntax.StmtSet
	}{
		{
			&mig.DropPrimary{Driver: internal.PSQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP CONSTRAINT", Value: "key"},
		},
		{
			&mig.DropPrimary{Driver: internal.MySQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP PRIMARY KEY"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropPrimary.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestDropForeign_Build(t *testing.T) {
	testCases := []struct {
		DropForeign *mig.DropForeign
		Expected    *syntax.StmtSet
	}{
		{
			&mig.DropForeign{Driver: internal.PSQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP CONSTRAINT", Value: "key"},
		},
		{
			&mig.DropForeign{Driver: internal.MySQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP FOREIGN KEY", Value: "key"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropForeign.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestDropUnique_Build(t *testing.T) {
	testCases := []struct {
		DropUnique *mig.DropUnique
		Expected   *syntax.StmtSet
	}{
		{
			&mig.DropUnique{Driver: internal.PSQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP CONSTRAINT", Value: "key"},
		},
		{
			&mig.DropUnique{Driver: internal.MySQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP INDEX", Value: "key"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropUnique.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
