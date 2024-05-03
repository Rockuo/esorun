package value

type Value interface {
	GetType() Type
	String() string

	SetValue(v interface{})
	GetValue() interface{}
}
