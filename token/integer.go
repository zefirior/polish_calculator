package token

import "strconv"

type IntegerToken struct {
	rawValue string
}

func NewIntegerToken(value string) *IntegerToken {
	return &IntegerToken{value}
}

func (t *IntegerToken) Value() int {
	r, _ := strconv.Atoi(t.rawValue)
	return r
}

func (t *IntegerToken) GetType() TokenType {
	return INTEGER
}
