package token

type MinusToken struct {}

func (t *MinusToken) GetType() TokenType {
	return MINUS
}

