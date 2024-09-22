package ast

import (
	"errors"
	"log/slog"
	"tshint/lexer/token"
)

type astVisitor struct {
	Prog  *Program
	index uint
}

// VisitAny implements Visitor.
func (a astVisitor) VisitAny(*Any) error {
	panic("unimplemented")
}

// VisitBoolean implements Visitor.
func (a astVisitor) VisitBoolean(*Boolean) error {
	panic("unimplemented")
}

// VisitFunction implements Visitor.
func (a astVisitor) VisitFunction(*Function) error {
	panic("unimplemented")
}

// VisitIdentifier implements Visitor.
func (a astVisitor) VisitIdentifier(*Identifier) error {
	panic("unimplemented")
}

// VisitIntersection implements Visitor.
func (a astVisitor) VisitIntersection(*Intersection) error {
	panic("unimplemented")
}

// VisitNever implements Visitor.
func (a astVisitor) VisitNever(*Never) error {
	panic("unimplemented")
}

// VisitNull implements Visitor.
func (a astVisitor) VisitNull(*Null) error {
	panic("unimplemented")
}

// VisitNumber implements Visitor.
func (a astVisitor) VisitNumber(*Number) error {
	panic("unimplemented")
}

// VisitProgram implements Visitor.
func (a astVisitor) VisitProgram(*Program) error {
	panic("unimplemented")
}

// VisitString implements Visitor.
func (a astVisitor) VisitString(*String) error {
	panic("unimplemented")
}

// VisitUndefined implements Visitor.
func (a astVisitor) VisitUndefined(*Undefined) error {
	panic("unimplemented")
}

// VisitUnion implements Visitor.
func (a astVisitor) VisitUnion(*Union) error {
	panic("unimplemented")
}

// VisitUnknown implements Visitor.
func (a astVisitor) VisitUnknown(*Unknown) error {
	panic("unimplemented")
}

func (a *astVisitor) IncIndex() {
	a.index++
}

var ErrFrom = errors.New("failed to convert the tokens into an AST")

func From(sourceCode []byte, chanToken <-chan token.Token) (*Program, error) {
	visitor := astVisitor{}
	for t := range chanToken {
		content := t.Content(sourceCode)
		slog.Info(
			"Token",
			"value", t,
			"content", content,
		)

		switch t.Kind {
		case "const":
			n := &Number{ByteRange: t.ByteRange}
			if err := n.Accept(visitor); err != nil {
				return nil, errors.Join(ErrFrom, err)
			}
		case "identifier":
		case ":":
		case "number":
			n := &Number{ByteRange: t.ByteRange}
			if err := n.Accept(visitor); err != nil {
				return nil, errors.Join(ErrFrom, err)
			}
		case "=":
		case ";":
		}
	}

	return visitor.Prog, nil
}
