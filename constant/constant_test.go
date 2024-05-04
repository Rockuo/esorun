package constant

import "testing"
import "github.com/stretchr/testify/assert"

func TestParseConstant(t *testing.T) {
	var r Constant
	var e error

	r, e = ParseConstant("STD_IN")
	assert.Equal(t, Constant{ChanelType, StdIn}, r)
	assert.Nil(t, e)

	r, e = ParseConstant("STD_OUT")
	assert.Equal(t, Constant{ChanelType, StdOut}, r)
	assert.Nil(t, e)

	r, e = ParseConstant("STD_ERR")
	assert.Equal(t, Constant{ChanelType, StdErr}, r)
	assert.Nil(t, e)

	r, e = ParseConstant("CMD")
	assert.Equal(t, Constant{GlobalType, Cmd}, r)
	assert.Nil(t, e)

	r, e = ParseConstant("XX")
	assert.Equal(t, Constant{InvalidType, ""}, r)
	assert.Error(t, e)
}
