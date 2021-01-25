package syntax_test

import (
	"testing"

	"github.com/champon1020/mgorm/internal"
	"github.com/champon1020/mgorm/syntax"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestWhen_Name(t *testing.T) {
	w := new(syntax.When)
	assert.Equal(t, "WHEN", syntax.WhenName(w))
}

func TestWhen_Build(t *testing.T) {
	testCases := []struct {
		When   *syntax.When
		Result *syntax.StmtSet
	}{
		{
			&syntax.When{Expr: "lhs = rhs"},
			&syntax.StmtSet{Clause: "WHEN", Value: "lhs = rhs"},
		},
		{
			&syntax.When{Expr: "lhs = ?", Values: []interface{}{10}},
			&syntax.StmtSet{Clause: "WHEN", Value: "lhs = 10"},
		},
		{
			&syntax.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
			&syntax.StmtSet{Clause: "WHEN", Value: `lhs1 = 10 AND lhs2 = "str"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.When.Build()
		if err != nil {
			t.Errorf("%v\n", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewWhen(t *testing.T) {
	testCases := []struct {
		Expr   string
		Values []interface{}
		Result *syntax.When
	}{
		{"lhs = rhs", nil, &syntax.When{Expr: "lhs = rhs"}},
		{"lhs = ?", []interface{}{10}, &syntax.When{Expr: "lhs = ?", Values: []interface{}{10}}},
		{
			"lhs1 = ? AND lhs2 = ?",
			[]interface{}{10, "str"},
			&syntax.When{Expr: "lhs1 = ? AND lhs2 = ?", Values: []interface{}{10, "str"}},
		},
	}

	for _, testCase := range testCases {
		res := syntax.NewWhen(testCase.Expr, testCase.Values...)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestThen_Name(t *testing.T) {
	w := new(syntax.Then)
	assert.Equal(t, "THEN", syntax.ThenName(w))
}

func TestThen_Build(t *testing.T) {
	testCases := []struct {
		Then   *syntax.Then
		Result *syntax.StmtSet
	}{
		{
			&syntax.Then{Value: 10},
			&syntax.StmtSet{Clause: "THEN", Value: "10"},
		},
		{
			&syntax.Then{Value: "str"},
			&syntax.StmtSet{Clause: "THEN", Value: `"str"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Then.Build()
		if err != nil {
			t.Errorf("%v\n", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewThen(t *testing.T) {
	testCases := []struct {
		Value  interface{}
		Result *syntax.Then
	}{
		{10, &syntax.Then{Value: 10}},
		{"str", &syntax.Then{Value: "str"}},
	}

	for _, testCase := range testCases {
		res := syntax.NewThen(testCase.Value)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestElse_Name(t *testing.T) {
	w := new(syntax.Else)
	assert.Equal(t, "ELSE", syntax.ElseName(w))
}

func TestElse_Build(t *testing.T) {
	testCases := []struct {
		Else   *syntax.Else
		Result *syntax.StmtSet
	}{
		{
			&syntax.Else{Value: 10},
			&syntax.StmtSet{Clause: "ELSE", Value: "10"},
		},
		{
			&syntax.Else{Value: "str"},
			&syntax.StmtSet{Clause: "ELSE", Value: `"str"`},
		},
	}

	for _, testCase := range testCases {
		res, err := testCase.Else.Build()
		if err != nil {
			t.Errorf("%v\n", err)
			continue
		}
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestNewElse(t *testing.T) {
	testCases := []struct {
		Value  interface{}
		Result *syntax.Else
	}{
		{10, &syntax.Else{Value: 10}},
		{"str", &syntax.Else{Value: "str"}},
	}

	for _, testCase := range testCases {
		res := syntax.NewElse(testCase.Value)
		if diff := cmp.Diff(res, testCase.Result); diff != "" {
			internal.PrintTestDiff(t, diff)
		}
	}
}
