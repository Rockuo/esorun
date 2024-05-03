package value

import (
	"testing"
)

func TestUintValue(t *testing.T) {
	valueTest(
		t,
		BaseType{TypeNameUint},
		func() Value {
			return NewUintValue(123)
		},
		"123",
		uint(100),
	)
}
