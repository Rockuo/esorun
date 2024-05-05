package parser

import (
	"esorun/instructions"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestParseConstant(t *testing.T) {
	var r instructions.Constant
	var e error

	r, e = ParseConstant("STD_IN")
	assert.Equal(t, instructions.Constant{instructions.ChanelType, instructions.StdIn}, r)
	assert.Nil(t, e)

	r, e = ParseConstant("STD_OUT")
	assert.Equal(t, instructions.Constant{instructions.ChanelType, instructions.StdOut}, r)
	assert.Nil(t, e)

	r, e = ParseConstant("STD_ERR")
	assert.Equal(t, instructions.Constant{instructions.ChanelType, instructions.StdErr}, r)
	assert.Nil(t, e)

	r, e = ParseConstant("CMD")
	assert.Equal(t, instructions.Constant{instructions.GlobalType, instructions.Cmd}, r)
	assert.Nil(t, e)

	r, e = ParseConstant("XX")
	assert.Equal(t, instructions.Constant{instructions.InvalidType, ""}, r)
	assert.Error(t, e)
}
