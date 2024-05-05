package instructions

import (
	"esorun/value"
)

const (
	InstArgVariantVariable = iota
	InstArgVariantValue
	InstArgVariantConst
	InstArgVariantType
)

type InstArgVariant int

type InstructionArgument struct {
	Variant  InstArgVariant
	Variable *VariableDescriptor
	Value    *value.StringValue
	Constant *Constant
	TypeDef  *Type
}
