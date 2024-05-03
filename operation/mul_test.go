package operation

import (
	. "esorun/value"
	"testing"
)

func TestOpMul(t *testing.T) {
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
				NewByteValue(20),
				NewBoolValue(false),
			},
			{
				NewFloatValue(1.5),
				NewFloatValue(2),
				NewFloatValue(3),
				NewBoolValue(false),
			},
			{
				NewIntValue(10),
				NewIntValue(-2),
				NewIntValue(-20),
				NewBoolValue(false),
			},
			{
				NewStringValue("abc"),
				nil,
				nil,
				nil,
			},
			{
				NewUintValue(10),
				NewUintValue(2),
				NewUintValue(20),
				NewBoolValue(false),
			},
		},
		OpMul,
	)
}
