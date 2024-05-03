package value

type StringValue string

func NewStringValue(val string) *StringValue {
	v := StringValue(val)
	return &v
}

func (*StringValue) GetType() Type {
	return BaseType{TypeNameString}
}

func (v *StringValue) String() string {
	return string(*v)
}

func (v *StringValue) SetValue(i interface{}) {
	*v = StringValue(i.(string))
}

func (v *StringValue) GetValue() interface{} {
	return string(*v)
}
