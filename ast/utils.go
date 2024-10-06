package ast

import (
	"cmp"
	"slices"
)

func uniques[Item cmp.Ordered, S ~[]Item](s S) S {
	slices.Sort(s)
	return slices.Compact(s)
}
