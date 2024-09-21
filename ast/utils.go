package ast

import (
	"cmp"
	"slices"
)

func uniques[Item cmp.Ordered, S ~[]Item](s S) S {
	slices.Sort(s)
	uniqItems := slices.Compact(s)

	return uniqItems
}
