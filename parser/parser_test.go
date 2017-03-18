package parser

import (
	"testing"

	"github.com/gregbair/bfgo/ast"
	"github.com/gregbair/bfgo/lexer"
)

func TestParsingPointerMovements(t *testing.T) {
	input := "<<<>>>"

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()

	if len(input) != len(prog.Statements) {
		t.Fatalf("Wrong count. Expected %d, got %d", len(input), len(prog.Statements))
	}

	stmts := prog.Statements
	for i := range input {
		var dir *ast.PointerStatement
		if input[i] == '<' {
			dir = leftPointer
		} else {
			dir = rightPointer
		}
		testStatement(t, stmts[i], dir)
	}
}

func TestParsingManipulation(t *testing.T) {
	input := "+-+-+-"

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()

	if len(input) != len(prog.Statements) {
		t.Fatalf("Wrong count. Expected %d, got %d", len(input), len(prog.Statements))
	}

	stmts := prog.Statements
	for i := range input {
		var dir *ast.ManipulationStatement
		if input[i] == '+' {
			dir = upPointer
		} else {
			dir = downPointer
		}
		testStatement(t, stmts[i], dir)

	}
}

func testStatement(t *testing.T, actual ast.Statement, expected ast.Statement) {
	if actual != expected {
		t.Fatalf("Wrong! expected %v, got %v", actual, expected)
	}
}

func TestLoops(t *testing.T) {
	input := "[>>++<<--]++"

	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()

	if len(prog.Statements) != 3 {
		t.Fatalf("Wrong count. Expected 3, got %d", len(prog.Statements))
	}

	loop, ok := prog.Statements[0].(*ast.LoopStatement)
	if !ok {
		t.Fatalf("Wrong! expected *ast.LoopStatement, got %T", prog.Statements[0])
	}

	for i := 0; i < 4; i++ {
		stmt := loop.Statements[i]
		switch input[i+1] {
		case '>':
			testStatement(t, stmt, rightPointer)
		case '<':
			testStatement(t, stmt, leftPointer)
		case '+':
			testStatement(t, stmt, upPointer)
		case '-':
			testStatement(t, stmt, downPointer)
		}
	}
}
