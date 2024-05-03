package value

import "fmt"

type FloatValue float64

func NewFloatValue(val float64) *FloatValue {
	v := FloatValue(val)
	return &v
}

func (*FloatValue) GetType() Type {
	return BaseType{TypeNameFloat}
}

func (v *FloatValue) String() string {
	return fmt.Sprintf("%f", float64(*v))
}

func (v *FloatValue) SetValue(i interface{}) {
	*v = FloatValue(i.(float64))
}

func (v *FloatValue) GetValue() interface{} {
	return float64(*v)
}
