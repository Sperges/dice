package main

import (
	"roll/src/parser"
	"roll/src/token"
	"testing"
)

func TestSpace(t *testing.T) {
	stream := token.NewStream("2 + 2")
	if ast, err := parser.ParseStream(stream, 0); err != nil {
		t.Error(err)
	} else {
		eval := ast.Evaluate()

		if len(eval) > 1 {
			t.Errorf("len(eval) = %d; want len(eval) > 1", len(eval))
		}

		if eval[0] != 4 {
			t.Errorf("eval[0] = %d; want eval[0] = 4", eval)
		}
	}
}

func TestX(t *testing.T) {
	stream := token.NewStream("1d4")
	if ast, err := parser.ParseStream(stream, 0); err != nil {
		t.Error(err)
	} else {
		eval := ast.Evaluate()

		if len(eval) > 1 {
			t.Errorf("len(eval) = %d; want len(eval) > 1", len(eval))
		}

		if eval[0] > 4 || eval[0] < 0 {
			t.Errorf("eval[0] = %d; want 0 < eval[0] < 4", eval)
		}
	}
}

func TestUnaryD(t *testing.T) {
	stream := token.NewStream("1+d4")
	if ast, err := parser.ParseStream(stream, 0); err != nil {
		t.Error(err)
	} else {
		eval := ast.Evaluate()

		if len(eval) > 1 {
			t.Errorf("len(eval) = %d; want len(eval) > 1", len(eval))
		}

		if eval[0] > 5 || eval[0] < 1 {
			t.Errorf("eval[0] = %d; want 1 < eval[0] < 5", eval)
		}
	}
}

func TestKeepHighStream(t *testing.T) {
	got := token.NewStream("2d4kh1")
	if got.Tokens()[3].Kind != token.KeepHigh {
		t.Errorf("got.tokens[3] = %s; want KeepHigh", got.Tokens()[3].Kind.String())
	}
}

func TestKeepHighAST(t *testing.T) {
	stream := token.NewStream("2d4kh1")
	if ast, err := parser.ParseStream(stream, 0); err != nil {
		t.Error(err)
	} else {
		if ast.Kind() != parser.KeepHigh {
			t.Errorf("ast.lhs.kind = %s; want KeepHighNode", ast.Kind())
		}
	}

}
