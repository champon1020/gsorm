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

func TestDropPK_Build(t *testing.T) {
	testCases := []struct {
		DropPK   *mig.DropPK
		Expected *syntax.StmtSet
	}{
		{
			&mig.DropPK{Driver: internal.PSQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP CONSTRAINT", Value: "key"},
		},
		{
			&mig.DropPK{Driver: internal.MySQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP PRIMARY KEY"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropPK.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestDropFK_Build(t *testing.T) {
	testCases := []struct {
		DropFK   *mig.DropFK
		Expected *syntax.StmtSet
	}{
		{
			&mig.DropFK{Driver: internal.PSQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP CONSTRAINT", Value: "key"},
		},
		{
			&mig.DropFK{Driver: internal.MySQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP FOREIGN KEY", Value: "key"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropFK.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestDropUC_Build(t *testing.T) {
	testCases := []struct {
		DropUC   *mig.DropUC
		Expected *syntax.StmtSet
	}{
		{
			&mig.DropUC{Driver: internal.PSQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP CONSTRAINT", Value: "key"},
		},
		{
			&mig.DropUC{Driver: internal.MySQL, Key: "key"},
			&syntax.StmtSet{Keyword: "DROP INDEX", Value: "key"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.DropUC.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
