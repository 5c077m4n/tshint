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

var ErrWalk = errors.New("failed to walk the AST")

func Walk(sourceCode []byte) error {
	parser.SetLanguage(typescript.GetLanguage())

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	tree, err := parser.ParseCtx(ctx, nil, sourceCode)
	if err != nil {
		return errors.Join(ErrWalk, err)
	}

	for n := range TreeLeafIter(tree) {
		slog.Info(
			"node",
			// "value", n,
			// "type", n.Type(),
			"content", n.Content(sourceCode),
		)
	}

	return nil
}
