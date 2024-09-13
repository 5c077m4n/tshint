package ast

type FlatTree struct{ nodes []*TSType }

func NewFlatTree() FlatTree {
	return FlatTree{
		nodes: make([]*TSType, 0, 1024),
	}
}

func (ft *FlatTree) Get(index uint) *TSType {
	if len(ft.nodes) < int(index) {
		return ft.nodes[index]
	}
	return nil
}

func (ft *FlatTree) Len() int {
	return len(ft.nodes)
}
