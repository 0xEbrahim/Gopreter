package lexer

import "Gopreter/token"

type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
}

func New(code string) *Lexer {
	l := &Lexer{
		input: []rune(code),
	}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	runes := make([]rune, 0)
	runes = append(runes, l.ch)
	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, runes)
	case '+':
		tok = token.NewToken(token.PLUS, runes)
	case ',':
		tok = token.NewToken(token.COMMA, runes)
	case ';':
		tok = token.NewToken(token.SEMICOLON, runes)
	case '(':
		tok = token.NewToken(token.LPAREN, runes)
	case ')':
		tok = token.NewToken(token.RPAREN, runes)
	case '{':
		tok = token.NewToken(token.LBRACE, runes)
	case '}':
		tok = token.NewToken(token.RBRACE, runes)
	case 0:
		tok = token.NewToken(token.EOF, []rune(""))
	}
	l.ReadChar()
	return tok
}
