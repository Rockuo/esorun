package instructions

import (
	"esorun/constant"
	"esorun/value"
	"esorun/variable"
)

type InstructionArgument struct {
	variable *variable.Descriptor
	value    *value.StringValue
	constant *constant.Constant
}
