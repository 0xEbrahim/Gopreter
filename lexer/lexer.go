package lexer

import "unicode/utf8"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(code string) *Lexer {
	return &Lexer{input: code}
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= utf8.RuneCountInString(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}
