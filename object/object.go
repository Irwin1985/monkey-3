package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gcoka/monkey/ast"
)

// Object type internal string expression.
const (
	ErrorObject       = "ERROR"
	NullObject        = "NULL"
	IntegerObject     = "INTEGER"
	BooleanObject     = "BOOLEAN"
	ReturnValueObject = "RETURN_VALUE"
	FunctionObject    = "FUNCTION"
	StringObject      = "STRING"
	BuiltinObject     = "BUILTIN"
)

// Type is object Type
type Type string

// Object has type and value representation.
type Object interface {
	Type() Type
	Inspect() string
}

// BuiltinFunction is builtin function.
type BuiltinFunction func(args ...Object) Object

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

// Error is an error object.
type Error struct {
	Message string
}

// Type returns error type.
func (e *Error) Type() Type { return ErrorObject }

// Inspect returns error value.
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Function is a function object.
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns function type.
func (f *Function) Type() Type { return FunctionObject }

// Inspect returns function value.
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String is a string value object.
type String struct {
	Value string
}

// Inspect returns object's value.
func (s *String) Inspect() string { return s.Value }

// Type returns object's type.
func (s *String) Type() Type { return StringObject }

// Builtin is a builtin function object.
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns object's type.
func (b *Builtin) Type() Type { return BuiltinObject }

// Inspect returns object's value.
func (b *Builtin) Inspect() string { return "builtin function" }
