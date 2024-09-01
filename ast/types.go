// Package ast type structs here
package ast

// TSType common functions for all types
type TSType interface {
	tsType()
	IsNil() bool
}

// Number number type
type Number struct{}

func (n *Number) tsType() {}

// IsNil nil check
func (n *Number) IsNil() bool {
	return n == nil
}

// Boolean bool type
type Boolean struct{}

func (b *Boolean) tsType() {}

// IsNil nil check
func (b *Boolean) IsNil() bool {
	return b == nil
}

// String string type
type String struct{}

func (s *String) tsType() {}

// IsNil nil check
func (s *String) IsNil() bool {
	return s == nil
}

// Null `null` type
type Null struct{}

func (n *Null) tsType() {}

// IsNil nil check
func (n *Null) IsNil() bool {
	return n == nil
}

// Undefined `undefined` type
type Undefined struct{}

func (u *Undefined) tsType() {}

// IsNil nil check
func (u *Undefined) IsNil() bool {
	return u == nil
}

// Any `any` type
type Any struct{}

func (a *Any) tsType() {}

// IsNil nil check
func (a *Any) IsNil() bool {
	return a == nil
}

// Unknown `unknown` type
type Unknown struct{}

func (u *Unknown) tsType() {}

// IsNil nil check
func (u *Unknown) IsNil() bool {
	return u == nil
}

// Never `never` type
type Never struct{}

func (n *Never) tsType() {}

// IsNil nil check
func (n *Never) IsNil() bool {
	return n == nil
}

// Union discriminated union type (`|`)
type Union struct {
	Left  *TSType
	Right *TSType
}

func (u *Union) tsType() {}

// IsNil nil check
func (u *Union) IsNil() bool {
	return u == nil
}

// Intersection type (`&`)
type Intersection struct {
	Left  *TSType
	Right *TSType
}

func (i *Intersection) tsType() {}

// IsNil nil check
func (i *Intersection) IsNil() bool {
	return i == nil
}
