package internal_test

import (
	"errors"
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/stretchr/testify/assert"
)

const (
	OpTest1 = "test1.Op"
	OpTest2 = "test2.Op"
)

func TestCmpError(t *testing.T) {
	testCases := []struct {
		Actual   error
		Expected error
		Result   string
	}{
		{
			internal.NewError(OpTest1, internal.KindBasic, errors.New("test1 is occurred")),
			internal.NewError(OpTest2, internal.KindType, errors.New("test2 is occurred")),
			"\nOp:\n  got : test1.Op\n  want: test2.Op\nKind:\n  got : BasicError\n  want: TypeError\nErr:\n  got : test1 is occurred\n  want: test2 is occurred",
		},
	}

	for _, testCase := range testCases {
		diff := internal.CmpError(*testCase.Actual.(*internal.Error), *testCase.Expected.(*internal.Error))
		assert.Equal(t, testCase.Result, diff)
	}
}
