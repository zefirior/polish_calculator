package token

type LogicToken struct {}

func (t *LogicToken) GetType() TokenType {
	return LOGIC
}
