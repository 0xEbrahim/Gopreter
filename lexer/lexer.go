package lexer

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
