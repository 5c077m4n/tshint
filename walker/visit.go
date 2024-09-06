package walker

import (
	"tshint/ast"
)

type TSVisitor struct {
	tree ast.TSType
}

func (tsV *TSVisitor) VisitNumber(n *ast.Number) error {
	tsV.tree = n
	return nil
}

func (tsV *TSVisitor) VisitString(s *ast.String) error {
	tsV.tree = s
	return nil
}

func (tsV *TSVisitor) VisitBoolean(b *ast.Boolean) error {
	tsV.tree = b
	return nil
}

func (tsV *TSVisitor) VisitNull(n *ast.Null) error {
	tsV.tree = n
	return nil
}

func (tsV *TSVisitor) VisitUndefined(u *ast.Undefined) error {
	tsV.tree = u
	return nil
}

func (tsV *TSVisitor) VisitAny(a *ast.Any) error {
	tsV.tree = a
	return nil
}

func (tsV *TSVisitor) VisitUnknown(u *ast.Unknown) error {
	tsV.tree = u
	return nil
}

func (tsV *TSVisitor) VisitNever(n *ast.Never) error {
	tsV.tree = n
	return nil
}

func (tsV *TSVisitor) VisitUnion(u *ast.Union) error {
	tsV.tree = u
	return nil
}

func (tsV *TSVisitor) VisitIntersection(i *ast.Intersection) error {
	tsV.tree = i
	return nil
}

func (tsV *TSVisitor) GetTree() ast.TSType {
	return tsV.tree
}
