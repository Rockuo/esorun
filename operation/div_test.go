package operation

import (
	. "esorun/value"
	"testing"
)

func TestOpDiv(t *testing.T) {
	testOperation(
		t,
		[]testValue{
			{
				NewBoolValue(true),
				nil,
				nil,
				nil,
			},
			{
				NewByteValue(10),
				NewByteValue(2),
				NewByteValue(5),
				NewBoolValue(false),
			},
			{
				NewFloatValue(3),
				NewFloatValue(2),
				NewFloatValue(1.5),
				NewBoolValue(false),
			},
			{
				NewIntValue(-20),
				NewIntValue(-2),
				NewIntValue(10),
				NewBoolValue(false),
			},
			{
				NewStringValue("abc"),
				nil,
				nil,
				nil,
			},
			{
				NewUintValue(20),
				NewUintValue(2),
				NewUintValue(10),
				NewBoolValue(false),
			},
		},
		OpDiv,
	)
}
