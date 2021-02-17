package mig_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/mig"
	"github.com/google/go-cmp/cmp"
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
		actual := testCase.CreateDB.Build()
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestCreateTable_Build(t *testing.T) {
	testCases := []struct {
		CreateTable *mig.CreateTable
		Expected    *syntax.StmtSet
	}{
		{
			&mig.CreateTable{Table: "table"},
			&syntax.StmtSet{Keyword: "CREATE TABLE", Value: "table"},
		},
	}

	for _, testCase := range testCases {
		actual := testCase.CreateTable.Build()
		if diff := cmp.Diff(testCase.Expected, actual); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
