package internal_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	var (
		n0 int    = 1
		n1 int8   = 2
		n2 int16  = 3
		n3 int32  = 4
		n4 int64  = 5
		u0 uint   = 6
		u1 uint8  = 7
		u2 uint16 = 8
		u3 uint32 = 9
		u4 uint64 = 10
	)

	testCases := []struct {
		Value  interface{}
		Result string
	}{
		{"rhs", "rhs"},
		{n0, "1"},
		{n1, "2"},
		{n2, "3"},
		{n3, "4"},
		{n4, "5"},
		{u0, "6"},
		{u1, "7"},
		{u2, "8"},
		{u3, "9"},
		{u4, "10"},
		{true, "true"},
	}

	for _, testCase := range testCases {
		res, _ := internal.ToString(testCase.Value)
		assert.Equal(t, testCase.Result, res)
	}
}

/*
func TestToString_Fail(t *testing.T) {
	testCases := []struct {
		Value     interface{}
		ErrorCode int
	}{
		{map[string]string{"key": "value"}, ErrInvalidType},
		{[]int{1, 2}, ErrInvalidType},
		{[2]int{1, 2}, ErrInvalidType},
	}

	for _, testCase := range testCases {
		_, err := internal.toString(testCase.Value)
		assert.Equal(t, testCase.ErrorCode, err.(Error).Code)
	}
}
*/
