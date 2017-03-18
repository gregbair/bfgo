package lexer

import "github.com/gregbair/bfgo/token"

// Lexer turns strings into tokens
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New returns a new lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken advances the tokens and returns the current one.
func (l *Lexer) NextToken() *token.Token {
	l.skipNonTokens()
	tok := token.NewFromByte(l.ch)
	l.readChar()
	return tok
}

func (l *Lexer) skipNonTokens() {
	for token.NewFromByte(l.ch).Type == token.None {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = token.EOFLITERAL
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}
