package ast

import "github.com/gregbair/bfgo/token"

// Node is the base Statement
type Node interface {
	TokenLiteral() byte
}

type Statement interface {
	Node
	statementNode()
}

// Program is the root Node
type Program struct {
	Node
	Statements []Statement
}

// TokenLiteral returns the first literal on the statement.
func (p *Program) TokenLiteral() byte {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return byte(0)
}

// PointerDirection is either left or right
type PointerDirection int

// PointerDirection is either left or right
const (
	Left PointerDirection = iota
	Right
)

// PointerStatement represents a direction of a pointer (< or >)
type PointerStatement struct {
	Token     *token.Token
	Direction PointerDirection
}

func (ps *PointerStatement) statementNode() {}

// TokenLiteral returns the literal string of the token.
func (ps *PointerStatement) TokenLiteral() byte { return ps.Token.Literal }

type ManipulationDirection int

const (
	Up ManipulationDirection = iota
	Down
)

type ManipulationStatement struct {
	Token     *token.Token
	Direction ManipulationDirection
}

func (ms *ManipulationStatement) statementNode() {}

// TokenLiteral returns the literal string of the token.
func (ms *ManipulationStatement) TokenLiteral() byte { return ms.Token.Literal }

type LoopStatement struct {
	Token      *token.Token
	Statements []Statement
}

func (ls *LoopStatement) statementNode() {}

// TokenLiteral returns the literal string of the token.
func (ls *LoopStatement) TokenLiteral() byte { return ls.Token.Literal }
