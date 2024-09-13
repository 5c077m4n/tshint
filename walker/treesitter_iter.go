package walker

import (
	"iter"

	treesitter "github.com/smacker/go-tree-sitter"
)

// TreeLeafIter convert the Treesitter node tree into an iterator
func TreeLeafIter[n *treesitter.Node](t *treesitter.Tree) iter.Seq[n] {
	return func(yield func(n) bool) {
		node := t.RootNode()

		for {
			if node.ChildCount() != 0 {
				node = node.Child(0)
				continue
			}
			if sibling := node.NextSibling(); sibling != nil {
				if node.ChildCount() == 0 && !yield(node) {
					return
				}

				node = sibling
				continue
			}

			for {
				if node.ChildCount() == 0 && !yield(node) {
					return
				}

				if node = node.Parent(); node == nil {
					return
				}
				if sibling := node.NextSibling(); sibling != nil {
					if node.ChildCount() == 0 && !yield(node) {
						return
					}

					node = sibling
					break
				}
			}
		}
	}
}
