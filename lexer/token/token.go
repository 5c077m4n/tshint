package token

import (
	treesitter "github.com/tree-sitter/go-tree-sitter"
)

type Token struct {
	Kind      string
	ByteRange [2]uint
	IsError   bool
}

func (t *Token) Location(_sourceCode []byte) [2][2]uint {
	return [2][2]uint{{0, 0}, {0, 0}}
}

func (t *Token) Content(sourceCode []byte) []byte {
	return sourceCode[t.ByteRange[0]:t.ByteRange[1]]
}

func From(node *treesitter.Node) Token {
	start, end := node.ByteRange()
	return Token{
		Kind:      node.Kind(),
		ByteRange: [2]uint{start, end},
		IsError:   node.IsError(),
	}
}
