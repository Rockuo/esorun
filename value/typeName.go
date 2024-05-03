package value

type TypeName string

const (
	TypeNameByte       = "byte"
	TypeNameString     = "string"
	TypeNameInt        = "int"
	TypeNameUint       = "uint"
	TypeNameFloat      = "float"
	TypeNameBool       = "bool"
	TypeNameArray      = "array"
	TypeNameDictionary = "dictionary"
)

func (tn TypeName) isValid() bool {
	switch tn {
	case TypeNameByte,
		TypeNameString,
		TypeNameInt,
		TypeNameUint,
		TypeNameFloat,
		TypeNameBool,
		TypeNameArray,
		TypeNameDictionary:
		return true
	}
	return false
}
