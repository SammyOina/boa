package object

import "fmt"

type ObjectType int

const (
	_ ObjectType = iota
	INTEGER
	BOOLEAN
	NULL
	RETURN
)

var ObjectString = map[ObjectType]string{
	INTEGER: "integer",
	BOOLEAN: "bool",
	NULL:    "null",
	RETURN:  "return",
}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Null struct{}

func (n *Null) Type() ObjectType {
	return NULL
}

func (n *Null) Inspect() string {
	return "null"
}

type Return struct {
	Value Object
}

func (r *Return) Type() ObjectType {
	return RETURN
}

func (r *Return) Inspect() string {
	return r.Value.Inspect()
}
