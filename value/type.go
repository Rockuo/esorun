package value

type Type interface {
	getName() TypeName
	isValid() bool
}

type ScalarType struct {
	name TypeName
}

func (st ScalarType) getName() TypeName {
	return st.name
}

func (st ScalarType) isValid() bool {
	return st.name != TypeNameArray && st.name != TypeNameDictionary
}

type ArrayType struct {
	ScalarType
	itemType Type
}

func (at ArrayType) getName() string {
	return string(at.name)
}

func (at ArrayType) isValid() bool {
	return at.name == TypeNameArray && at.itemType.isValid()
}

type DictionaryType struct {
	ScalarType
	keyType  ScalarType
	itemType Type
}

func (dt DictionaryType) getName() string {
	return string(dt.name)
}

func (dt DictionaryType) isValid() bool {
	return dt.name == TypeNameDictionary &&
		dt.keyType.isValid() &&
		dt.itemType.isValid()
}
