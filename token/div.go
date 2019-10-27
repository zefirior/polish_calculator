package token

type DivToken struct {}

func (t *DivToken) GetType() TokenType {
	return DIV
}
