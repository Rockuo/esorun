package operation

import . "esorun/value"
import "fmt"

func OpDiv(dst Value, arg Value) error {
	var err = checkSameType(dst, arg)
	if err != nil {
		return err
	}

	switch dst.(type) {
	case *ByteValue:
		divByte(dst.(*ByteValue), arg.(*ByteValue))
	case *FloatValue:
		divFloat(dst.(*FloatValue), arg.(*FloatValue))
	case *IntValue:
		divInt(dst.(*IntValue), arg.(*IntValue))
	case *UintValue:
		divUint(dst.(*UintValue), arg.(*UintValue))
	default:
		return fmt.Errorf("Div not supported on type %v", dst.GetType().GetName())
	}
	return nil
}

func divByte(dst *ByteValue, arg *ByteValue)    { *dst /= *arg }
func divFloat(dst *FloatValue, arg *FloatValue) { *dst /= *arg }
func divInt(dst *IntValue, arg *IntValue)       { *dst /= *arg }
func divUint(dst *UintValue, arg *UintValue)    { *dst /= *arg }
