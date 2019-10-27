package token

type TokenType int

const (
	INTEGER TokenType = iota
	PLUS
	MINUS
	MUL
	DIV
	LT
	EQ
	LOGIC
	VARIABLE
)

type Token interface {
	GetType() TokenType
}
