package value

import "fmt"

type ByteValue byte

func NewByteValue(val byte) *ByteValue {
	v := ByteValue(val)
	return &v
}

func (*ByteValue) GetType() Type {
	return BaseType{TypeNameByte}
}

func (v *ByteValue) String() string {
	return fmt.Sprintf("%c", byte(*v))
}

func (v *ByteValue) SetValue(i interface{}) {
	*v = ByteValue(i.(byte))
}

func (v *ByteValue) GetValue() interface{} {
	return byte(*v)
}
