package lexer

import (
	"context"
	"errors"
	"log/slog"
	"time"
	"tshint/lexer/token"

	treesitter "github.com/tree-sitter/go-tree-sitter"
	treesitterTypescript "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

var ErrTokenize = errors.New("failed to tokenize the input")

func Tokenize(sourceCode []byte, chanToken chan<- token.Token) error {
	defer close(chanToken)

	parser := treesitter.NewParser()
	defer parser.Close()

	err := parser.SetLanguage(treesitter.NewLanguage(treesitterTypescript.LanguageTypescript()))
	if err != nil {
		return errors.Join(ErrTokenize, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	tree := parser.ParseCtx(ctx, sourceCode, nil)
	defer tree.Close()

	cursor := tree.Walk()
	defer cursor.Close()

	for node := range treesitterIter(cursor) {
		t := token.From(node)
		if t.IsError {
			slog.Warn(
				"Syntax error",
				"token", t,
			)
		}

		if node.ChildCount() == 0 {
			chanToken <- t
		}
	}

	return nil
}
