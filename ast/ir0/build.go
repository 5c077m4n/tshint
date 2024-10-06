package ir0

import (
	"context"
	"errors"
	"log/slog"
	"time"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_typescript "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

var (
	ErrCreateParser = errors.New("could not create the requested parser")
	ErrBuildAST     = errors.New("failed to walk the input's AST")
)

func createTSParser() (*tree_sitter.Parser, error) {
	parser := tree_sitter.NewParser()

	err := parser.SetLanguage(
		tree_sitter.NewLanguage(tree_sitter_typescript.LanguageTypescript()),
	)
	if err != nil {
		return nil, errors.Join(ErrCreateParser, err)
	}

	return parser, nil
}

func Build(sourceCode []byte) (*Block, error) {
	parser, err := createTSParser()
	if err != nil {
		return nil, errors.Join(ErrBuildAST, err)
	}
	defer parser.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	tree := parser.ParseCtx(ctx, sourceCode, nil)
	defer tree.Close()

	cursor := tree.Walk()
	defer cursor.Close()

	visitor := &astVisitor{Program: &Block{}}

	for node := range treesitterIter(cursor) {
		content := sourceCode[node.StartByte():node.EndByte()]

		if node.IsNamed() && node.ChildCount() == 0 {
			slog.Info(
				"TS Node",
				"named", node.IsNamed(),
				"extra", node.IsExtra(),
				"missing", node.IsMissing(),
				"kind", node.Kind(),
				"content", content,
			)

			astNode, err := From(node, sourceCode)
			if err != nil {
				return nil, errors.Join(ErrBuildAST, err)
			}

			if err := astNode.Accept(visitor); err != nil {
				return nil, errors.Join(ErrBuildAST, err)
			}
		}
	}

	return visitor.Program, nil
}
