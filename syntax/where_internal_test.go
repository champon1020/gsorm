package syntax

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestWhere_Name(t *testing.T) {
	w := new(Where)
	assert.Equal(t, "WHERE", w.name())
}

func TestAnd_Name(t *testing.T) {
	a := new(And)
	assert.Equal(t, "AND", a.name())
}

func TestOr_Name(t *testing.T) {
	o := new(Or)
	assert.Equal(t, "OR", o.name())
}

func TestBuildStmtSet(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *StmtSet
	}{
		{
			"lhs = rhs",
			[]interface{}{},
			&StmtSet{Value: "lhs = rhs"},
		},
		{
			"lhs = ?",
			[]interface{}{"rhs"},
			&StmtSet{Value: "lhs = rhs"},
		},
		{
			"lhs = ?",
			[]interface{}{100},
			&StmtSet{Value: "lhs = 100"},
		},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{"rhs", 100},
			&StmtSet{Value: "lhs1 = rhs AND lhs2 = 100"},
		},
		{
			"IN lhs (?, ?, ?)",
			[]interface{}{"rhs", 100, true},
			&StmtSet{Value: "IN lhs (rhs, 100, true)"},
		},
		{
			"lhs LIKE %%?%%",
			[]interface{}{"rhs"},
			&StmtSet{Value: "lhs LIKE %rhs%"},
		},
		{
			"lhs BETWEEN ? AND ?",
			[]interface{}{10, 100},
			&StmtSet{Value: "lhs BETWEEN 10 AND 100"},
		},
	}

	for _, testCase := range testCases {
		res, _ := buildStmtSet(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			PrintTestDiff(t, diff)
		}
	}
}

func TestBuildStmtSet_Fail(t *testing.T) {
	testCases := []struct {
		Expr      string
		Values    []interface{}
		ErrorCode int
	}{
		{"lhs = ? AND rhs = ?", []interface{}{10}, ErrInvalidLen},
		{"lhs = ?", []interface{}{[]string{}}, ErrInvalidType},
	}

	for _, testCase := range testCases {
		_, err := buildStmtSet(testCase.Expr, testCase.Values...)
		if err == nil {
			t.Errorf("Error is not occurred with: %v, %v\n", testCase.Expr, testCase.Values)
			continue
		}
		assert.Equal(t, testCase.ErrorCode, err.(Error).Code)
	}
}
