package value

import "fmt"

type IntValue int

func NewIntValue(val int) *IntValue {
	v := IntValue(val)
	return &v
}

func (*IntValue) GetType() Type {
	return BaseType{TypeNameInt}
}

func (v *IntValue) String() string {
	return fmt.Sprintf("%d", int(*v))
}

func (v *IntValue) SetValue(i interface{}) {
	*v = IntValue(i.(int))
}

func (v *IntValue) GetValue() interface{} {
	return int(*v)
}
