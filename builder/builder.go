package builder

import (
	"price_calculation/expression"
	"price_calculation/token"
)

type Builder struct {
	stack []expression.Expression
}

func (b *Builder) Build(tokens []token.Token) expression.Expression {
	if len(tokens) == 0 {
		return nil
	}

	b.clear()
	for _, _token := range tokens {
		switch _token.GetType() {
		case token.INTEGER:
			b.processInteger(_token)
		case token.VARIABLE:
			b.processVariable(_token)
		case token.PLUS:
			b.processPlus(_token)
		case token.MINUS:
			b.processMinus(_token)
		case token.MUL:
			b.processMul(_token)
		case token.DIV:
			b.processDiv(_token)
		case token.LT:
			b.processLt(_token)
		case token.EQ:
			b.processEq(_token)
		case token.LOGIC:
			b.processLogic(_token)
		}
	}
	return b.pop()
}

func (b *Builder) processInteger(_token token.Token) {
	value := _token.(*token.IntegerToken).Value()
	b.stack = append(b.stack, expression.NewConstExpr(value))
}

func (b *Builder) processVariable(_token token.Token) {
	value := _token.(*token.VariableToken).Value()
	b.stack = append(b.stack, expression.NewVariableExpr(value))
}

func (b *Builder) processPlus(_token token.Token) {
	right := b.pop()
	left := b.pop()
	b.stack = append(b.stack, expression.NewPlusExpr(left, right))
}

func (b *Builder) processMinus(_token token.Token) {
	right := b.pop()
	left := b.pop()
	b.stack = append(b.stack, expression.NewMinusExpr(left, right))
}

func (b *Builder) processMul(_token token.Token) {
	right := b.pop()
	left := b.pop()
	b.stack = append(b.stack, expression.NewMulExpr(left, right))
}

func (b *Builder) processDiv(_token token.Token) {
	right := b.pop()
	left := b.pop()
	b.stack = append(b.stack, expression.NewDivExpr(left, right))
}

func (b *Builder) processLt(_token token.Token) {
	right := b.pop()
	left := b.pop()
	b.stack = append(b.stack, expression.NewLtExpr(left, right))
}

func (b *Builder) processEq(_token token.Token) {
	right := b.pop()
	left := b.pop()
	b.stack = append(b.stack, expression.NewEqExpr(left, right))
}

func (b *Builder) processLogic(_token token.Token) {
	right := b.pop()
	left := b.pop()
	logic := b.pop()
	b.stack = append(b.stack, expression.NewLogicExpr(logic, left, right))
}

func (b *Builder) pop() (expr expression.Expression) {
	end := len(b.stack) - 1
	expr = b.stack[end]
	b.stack = b.stack[:end]
	return
}

func (b *Builder) clear() {
	b.stack = b.stack[:0]
}
