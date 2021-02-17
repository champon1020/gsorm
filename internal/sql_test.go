package internal_test

import (
	"reflect"
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/stretchr/testify/assert"
)

func TestSQL_String(t *testing.T) {
	testCases := []struct {
		SQL    internal.SQL
		Result reflect.Kind
	}{
		{"Test", reflect.String},
	}

	for _, testCase := range testCases {
		sql := testCase.SQL.String()
		assert.Equal(t, testCase.Result, reflect.TypeOf(sql).Kind())
	}
}

func TestSQL_Write(t *testing.T) {
	testCases := []struct {
		SQL    internal.SQL
		Str    string
		Result internal.SQL
	}{
		{"test", "add", "test add"},
		{"", "add", "add"},
		{"(test", ")", "(test)"},
		{"test", ",", "test,"},
		{"(", "test", "(test"},
	}

	for _, testCase := range testCases {
		testCase.SQL.Write(testCase.Str)
		assert.Equal(t, testCase.Result, testCase.SQL)
	}
}
