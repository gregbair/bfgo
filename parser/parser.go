package parser

import (
	"github.com/gregbair/bfgo/ast"
	"github.com/gregbair/bfgo/lexer"
	"github.com/gregbair/bfgo/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  *token.Token
	peekToken *token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

var rightPointer = &ast.PointerStatement{Token: token.NewFromByte('>'), Direction: ast.Right}
var leftPointer = &ast.PointerStatement{Token: token.NewFromByte('<'), Direction: ast.Left}
var upPointer = &ast.ManipulationStatement{Token: token.NewFromByte('+'), Direction: ast.Up}
var downPointer = &ast.ManipulationStatement{Token: token.NewFromByte('-'), Direction: ast.Down}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LessThan:
		return leftPointer
	case token.GreaterThan:
		return rightPointer
	case token.Plus:
		return upPointer
	case token.Minus:
		return downPointer
	case token.LBracket:
		return p.parseLoop()
	default:
		return nil
	}
}

func (p *Parser) parseLoop() *ast.LoopStatement {
	loop := &ast.LoopStatement{Token: p.curToken}
	p.nextToken()

	for p.curToken.Type != token.RBracket {
		loop.Statements = append(loop.Statements, p.parseStatement())
		p.nextToken()
	}
	return loop
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.nextToken()
	}

	return prog
}
