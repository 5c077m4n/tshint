// Package ast type structs here
package ast

import "errors"

// TSType common functions for all types
type TSType interface {
	IsNil() bool
	Accept(Visitor) error
}

var ErrAccept = errors.New("failed to accept the AST node")

func wrapAcceptErr(err error) error {
	if err != nil {
		return errors.Join(ErrAccept, err)
	}
	return nil
}

// Number number type
type Number struct{}

// IsNil nil check
func (n *Number) IsNil() bool {
	return n == nil
}

// Accept visitor fn
func (n *Number) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitNumber(n))
}

// Boolean bool type
type Boolean struct{}

// IsNil nil check
func (b *Boolean) IsNil() bool {
	return b == nil
}

// Accept visitor fn
func (b *Boolean) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitBoolean(b))
}

// String string type
type String struct{}

// IsNil nil check
func (s *String) IsNil() bool {
	return s == nil
}

// Accept visitor fn
func (s *String) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitString(s))
}

// Null `null` type
type Null struct{}

// IsNil nil check
func (n *Null) IsNil() bool {
	return n == nil
}

// Accept visitor fn
func (n *Null) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitNull(n))
}

// Undefined `undefined` type
type Undefined struct{}

// IsNil nil check
func (u *Undefined) IsNil() bool {
	return u == nil
}

// Accept visitor fn
func (u *Undefined) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitUndefined(u))
}

// Any `any` type
type Any struct{}

// IsNil nil check
func (a *Any) IsNil() bool {
	return a == nil
}

// Accept visitor fn
func (a *Any) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitAny(a))
}

// Unknown `unknown` type
type Unknown struct{}

// IsNil nil check
func (u *Unknown) IsNil() bool {
	return u == nil
}

// Accept visitor fn
func (u *Unknown) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitUnknown(u))
}

// Never `never` type
type Never struct{}

// IsNil nil check
func (n *Never) IsNil() bool {
	return n == nil
}

// Accept visitor fn
func (n *Never) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitNever(n))
}

// Union discriminated union type (`|`)
type Union struct {
	Left  TSType
	Right TSType
}

// IsNil nil check
func (u *Union) IsNil() bool {
	return u == nil
}

// Accept visitor fn
func (u *Union) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitUnion(u))
}

// Intersection type (`&`)
type Intersection struct {
	Left  TSType
	Right TSType
}

// IsNil nil check
func (i *Intersection) IsNil() bool {
	return i == nil
}

// Accept visitor fn
func (i *Intersection) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitIntersection(i))
}
