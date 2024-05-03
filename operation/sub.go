package operation

import . "esorun/value"
import "fmt"

func OpSub(dst Value, arg Value) error {
	var err = checkSameType(dst, arg)
	if err != nil {
		return err
	}

	switch dst.(type) {
	case *ByteValue:
		subByte(dst.(*ByteValue), arg.(*ByteValue))
	case *FloatValue:
		subFloat(dst.(*FloatValue), arg.(*FloatValue))
	case *IntValue:
		subInt(dst.(*IntValue), arg.(*IntValue))
	case *UintValue:
		subUint(dst.(*UintValue), arg.(*UintValue))
	default:
		return fmt.Errorf("Sub not supported on type %v", dst.GetType().GetName())
	}
	return nil
}

func subByte(dst *ByteValue, arg *ByteValue)    { *dst -= *arg }
func subFloat(dst *FloatValue, arg *FloatValue) { *dst -= *arg }
func subInt(dst *IntValue, arg *IntValue)       { *dst -= *arg }
func subUint(dst *UintValue, arg *UintValue)    { *dst -= *arg }
