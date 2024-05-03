package operation

import (
	"fmt"
	"strconv"
)
import . "esorun/value"

func Cast(value Value, newType Type) (res Value, err error) {
	if value.GetType() == newType {
		return value, nil
	}

	// from string
	if value.GetType().GetName() == TypeNameString {
		switch newType.GetName() {
		case TypeNameByte:
			if len(value.String()) == 1 {
				arr := []byte(value.String())
				return NewByteValue(arr[0]), nil
			}
			return nil, fmt.Errorf("string must have length 1 to covnert to byte")
		case TypeNameInt:
			v, err := strconv.ParseInt(value.String(), 10, 64)
			if err != nil {
				return nil, err
			}
			return NewIntValue(int(v)), nil
		case TypeNameUint:
			v, err := strconv.ParseUint(value.String(), 10, 64)
			if err != nil {
				return nil, err
			}
			return NewUintValue(uint(v)), nil
		case TypeNameFloat:
			v, err := strconv.ParseFloat(value.String(), 64)
			if err != nil {
				return nil, err
			}
			return NewFloatValue(v), nil
		case TypeNameBool:
			if s := value.String(); s == "true" {
				return NewBoolValue(true), nil
			} else if s := value.String(); s == "false" {
				return NewBoolValue(true), nil
			}
			return nil, fmt.Errorf("bool or string must be true or false")
		}
	}

	// non complex types
	switch newType.GetName() {
	case TypeNameString:
		return NewStringValue(value.String()), nil
	case TypeNameByte:
		switch value.GetType().GetName() {
		case TypeNameInt:
			return NewByteValue(byte(value.GetValue().(int))), nil
		case TypeNameUint:
			return NewByteValue(byte(value.GetValue().(uint))), nil
		case TypeNameFloat:
			return NewByteValue(byte(value.GetValue().(float64))), nil
		}
		return nil, fmt.Errorf(
			"%v cannot be converted to byte",
			value.GetType().GetName(),
		)
	case TypeNameInt:
		switch value.GetType().GetName() {
		case TypeNameByte:
			return NewIntValue(int(value.GetValue().(byte))), nil
		case TypeNameUint:
			return NewIntValue(int(value.GetValue().(uint))), nil
		case TypeNameFloat:
			return NewIntValue(int(value.GetValue().(float64))), nil
		}
		return nil, fmt.Errorf(
			"%v cannot be converted to int",
			value.GetType().GetName(),
		)
	case TypeNameUint:
		switch value.GetType().GetName() {
		case TypeNameByte:
			return NewUintValue(uint(value.GetValue().(byte))), nil
		case TypeNameInt:
			return NewUintValue(uint(value.GetValue().(int))), nil
		case TypeNameFloat:
			return NewUintValue(uint(value.GetValue().(float64))), nil
		}
		return nil, fmt.Errorf(
			"%v cannot be converted to uint",
			value.GetType().GetName(),
		)
	case TypeNameFloat:
		switch value.GetType().GetName() {
		case TypeNameByte:
			return NewFloatValue(float64(value.GetValue().(byte))), nil
		case TypeNameInt:
			return NewFloatValue(float64(value.GetValue().(int))), nil
		case TypeNameUint:
			return NewFloatValue(float64(value.GetValue().(uint))), nil
		}
		return nil, fmt.Errorf(
			"%v cannot be converted to float",
			value.GetType().GetName(),
		)
	}

	return nil, fmt.Errorf(
		"%v cannot be converted to %v",
		value.GetType().GetName(),
		newType.GetName(),
	)
}
