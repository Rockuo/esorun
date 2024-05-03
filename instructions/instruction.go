package instructions

import "fmt"

type InstructionName string

const (
	// creation
	InstNewVar  InstructionName = "VAR"
	InstNewIVar InstructionName = "IVAR"

	// modification
	InstSet  InstructionName = "SET"
	InstCast InstructionName = "CAST"

	// operation
	InstAdd InstructionName = "ADD"
	InstSub InstructionName = "SUB"
	InstMul InstructionName = "MUL"
	InstDiv InstructionName = "DIV"

	// condition
	InstEquals      InstructionName = "CMPE"
	InstNotEquals   InstructionName = "CMPNE"
	InstGreaterThan InstructionName = "CMPGR"
	InstLessThan    InstructionName = "CMPLE"
	InstIsType      InstructionName = "CMPT"

	// controls
	InstJump  InstructionName = "JMP"
	InstLabel InstructionName = "LBL"

	// IO
	InstWrite InstructionName = "WRITE"
	InstRead  InstructionName = "READ"
)

func (name InstructionName) isValid() bool {
	switch name {
	case
		InstNewVar,
		InstNewIVar,
		InstSet,
		InstCast,
		InstAdd,
		InstSub,
		InstMul,
		InstDiv,
		InstEquals,
		InstNotEquals,
		InstGreaterThan,
		InstLessThan,
		InstIsType,
		InstJump,
		InstLabel,
		InstWrite,
		InstRead:
		return true
	}
	return false
}

type Instruction struct {
	name   InstructionName
	value1 InstructionArgument
	value2 InstructionArgument
}

func NewInstructionFromString(instructionStr string) (instruction Instruction, err error) {
	var value1 string
	var value2 string
	_, err = fmt.Sscanf(instructionStr, "%v %v %v", &instruction.name, &value1, &value2)
	if err == nil {
		return
	}
	if !instruction.name.isValid() {
		err = fmt.Errorf("invalid instruction: %s", instructionStr)
		return
	}

	instruction.value1 = *NewInstructionArgumentFromString(value1)
	instruction.value2 = *NewInstructionArgumentFromString(value2)

	return
}
