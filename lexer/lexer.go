package lexer

import (
	"context"
	"errors"
	"log/slog"
	"time"

	treesitter "github.com/tree-sitter/go-tree-sitter"
	treesitterTypescript "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

var ErrTokenize = errors.New("failed to tokenize the input")

func Tokenize(sourceCode []byte) error {
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

	for leaf := range treeLeafIter(cursor) {
		start, end := leaf.ByteRange()
		content := sourceCode[start:end]

		slog.Info(
			"node",
			"kind", leaf.Kind(),
			"content", content,
		)
	}

	return nil
}
