package parser

import (
	"price_calculation/token"
	"strconv"
	"strings"
)

type Parser struct {}

func (p *Parser) Parse(input string) (error, []token.Token) {
	tokens := make([]token.Token, 0)
	for _, value := range strings.Split(input, " ") {
		tokens = append(tokens, getToken(value))
	}
	return nil, tokens
}

func getToken(value string) token.Token {
	switch value {
	case "+":
		return &token.PlusToken{}
	case "-":
		return &token.MinusToken{}
	case "*":
		return &token.MulToken{}
	case "/":
		return &token.DivToken{}
	case "<":
		return &token.LtToken{}
	case "=":
		return &token.EqToken{}
	case "?":
		return &token.LogicToken{}
	default:
		if _, err := strconv.Atoi(value); err == nil {
			return token.NewIntegerToken(value)
		}
		return token.NewVariableToken(value)
	}
}