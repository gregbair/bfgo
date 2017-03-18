package lexer

import (
	"testing"

	"github.com/gregbair/bfgo/token"
)

func TestNextToken(t *testing.T) {
	input := `<>+-.,[]`

	testCases := []*token.Token{
		token.NewFromByte('<'),
		token.NewFromByte('>'),
		token.NewFromByte('+'),
		token.NewFromByte('-'),
		token.NewFromByte('.'),
		token.NewFromByte(','),
		token.NewFromByte('['),
		token.NewFromByte(']'),
		token.NewFromByte(token.EOFLITERAL),
	}

	l := New(input)
	for _, tt := range testCases {
		tok := l.NextToken()
		if tt != tok {
			t.Errorf("Unexpected token: expected %#v, got %#v", tt, tok)
		}
	}
}
