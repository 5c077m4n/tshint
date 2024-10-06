package ir0

type astVisitor struct {
	Program *Block
	index   uint
}

func (a *astVisitor) GetCurrentStatement() Node {
	return a.Program.Statements[a.index]
}

// IncIndex increase the current statement index
func (a *astVisitor) IncIndex() {
	a.index++
}

// VisitAny implements Visitor.
func (a *astVisitor) VisitAny(any *Any) error {
	a.Program.Statements = append(a.Program.Statements, any)
	return nil
}

// VisitBoolean implements Visitor.
func (a *astVisitor) VisitBoolean(b *Boolean) error {
	a.Program.Statements = append(a.Program.Statements, b)
	return nil
}

// VisitBooleanPrimitive implements Visitor.
func (a *astVisitor) VisitBooleanPrimitive(b *BooleanPrimitive) error {
	a.Program.Statements = append(a.Program.Statements, b)
	return nil
}

// VisitFunction implements Visitor.
func (a *astVisitor) VisitFunction(*Function) error {
	panic("unimplemented")
}

// VisitIdentifier implements Visitor.
func (a *astVisitor) VisitIdentifier(i *Identifier) error {
	a.Program.Statements = append(a.Program.Statements, i)
	return nil
}

// VisitIntersection implements Visitor.
func (a *astVisitor) VisitIntersection(*Intersection) error {
	panic("unimplemented")
}

// VisitNever implements Visitor.
func (a *astVisitor) VisitNever(n *Never) error {
	a.Program.Statements = append(a.Program.Statements, n)
	return nil
}

// VisitNull implements Visitor.
func (a *astVisitor) VisitNull(n *Null) error {
	a.Program.Statements = append(a.Program.Statements, n)
	return nil
}

// VisitNumber implements Visitor.
func (a *astVisitor) VisitNumber(n *Number) error {
	a.Program.Statements = append(a.Program.Statements, n)
	return nil
}

// VisitNumber implements Visitor.
func (a *astVisitor) VisitNumberPrimitive(n *NumberPrimitive) error {
	a.Program.Statements = append(a.Program.Statements, n)
	return nil
}

// VisitProgram implements Visitor.
func (a *astVisitor) VisitProgram(*Block) error {
	panic("unimplemented")
}

// VisitString implements Visitor.
func (a *astVisitor) VisitString(s *String) error {
	a.Program.Statements = append(a.Program.Statements, s)
	return nil
}

// VisitStringPrimitive implements Visitor.
func (a *astVisitor) VisitStringPrimitive(s *StringPrimitive) error {
	a.Program.Statements = append(a.Program.Statements, s)
	return nil
}

// VisitUndefined implements Visitor.
func (a *astVisitor) VisitUndefined(u *Undefined) error {
	a.Program.Statements = append(a.Program.Statements, u)
	return nil
}

// VisitUnion implements Visitor.
func (a *astVisitor) VisitUnion(*Union) error {
	panic("unimplemented")
}

// VisitUnknown implements Visitor.
func (a *astVisitor) VisitUnknown(u *Unknown) error {
	a.Program.Statements = append(a.Program.Statements, u)
	return nil
}

// VisitUnknown implements Visitor.
func (a *astVisitor) VisitIllegal(i *Illegal) error {
	a.Program.Statements = append(a.Program.Statements, i)
	return nil
}
