// Package main
package main

import (
	"flag"
	"log/slog"
	"tshint/lexer"
)

func main() {
	source := flag.String("eval", "", "Source code to evaluate")
	flag.Parse()

	sourceCode := []byte(*source)

	tokenIter, err := lexer.Tokenize(sourceCode)
	if err != nil {
		panic(err)
	}

	for token := range tokenIter {
		slog.Info(
			"node",
			"type", token.Type(),
			"content", token.Content(sourceCode),
		)
	}
}
