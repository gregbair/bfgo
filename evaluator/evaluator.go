package evaluator

import "github.com/gregbair/bfgo/ast"

type environment struct {
	state  []int
	curPos int
}

func newEnvironment() *environment {
	env := &environment{}

	env.state = make([]int, 30000)
	env.curPos = 0
	return env
}

func Eval(prog ast.Node) *environment {
	env := newEnvironment()
	return evalWithEnvironment(prog, env)
}

func evalWithEnvironment(prog ast.Node, env *environment) *environment {
	switch node := prog.(type) {
	case *ast.Program:
		evalProgram(node, env)
	case *ast.PointerStatement:
		evalPointerStatement(node, env)
	case *ast.ManipulationStatement:
		evalManipulationStatement(node, env)
	}

	return env
}

func evalProgram(prog *ast.Program, env *environment) {
	for _, s := range prog.Statements {
		evalWithEnvironment(s, env)
	}
}

func evalPointerStatement(pointer *ast.PointerStatement, env *environment) {
	switch pointer.Direction {
	case ast.Left:
		if env.curPos > 0 {
			env.curPos--
		}
	case ast.Right:
		if env.curPos < len(env.state) {
			env.curPos++
		}
	}
}

func evalManipulationStatement(stmt *ast.ManipulationStatement, env *environment) {
	switch stmt.Direction {
	case ast.Up:
		env.state[env.curPos]++
	case ast.Down:
		if env.state[env.curPos] > 0 {
			env.state[env.curPos]--
		}
	}
}
