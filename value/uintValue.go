package value

import "fmt"

type UintValue int

func NewUintValue(val uint) *UintValue {
	v := UintValue(val)
	return &v
}

func (*UintValue) GetType() Type {
	return BaseType{TypeNameUint}
}

func (v *UintValue) String() string {
	return fmt.Sprintf("%d", *v)
}

func (v *UintValue) SetValue(i interface{}) {
	*v = UintValue(i.(uint))
}

func (v *UintValue) GetValue() interface{} {
	return uint(*v)
}
