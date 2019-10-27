package expression

import (
	"github.com/stretchr/testify/assert"
	"price_calculation/storage"
	"testing"
)

type testCase struct {
	expected, left, right, logic int
}

func TestExpression_plus_evaluate(t *testing.T) {
	s := storage.NewStorage()
	expr := PlusExpr{
		&ConstExpr{10},
		&ConstExpr{10},
	}
	assert.Equal(t, expr.Evaluate(s), 20)
}

func TestExpression_deep_evaluate(t *testing.T) {
	s := storage.NewStorage()
	expr := PlusExpr{
		&PlusExpr{
			&ConstExpr{10},
			&ConstExpr{10},
		},
		&ConstExpr{10},
	}
	assert.Equal(t, expr.Evaluate(s), 30)
}

func TestExpression_mul_evaluate(t *testing.T) {
	s := storage.NewStorage()
	expr := MulExpr{
		&ConstExpr{10},
		&ConstExpr{10},
	}
	assert.Equal(t, expr.Evaluate(s), 100)
}

func TestExpression_div_evaluate(t *testing.T) {
	s := storage.NewStorage()

	cases := []testCase{
		{expected: 0, left: 1, right: 10},
		{expected: 1, left: 19, right: 10},
		{expected: 2, left: 20, right: 10},
		{expected: 2, left: 21, right: 10},
	}

	for _, _case := range cases {
		expr := DivExpr{
			&ConstExpr{_case.left},
			&ConstExpr{_case.right},
		}
		assert.Equal(t, expr.Evaluate(s), _case.expected)
	}
}

func TestExpression_eq_evaluate(t *testing.T) {
	s := storage.NewStorage()

	cases := []testCase{
		{expected: 0, left: 9, right: 10},
		{expected: 1, left: 10, right: 10},
		{expected: 0, left: 11, right: 10},
	}

	for _, _case := range cases {
		expr := EqExpr{
			&ConstExpr{_case.left},
			&ConstExpr{_case.right},
		}
		assert.Equal(t, expr.Evaluate(s), _case.expected)
	}
}

func TestExpression_lt_evaluate(t *testing.T) {
	s := storage.NewStorage()

	cases := []testCase{
		{expected: 1, left: 9, right: 10},
		{expected: 0, left: 10, right: 10},
		{expected: 0, left: 11, right: 10},
	}

	for _, _case := range cases {
		expr := LtExpr{
			&ConstExpr{_case.left},
			&ConstExpr{_case.right},
		}
		assert.Equal(t, expr.Evaluate(s), _case.expected)
	}
}

func TestExpression_logic_evaluate(t *testing.T) {
	s := storage.NewStorage()

	cases := []testCase{
		{expected: 9, logic: 1, left: 9, right: 10},
		{expected: 10, logic: 0, left: 9, right: 10},
	}

	for _, _case := range cases {
		expr := LogicExpr{
			&ConstExpr{_case.logic},
			&ConstExpr{_case.left},
			&ConstExpr{_case.right},
		}
		assert.Equal(t, expr.Evaluate(s), _case.expected)
	}
}

func TestExpression_variable_evaluate(t *testing.T) {
	cases := []testCase{
		{expected: 2, left: 1},
		{expected: 3, left: 2},
	}

	for _, _case := range cases {
		s := storage.NewStorage()
		s.Set("a", _case.left)

		expr := PlusExpr{
			&VariableExpr{"a"},
			&ConstExpr{1},
		}
		assert.Equal(t, expr.Evaluate(s), _case.expected)
	}
}