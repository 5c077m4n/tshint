package lexer

import (
	"context"
	"errors"
	"iter"
	"time"

	treesitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

var parser = func() *treesitter.Parser {
	p := treesitter.NewParser()
	p.SetLanguage(typescript.GetLanguage())
	return p
}()

var ErrTokenize = errors.New("failed to tokenize the input")

func Tokenize(sourceCode []byte) (iter.Seq[*treesitter.Node], error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	tree, err := parser.ParseCtx(ctx, nil, sourceCode)
	if err != nil {
		return nil, errors.Join(ErrTokenize, err)
	}

	return TreeLeafIter(tree), nil
}
