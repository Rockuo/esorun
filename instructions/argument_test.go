package instructions

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewInstructionArgumentFromString(t *testing.T) {
	argument := NewInstructionArgumentFromString("@@variable")
	assert.Equal(t, InstructionArgument{isVariable: true, isInternal: true, valueDefinition: "variable"}, *argument)

	argument = NewInstructionArgumentFromString("@variable")
	assert.Equal(t, InstructionArgument{isVariable: true, isInternal: false, valueDefinition: "variable"}, *argument)

	argument = NewInstructionArgumentFromString("valueDefinition")
	assert.Equal(t, InstructionArgument{isVariable: false, isInternal: false, valueDefinition: "valueDefinition"}, *argument)
}
