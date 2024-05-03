package value

import "fmt"

type ByteValue byte

func NewByteValue(val byte) *ByteValue {
	v := ByteValue(val)
	return &v
}

func (ByteValue) getType() Type {
	return ScalarType{TypeNameByte}
}

func (v ByteValue) toString() string {
	return fmt.Sprintf("%c", v)
}

func (v *ByteValue) setValue(i interface{}) {
	*v = ByteValue(i.(byte))
}

func (v ByteValue) getValue() interface{} {
	return byte(v)
}

func (v *ByteValue) opAdd(value Value) error {
	var val, ok = value.(*ByteValue)
	if !ok {
		return fmt.Errorf("Values are not same type")
	}
	*v += *val
	return nil
}

func (v *ByteValue) opSub(value Value) error {
	var val, ok = value.(*ByteValue)
	if !ok {
		return fmt.Errorf("Values are not same type")
	}
	*v -= *val
	return nil
}

func (v *ByteValue) opMul(value Value) error {
	var val, ok = value.(*ByteValue)
	if !ok {
		return fmt.Errorf("Values are not same type")
	}
	*v *= *val
	return nil
}

func (v *ByteValue) opDiv(value Value) error {
	var val, ok = value.(*ByteValue)
	if !ok {
		return fmt.Errorf("Values are not same type")
	}
	*v /= *val
	return nil
}

func (v ByteValue) cmpEquals(value Value) (isTrue bool, err error) {
	var val, ok = value.(*ByteValue)
	if !ok {
		return false, fmt.Errorf("Values are not same type")
	}
	return v == *val, nil
}

func (v ByteValue) cmpNotEquals(value Value) (isTrue bool, err error) {
	var val, ok = value.(*ByteValue)
	if !ok {
		return false, fmt.Errorf("Values are not same type")
	}
	return v != *val, nil
}

func (v ByteValue) cmpGreaterThan(value Value) (isTrue bool, err error) {
	var val, ok = value.(*ByteValue)
	if !ok {
		return false, fmt.Errorf("Values are not same type")
	}
	return v > *val, nil
}

func (v ByteValue) cmpLessTan(value Value) (isTrue bool, err error) {
	var val, ok = value.(*ByteValue)
	if !ok {
		return false, fmt.Errorf("Values are not same type")
	}
	return v < *val, nil
}

func (ByteValue) cmpIsType(t TypeName) bool {
	return t == TypeNameByte
}
