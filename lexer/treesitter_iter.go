package lexer

import (
	"iter"

	treesitter "github.com/tree-sitter/go-tree-sitter"
)

// treesitterIter convert the Treesitter node tree into an iterator
// cite: https://github.com/baz-scm/tree-sitter-traversal/blob/6c063b0e2accff29e35af6057484c033e0733cea/src/lib.rs#L142
func treesitterIter[N *treesitter.Node](c *treesitter.TreeCursor) iter.Seq[N] {
	return func(yield func(N) bool) {
		for {
			if c.GotoFirstChild() {
				continue
			}

			node := c.Node()
			if c.GotoNextSibling() {
				if !yield(node) {
					return
				}
				continue
			}

			for {
				if !yield(c.Node()) {
					return
				}

				if !c.GotoParent() {
					return
				}

				node := c.Node()
				if c.GotoNextSibling() {
					if !yield(node) {
						return
					}
					break
				}
			}
		}
	}
}
