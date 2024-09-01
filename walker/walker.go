package walker

import (
	"context"
	"errors"
	"log/slog"
	"time"

	treesitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

var parser = treesitter.NewParser()

func visit(node *treesitter.Node, source []byte) error {
	if node == nil || node.IsNull() {
		return nil
	}

	slog.Info(
		"node",
		"value", node,
		"type", node.Type(),
		"content", node.Content(source),
		"# of children", node.ChildCount(),
	)
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if err := visit(child, source); err != nil {
			return err
		}
	}

	return nil
}

var ErrWalk = errors.New("failed to walk the AST")

func Walk(sourceCode []byte) error {
	parser.SetLanguage(typescript.GetLanguage())

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	tree, err := parser.ParseCtx(ctx, nil, sourceCode)
	if err != nil {
		return errors.Join(ErrWalk, err)
	}

	root := tree.RootNode()
	if err := visit(root, sourceCode); err != nil {
		return errors.Join(ErrWalk, err)
	}

	return nil
}
