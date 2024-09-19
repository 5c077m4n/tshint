// Package main
package main

import (
	"flag"
	"tshint/lexer"
)

func main() {
	source := flag.String("eval", "", "Source code to evaluate")
	flag.Parse()

	sourceCode := []byte(*source)

	if err := lexer.Tokenize(sourceCode); err != nil {
		panic(err)
	}
}
