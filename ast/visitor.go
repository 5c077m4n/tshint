package ast

type Visitor interface {
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
}
