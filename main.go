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

	tokenList, err := lexer.Tokenize(sourceCode)
	if err != nil {
		panic(err)
	}

	for _, token := range tokenList {
		slog.Info(
			"Token",
			"kind", token.Kind,
			"content", token.Content(sourceCode),
		)
	}
}
