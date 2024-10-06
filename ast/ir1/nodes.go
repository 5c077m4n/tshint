// Package ir1 Medium-level Intermidiate Representation
package ir1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type (
	ByteRange = [2]uint
	// Position the position of the node in the source code
	Position struct {
		columnStart uint
		columnEnd   uint
		lineStart   uint
		lineEnd     uint
	}
	// Node common functions for all types
	Node interface {
		IsNil() bool
		Accept(Visitor) error
		Content([]byte) []byte
		Location([]byte) Position
		String() string
	}
)

var ErrAccept = errors.New("failed to accept the AST node")

func wrapAcceptErr(err error) error {
	if err != nil {
		return errors.Join(ErrAccept, err)
	}
	return nil
}

// Block tree root (should only be one at the top)
type Block struct {
	Statements []Node
	ByteRange
}

// IsNil nil check
func (p *Block) IsNil() bool {
	return p == nil
}

// Accept visitor fn
func (p *Block) Accept(v Visitor) error {
	return wrapAcceptErr(v.VisitProgram(p))
}

// Content returns the node's content
func (p *Block) Content(source []byte) []byte {
	return source[p.ByteRange[0]:p.ByteRange[1]]
}

// Location returns the node's position in the source code
func (p *Block) Location(source []byte) Position {
	sourceLines := strings.Split(string(source), "\n")
	return Position{
		columnStart: 0,
		columnEnd:   uint(len(sourceLines[len(sourceLines)-1])),
		lineStart:   0,
		lineEnd:     uint(len(sourceLines)),
	}
}

func (p *Block) String() string {
	progStr := make([]string, 0, len(p.Statements))
	for _, n := range p.Statements {
		progStr = append(progStr, n.String())
	}

	return strings.Join(progStr, "\n")
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

// Content returns the node's content
func (n *Number) Content(source []byte) []byte {
	return source[n.ByteRange[0]:n.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (n *Number) Location(_source []byte) Position {
	return Position{}
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

// Content returns the node's content
func (b *Boolean) Content(source []byte) []byte {
	return source[b.ByteRange[0]:b.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (b *Boolean) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
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

// Content returns the node's content
func (s *String) Content(source []byte) []byte {
	return source[s.ByteRange[0]:s.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (s *String) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
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

// Content returns the node's content
func (n *Null) Content(source []byte) []byte {
	return source[n.ByteRange[0]:n.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (n *Null) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
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

// Content returns the node's content
func (u *Undefined) Content(source []byte) []byte {
	return source[u.ByteRange[0]:u.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (u *Undefined) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
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

// Content returns the node's content
func (a *Any) Content(source []byte) []byte {
	return source[a.ByteRange[0]:a.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (a *Any) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
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

// Content returns the node's content
func (u *Unknown) Content(source []byte) []byte {
	return source[u.ByteRange[0]:u.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (u *Unknown) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
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

// Content returns the node's content
func (n *Never) Content(source []byte) []byte {
	return source[n.ByteRange[0]:n.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (n *Never) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
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

// Content returns the node's content
func (u *Union) Content(source []byte) []byte {
	return source[u.ByteRange[0]:u.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (u *Union) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
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

// Content returns the node's content
func (i *Intersection) Content(source []byte) []byte {
	return source[i.ByteRange[0]:i.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (i *Intersection) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
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
	Body    Node
	IsArrow bool
	ByteRange
}

// IsNil nil check
func (f *Function) IsNil() bool {
	return f == nil
}

// Content returns the node's content
func (f *Function) Content(source []byte) []byte {
	return source[f.ByteRange[0]:f.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (f *Function) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
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

// Content returns the node's content
func (i Identifier) Content(source []byte) []byte {
	return source[i.ByteRange[0]:i.ByteRange[1]]
}

// Location returns the node's position in the source code
//
// TODO: implement
func (i *Identifier) Location(_source []byte) Position {
	return Position{
		columnStart: 0,
		columnEnd:   0,
		lineStart:   0,
		lineEnd:     0,
	}
}

// String strigify
func (i *Identifier) String() string {
	initKeyword := "let"
	if i.IsConst {
		initKeyword = "const"
	}

	if i.Value == nil {
		return fmt.Sprintf("%s %s: %s;", initKeyword, i.Name, i.Type)
	}

	return fmt.Sprintf("%s %s: %s = %s;", initKeyword, i.Name, i.Type, i.Value)
}
