package instructions

const InstructionArgumentVariablePrefix = '@'

type InstructionArgument struct {
	isVariable      bool
	isInternal      bool
	valueDefinition string
}

func NewInstructionArgumentFromString(interpretation string) *InstructionArgument {
	var argument = &InstructionArgument{}
	var length = len(interpretation)

	argument.valueDefinition = interpretation
	argument.isVariable = length >= 2 && interpretation[0] == InstructionArgumentVariablePrefix

	if argument.isVariable {
		argument.isInternal = length >= 3 && interpretation[1] == InstructionArgumentVariablePrefix
		argument.valueDefinition = argument.valueDefinition[1:]
		if argument.isInternal {
			argument.valueDefinition = argument.valueDefinition[1:]
		}
	}
	return argument
}
