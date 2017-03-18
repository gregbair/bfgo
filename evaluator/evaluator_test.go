package evaluator

import (
	"testing"

	"github.com/gregbair/bfgo/lexer"
	"github.com/gregbair/bfgo/parser"
)

func TestEnvironment(t *testing.T) {
	env := newEnvironment()
	state := env.state

	if len(state) != 30000 {
		t.Fatalf("Wrong length of state. Expected 30000, got %d", len(state))
	}

	foundNotZero := false

	for _, i := range state {
		if i != 0 {
			foundNotZero = true
			break
		}
	}

	if foundNotZero {
		t.Fatal("Found a non-zero initial value!")
	}

	if env.curPos != 0 {
		t.Fatalf("Wrong initial position. Got %d", env.curPos)
	}
}

func TestSimple(t *testing.T) {
	input := "++>++-<->++" // end should be [1][3]
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()

	env := Eval(prog)

	state := env.state

	if state[0] != 1 {
		t.Errorf("Expected state[0] to be 1. Got %d", state[0])
	}

	if state[1] != 3 {
		t.Errorf("Expected state[1] to be 3. Got %d", state[1])
	}
}

func TestLoop(t *testing.T) {
	input := `+++++ +++++            initialize counter (cell #0) to 10
[                       use loop to set 70/100/30/10
    > +++++ ++              add  7 to cell #1
    > +++++ +++++           add 10 to cell #2
    > +++                   add  3 to cell #3
    > +                     add  1 to cell #4
<<<< -                  decrement counter (cell #0)
]`

	expectedState := []int{0, 70, 100, 30, 10, 0}

	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()

	env := Eval(prog)

	state := env.state

	for i, v := range expectedState {
		if state[i] != v {
			t.Errorf("State wrong at index %d, expected %d, got %d", i, v, state[i])
		}
	}
}
