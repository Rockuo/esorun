package instructions

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

func (name InstructionName) IsValid() bool {
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
	Name InstructionName
	Arg1 InstructionArgument
	Arg2 InstructionArgument
}
