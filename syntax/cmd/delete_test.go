package cmd_test

import (
	"testing"

	"github.com/champon1020/mgorm/syntax"
	"github.com/champon1020/mgorm/syntax/cmd"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestDelete_String(t *testing.T) {
	d := new(cmd.Delete)
	assert.Equal(t, "DELETE()", d.String())
}

func TestDelete_Build(t *testing.T) {
	testCases := []struct {
		Delete *cmd.Delete
		Result *syntax.StmtSet
	}{
		{&cmd.Delete{}, &syntax.StmtSet{Keyword: "DELETE"}},
	}

	for _, testCase := range testCases {
		res := testCase.Delete.Build()
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestNewDelete(t *testing.T) {
	testCases := []struct {
		Result *cmd.Delete
	}{
		{&cmd.Delete{}},
	}

	for _, testCase := range testCases {
		res := cmd.NewDelete()
		if diff := cmp.Diff(testCase.Result, res); diff != "" {
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
