package value

import "fmt"

type BoolValue bool

func NewBoolValue(val bool) *BoolValue {
	v := BoolValue(val)
	return &v
}

func (*BoolValue) GetType() Type {
	return BaseType{TypeNameBool}
}

func (v *BoolValue) String() string {
	return fmt.Sprintf("%v", bool(*v))
}

func (v *BoolValue) SetValue(i interface{}) {
	*v = BoolValue(i.(bool))
}

func (v *BoolValue) GetValue() interface{} {
	return bool(*v)
}
