package operation

import . "esorun/value"
import "fmt"

func OpSize(dst Value, arg Value) error {
	var result int

	switch arg.(type) {
	case *ArrayValue:
		result = arg.(*ArrayValue).GetSize()
	case *DictionaryValue:
		result = arg.(*DictionaryValue).GetSize()
	case *StringValue:
		result = arg.(*StringValue).GetSize()
	default:
		return fmt.Errorf("size not supported on type %v", arg.GetType().GetName())
	}

	if dst.GetType().GetName() == TypeNameInt {
		dst.SetValue(result)
	} else if dst.GetType().GetName() == TypeNameUint {
		dst.SetValue(uint(result))
	} else {
		return fmt.Errorf("size cannot be assigned to variable type %v", dst.GetType().GetName())
	}

	return nil
}
