package parser

import (
	. "esorun/instructions"
	"esorun/value"
	"fmt"
)

const valueStart = '!'

func ParseArgument(argStr string) (argument *InstructionArgument, err error) {
	argChars := []byte(argStr)

	argument = &InstructionArgument{}
	argVar, argErr := ParseVariable(argChars)
	if argErr == nil {
		argument.Variable = argVar
		argument.Variant = InstArgVariantVariable
		return argument, nil
	}

	constVar, constErr := ParseConstant(argStr)
	if constErr == nil {
		argument.Constant = &constVar
		argument.Variant = InstArgVariantConst
		return argument, nil
	}

	typeVar, typeErr := ParseType(argStr)
	if typeErr == nil {
		argument.TypeDef = typeVar
		argument.Variant = InstArgVariantType
		return argument, nil
	}

	if argChars[0] == valueStart {
		argument.Value = value.NewStringValue(string(argChars[1:]))
		argument.Variant = InstArgVariantValue
		return argument, nil
	}

	return nil, fmt.Errorf("argument \"%v\" is neither Variable constant nor value", argStr)
}
