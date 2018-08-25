package object

import "fmt"

const (
	nullObject    = "NULL"
	integerObject = "INTEGER"
	booleanObject = "BOOLEAN"
)

// Type is object Type
type Type string

// Object has type and value representation.
type Object interface {
	Type() Type
	Inspect() string
}

// Null is a null value object.
type Null struct{}

// Inspect returns object's value.
func (n *Null) Inspect() string { return "null" }

// Type returns object's type.
func (n *Null) Type() Type { return nullObject }

// Integer is an integer value object.
type Integer struct {
	Value int64
}

// Inspect returns object's value.
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type returns object's type.
func (i *Integer) Type() Type { return integerObject }

// Boolean is a boolean value object.
type Boolean struct {
	Value bool
}

// Inspect returns object's value.
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type returns object's type.
func (b *Boolean) Type() Type { return booleanObject }
