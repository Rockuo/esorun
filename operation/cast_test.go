package operation

import (
	. "esorun/value"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestCast(t *testing.T) {
	toTypeTest(t, []Value{
		NewStringValue("@"),
		NewIntValue(64),
		NewUintValue(64),
		NewFloatValue(64),
	}, BaseType{Name: TypeNameByte}, NewByteValue(64))

	toTypeTest(t, []Value{
		NewIntValue(64),
		NewUintValue(64),
	}, BaseType{Name: TypeNameString}, NewStringValue("64"))

	toTypeTest(t, []Value{
		NewStringValue("64"),
		NewUintValue(64),
		NewFloatValue(64),
	}, BaseType{Name: TypeNameInt}, NewIntValue(64))

	toTypeTest(t, []Value{
		NewStringValue("64"),
		NewIntValue(64),
		NewFloatValue(64),
	}, BaseType{Name: TypeNameUint}, NewUintValue(64))

	toTypeTest(t, []Value{
		NewStringValue("64"),
		NewIntValue(64),
		NewUintValue(64),
	}, BaseType{Name: TypeNameFloat}, NewFloatValue(64))

}

func toTypeTest(t *testing.T, values []Value, tp Type, c Value) {
	for _, v := range values {
		res, err := Cast(v, tp)
		assert.Equal(t, c, res)
		assert.Nil(t, err)
	}
}
