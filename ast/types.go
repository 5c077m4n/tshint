// Package ast type structs here
package ast

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Node common functions for all types
type Node interface {
	IsNil() bool
	Accept(Visitor) error
	String() string
}

var ErrAccept = errors.New("failed to accept the AST node")

func wrapAcceptErr(err error) error {
	if err != nil {
		return errors.Join(ErrAccept, err)
	}
	return nil
}

type ByteRange [2]uint

// Program tree root (should only be one at the top)
type Program struct {
	Children []Node
	ByteRange
}

// IsNil nil check
func (p *Program) IsNil() bool {
	return p == nil
}

// Accept visitor fn
func (p *Program) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitProgram(p))
}

// Number number type
type Number struct {
	Values []float32
	ByteRange
}

// IsNil nil check
func (n *Number) IsNil() bool {
	return n == nil
}

// Accept visitor fn
func (n *Number) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitNumber(n))
}

// String strigify
func (n *Number) String() string {
	if len(n.Values) == 0 {
		return "number"
	}

	uniqueValues := uniques(n.Values)
	values := make([]string, 0, len(uniqueValues))

	for _, v := range uniqueValues {
		values = append(values, strconv.FormatFloat(float64(v), 'g', -1, 32))
	}
	return strings.Join(values, " | ")
}

// Boolean bool type
type Boolean struct {
	Values []bool
	ByteRange
}

// IsNil nil check
func (b *Boolean) IsNil() bool {
	return b == nil
}

// Accept visitor fn
func (b *Boolean) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitBoolean(b))
}

// String strigify
func (b *Boolean) String() string {
	if len(b.Values) == 0 {
		return "boolean"
	}

	values := make([]string, 0, len(b.Values))
	for _, v := range b.Values {
		if v {
			values = append(values, "true")
		} else {
			values = append(values, "false")
		}
	}

	values = uniques(values)
	return strings.Join(values, " | ")
}

// String string type
type String struct {
	Values []string
	ByteRange
}

// IsNil nil check
func (s *String) IsNil() bool {
	return s == nil
}

// Accept visitor fn
func (s *String) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitString(s))
}

// String strigify
func (s *String) String() string {
	if len(s.Values) == 0 {
		return "string"
	}

	return strings.Join(uniques(s.Values), " | ")
}

// Null `null` type
type Null struct {
	ByteRange
}

// IsNil nil check
func (n *Null) IsNil() bool {
	return n == nil
}

// Accept visitor fn
func (n *Null) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitNull(n))
}

// String strigify
func (n *Null) String() string {
	return "null"
}

// Undefined `undefined` type
type Undefined struct {
	ByteRange
}

// IsNil nil check
func (u *Undefined) IsNil() bool {
	return u == nil
}

// Accept visitor fn
func (u *Undefined) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitUndefined(u))
}

// String strigify
func (u *Undefined) String() string {
	return "undefined"
}

// Any `any` type
type Any struct {
	ByteRange
}

// IsNil nil check
func (a *Any) IsNil() bool {
	return a == nil
}

// Accept visitor fn
func (a *Any) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitAny(a))
}

// String strigify
func (a *Any) String() string {
	return "any"
}

// Unknown `unknown` type
type Unknown struct {
	ByteRange
}

// IsNil nil check
func (u *Unknown) IsNil() bool {
	return u == nil
}

// Accept visitor fn
func (u *Unknown) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitUnknown(u))
}

// String strigify
func (u *Unknown) String() string {
	return "unknown"
}

// Never `never` type
type Never struct {
	ByteRange
}

// IsNil nil check
func (n *Never) IsNil() bool {
	return n == nil
}

// Accept visitor fn
func (n *Never) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitNever(n))
}

// String strigify
func (n *Never) String() string {
	return "never"
}

// Union discriminated union type (`|`)
type Union struct {
	Left  Node
	Right Node
	ByteRange
}

// IsNil nil check
func (u *Union) IsNil() bool {
	return u == nil
}

// Accept visitor fn
func (u *Union) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitUnion(u))
}

// String strigify
func (u *Union) String() string {
	return strings.Join(
		[]string{u.Left.String(), u.Right.String()},
		" | ",
	)
}

// Intersection type (`&`)
type Intersection struct {
	Left  Node
	Right Node
	ByteRange
}

// IsNil nil check
func (i *Intersection) IsNil() bool {
	return i == nil
}

// Accept visitor fn
func (i *Intersection) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitIntersection(i))
}

// String strigify
func (i *Intersection) String() string {
	return strings.Join(
		[]string{i.Left.String(), i.Right.String()},
		" & ",
	)
}

// Function func type
type Function struct {
	Name    string
	Inputs  []Node
	Output  Node
	IsArrow bool
	ByteRange
}

// IsNil nil check
func (f *Function) IsNil() bool {
	return f == nil
}

// Accept visitor fn
func (f *Function) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitFunction(f))
}

// String strigify
func (f *Function) String() string {
	if f.IsArrow {
		return fmt.Sprintf(
			"const %s = (%s): %s => { /*...*/ }",
			f.Name, f.Inputs, f.Output,
		)
	}
	return fmt.Sprintf(
		"function %s(%s): %s { /*...*/ }",
		f.Name, f.Inputs, f.Output,
	)
}

// Identifier ident type
type Identifier struct {
	Name    string
	IsConst bool
	Type    Node
	Value   Node
	ByteRange
}

// IsNil nil check
func (i *Identifier) IsNil() bool {
	return i == nil
}

// Accept visitor fn
func (i *Identifier) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitIdentifier(i))
}

// String strigify
func (i *Identifier) String() string {
	initKW := "let"
	if i.IsConst {
		initKW = "const"
	}

	if i.Value == nil {
		return fmt.Sprintf("%s %s: %s;", initKW, i.Name, i.Type)
	}

	return fmt.Sprintf("%s %s: %s = %s;", initKW, i.Name, i.Type, i.Value)
}
