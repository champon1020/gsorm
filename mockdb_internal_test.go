package mgorm

import (
	"errors"
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/stretchr/testify/assert"
)

var (
	opTest1 internal.Op = "mgorm.Test1"
	opTest2 internal.Op = "mgorm.Test2"
	opTest3 internal.Op = "mgorm.Test3"
	opTest4 internal.Op = "mgorm.Test4"
)

func TestMockDB_addExecuted(t *testing.T) {
	testCases := []struct {
		MockDB *MockDB
		Called []*opArgs
		Result *MockDB
	}{
		{
			&MockDB{},
			[]*opArgs{{op: opTest1, args: []interface{}{10, "str"}}},
			&MockDB{actual: [][]*opArgs{[]*opArgs{{op: opTest1, args: []interface{}{10, "str"}}}}},
		},
		{
			&MockDB{},
			[]*opArgs{
				{op: opTest1, args: []interface{}{10, "str1"}},
				{op: opTest2, args: []interface{}{100, "str2"}},
			},
			&MockDB{actual: [][]*opArgs{[]*opArgs{
				{op: opTest1, args: []interface{}{10, "str1"}},
				{op: opTest2, args: []interface{}{100, "str2"}},
			}}},
		},
	}

	for _, testCase := range testCases {
		testCase.MockDB.addExecuted(testCase.Called)
		assert.Equal(t, testCase.Result.actual, testCase.MockDB.actual)
	}
}

func TestMockDB_AddExpected(t *testing.T) {
	testCases := []struct {
		MockDB *MockDB
		Stmt   *Stmt
		Result *MockDB
	}{
		{
			&MockDB{},
			&Stmt{called: []*opArgs{{op: opTest1, args: []interface{}{10, "str"}}}},
			&MockDB{expected: [][]*opArgs{[]*opArgs{{op: opTest1, args: []interface{}{10, "str"}}}}},
		},
		{
			&MockDB{},
			&Stmt{called: []*opArgs{
				{op: opTest1, args: []interface{}{10, "str1"}},
				{op: opTest2, args: []interface{}{100, "str2"}},
			}},
			&MockDB{expected: [][]*opArgs{[]*opArgs{
				{op: opTest1, args: []interface{}{10, "str1"}},
				{op: opTest2, args: []interface{}{100, "str2"}},
			}}},
		},
	}

	for _, testCase := range testCases {
		testCase.MockDB.AddExpected(testCase.Stmt)
		assert.Equal(t, testCase.Result.expected, testCase.MockDB.expected)
	}
}

func TestMockDB_Result(t *testing.T) {
	testCases := []struct {
		MockDB *MockDB
	}{
		{
			&MockDB{
				expected: [][]*opArgs{
					[]*opArgs{{op: opTest1, args: []interface{}{10, "str"}}},
				},
				actual: [][]*opArgs{
					[]*opArgs{{op: opTest1, args: []interface{}{10, "str"}}},
				},
			},
		},
		{
			&MockDB{
				expected: [][]*opArgs{
					[]*opArgs{{op: opTest1, args: []interface{}{10, "str", []string{"s1", "s2"}}}},
				},
				actual: [][]*opArgs{
					[]*opArgs{{op: opTest1, args: []interface{}{10, "str", []string{"s1", "s2"}}}},
				},
			},
		},
		{
			&MockDB{
				expected: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
						{op: opTest2, args: []interface{}{100, "str2"}},
					},
				},
				actual: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
						{op: opTest2, args: []interface{}{100, "str2"}},
					},
				},
			},
		},
		{
			&MockDB{
				expected: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
						{op: opTest2, args: []interface{}{100, "str2"}},
					},
					[]*opArgs{
						{op: opTest3, args: []interface{}{1000, "str3"}},
						{op: opTest4, args: []interface{}{10000, "str4"}},
					},
				},
				actual: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
						{op: opTest2, args: []interface{}{100, "str2"}},
					},
					[]*opArgs{
						{op: opTest3, args: []interface{}{1000, "str3"}},
						{op: opTest4, args: []interface{}{10000, "str4"}},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		if err := testCase.MockDB.Result(); err != nil {
			t.Errorf("Error was occurred:\n  %v", err)
		}
	}
}

func TestMockDB_Result_Fail(t *testing.T) {
	testCases := []struct {
		MockDB *MockDB
		Error  error
	}{
		{
			&MockDB{
				expected: [][]*opArgs{
					[]*opArgs{{op: opTest1, args: []interface{}{10, "str1"}}},
				},
				actual: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
						{op: opTest2, args: []interface{}{100, "str2"}},
					},
				},
			},
			errors.New(`Test1(10, "str1").Test2(100, "str2") was executed, but Test1(10, "str1") is expected`),
		},
		{
			&MockDB{
				expected: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
						{op: opTest2, args: []interface{}{100, "str2"}},
					},
				},
				actual: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
					},
				},
			},
			errors.New(`Test1(10, "str1") was executed, but Test1(10, "str1").Test2(100, "str2") is expected`),
		},
		{
			&MockDB{
				expected: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
					},
				},
				actual: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
					},
					[]*opArgs{
						{op: opTest2, args: []interface{}{100, "str2"}},
					},
				},
			},
			errors.New(`Test2(100, "str2") was executed, but not expected`),
		},
		{
			&MockDB{
				expected: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
					},
					[]*opArgs{
						{op: opTest2, args: []interface{}{100, "str2"}},
					},
				},
				actual: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
					},
				},
			},
			errors.New(`no query was executed, but Test2(100, "str2") is expected`),
		},
		{
			&MockDB{
				expected: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
					},
				},
				actual: [][]*opArgs{
					[]*opArgs{
						{op: opTest2, args: []interface{}{10, "str1"}},
					},
				},
			},
			errors.New(`Test2(10, "str1") was executed, but Test1(10, "str1") is expected`),
		},
		{
			&MockDB{
				expected: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{10, "str1"}},
					},
				},
				actual: [][]*opArgs{
					[]*opArgs{
						{op: opTest1, args: []interface{}{100, "str2"}},
					},
				},
			},
			errors.New(`Test1(100, "str2") was executed, but Test1(10, "str1") is expected`),
		},
	}

	for _, testCase := range testCases {
		err := testCase.MockDB.Result()
		if err.Error() != testCase.Error.Error() {
			t.Errorf("\nGot : %v\nWant: %v\n", err, testCase.Error)
		}
	}
}

func TestOpArgsToQueryString(t *testing.T) {
	testCases := []struct {
		OpArgs []*opArgs
		Result string
	}{
		{
			[]*opArgs{{op: opTest1, args: []interface{}{10, "str"}}},
			`Test1(10, "str")`,
		},
		{
			[]*opArgs{
				{op: opTest1, args: []interface{}{10, "str1"}},
				{op: opTest2, args: []interface{}{100, "str2"}},
			},
			`Test1(10, "str1").Test2(100, "str2")`,
		},
		{
			[]*opArgs{
				{op: opTest1, args: []interface{}{10, "str1"}},
				{op: opTest2, args: []interface{}{100, []string{"s1", "s2"}}},
			},
			`Test1(10, "str1").Test2(100, ["s1" "s2"])`,
		},
		{
			[]*opArgs{
				{op: opTest1, args: []interface{}{10, "str1"}},
				{op: opTest2, args: []interface{}{[]int{10, 100}, []string{"s1", "s2"}, []bool{true, false}}},
			},
			`Test1(10, "str1").Test2([10 100], ["s1" "s2"], [true false])`,
		},
	}

	for _, testCase := range testCases {
		res := opArgsToQueryString(testCase.OpArgs)
		assert.Equal(t, testCase.Result, res)
	}
}
