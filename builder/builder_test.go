package builder

import (
	"github.com/stretchr/testify/assert"
	"price_calculation/expression"
	"price_calculation/parser"
	"price_calculation/storage"
	"testing"
)

func TestBuilder_pop(t *testing.T) {
	var expr *expression.ConstExpr

	b := Builder{stack: []expression.Expression{
		expression.NewConstExpr(10),
		expression.NewConstExpr(20),
	}}

	expr = b.pop().(*expression.ConstExpr)
	assert.Equal(t, 20, expr.Value(), "Must returned 20")

	expr = b.pop().(*expression.ConstExpr)
	assert.Equal(t, 10, expr.Value(), "Must returned 10")
}

type tester struct {
	t *testing.T
	expr expression.Expression
}

func (t *tester) prepare(input string) {
	p, b := parser.Parser{}, Builder{}
	_, tokens := p.Parse(input)
	t.expr = b.Build(tokens)
}

func (t *tester) assertCalculation(variables string, expected int) {
	s := storage.FromString(variables)
	assert.Equal(t.t, expected, t.expr.Evaluate(s))
}

func TestBuilder_plus(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("10 20 +")
	_tester.assertCalculation("", 30)
}

func TestBuilder_minus(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("10 20 -")
	_tester.assertCalculation("", -10)
}

func TestBuilder_mul(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("10 20 *")
	_tester.assertCalculation("", 200)
}

func TestBuilder_div(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("41 10 /")
	_tester.assertCalculation("", 4)
}

func TestBuilder_lt(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("100 10 <")
	_tester.assertCalculation("", 0)
}

func TestBuilder_eq(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("10 10 =")
	_tester.assertCalculation("", 1)
}

func TestBuilder_ne(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("100 10 =")
	_tester.assertCalculation("", 0)
}

func TestBuilder_logic_left(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("1 100 10 ?")
	_tester.assertCalculation("", 100)
}

func TestBuilder_logic_right(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("0 100 10 ?")
	_tester.assertCalculation("", 10)
}

func TestBuilder_variable(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("a b +")
	_tester.assertCalculation("100 10", 110)
	_tester.assertCalculation("101 10", 111)
}

func TestBuilder_complex_trie(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("a a * 43 + 2 * b 2 / -") // (a**2 + 43) * 2 - b / 2
	_tester.assertCalculation("1 2", 87)
	_tester.assertCalculation("2 4", 92)
}

func TestBuilder_big_stack(t *testing.T) {
	_tester := tester{t: t}
	_tester.prepare("a a a a a a a a a a a a a a a a + + + + + + + + + + + + + + +") // (a**2 + 43) * 2 - b / 2
	_tester.assertCalculation("1", 16)
}
