package value

import (
	"testing"
)

func TestByteValue(t *testing.T) {
	valueTest(
		t,
		BaseType{TypeNameByte},
		func() Value {
			return NewByteValue(64)
		},
		"@",
		byte(100),
	)
}
