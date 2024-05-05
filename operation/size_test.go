package operation

import (
	. "esorun/value"
	"testing"
)

func TestOpSize(t *testing.T) {
	arr := NewArrayValue(BaseType{TypeNameInt})
	arr.AddItem(NewIntValue(1))
	arr.AddItem(NewIntValue(1))
	arr.AddItem(NewIntValue(1))
	arr.AddItem(NewIntValue(1))
	dic := NewDictionaryValue(BaseType{TypeNameInt})
	dic.SetItem("a", NewIntValue(1))
	dic.SetItem("b", NewIntValue(1))
	testOperation(
		t,
		[]testValue{
			{
				NewIntValue(1),
				NewStringValue("aaa"),
				NewIntValue(3),
				NewIntValue(2),
			},
			{
				NewIntValue(1),
				arr,
				NewIntValue(4),
				nil,
			},
			{
				NewIntValue(1),
				dic,
				NewIntValue(2),
				nil,
			},
		},
		OpSize,
	)
}
