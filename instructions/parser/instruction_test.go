package parser

import (
	"esorun/instructions"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	names := []string{
		"VAR", "IVAR", "SET", "CAST", "ADD", "SUB", "MUL", "DIV", "SIZE", "CMPE", "CMPNE", "CMPGR",
		"CMPLE", "CMPT", "JMP", "LBL", "WRITE", "READ",
	}
	variales := []string{"@basic", "@@basic", "@arrDic[@basic]", "@@arrDic[@b@asic]"}
	values := []string{"#aa a", "#1"}

	for _, name := range names {
		for _, variable := range variales {
			for _, value := range values {
				inst, err := ParseInstruction(fmt.Sprintf("%v %v %v", name, variable, value))
				assert.Nil(t, err)
				assert.NotNil(t, inst)
				assert.Equal(t, instructions.InstructionName(name), inst.Name)
				assert.NotNil(t, inst.Arg1)
				assert.NotNil(t, instructions.InstArgVariant(instructions.InstArgVariantVariable), inst.Arg1.Variant)
				assert.NotNil(t, inst.Arg2)
				assert.NotNil(t, instructions.InstArgVariant(instructions.InstArgVariantValue), inst.Arg1.Value)
			}
		}
	}

	inst, err := ParseInstruction("IVAR [byte] !memory")
	assert.Nil(t, err)
	assert.NotNil(t, inst)
	assert.Equal(t, instructions.InstructionName("IVAR"), inst.Name)
	assert.NotNil(t, inst.Arg1)
	assert.NotNil(t, instructions.InstArgVariant(instructions.InstArgVariantType), inst.Arg1.Variant)
	assert.NotNil(t, inst.Arg2)
	assert.NotNil(t, instructions.InstArgVariant(instructions.InstArgVariantValue), inst.Arg1.Value)

}
