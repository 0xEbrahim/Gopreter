package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	PLUS      = "+"
	COMMA     = ","
	ASSIGN    = "="
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal []rune
}

func NewToken(Type TokenType, Literal []rune) Token {
	return Token{Type: Type, Literal: Literal}
}
