package operation

import (
	. "esorun/value"
	"testing"
)

func TestOpSub(t *testing.T) {
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
				NewByteValue(20),
				NewByteValue(10),
				NewByteValue(10),
				NewBoolValue(false),
			},
			{
				NewFloatValue(20),
				NewFloatValue(10),
				NewFloatValue(10),
				NewBoolValue(false),
			},
			{
				NewIntValue(20),
				NewIntValue(10),
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
				NewUintValue(10),
				NewUintValue(10),
				NewBoolValue(false),
			},
		},
		OpSub,
	)
}
