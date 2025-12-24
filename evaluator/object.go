package evaluator

import "fmt"

type ObjectType string

const (
	I32_OBJ     = "I32"
	I64_OBJ     = "I64"
	STRING_OBJ  = "STRING"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer object
type Integer struct {
	Value int64
	IntType ObjectType
}

func (i *Integer) Type() ObjectType { return i.IntType }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// String object
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

// Boolean object
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null object
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// Error object
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return "ERROR" }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

var (
	NULL = &Null{}
)
