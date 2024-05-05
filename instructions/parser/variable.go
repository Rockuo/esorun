package parser

import (
	. "esorun/instructions"
	"fmt"
	"slices"
)

// todo user error.New in app?

const variablePrefix = '@'
const arrayDicStart = '['
const arrayDicEnd = ']'

func ParseVariable(chars []byte) (descriptor *VariableDescriptor, err error) {
	var startIndicator byte
	var version DescriptorVersion = VersionBasic
	last := len(chars) - 1
	if chars[last] == arrayDicEnd {
		startIndicator = arrayDicStart
		version = VersionArrayDic
	}

	if startIndicator == 0 {
		descriptor, err = parseVariableName(chars)
		if err != nil {
			return nil, err
		}
		descriptor.Version = version
		return descriptor, nil
	}

	accessStartIndex := slices.Index(chars, startIndicator)
	descriptor, err = parseVariableName(chars[:accessStartIndex])
	if err != nil {
		return nil, err
	}
	descriptor.Version = version
	descriptor.InnerDescriptor, err = parseVariableName(chars[accessStartIndex+1 : last])
	if err != nil {
		return nil, err
	}
	descriptor.InnerDescriptor.Version = VersionBasic
	return descriptor, nil
}

func parseVariableName(chars []byte) (variable *VariableDescriptor, err error) {
	variable = &VariableDescriptor{}

	if len(chars) < 1 {
		return nil, fmt.Errorf("not enough characters")
	} else if chars[0] != variablePrefix {
		return nil, fmt.Errorf("token is not variable")
	}

	start := 1
	if chars[1] == variablePrefix {
		start++
		variable.IsInternal = true
	}
	variable.Name = string(chars[start:])
	return variable, nil

}
