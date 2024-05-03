package value

import (
	"testing"
)

func TestByteValue(t *testing.T) {
	valueTest(
		t,
		ScalarType{TypeNameByte},
		func() Value {
			return NewByteValue(64)
		},
		"@",
		byte(100),
		map[string]map[string]Value{
			"add": {
				"before":   NewByteValue(10),
				"argument": NewByteValue(20),
				"after":    NewByteValue(30),
			},
			"sub": {
				"before":   NewByteValue(20),
				"argument": NewByteValue(10),
				"after":    NewByteValue(10),
			},
			"mul": {
				"before":   NewByteValue(10),
				"argument": NewByteValue(2),
				"after":    NewByteValue(40),
			},
			"div": {
				"before":   NewByteValue(10),
				"argument": NewByteValue(5),
				"after":    NewByteValue(2),
			},
		},
		NewStringValue("a"),
		NewByteValue(1),
		NewByteValue(70),
		NewByteValue(60),
	)
}
