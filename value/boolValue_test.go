package value

import (
	"testing"
)

func TestBoolValue(t *testing.T) {
	valueTest(
		t,
		BaseType{TypeNameBool},
		func() Value {
			return NewBoolValue(true)
		},
		"true",
		false,
	)
}
