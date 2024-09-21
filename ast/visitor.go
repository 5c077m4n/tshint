package ast

type Visitor interface {
	VisitProgram(*Program) error
	VisitNumber(*Number) error
	VisitBoolean(*Boolean) error
	VisitString(*String) error
	VisitNull(*Null) error
	VisitUndefined(*Undefined) error
	VisitAny(*Any) error
	VisitUnknown(*Unknown) error
	VisitNever(*Never) error
	VisitUnion(*Union) error
	VisitIntersection(*Intersection) error
	VisitFunction(*Function) error
	VisitIdentifier(*Identifier) error
}
