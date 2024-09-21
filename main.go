// Package main
package main

import (
	"flag"
	"log/slog"
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

	for token := range chanToken {
		slog.Info(
			"Token",
			"kind", token.Kind,
			"content", token.Content(sourceCode),
		)
	}
}
