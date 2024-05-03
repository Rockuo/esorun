package operation

import (
	. "esorun/value"
	"fmt"
)

func OpCmpEquals(v1 Value, v2 Value) (isTrue bool, err error) {
	isTrue = false
	err = checkSameType(v1, v2)
	if err != nil {
		return
	}
	switch v1.GetType().GetName() {
	case TypeNameByte:
		isTrue = *v1.(*ByteValue) == *v2.(*ByteValue)
	case TypeNameString:
		isTrue = *v1.(*StringValue) == *v2.(*StringValue)
	case TypeNameInt:
		isTrue = *v1.(*IntValue) == *v2.(*IntValue)
	case TypeNameUint:
		isTrue = *v1.(*UintValue) == *v2.(*UintValue)
	case TypeNameFloat:
		isTrue = *v1.(*FloatValue) == *v2.(*FloatValue)
	case TypeNameBool:
		isTrue = *v1.(*BoolValue) == *v2.(*BoolValue)
	//case TypeNameArray:
	//case TypeNameDictionary:
	default:
		panic("Unhandeled type " + v1.GetType().GetName())
	}
	return
}

func OpCmpNotEquals(v1, v2 Value) (isTrue bool, err error) {
	isFalse, err := OpCmpEquals(v1, v2)
	if err != nil {
		return false, err
	}
	return !isFalse, nil
}

func OpCmpGreaterThen(v1, v2 Value) (isTrue bool, err error) {
	isTrue = false
	err = checkSameType(v1, v2)
	if err != nil {
		return
	}
	switch v1.GetType().GetName() {
	case TypeNameByte:
		isTrue = *v1.(*ByteValue) > *v2.(*ByteValue)
	case TypeNameInt:
		isTrue = *v1.(*IntValue) > *v2.(*IntValue)
	case TypeNameUint:
		isTrue = *v1.(*UintValue) > *v2.(*UintValue)
	case TypeNameFloat:
		isTrue = *v1.(*FloatValue) > *v2.(*FloatValue)
	case TypeNameString,
		TypeNameBool:
		err = fmt.Errorf("Invalid opareation Greater than on type %v", v1.GetType().GetName())
	//case TypeNameArray:
	//case TypeNameDictionary:
	default:
		panic("Unhandeled type " + v1.GetType().GetName())
	}
	return
}

func OpCmpLessThen(v1, v2 Value) (isTrue bool, err error) {
	isTrue = false
	err = checkSameType(v1, v2)
	if err != nil {
		return
	}
	switch v1.GetType().GetName() {
	case TypeNameByte:
		isTrue = *v1.(*ByteValue) < *v2.(*ByteValue)
	case TypeNameInt:
		isTrue = *v1.(*IntValue) < *v2.(*IntValue)
	case TypeNameUint:
		isTrue = *v1.(*UintValue) < *v2.(*UintValue)
	case TypeNameFloat:
		isTrue = *v1.(*FloatValue) < *v2.(*FloatValue)
	case TypeNameString,
		TypeNameBool:
		err = fmt.Errorf("Invalid opareation Greater than on type %v", v1.GetType().GetName())
	//case TypeNameArray:
	//case TypeNameDictionary:
	default:
		panic("Unhandeled type " + v1.GetType().GetName())
	}
	return
}

func OpCmpIsType(v Value, t Type) bool {
	return t == v.GetType()
}
