package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCodeFromString(t *testing.T) {
	code, err := ParseCodeFromString(`
	; test Code
	IVAR [byte] !memory
	IVAR uint !pointer
	IVAR uint !zero

	; start
	ADD @@memory[@@pointer] !8
	`)
	assert.Nil(t, err)
	assert.NotNil(t, code)
	assert.Equal(t, 4, len(code))

	code, err = ParseCodeFromFile("../../testFiles/code_parse_loader.eso")
	assert.Nil(t, err)
	assert.NotNil(t, code)
	assert.Equal(t, 4, len(code))
}
