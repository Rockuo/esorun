package value

type Value interface {
	getType() Type
	toString() string

	setValue(v interface{})
	getValue() interface{}

	opAdd(v Value) error
	opSub(v Value) error
	opMul(v Value) error
	opDiv(v Value) error

	cmpEquals(v Value) (isTrue bool, err error)
	cmpNotEquals(v Value) (isTrue bool, err error)
	cmpGreaterThan(v Value) (isTrue bool, err error)
	cmpLessTan(v Value) (isTrue bool, err error)
	cmpIsType(t TypeName) bool
}
