// Package main
package main

import (
	"flag"
	"log/slog"
	"tshint/ast/ir0"
)

func main() {
	source := flag.String("eval", "", "Source code to evaluate")
	flag.Parse()

	sourceCode := []byte(*source)

	tree, err := ir0.Build(sourceCode)
	if err != nil {
		panic(err)
	}

	slog.Info(
		"Tree",
		"value", tree,
	)
}
