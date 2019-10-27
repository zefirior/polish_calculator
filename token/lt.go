package token

type LtToken struct {}

func (t *LtToken) GetType() TokenType {
	return LT
}
