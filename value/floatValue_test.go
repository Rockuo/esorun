package value

import (
	"testing"
)

func TestFloatValue(t *testing.T) {
	valueTest(
		t,
		BaseType{TypeNameFloat},
		func() Value {
			return NewFloatValue(12.34)
		},
		"12.340000",
		float64(100),
	)
}
