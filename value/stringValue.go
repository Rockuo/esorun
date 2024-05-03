package value

import "fmt"

type StringValue string

func NewStringValue(val string) *StringValue {
	v := StringValue(val)
	return &v
}

func (StringValue) getType() Type {
	return ScalarType{TypeNameString}
}

func (v *StringValue) toString() string {
	return string(*v)
}

func (v *StringValue) setValue(i interface{}) {
	*v = StringValue(i.(string))
}

func (v StringValue) getValue() interface{} {
	return string(v)
}

func (v *StringValue) opAdd(value Value) error {
	var val, ok = value.(*StringValue)
	if !ok {
		return fmt.Errorf("Values are not same type")
	}
	*v += *val
	return nil
}

func (v *StringValue) opSub(value Value) error {
	return fmt.Errorf("Invalid opareation")
}

func (v *StringValue) opMul(value Value) error {
	return fmt.Errorf("Invalid opareation")
}

func (v *StringValue) opDiv(value Value) error {
	return fmt.Errorf("Invalid opareation")
}

func (v StringValue) cmpEquals(value Value) (isTrue bool, err error) {
	var val, ok = value.(*StringValue)
	if !ok {
		return false, fmt.Errorf("Values are not same type")
	}
	return v == *val, nil
}

func (v StringValue) cmpNotEquals(value Value) (isTrue bool, err error) {
	isFalse, err := v.cmpEquals(value)
	if err != nil {
		return false, err
	}
	return !isFalse, err
}

func (v StringValue) cmpGreaterThan(value Value) (isTrue bool, err error) {
	return false, fmt.Errorf("Invalid opareation")
}

func (v StringValue) cmpLessTan(value Value) (isTrue bool, err error) {
	return false, fmt.Errorf("Invalid opareation")
}

func (StringValue) cmpIsType(t TypeName) bool {
	return t == TypeNameString
}
