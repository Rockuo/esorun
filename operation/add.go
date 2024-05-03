package operation

import . "esorun/value"
import "fmt"

func OpAdd(dst Value, arg Value) error {
	var err = checkSameType(dst, arg)
	if err != nil {
		return err
	}

	switch dst.(type) {
	case *ByteValue:
		addByte(dst.(*ByteValue), arg.(*ByteValue))
	case *FloatValue:
		addFloat(dst.(*FloatValue), arg.(*FloatValue))
	case *IntValue:
		addInt(dst.(*IntValue), arg.(*IntValue))
	case *UintValue:
		addUint(dst.(*UintValue), arg.(*UintValue))
	case *StringValue:
		addString(dst.(*StringValue), arg.(*StringValue))
	default:
		return fmt.Errorf("Add not supported on type %v", dst.GetType().GetName())
	}
	return nil
}

func addByte(dst *ByteValue, arg *ByteValue)       { *dst += *arg }
func addFloat(dst *FloatValue, arg *FloatValue)    { *dst += *arg }
func addInt(dst *IntValue, arg *IntValue)          { *dst += *arg }
func addUint(dst *UintValue, arg *UintValue)       { *dst += *arg }
func addString(dst *StringValue, arg *StringValue) { *dst += *arg }
