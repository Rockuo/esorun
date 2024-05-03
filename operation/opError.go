package operation

import (
	. "esorun/value"
	"fmt"
)

func checkSameType(v1 Value, v2 Value) error {
	var ok bool
	switch v1.(type) {
	case *BoolValue:
		_, ok = v2.(*BoolValue)
	case *ByteValue:
		_, ok = v2.(*ByteValue)
	case *FloatValue:
		_, ok = v2.(*FloatValue)
	case *IntValue:
		_, ok = v2.(*IntValue)
	case *StringValue:
		_, ok = v2.(*StringValue)
	case *UintValue:
		_, ok = v2.(*UintValue)
	default:
		panic("Unimplemented type")
	}
	if ok {
		return nil
	}
	return fmt.Errorf("Types are not same")
}
