package token

// Type is the "type" of token we've got.
type Type int

// The different token types
const (
	None Type = iota
	Plus
	Minus
	LessThan
	GreaterThan
	Dot
	Comma
	LBracket
	RBracket
	EOF
)

// EOFLITERAL represents the Literal for a EOF token
const EOFLITERAL = byte(1)

// NONELITERAL represents the Literal for a None token
const NONELITERAL = byte(0)

var tokenMap = map[byte]*Token{
	'+':         newFromType(Plus, '+'),
	'-':         newFromType(Minus, '-'),
	'<':         newFromType(LessThan, '<'),
	'>':         newFromType(GreaterThan, '>'),
	'.':         newFromType(Dot, '.'),
	',':         newFromType(Comma, ','),
	'[':         newFromType(LBracket, '['),
	']':         newFromType(RBracket, ']'),
	EOFLITERAL:  newFromType(EOF, EOFLITERAL),
	NONELITERAL: newFromType(None, NONELITERAL),
}

// Token represents a token in BF.
type Token struct {
	Type    Type
	Literal byte
}

func newFromType(t Type, l byte) *Token {
	return &Token{Type: t, Literal: l}
}

// NewFromByte returns a Token pointer based on the given byte.
func NewFromByte(b byte) *Token {
	if tok, ok := tokenMap[b]; ok {
		return tok
	}
	return &Token{Type: None}
}
