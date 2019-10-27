package token

type MulToken struct {}

func (t *MulToken) GetType() TokenType {
	return MUL
}
