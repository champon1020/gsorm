package errors_test

import (
	"testing"

	"github.com/champon1020/mgorm/errors"
	"gotest.tools/v3/assert"
)

func TestError_Is(t *testing.T) {
	testCases := []struct {
		ExpectedError *errors.Error
		ActualError   *errors.Error
		Result        bool
	}{
		{
			&errors.Error{Msg: "Type is invalid", Code: errors.InvalidTypeError},
			&errors.Error{Msg: "Type is invalid", Code: errors.InvalidTypeError},
			true,
		},
		{
			&errors.Error{Msg: "Type is invalid", Code: errors.InvalidTypeError},
			&errors.Error{Msg: "Value is invalid", Code: errors.InvalidValueError},
			false,
		},
	}

	for _, testCase := range testCases {
		is := testCase.ExpectedError.Is(testCase.ActualError)
		assert.Equal(t, testCase.Result, is)
	}
}
