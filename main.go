// Package main
package main

import (
	"flag"
	"log/slog"
	"tshint/ast"
	"tshint/lexer"
	"tshint/lexer/token"
)

func main() {
	source := flag.String("eval", "", "Source code to evaluate")
	flag.Parse()

	sourceCode := []byte(*source)

	chanToken := make(chan token.Token)
	go func() {
		err := lexer.Tokenize(sourceCode, chanToken)
		if err != nil {
			panic(err)
		}
	}()

	tree, err := ast.From(sourceCode, chanToken)
	if err != nil {
		panic(err)
	}

	slog.Info(
		"AST",
		"tree", tree,
	)
}
