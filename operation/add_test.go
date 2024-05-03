package operation

import (
	. "esorun/value"
	"testing"
)

func TestOpAdd(t *testing.T) {
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
				NewByteValue(20),
				NewByteValue(30),
				NewBoolValue(false),
			},
			{
				NewFloatValue(10),
				NewFloatValue(20),
				NewFloatValue(30),
				NewBoolValue(false),
			},
			{
				NewIntValue(10),
				NewIntValue(20),
				NewIntValue(30),
				NewBoolValue(false),
			},
			{
				NewStringValue("abc"),
				NewStringValue("d"),
				NewStringValue("abcd"),
				NewBoolValue(false),
			},
			{
				NewUintValue(10),
				NewUintValue(20),
				NewUintValue(30),
				NewBoolValue(false),
			},
		},
		OpAdd,
	)
}
