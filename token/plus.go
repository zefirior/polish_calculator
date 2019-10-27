package token

type PlusToken struct {}

func (t *PlusToken) GetType() TokenType {
	return PLUS
}
