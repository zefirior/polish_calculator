package token

type VariableToken struct {
	value string
}

func NewVariableToken(value string) *VariableToken {
	return &VariableToken{value}
}

func (t *VariableToken) Value() string {
	return t.value
}

func (t *VariableToken) GetType() TokenType {
	return VARIABLE
}
