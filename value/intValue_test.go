package value

import (
	"testing"
)

func TestIntValue(t *testing.T) {
	valueTest(
		t,
		BaseType{TypeNameInt},
		func() Value {
			return NewIntValue(123)
		},
		"123",
		100,
	)
}
