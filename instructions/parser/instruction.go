package parser

import (
	"esorun/instructions"
	"fmt"
	"regexp"
)

var r, _ = regexp.Compile(`(\S*) (\S*) (.*)`)

func ParseInstruction(row string) (instruction *instructions.Instruction, err error) {
	matches := r.FindStringSubmatch(row)

	if len(matches) < 4 {
		return nil, fmt.Errorf("invalid instruction: %v", row)
	}

	instruction = &instructions.Instruction{}
	instruction.Name = instructions.InstructionName(matches[1])
	if !instruction.Name.IsValid() {
		return nil, fmt.Errorf("invalid instruction name %s", instruction.Name)
	}

	arg1, arg1Err := ParseArgument(matches[2])
	if arg1Err != nil {
		return nil, fmt.Errorf("invalid arg1 %s", arg1Err)
	}

	arg2, arg2Err := ParseArgument(matches[2])
	if arg2Err != nil {
		return nil, fmt.Errorf("invalid arg2 %s", arg2Err)
	}

	instruction.Arg1 = arg1
	instruction.Arg2 = arg2

	return instruction, nil
}
