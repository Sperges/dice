package main

import (
	"os"
	"roll/src/parser"
	"roll/src/token"
)

func main() {
	stream := token.NewStream(os.Args[1])
	// println(stream.String())
	if ast, err := parser.ParseStream(stream, -1); err != nil {
		println(err.Error())
	} else {
		// println(ast.String())
		println(ast.EvaluateAndSum())
	}
}
