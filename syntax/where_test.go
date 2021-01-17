package syntax_test

import (
	"errors"
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestWhere_Name(t *testing.T) {
	w := new(syntax.Where)
	assert.Equal(t, "WHERE", syntax.WhereName(w))
}

func TestWhere_Build(t *testing.T) {
	testCases := []struct {
		Where  *syntax.Where
		Result *syntax.StmtSet
	}{
		{
			&syntax.Where{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "WHERE", Value: "lhs = 10"},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Where.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewWhere(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.Where
	}{
		{"lhs = ?", []interface{}{"rhs"}, &syntax.Where{Expr: "lhs = ?", Values: []interface{}{"rhs"}}},
	}

	for _, testCase := range testCases {
		res := syntax.NewWhere(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestAnd_Name(t *testing.T) {
	a := new(syntax.And)
	assert.Equal(t, "AND", syntax.AndName(a))
}

func TestAnd_Build(t *testing.T) {
	testCases := []struct {
		And    *syntax.And
		Result *syntax.StmtSet
	}{
		{
			&syntax.And{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "AND", Value: "lhs = 10", Parens: true},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.And.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewAdd(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.And
	}{
		{"lhs = ?", []interface{}{"rhs"}, &syntax.And{Expr: "lhs = ?", Values: []interface{}{"rhs"}}},
	}

	for _, testCase := range testCases {
		res := syntax.NewAnd(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestOr_Name(t *testing.T) {
	o := new(syntax.Or)
	assert.Equal(t, "OR", syntax.OrName(o))
}

func TestOr_Build(t *testing.T) {
	testCases := []struct {
		Or     *syntax.Or
		Result *syntax.StmtSet
	}{
		{
			&syntax.Or{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "OR", Value: "lhs = 10", Parens: true},
		},
	}

	for _, testCase := range testCases {
		res, _ := testCase.Or.Build()
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewOr(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.Or
	}{
		{"lhs = ?", []interface{}{"rhs"}, &syntax.Or{Expr: "lhs = ?", Values: []interface{}{"rhs"}}},
	}

	for _, testCase := range testCases {
		res := syntax.NewOr(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestBuildStmtSet(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.StmtSet
	}{
		{
			"lhs = rhs",
			[]interface{}{},
			&syntax.StmtSet{Value: "lhs = rhs"},
		},
		{
			"lhs = ?",
			[]interface{}{"rhs"},
			&syntax.StmtSet{Value: "lhs = rhs"},
		},
		{
			"lhs = ?",
			[]interface{}{100},
			&syntax.StmtSet{Value: "lhs = 100"},
		},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{"rhs", 100},
			&syntax.StmtSet{Value: "lhs1 = rhs AND lhs2 = 100"},
		},
		{
			"IN lhs (?, ?, ?)",
			[]interface{}{"rhs", 100, true},
			&syntax.StmtSet{Value: "IN lhs (rhs, 100, true)"},
		},
		{
			"lhs LIKE %%?%%",
			[]interface{}{"rhs"},
			&syntax.StmtSet{Value: "lhs LIKE %rhs%"},
		},
		{
			"lhs BETWEEN ? AND ?",
			[]interface{}{10, 100},
			&syntax.StmtSet{Value: "lhs BETWEEN 10 AND 100"},
		},
	}

	for _, testCase := range testCases {
		res, _ := syntax.BuildStmtSet(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestBuildStmtSet_Fail(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Error  error
	}{
		{
			"lhs = ? AND rhs = ?",
			[]interface{}{10},
			internal.NewError(
				syntax.OpBuildStmtSet,
				internal.KindBasic,
				errors.New("Length of values is not valid"),
			),
		},
		{
			"lhs = ?",
			[]interface{}{[]string{}},
			internal.NewError(
				internal.OpToString,
				internal.KindType,
				errors.New("type is invalid"),
			),
		},
	}

	for _, testCase := range testCases {
		_, err := syntax.BuildStmtSet(testCase.Expr, testCase.Values...)
		if err == nil {
			t.Errorf("Error is not occurred")
			continue
		}

		e, ok := err.(*internal.Error)
		if !ok {
			t.Errorf("Error type is invalid")
			continue
		}

		if diff := internal.CmpError(*e, *testCase.Error.(*internal.Error)); diff != "" {
			t.Errorf(diff)
		}
	}
}
