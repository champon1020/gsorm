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

func TestError_Error(t *testing.T) {
	err := internal.NewError(
		OpTest1,
		internal.KindBasic,
		errors.New("test error"),
	)
	expected := "[BasicError] test1.Op: test error"
	actual := err.Error()
	assert.Equal(t, expected, actual)
}

func TestCmpError(t *testing.T) {
	testCases := []struct {
		Actual   error
		Expected error
		Result   string
	}{
		{
			internal.NewError(OpTest2, internal.KindType, errors.New("test2 is occurred")),
			internal.NewError(OpTest1, internal.KindBasic, errors.New("test1 is occurred")),
			"\nOp:\n  got : test1.Op\n  want: test2.Op\n" +
				"Kind:\n  got : BasicError\n  want: TypeError\n" +
				"Err:\n  got : test1 is occurred\n  want: test2 is occurred",
		},
	}

	for _, testCase := range testCases {
		diff := internal.CmpError(testCase.Actual.(*internal.Error), testCase.Expected.(*internal.Error))
		assert.Equal(t, testCase.Result, diff)
	}
}
