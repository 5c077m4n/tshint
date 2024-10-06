package ir0

type Visitor interface {
	VisitProgram(*Block) error
	VisitNumber(*Number) error
	VisitNumberPrimitive(*NumberPrimitive) error
	VisitBoolean(*Boolean) error
	VisitBooleanPrimitive(*BooleanPrimitive) error
	VisitString(*String) error
	VisitStringPrimitive(*StringPrimitive) error
	VisitNull(*Null) error
	VisitUndefined(*Undefined) error
	VisitAny(*Any) error
	VisitUnknown(*Unknown) error
	VisitNever(*Never) error
	VisitUnion(*Union) error
	VisitIntersection(*Intersection) error
	VisitFunction(*Function) error
	VisitIdentifier(*Identifier) error
	VisitIllegal(*Illegal) error
}
