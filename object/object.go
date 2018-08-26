package object

import "fmt"

// Object type internal string expression.
const (
	NullObject        = "NULL"
	IntegerObject     = "INTEGER"
	BooleanObject     = "BOOLEAN"
	ReturnValueObject = "RETURN_VALUE"
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
func (n *Null) Type() Type { return NullObject }

// Integer is an integer value objectB
type Integer struct {
	Value int64
}

// Inspect returns object's value.
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type returns object's type.
func (i *Integer) Type() Type { return IntegerObject }

// Boolean is a boolean value object.
type Boolean struct {
	Value bool
}

// Inspect returns object's value.
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type returns object's type.
func (b *Boolean) Type() Type { return BooleanObject }

// ReturnValue is a return value object.
type ReturnValue struct {
	Value Object
}

// Inspect returns wrapped object's value.
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Type returns object's type.
func (rv *ReturnValue) Type() Type { return ReturnValueObject }
