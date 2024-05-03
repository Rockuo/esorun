package value

import (
	"testing"
)

func TestStringValue(t *testing.T) {
	valueTest(
		t,
		BaseType{TypeNameString},
		func() Value {
			return NewStringValue("abc")
		},
		"abc",
		"ab",
	)
}
