package ast

type Type interface {
	HasPosition
	isType()
}

// NominalType represents a base type (e.g. boolean, integer, etc.)

type NominalType struct {
	Identifier string
	Pos        *Position
}

func (*NominalType) isType() {}

func (t *NominalType) StartPosition() *Position {
	return t.Pos
}

func (t *NominalType) EndPosition() *Position {
	return t.Pos
}

// VariableSizedType is a variable sized array type

type VariableSizedType struct {
	Type
	StartPos *Position
	EndPos   *Position
}

func (*VariableSizedType) isType() {}

func (t *VariableSizedType) StartPosition() *Position {
	return t.StartPos
}

func (t *VariableSizedType) EndPosition() *Position {
	return t.EndPos
}

// ConstantSizedType is a constant sized array type

type ConstantSizedType struct {
	Type
	Size     int
	StartPos *Position
	EndPos   *Position
}

func (*ConstantSizedType) isType() {}

func (t *ConstantSizedType) StartPosition() *Position {
	return t.StartPos
}

func (t *ConstantSizedType) EndPosition() *Position {
	return t.EndPos
}

// FunctionType

type FunctionType struct {
	ParameterTypes []Type
	ReturnType     Type
	StartPos       *Position
	EndPos         *Position
}

func (*FunctionType) isType() {}

func (t *FunctionType) StartPosition() *Position {
	return t.StartPos
}

func (t *FunctionType) EndPosition() *Position {
	return t.EndPos
}
