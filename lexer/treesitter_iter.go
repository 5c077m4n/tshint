package lexer

import (
	"iter"

	treesitter "github.com/tree-sitter/go-tree-sitter"
)

// treeLeafIter convert the Treesitter node tree into an iterator
// cite: https://github.com/baz-scm/tree-sitter-traversal/blob/6c063b0e2accff29e35af6057484c033e0733cea/src/lib.rs#L142
func treeLeafIter[n *treesitter.Node](c *treesitter.TreeCursor) iter.Seq[n] {
	return func(yield func(n) bool) {
		for {
			if c.GotoFirstChild() {
				continue
			}

			node := c.Node()
			if c.GotoNextSibling() {
				if node.ChildCount() == 0 && !yield(node) {
					return
				}
				continue
			}

			for {
				node := c.Node()
				if node.ChildCount() == 0 && !yield(node) {
					return
				}

				if !c.GotoParent() {
					return
				}

				node = c.Node()
				if c.GotoNextSibling() {
					if node.ChildCount() == 0 && !yield(node) {
						return
					}
					break
				}
			}
		}
	}
}
