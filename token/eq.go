package token

type EqToken struct {}

func (t *EqToken) GetType() TokenType {
	return EQ
}
