package operation

import . "esorun/value"
import "fmt"

func OpMul(dst Value, arg Value) error {
	var err = checkSameType(dst, arg)
	if err != nil {
		return err
	}

	switch dst.(type) {
	case *ByteValue:
		mulByte(dst.(*ByteValue), arg.(*ByteValue))
	case *FloatValue:
		mulFloat(dst.(*FloatValue), arg.(*FloatValue))
	case *IntValue:
		mulInt(dst.(*IntValue), arg.(*IntValue))
	case *UintValue:
		mulUint(dst.(*UintValue), arg.(*UintValue))
	default:
		return fmt.Errorf("Mul not supported on type %v", dst.GetType().GetName())
	}
	return nil
}

func mulByte(dst *ByteValue, arg *ByteValue)    { *dst *= *arg }
func mulFloat(dst *FloatValue, arg *FloatValue) { *dst *= *arg }
func mulInt(dst *IntValue, arg *IntValue)       { *dst *= *arg }
func mulUint(dst *UintValue, arg *UintValue)    { *dst *= *arg }
