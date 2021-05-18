package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestCreateDB_Build(t *testing.T) {
	testCases := []struct {
		CreateDB *mig.CreateDB
		Expected *syntax.StmtSet
	}{
		{
			&mig.CreateDB{DBName: "database"},
			&syntax.StmtSet{Keyword: "CREATE DATABASE", Value: "database"},
		},
	}

	for _, testCase := range testCases {
		actual, err := testCase.CreateDB.Build()
		if err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestCreateDB_String(t *testing.T) {
	testCases := []struct {
		CreateDB *mig.CreateDB
		Expected string
	}{
		{
			&mig.CreateDB{DBName: "database"},
			`CREATE DATABASE(database)`,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.CreateDB.String()
		assert.Equal(t, testCase.Expected, actual)
	}
}
