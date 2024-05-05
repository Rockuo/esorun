package instructions

const (
	TypeVariantByte       = "byte"
	TypeVariantBool       = "bool"
	TypeVariantString     = "string"
	TypeVariantInt        = "int"
	TypeVariantUint       = "uint"
	TypeVariantFloat      = "float"
	TypeVariantArray      = "array"
	TypeVariantDictionary = "dictionary"
)

type TypeVariant string

const (
	KeyTypeInt    = "int"
	KeyTypeString = "string"
)

type KeyType string

type Type struct {
	Variant   TypeVariant
	KeyType   *KeyType
	ValueType *Type
}
