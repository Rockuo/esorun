package value

import (
	"testing"
)

func TestStringValue(t *testing.T) {
	valueTest(
		t,
		ScalarType{TypeNameString},
		func() Value {
			return NewStringValue("abc")
		},
		"abc",
		"ab",
		map[string]map[string]Value{
			"add": {
				"before":   NewStringValue("abc"),
				"argument": NewStringValue("d"),
				"after":    NewStringValue("abcd"),
			},
			"sub": {
				"before":   NewStringValue("abc"),
				"argument": NewStringValue("d"),
				"after":    nil,
			},
			"mul": {
				"before":   NewStringValue("abc"),
				"argument": NewStringValue("d"),
				"after":    nil,
			},
			"div": {
				"before":   NewStringValue("abc"),
				"argument": NewStringValue("d"),
				"after":    nil,
			},
		},
		NewByteValue(1),
		NewStringValue("cde"),
		nil,
		nil,
	)
}
