package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(code string) *Lexer {
	return &Lexer{input: code}
}
