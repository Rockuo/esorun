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
		return nil, fmt.Errorf("invalid instruction")
	}

	instruction = &instructions.Instruction{}
	instruction.Name = instructions.InstructionName(matches[1])
	if !instruction.Name.IsValid() {
		return nil, fmt.Errorf("invalid instruction name %s", instruction.Name)
	}
	//arg1 := Parsematches[2]
	// todo

	return nil, nil
}
